# CloudflareEmailSecurity SDK

from cloudflareemailsecurity_sdk.utility.voxgig_struct import voxgig_struct as vs
from cloudflareemailsecurity_sdk.core.utility_type import CloudflareEmailSecurityUtility
from cloudflareemailsecurity_sdk.core.spec import CloudflareEmailSecuritySpec
from cloudflareemailsecurity_sdk.core import helpers

# Load utility registration (populates Utility._registrar)
from cloudflareemailsecurity_sdk.utility import register

# Load features
from cloudflareemailsecurity_sdk.feature.base_feature import CloudflareEmailSecurityBaseFeature
from cloudflareemailsecurity_sdk.features import _has_feature, _make_feature


class CloudflareEmailSecuritySDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = CloudflareEmailSecurityUtility()
        self._utility = utility

        from cloudflareemailsecurity_sdk.config import shared_config
        config = shared_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        # Extension feature INSTANCES come from the RAW construction
        # options - extend is consumed exactly once, here. make_options
        # strips the key before cloning (vs.clone flattens arbitrary
        # objects), so self.options never carries the instances.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        extend = options.get("extend") if isinstance(options, dict) else None
        if not isinstance(extend, list):
            extend = []
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        # An active name with no generated feature class is
                        # legal when an extend-supplied instance carries that
                        # name (station's adopt path): the instance is added
                        # below, positioned by its own __after__ entry, so
                        # skip it here rather than add a BaseFeature stray
                        # that would silently shift feature positions.
                        if not _has_feature(fname) and any(
                            fname == (f.get("name") if isinstance(f, dict)
                                      else getattr(f, "name", None))
                            for f in extend
                        ):
                            continue
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        for f in extend:
            if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return CloudflareEmailSecurityUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = CloudflareEmailSecuritySpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    # Raw endpoint access is operator-controllable, like every entity op.
    # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    # either one reaches the same endpoint.
    def direct(self, fetchargs=None):
        if not self._op_allowed("direct"):
            return self._op_denied("direct")

        return self._raw_request(fetchargs)

    # Is this raw-access op permitted by the SDK's allow.op option?
    def _op_allowed(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return isinstance(allow_op, str) and op in allow_op

    def _op_denied(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return {
            "ok": False,
            "err": Exception(
                "CloudflareEmailSecuritySDK: " + op + ": operation not allowed by"
                ' SDK option allow.op value: "' + str(allow_op) + '"'),
        }

    # Ungated request path shared by direct and graphql, each of which checks
    # its own allow.op token first. Private, rather than a flag on fetchargs:
    # a caller-supplied marker would let anyone opt straight back out of the
    # gate by passing it.
    def _raw_request(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }

    # Raw GraphQL access: the pressure valve that makes the generated
    # surface's deliberate omissions (per-call selection sets, typed filter
    # builders, batching, subscriptions) livable — the whole schema stays
    # reachable.
    #
    # Thin wrapper over the same prepare/fetch path direct uses, with the one
    # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
    # as a top-level `errors` array, so status alone would report a failed
    # query as ok.
    #
    # NOTE: like direct, this bypasses the feature pipeline — no retry,
    # ratelimit or paging features apply.
    def graphql(self, query, variables=None, ctrl=None):
        if not self._op_allowed("graphql"):
            return self._op_denied("graphql")

        res = self._raw_request({
            "method": "POST",
            "headers": {"content-type": "application/json"},
            "body": {"query": query, "variables": variables or {}},
            "ctrl": ctrl or {},
        })

        # Errors are read BEFORE any status check: a GraphQL parse or
        # validation failure comes back as HTTP 400 carrying the standard
        # { errors: [...] } body, and the raw path represents a non-2xx as
        # ok:False with no err — so returning early on status would discard
        # the server's own diagnostics, which are the only useful part of
        # that response.
        errors = vs.getpath(res, "data.errors")

        if isinstance(errors, list) and 0 < len(errors):
            first = errors[0] if isinstance(errors[0], dict) else {}
            msg = first.get("message") or "graphql error"
            res["ok"] = False
            res["err"] = Exception("CloudflareEmailSecuritySDK: graphql: " + str(msg))
            res["graphql"] = errors

        return res


    def EmailAccountRulesPlanResponseSingle(self, data=None) -> "EmailAccountRulesPlanResponseSingleEntity":
        """Entity factory: client.EmailAccountRulesPlanResponseSingle().list() / client.EmailAccountRulesPlanResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_account_rules_plan_response_single_entity import EmailAccountRulesPlanResponseSingleEntity
        return EmailAccountRulesPlanResponseSingleEntity(self, data)


    def EmailAccountRulesResponseCollection(self, data=None) -> "EmailAccountRulesResponseCollectionEntity":
        """Entity factory: client.EmailAccountRulesResponseCollection().list() / client.EmailAccountRulesResponseCollection().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_account_rules_response_collection_entity import EmailAccountRulesResponseCollectionEntity
        return EmailAccountRulesResponseCollectionEntity(self, data)


    def EmailAuth(self, data=None) -> "EmailAuthEntity":
        """Entity factory: client.EmailAuth().list() / client.EmailAuth().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_auth_entity import EmailAuthEntity
        return EmailAuthEntity(self, data)


    def EmailCatchAllRuleResponseSingle(self, data=None) -> "EmailCatchAllRuleResponseSingleEntity":
        """Entity factory: client.EmailCatchAllRuleResponseSingle().list() / client.EmailCatchAllRuleResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_catch_all_rule_response_single_entity import EmailCatchAllRuleResponseSingleEntity
        return EmailCatchAllRuleResponseSingleEntity(self, data)


    def EmailDestinationAddressResponseSingle(self, data=None) -> "EmailDestinationAddressResponseSingleEntity":
        """Entity factory: client.EmailDestinationAddressResponseSingle().list() / client.EmailDestinationAddressResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_destination_address_response_single_entity import EmailDestinationAddressResponseSingleEntity
        return EmailDestinationAddressResponseSingleEntity(self, data)


    def EmailDestinationAddressesResponseCollection(self, data=None) -> "EmailDestinationAddressesResponseCollectionEntity":
        """Entity factory: client.EmailDestinationAddressesResponseCollection().list() / client.EmailDestinationAddressesResponseCollection().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_destination_addresses_response_collection_entity import EmailDestinationAddressesResponseCollectionEntity
        return EmailDestinationAddressesResponseCollectionEntity(self, data)


    def EmailDnsSettingsResponseCollection(self, data=None) -> "EmailDnsSettingsResponseCollectionEntity":
        """Entity factory: client.EmailDnsSettingsResponseCollection().list() / client.EmailDnsSettingsResponseCollection().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_dns_settings_response_collection_entity import EmailDnsSettingsResponseCollectionEntity
        return EmailDnsSettingsResponseCollectionEntity(self, data)


    def EmailEmailSettingsResponseSingle(self, data=None) -> "EmailEmailSettingsResponseSingleEntity":
        """Entity factory: client.EmailEmailSettingsResponseSingle().list() / client.EmailEmailSettingsResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_email_settings_response_single_entity import EmailEmailSettingsResponseSingleEntity
        return EmailEmailSettingsResponseSingleEntity(self, data)


    def EmailRoutingSuppression(self, data=None) -> "EmailRoutingSuppressionEntity":
        """Entity factory: client.EmailRoutingSuppression().list() / client.EmailRoutingSuppression().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_routing_suppression_entity import EmailRoutingSuppressionEntity
        return EmailRoutingSuppressionEntity(self, data)


    def EmailRuleResponseSingle(self, data=None) -> "EmailRuleResponseSingleEntity":
        """Entity factory: client.EmailRuleResponseSingle().list() / client.EmailRuleResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_rule_response_single_entity import EmailRuleResponseSingleEntity
        return EmailRuleResponseSingleEntity(self, data)


    def EmailRulesResponseCollection(self, data=None) -> "EmailRulesResponseCollectionEntity":
        """Entity factory: client.EmailRulesResponseCollection().list() / client.EmailRulesResponseCollection().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_rules_response_collection_entity import EmailRulesResponseCollectionEntity
        return EmailRulesResponseCollectionEntity(self, data)


    def EmailSecurity(self, data=None) -> "EmailSecurityEntity":
        """Entity factory: client.EmailSecurity().list() / client.EmailSecurity().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_security_entity import EmailSecurityEntity
        return EmailSecurityEntity(self, data)


    def EmailSecurityAnalytics(self, data=None) -> "EmailSecurityAnalyticsEntity":
        """Entity factory: client.EmailSecurityAnalytics().list() / client.EmailSecurityAnalytics().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_security_analytics_entity import EmailSecurityAnalyticsEntity
        return EmailSecurityAnalyticsEntity(self, data)


    def EmailSecuritySetting(self, data=None) -> "EmailSecuritySettingEntity":
        """Entity factory: client.EmailSecuritySetting().list() / client.EmailSecuritySetting().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_security_setting_entity import EmailSecuritySettingEntity
        return EmailSecuritySettingEntity(self, data)


    def EmailSending(self, data=None) -> "EmailSendingEntity":
        """Entity factory: client.EmailSending().list() / client.EmailSending().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_entity import EmailSendingEntity
        return EmailSendingEntity(self, data)


    def EmailSendingLimitsResponseSingle(self, data=None) -> "EmailSendingLimitsResponseSingleEntity":
        """Entity factory: client.EmailSendingLimitsResponseSingle().list() / client.EmailSendingLimitsResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_limits_response_single_entity import EmailSendingLimitsResponseSingleEntity
        return EmailSendingLimitsResponseSingleEntity(self, data)


    def EmailSendingReputationResponseSingle(self, data=None) -> "EmailSendingReputationResponseSingleEntity":
        """Entity factory: client.EmailSendingReputationResponseSingle().list() / client.EmailSendingReputationResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_reputation_response_single_entity import EmailSendingReputationResponseSingleEntity
        return EmailSendingReputationResponseSingleEntity(self, data)


    def EmailSendingSubdomainDnsStatus(self, data=None) -> "EmailSendingSubdomainDnsStatusEntity":
        """Entity factory: client.EmailSendingSubdomainDnsStatus().list() / client.EmailSendingSubdomainDnsStatus().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_dns_status_entity import EmailSendingSubdomainDnsStatusEntity
        return EmailSendingSubdomainDnsStatusEntity(self, data)


    def EmailSendingSubdomainPreview(self, data=None) -> "EmailSendingSubdomainPreviewEntity":
        """Entity factory: client.EmailSendingSubdomainPreview().list() / client.EmailSendingSubdomainPreview().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_preview_entity import EmailSendingSubdomainPreviewEntity
        return EmailSendingSubdomainPreviewEntity(self, data)


    def EmailSendingSubdomainReputationComplaint(self, data=None) -> "EmailSendingSubdomainReputationComplaintEntity":
        """Entity factory: client.EmailSendingSubdomainReputationComplaint().list() / client.EmailSendingSubdomainReputationComplaint().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_reputation_complaint_entity import EmailSendingSubdomainReputationComplaintEntity
        return EmailSendingSubdomainReputationComplaintEntity(self, data)


    def EmailSendingSubdomainResponseSingle(self, data=None) -> "EmailSendingSubdomainResponseSingleEntity":
        """Entity factory: client.EmailSendingSubdomainResponseSingle().list() / client.EmailSendingSubdomainResponseSingle().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_response_single_entity import EmailSendingSubdomainResponseSingleEntity
        return EmailSendingSubdomainResponseSingleEntity(self, data)


    def EmailSendingSubdomainsResponseCollection(self, data=None) -> "EmailSendingSubdomainsResponseCollectionEntity":
        """Entity factory: client.EmailSendingSubdomainsResponseCollection().list() / client.EmailSendingSubdomainsResponseCollection().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_subdomains_response_collection_entity import EmailSendingSubdomainsResponseCollectionEntity
        return EmailSendingSubdomainsResponseCollectionEntity(self, data)


    def EmailSendingSuppression(self, data=None) -> "EmailSendingSuppressionEntity":
        """Entity factory: client.EmailSendingSuppression().list() / client.EmailSendingSuppression().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.email_sending_suppression_entity import EmailSendingSuppressionEntity
        return EmailSendingSuppressionEntity(self, data)


    def Subdomain(self, data=None) -> "SubdomainEntity":
        """Entity factory: client.Subdomain().list() / client.Subdomain().load({"id": ...})."""
        from cloudflareemailsecurity_sdk.entity.subdomain_entity import SubdomainEntity
        return SubdomainEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "CloudflareEmailSecuritySDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from cloudflareemailsecurity_sdk.entity.email_account_rules_plan_response_single_entity import EmailAccountRulesPlanResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_account_rules_response_collection_entity import EmailAccountRulesResponseCollectionEntity
    from cloudflareemailsecurity_sdk.entity.email_auth_entity import EmailAuthEntity
    from cloudflareemailsecurity_sdk.entity.email_catch_all_rule_response_single_entity import EmailCatchAllRuleResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_destination_address_response_single_entity import EmailDestinationAddressResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_destination_addresses_response_collection_entity import EmailDestinationAddressesResponseCollectionEntity
    from cloudflareemailsecurity_sdk.entity.email_dns_settings_response_collection_entity import EmailDnsSettingsResponseCollectionEntity
    from cloudflareemailsecurity_sdk.entity.email_email_settings_response_single_entity import EmailEmailSettingsResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_routing_suppression_entity import EmailRoutingSuppressionEntity
    from cloudflareemailsecurity_sdk.entity.email_rule_response_single_entity import EmailRuleResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_rules_response_collection_entity import EmailRulesResponseCollectionEntity
    from cloudflareemailsecurity_sdk.entity.email_security_entity import EmailSecurityEntity
    from cloudflareemailsecurity_sdk.entity.email_security_analytics_entity import EmailSecurityAnalyticsEntity
    from cloudflareemailsecurity_sdk.entity.email_security_setting_entity import EmailSecuritySettingEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_entity import EmailSendingEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_limits_response_single_entity import EmailSendingLimitsResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_reputation_response_single_entity import EmailSendingReputationResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_dns_status_entity import EmailSendingSubdomainDnsStatusEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_preview_entity import EmailSendingSubdomainPreviewEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_reputation_complaint_entity import EmailSendingSubdomainReputationComplaintEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_subdomain_response_single_entity import EmailSendingSubdomainResponseSingleEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_subdomains_response_collection_entity import EmailSendingSubdomainsResponseCollectionEntity
    from cloudflareemailsecurity_sdk.entity.email_sending_suppression_entity import EmailSendingSuppressionEntity
    from cloudflareemailsecurity_sdk.entity.subdomain_entity import SubdomainEntity
