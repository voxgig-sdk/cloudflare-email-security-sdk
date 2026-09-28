<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class CloudflareEmailSecuritySDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new CloudflareEmailSecurityUtility();
        $this->_utility = $utility;

        $config = CloudflareEmailSecurityConfig::shared_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Feature INSTANCES supplied at construction (the station adopt
        // path) are read from the RAW construction options - extend is
        // consumed exactly once, here; make_options strips it from the
        // processed map so options_map() stays clean data.
        $extend_val = is_array($options["extend"] ?? null) ? $options["extend"] : [];

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = CloudflareEmailSecurityHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = CloudflareEmailSecurityHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        // An active name with no generated feature class is
                        // legal when an extend-supplied instance carries that
                        // name (station's adopt path): the instance is added
                        // below, positioned by its own __after__ entry, so
                        // skip it here rather than add a BaseFeature stray
                        // that would silently shift feature positions.
                        if (!CloudflareEmailSecurityFeatures::has_feature($fname)) {
                            foreach ($extend_val as $ef) {
                                if (is_object($ef) && method_exists($ef, 'get_name')
                                    && $fname === $ef->get_name()) {
                                    continue 2;
                                }
                            }
                        }
                        ($utility->feature_add)($this->_rootctx, CloudflareEmailSecurityFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        foreach ($extend_val as $f) {
            if (is_object($f) && method_exists($f, 'get_name')) {
                ($utility->feature_add)($this->_rootctx, $f);
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return CloudflareEmailSecurityUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = CloudflareEmailSecurityHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = CloudflareEmailSecurityHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = CloudflareEmailSecurityHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new CloudflareEmailSecuritySpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens,
    // since either one reaches the same endpoint.
    public function direct(array $fetchargs = []): mixed
    {
        if (!$this->op_allowed("direct")) {
            return $this->op_denied("direct");
        }

        return $this->raw_request($fetchargs);
    }

    // Is this raw-access op permitted by the SDK's allow.op option?
    private function op_allowed(string $op): bool
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return is_string($allow_op) && str_contains($allow_op, $op);
    }

    private function op_denied(string $op): array
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return [
            "ok" => false,
            "err" => new CloudflareEmailSecurityError($op . "_allow",
                "CloudflareEmailSecuritySDK: " . $op . ": operation not allowed by" .
                " SDK option allow.op value: \"" . (string)$allow_op . "\""),
        ];
    }

    // Ungated request path shared by direct and graphql, each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    private function raw_request(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = CloudflareEmailSecurityHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = CloudflareEmailSecurityHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }

    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path direct uses, with the
    // one thing raw direct cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report
    // a failed query as ok.
    //
    // NOTE: like direct, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    public function graphql(string $query, ?array $variables = null, ?array $ctrl = null): mixed
    {
        if (!$this->op_allowed("graphql")) {
            return $this->op_denied("graphql");
        }

        $res = $this->raw_request([
            "method" => "POST",
            "headers" => ["content-type" => "application/json"],
            "body" => ["query" => $query, "variables" => $variables ?? []],
            "ctrl" => $ctrl ?? [],
        ]);

        if (!is_array($res)) {
            return $res;
        }

        // Errors are read BEFORE any status check: a GraphQL parse or
        // validation failure comes back as HTTP 400 carrying the standard
        // { errors: [...] } body, and the raw path represents a non-2xx as
        // ok:false with no err — so returning early on status would discard
        // the server's own diagnostics, which are the only useful part of
        // that response.
        $errors = Struct::getpath($res, "data.errors");

        if (is_array($errors) && 0 < count($errors)) {
            $first = is_array($errors[0]) ? $errors[0] : [];
            $msg = $first["message"] ?? "";
            if (!is_string($msg) || "" === $msg) {
                $msg = "graphql error";
            }
            $res["ok"] = false;
            $res["err"] = new CloudflareEmailSecurityError("graphql_error",
                "CloudflareEmailSecuritySDK: graphql: " . $msg);
            $res["graphql"] = $errors;
        }

        return $res;
    }


    private $_email_account_rules_plan_response_single = null;

    // Canonical facade: $client->EmailAccountRulesPlanResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_account_rules_plan_response_single()
    // resolves here too.
    public function EmailAccountRulesPlanResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_account_rules_plan_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_account_rules_plan_response_single === null) {
                $this->_email_account_rules_plan_response_single = new EmailAccountRulesPlanResponseSingleEntity($this, null);
            }
            return $this->_email_account_rules_plan_response_single;
        }
        return new EmailAccountRulesPlanResponseSingleEntity($this, $data);
    }


    private $_email_account_rules_response_collection = null;

    // Canonical facade: $client->EmailAccountRulesResponseCollection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_account_rules_response_collection()
    // resolves here too.
    public function EmailAccountRulesResponseCollection($data = null)
    {
        require_once __DIR__ . '/entity/email_account_rules_response_collection_entity.php';
        if ($data === null) {
            if ($this->_email_account_rules_response_collection === null) {
                $this->_email_account_rules_response_collection = new EmailAccountRulesResponseCollectionEntity($this, null);
            }
            return $this->_email_account_rules_response_collection;
        }
        return new EmailAccountRulesResponseCollectionEntity($this, $data);
    }


    private $_email_auth = null;

    // Canonical facade: $client->EmailAuth()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_auth()
    // resolves here too.
    public function EmailAuth($data = null)
    {
        require_once __DIR__ . '/entity/email_auth_entity.php';
        if ($data === null) {
            if ($this->_email_auth === null) {
                $this->_email_auth = new EmailAuthEntity($this, null);
            }
            return $this->_email_auth;
        }
        return new EmailAuthEntity($this, $data);
    }


    private $_email_catch_all_rule_response_single = null;

    // Canonical facade: $client->EmailCatchAllRuleResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_catch_all_rule_response_single()
    // resolves here too.
    public function EmailCatchAllRuleResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_catch_all_rule_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_catch_all_rule_response_single === null) {
                $this->_email_catch_all_rule_response_single = new EmailCatchAllRuleResponseSingleEntity($this, null);
            }
            return $this->_email_catch_all_rule_response_single;
        }
        return new EmailCatchAllRuleResponseSingleEntity($this, $data);
    }


    private $_email_destination_address_response_single = null;

    // Canonical facade: $client->EmailDestinationAddressResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_destination_address_response_single()
    // resolves here too.
    public function EmailDestinationAddressResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_destination_address_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_destination_address_response_single === null) {
                $this->_email_destination_address_response_single = new EmailDestinationAddressResponseSingleEntity($this, null);
            }
            return $this->_email_destination_address_response_single;
        }
        return new EmailDestinationAddressResponseSingleEntity($this, $data);
    }


    private $_email_destination_addresses_response_collection = null;

    // Canonical facade: $client->EmailDestinationAddressesResponseCollection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_destination_addresses_response_collection()
    // resolves here too.
    public function EmailDestinationAddressesResponseCollection($data = null)
    {
        require_once __DIR__ . '/entity/email_destination_addresses_response_collection_entity.php';
        if ($data === null) {
            if ($this->_email_destination_addresses_response_collection === null) {
                $this->_email_destination_addresses_response_collection = new EmailDestinationAddressesResponseCollectionEntity($this, null);
            }
            return $this->_email_destination_addresses_response_collection;
        }
        return new EmailDestinationAddressesResponseCollectionEntity($this, $data);
    }


    private $_email_dns_settings_response_collection = null;

    // Canonical facade: $client->EmailDnsSettingsResponseCollection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_dns_settings_response_collection()
    // resolves here too.
    public function EmailDnsSettingsResponseCollection($data = null)
    {
        require_once __DIR__ . '/entity/email_dns_settings_response_collection_entity.php';
        if ($data === null) {
            if ($this->_email_dns_settings_response_collection === null) {
                $this->_email_dns_settings_response_collection = new EmailDnsSettingsResponseCollectionEntity($this, null);
            }
            return $this->_email_dns_settings_response_collection;
        }
        return new EmailDnsSettingsResponseCollectionEntity($this, $data);
    }


    private $_email_email_settings_response_single = null;

    // Canonical facade: $client->EmailEmailSettingsResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_email_settings_response_single()
    // resolves here too.
    public function EmailEmailSettingsResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_email_settings_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_email_settings_response_single === null) {
                $this->_email_email_settings_response_single = new EmailEmailSettingsResponseSingleEntity($this, null);
            }
            return $this->_email_email_settings_response_single;
        }
        return new EmailEmailSettingsResponseSingleEntity($this, $data);
    }


    private $_email_routing_suppression = null;

    // Canonical facade: $client->EmailRoutingSuppression()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_routing_suppression()
    // resolves here too.
    public function EmailRoutingSuppression($data = null)
    {
        require_once __DIR__ . '/entity/email_routing_suppression_entity.php';
        if ($data === null) {
            if ($this->_email_routing_suppression === null) {
                $this->_email_routing_suppression = new EmailRoutingSuppressionEntity($this, null);
            }
            return $this->_email_routing_suppression;
        }
        return new EmailRoutingSuppressionEntity($this, $data);
    }


    private $_email_rule_response_single = null;

    // Canonical facade: $client->EmailRuleResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_rule_response_single()
    // resolves here too.
    public function EmailRuleResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_rule_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_rule_response_single === null) {
                $this->_email_rule_response_single = new EmailRuleResponseSingleEntity($this, null);
            }
            return $this->_email_rule_response_single;
        }
        return new EmailRuleResponseSingleEntity($this, $data);
    }


    private $_email_rules_response_collection = null;

    // Canonical facade: $client->EmailRulesResponseCollection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_rules_response_collection()
    // resolves here too.
    public function EmailRulesResponseCollection($data = null)
    {
        require_once __DIR__ . '/entity/email_rules_response_collection_entity.php';
        if ($data === null) {
            if ($this->_email_rules_response_collection === null) {
                $this->_email_rules_response_collection = new EmailRulesResponseCollectionEntity($this, null);
            }
            return $this->_email_rules_response_collection;
        }
        return new EmailRulesResponseCollectionEntity($this, $data);
    }


    private $_email_security = null;

    // Canonical facade: $client->EmailSecurity()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_security()
    // resolves here too.
    public function EmailSecurity($data = null)
    {
        require_once __DIR__ . '/entity/email_security_entity.php';
        if ($data === null) {
            if ($this->_email_security === null) {
                $this->_email_security = new EmailSecurityEntity($this, null);
            }
            return $this->_email_security;
        }
        return new EmailSecurityEntity($this, $data);
    }


    private $_email_security_analytics = null;

    // Canonical facade: $client->EmailSecurityAnalytics()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_security_analytics()
    // resolves here too.
    public function EmailSecurityAnalytics($data = null)
    {
        require_once __DIR__ . '/entity/email_security_analytics_entity.php';
        if ($data === null) {
            if ($this->_email_security_analytics === null) {
                $this->_email_security_analytics = new EmailSecurityAnalyticsEntity($this, null);
            }
            return $this->_email_security_analytics;
        }
        return new EmailSecurityAnalyticsEntity($this, $data);
    }


    private $_email_security_setting = null;

    // Canonical facade: $client->EmailSecuritySetting()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_security_setting()
    // resolves here too.
    public function EmailSecuritySetting($data = null)
    {
        require_once __DIR__ . '/entity/email_security_setting_entity.php';
        if ($data === null) {
            if ($this->_email_security_setting === null) {
                $this->_email_security_setting = new EmailSecuritySettingEntity($this, null);
            }
            return $this->_email_security_setting;
        }
        return new EmailSecuritySettingEntity($this, $data);
    }


    private $_email_sending = null;

    // Canonical facade: $client->EmailSending()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending()
    // resolves here too.
    public function EmailSending($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_entity.php';
        if ($data === null) {
            if ($this->_email_sending === null) {
                $this->_email_sending = new EmailSendingEntity($this, null);
            }
            return $this->_email_sending;
        }
        return new EmailSendingEntity($this, $data);
    }


    private $_email_sending_limits_response_single = null;

    // Canonical facade: $client->EmailSendingLimitsResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_limits_response_single()
    // resolves here too.
    public function EmailSendingLimitsResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_limits_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_sending_limits_response_single === null) {
                $this->_email_sending_limits_response_single = new EmailSendingLimitsResponseSingleEntity($this, null);
            }
            return $this->_email_sending_limits_response_single;
        }
        return new EmailSendingLimitsResponseSingleEntity($this, $data);
    }


    private $_email_sending_reputation_response_single = null;

    // Canonical facade: $client->EmailSendingReputationResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_reputation_response_single()
    // resolves here too.
    public function EmailSendingReputationResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_reputation_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_sending_reputation_response_single === null) {
                $this->_email_sending_reputation_response_single = new EmailSendingReputationResponseSingleEntity($this, null);
            }
            return $this->_email_sending_reputation_response_single;
        }
        return new EmailSendingReputationResponseSingleEntity($this, $data);
    }


    private $_email_sending_subdomain_dns_status = null;

    // Canonical facade: $client->EmailSendingSubdomainDnsStatus()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_subdomain_dns_status()
    // resolves here too.
    public function EmailSendingSubdomainDnsStatus($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_subdomain_dns_status_entity.php';
        if ($data === null) {
            if ($this->_email_sending_subdomain_dns_status === null) {
                $this->_email_sending_subdomain_dns_status = new EmailSendingSubdomainDnsStatusEntity($this, null);
            }
            return $this->_email_sending_subdomain_dns_status;
        }
        return new EmailSendingSubdomainDnsStatusEntity($this, $data);
    }


    private $_email_sending_subdomain_preview = null;

    // Canonical facade: $client->EmailSendingSubdomainPreview()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_subdomain_preview()
    // resolves here too.
    public function EmailSendingSubdomainPreview($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_subdomain_preview_entity.php';
        if ($data === null) {
            if ($this->_email_sending_subdomain_preview === null) {
                $this->_email_sending_subdomain_preview = new EmailSendingSubdomainPreviewEntity($this, null);
            }
            return $this->_email_sending_subdomain_preview;
        }
        return new EmailSendingSubdomainPreviewEntity($this, $data);
    }


    private $_email_sending_subdomain_reputation_complaint = null;

    // Canonical facade: $client->EmailSendingSubdomainReputationComplaint()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_subdomain_reputation_complaint()
    // resolves here too.
    public function EmailSendingSubdomainReputationComplaint($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_subdomain_reputation_complaint_entity.php';
        if ($data === null) {
            if ($this->_email_sending_subdomain_reputation_complaint === null) {
                $this->_email_sending_subdomain_reputation_complaint = new EmailSendingSubdomainReputationComplaintEntity($this, null);
            }
            return $this->_email_sending_subdomain_reputation_complaint;
        }
        return new EmailSendingSubdomainReputationComplaintEntity($this, $data);
    }


    private $_email_sending_subdomain_response_single = null;

    // Canonical facade: $client->EmailSendingSubdomainResponseSingle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_subdomain_response_single()
    // resolves here too.
    public function EmailSendingSubdomainResponseSingle($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_subdomain_response_single_entity.php';
        if ($data === null) {
            if ($this->_email_sending_subdomain_response_single === null) {
                $this->_email_sending_subdomain_response_single = new EmailSendingSubdomainResponseSingleEntity($this, null);
            }
            return $this->_email_sending_subdomain_response_single;
        }
        return new EmailSendingSubdomainResponseSingleEntity($this, $data);
    }


    private $_email_sending_subdomains_response_collection = null;

    // Canonical facade: $client->EmailSendingSubdomainsResponseCollection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_subdomains_response_collection()
    // resolves here too.
    public function EmailSendingSubdomainsResponseCollection($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_subdomains_response_collection_entity.php';
        if ($data === null) {
            if ($this->_email_sending_subdomains_response_collection === null) {
                $this->_email_sending_subdomains_response_collection = new EmailSendingSubdomainsResponseCollectionEntity($this, null);
            }
            return $this->_email_sending_subdomains_response_collection;
        }
        return new EmailSendingSubdomainsResponseCollectionEntity($this, $data);
    }


    private $_email_sending_suppression = null;

    // Canonical facade: $client->EmailSendingSuppression()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->email_sending_suppression()
    // resolves here too.
    public function EmailSendingSuppression($data = null)
    {
        require_once __DIR__ . '/entity/email_sending_suppression_entity.php';
        if ($data === null) {
            if ($this->_email_sending_suppression === null) {
                $this->_email_sending_suppression = new EmailSendingSuppressionEntity($this, null);
            }
            return $this->_email_sending_suppression;
        }
        return new EmailSendingSuppressionEntity($this, $data);
    }


    private $_subdomain = null;

    // Canonical facade: $client->Subdomain()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->subdomain()
    // resolves here too.
    public function Subdomain($data = null)
    {
        require_once __DIR__ . '/entity/subdomain_entity.php';
        if ($data === null) {
            if ($this->_subdomain === null) {
                $this->_subdomain = new SubdomainEntity($this, null);
            }
            return $this->_subdomain;
        }
        return new SubdomainEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new CloudflareEmailSecuritySDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
