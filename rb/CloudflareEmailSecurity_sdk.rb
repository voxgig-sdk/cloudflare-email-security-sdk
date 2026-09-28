# CloudflareEmailSecurity SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'CloudflareEmailSecurity_types'


class CloudflareEmailSecuritySDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = CloudflareEmailSecurityUtility.new
    @_utility = utility

    config = CloudflareEmailSecurityConfig.shared_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features in the resolved order (make_options puts an explicit array
    # order first, else defaults to test-first). Ordering matters: the `test`
    # feature installs the base mock transport and the transport features
    # (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
    # must be added before them to sit at the base of the chain.
    feature_opts = CloudflareEmailSecurityHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      featureorder = VoxgigStruct.getpath(@options, "__derived__.featureorder")
      if featureorder.is_a?(Array)
        featureorder.each do |fname|
          fopts = CloudflareEmailSecurityHelpers.to_map(feature_opts[fname])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, CloudflareEmailSecurityFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    CloudflareEmailSecurityUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = CloudflareEmailSecurityHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = CloudflareEmailSecurityHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = CloudflareEmailSecurityHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = CloudflareEmailSecuritySpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    # make_fetch_def returns a (fetchdef, err) tuple; destructure it and
    # return just the fetchdef Hash (raising on error) so callers — including
    # direct(), which indexes fetchdef["url"] — receive a Hash, mirroring the
    # ts/py prepare().
    fetchdef, fd_err = utility.make_fetch_def.call(ctx)
    raise fd_err if fd_err

    fetchdef
  end

  # Raw endpoint access is operator-controllable, like every entity op.
  # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  # either one reaches the same endpoint.
  def direct(fetchargs = {})
    return op_denied("direct") unless op_allowed?("direct")

    raw_request(fetchargs)
  end

  # Is this raw-access op permitted by the SDK's allow.op option?
  def op_allowed?(op)
    allow_op = VoxgigStruct.getpath(@options, "allow.op")
    allow_op.is_a?(String) && allow_op.include?(op)
  end

  def op_denied(op)
    allow_op = VoxgigStruct.getpath(@options, "allow.op")
    {
      "ok" => false,
      "err" => CloudflareEmailSecurityError.new(
        "#{op}_allow",
        "CloudflareEmailSecuritySDK: #{op}: operation not allowed by" \
        " SDK option allow.op value: \"#{allow_op}\""),
    }
  end

  # Ungated request path shared by direct and graphql, each of which checks
  # its own allow.op token first. Separate, rather than a flag on fetchargs:
  # a caller-supplied marker would let anyone opt straight back out of the
  # gate by passing it.
  def raw_request(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue CloudflareEmailSecurityError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = CloudflareEmailSecurityHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = CloudflareEmailSecurityHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end

  # Raw GraphQL access: the pressure valve that makes the generated surface's
  # deliberate omissions (per-call selection sets, typed filter builders,
  # batching, subscriptions) livable — the whole schema stays reachable.
  #
  # Thin wrapper over the same prepare/fetch path direct uses, with the one
  # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
  # as a top-level `errors` array, so status alone would report a failed
  # query as ok.
  #
  # NOTE: like direct, this bypasses the feature pipeline — no retry,
  # ratelimit or paging features apply.
  def graphql(query, variables = nil, ctrl = nil)
    return op_denied("graphql") unless op_allowed?("graphql")

    res = raw_request({
      "method" => "POST",
      "headers" => { "content-type" => "application/json" },
      "body" => { "query" => query, "variables" => variables || {} },
      "ctrl" => ctrl || {},
    })

    # Errors are read BEFORE any status check: a GraphQL parse or validation
    # failure comes back as HTTP 400 carrying the standard { errors: [...] }
    # body, and the raw path represents a non-2xx as ok:false with no err —
    # so returning early on status would discard the server's own
    # diagnostics, which are the only useful part of that response.
    errors = VoxgigStruct.getpath(res, "data.errors")

    if errors.is_a?(Array) && !errors.empty?
      first = errors[0].is_a?(Hash) ? errors[0] : {}
      msg = first["message"]
      msg = "graphql error" if msg.nil? || msg.to_s.empty?
      res["ok"] = false
      res["err"] = CloudflareEmailSecurityError.new(
        "graphql_error", "CloudflareEmailSecuritySDK: graphql: #{msg}")
      res["graphql"] = errors
    end

    res
  end


  # Canonical facade: client.EmailAccountRulesPlanResponseSingle.list / client.EmailAccountRulesPlanResponseSingle.load({ "id" => ... })
  def EmailAccountRulesPlanResponseSingle(data = nil)
    require_relative 'entity/email_account_rules_plan_response_single_entity'
    EmailAccountRulesPlanResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailAccountRulesResponseCollection.list / client.EmailAccountRulesResponseCollection.load({ "id" => ... })
  def EmailAccountRulesResponseCollection(data = nil)
    require_relative 'entity/email_account_rules_response_collection_entity'
    EmailAccountRulesResponseCollectionEntity.new(self, data)
  end


  # Canonical facade: client.EmailAuth.list / client.EmailAuth.load({ "id" => ... })
  def EmailAuth(data = nil)
    require_relative 'entity/email_auth_entity'
    EmailAuthEntity.new(self, data)
  end


  # Canonical facade: client.EmailCatchAllRuleResponseSingle.list / client.EmailCatchAllRuleResponseSingle.load({ "id" => ... })
  def EmailCatchAllRuleResponseSingle(data = nil)
    require_relative 'entity/email_catch_all_rule_response_single_entity'
    EmailCatchAllRuleResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailDestinationAddressResponseSingle.list / client.EmailDestinationAddressResponseSingle.load({ "id" => ... })
  def EmailDestinationAddressResponseSingle(data = nil)
    require_relative 'entity/email_destination_address_response_single_entity'
    EmailDestinationAddressResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailDestinationAddressesResponseCollection.list / client.EmailDestinationAddressesResponseCollection.load({ "id" => ... })
  def EmailDestinationAddressesResponseCollection(data = nil)
    require_relative 'entity/email_destination_addresses_response_collection_entity'
    EmailDestinationAddressesResponseCollectionEntity.new(self, data)
  end


  # Canonical facade: client.EmailDnsSettingsResponseCollection.list / client.EmailDnsSettingsResponseCollection.load({ "id" => ... })
  def EmailDnsSettingsResponseCollection(data = nil)
    require_relative 'entity/email_dns_settings_response_collection_entity'
    EmailDnsSettingsResponseCollectionEntity.new(self, data)
  end


  # Canonical facade: client.EmailEmailSettingsResponseSingle.list / client.EmailEmailSettingsResponseSingle.load({ "id" => ... })
  def EmailEmailSettingsResponseSingle(data = nil)
    require_relative 'entity/email_email_settings_response_single_entity'
    EmailEmailSettingsResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailRoutingSuppression.list / client.EmailRoutingSuppression.load({ "id" => ... })
  def EmailRoutingSuppression(data = nil)
    require_relative 'entity/email_routing_suppression_entity'
    EmailRoutingSuppressionEntity.new(self, data)
  end


  # Canonical facade: client.EmailRuleResponseSingle.list / client.EmailRuleResponseSingle.load({ "id" => ... })
  def EmailRuleResponseSingle(data = nil)
    require_relative 'entity/email_rule_response_single_entity'
    EmailRuleResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailRulesResponseCollection.list / client.EmailRulesResponseCollection.load({ "id" => ... })
  def EmailRulesResponseCollection(data = nil)
    require_relative 'entity/email_rules_response_collection_entity'
    EmailRulesResponseCollectionEntity.new(self, data)
  end


  # Canonical facade: client.EmailSecurity.list / client.EmailSecurity.load({ "id" => ... })
  def EmailSecurity(data = nil)
    require_relative 'entity/email_security_entity'
    EmailSecurityEntity.new(self, data)
  end


  # Canonical facade: client.EmailSecurityAnalytics.list / client.EmailSecurityAnalytics.load({ "id" => ... })
  def EmailSecurityAnalytics(data = nil)
    require_relative 'entity/email_security_analytics_entity'
    EmailSecurityAnalyticsEntity.new(self, data)
  end


  # Canonical facade: client.EmailSecuritySetting.list / client.EmailSecuritySetting.load({ "id" => ... })
  def EmailSecuritySetting(data = nil)
    require_relative 'entity/email_security_setting_entity'
    EmailSecuritySettingEntity.new(self, data)
  end


  # Canonical facade: client.EmailSending.list / client.EmailSending.load({ "id" => ... })
  def EmailSending(data = nil)
    require_relative 'entity/email_sending_entity'
    EmailSendingEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingLimitsResponseSingle.list / client.EmailSendingLimitsResponseSingle.load({ "id" => ... })
  def EmailSendingLimitsResponseSingle(data = nil)
    require_relative 'entity/email_sending_limits_response_single_entity'
    EmailSendingLimitsResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingReputationResponseSingle.list / client.EmailSendingReputationResponseSingle.load({ "id" => ... })
  def EmailSendingReputationResponseSingle(data = nil)
    require_relative 'entity/email_sending_reputation_response_single_entity'
    EmailSendingReputationResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSubdomainDnsStatus.list / client.EmailSendingSubdomainDnsStatus.load({ "id" => ... })
  def EmailSendingSubdomainDnsStatus(data = nil)
    require_relative 'entity/email_sending_subdomain_dns_status_entity'
    EmailSendingSubdomainDnsStatusEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSubdomainPreview.list / client.EmailSendingSubdomainPreview.load({ "id" => ... })
  def EmailSendingSubdomainPreview(data = nil)
    require_relative 'entity/email_sending_subdomain_preview_entity'
    EmailSendingSubdomainPreviewEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSubdomainReputationComplaint.list / client.EmailSendingSubdomainReputationComplaint.load({ "id" => ... })
  def EmailSendingSubdomainReputationComplaint(data = nil)
    require_relative 'entity/email_sending_subdomain_reputation_complaint_entity'
    EmailSendingSubdomainReputationComplaintEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSubdomainResponseSingle.list / client.EmailSendingSubdomainResponseSingle.load({ "id" => ... })
  def EmailSendingSubdomainResponseSingle(data = nil)
    require_relative 'entity/email_sending_subdomain_response_single_entity'
    EmailSendingSubdomainResponseSingleEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSubdomainsResponseCollection.list / client.EmailSendingSubdomainsResponseCollection.load({ "id" => ... })
  def EmailSendingSubdomainsResponseCollection(data = nil)
    require_relative 'entity/email_sending_subdomains_response_collection_entity'
    EmailSendingSubdomainsResponseCollectionEntity.new(self, data)
  end


  # Canonical facade: client.EmailSendingSuppression.list / client.EmailSendingSuppression.load({ "id" => ... })
  def EmailSendingSuppression(data = nil)
    require_relative 'entity/email_sending_suppression_entity'
    EmailSendingSuppressionEntity.new(self, data)
  end


  # Canonical facade: client.Subdomain.list / client.Subdomain.load({ "id" => ... })
  def Subdomain(data = nil)
    require_relative 'entity/subdomain_entity'
    SubdomainEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = CloudflareEmailSecuritySDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
