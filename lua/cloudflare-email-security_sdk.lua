-- CloudflareEmailSecurity SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Typed-model annotations (LuaLS ---@class); empty at runtime.
require("cloudflare-email-security_types")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local CloudflareEmailSecuritySDK = {}
CloudflareEmailSecuritySDK.__index = CloudflareEmailSecuritySDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

CloudflareEmailSecuritySDK._make_feature = _make_feature


function CloudflareEmailSecuritySDK.new(options)
  local self = setmetatable({}, CloudflareEmailSecuritySDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config_shared")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features in the resolved order (make_options puts an explicit list
  -- order first, else defaults to test-first). Ordering matters: the `test`
  -- feature installs the base mock transport and the transport features
  -- (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
  -- must be added before them to sit at the base of the chain.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local featureorder = vs.getpath(self.options, "__derived__.featureorder")
    if type(featureorder) == "table" then
      for _, fname in ipairs(featureorder) do
        local fopts = helpers.to_map(feature_opts[fname])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- CONSUMED, not kept. `extend` holds feature INSTANCES, and every shipped
  -- feature's init stores `self.client = ctx.client` - so leaving the list
  -- in self.options makes the options map CYCLIC (client.options.extend[1]
  -- .client == client), and options_map()'s vs.clone, which has no cycle
  -- guard, blew the stack on the first prepare_auth of any client built with
  -- an extend feature. The instances live on self.features from here on,
  -- which is the only place anything reads them; the SAME table is
  -- self._rootctx.options, so the root context loses the key too.
  self.options["extend"] = nil

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

    -- feature: test


  return self
end


function CloudflareEmailSecuritySDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function CloudflareEmailSecuritySDK:get_utility()
  return Utility.copy(self._utility)
end


function CloudflareEmailSecuritySDK:get_root_ctx()
  return self._rootctx
end


function CloudflareEmailSecuritySDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


-- Raw endpoint access is operator-controllable, like every entity op.
-- Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
-- either one reaches the same endpoint.
function CloudflareEmailSecuritySDK:direct(fetchargs)
  if not self:_op_allowed("direct") then
    return self:_op_denied("direct"), nil
  end

  return self:_raw_request(fetchargs)
end


-- Is this raw-access op permitted by the SDK's allow.op option?
function CloudflareEmailSecuritySDK:_op_allowed(op)
  local allow = vs.getpath(self.options, "allow.op")
  return type(allow) == "string" and allow:find(op, 1, true) ~= nil
end


function CloudflareEmailSecuritySDK:_op_denied(op)
  local allow = vs.getpath(self.options, "allow.op")
  if type(allow) ~= "string" then allow = "" end
  return {
    ok = false,
    err = "CloudflareEmailSecuritySDK: " .. op .. ": operation not allowed by" ..
      " SDK option allow.op value: \"" .. allow .. "\"",
  }
end


-- Ungated request path shared by direct and graphql, each of which checks its
-- own allow.op token first. Private, rather than a flag on fetchargs: a
-- caller-supplied marker would let anyone opt straight back out of the gate
-- by passing it.
function CloudflareEmailSecuritySDK:_raw_request(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end


-- Raw GraphQL access: the pressure valve that makes the generated surface's
-- deliberate omissions (per-call selection sets, typed filter builders,
-- batching, subscriptions) livable — the whole schema stays reachable.
--
-- Thin wrapper over the same prepare/fetch path direct uses, with the one
-- thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200 as
-- a top-level `errors` array, so status alone would report a failed query as
-- ok.
--
-- NOTE: like direct, this bypasses the feature pipeline — no retry, ratelimit
-- or paging features apply.
function CloudflareEmailSecuritySDK:graphql(query, variables, ctrl)
  if not self:_op_allowed("graphql") then
    return self:_op_denied("graphql"), nil
  end

  local res, err = self:_raw_request({
    method = "POST",
    headers = { ["content-type"] = "application/json" },
    body = {
      query = query,
      variables = type(variables) == "table" and variables or {},
    },
    ctrl = type(ctrl) == "table" and ctrl or {},
  })

  if err ~= nil or type(res) ~= "table" then
    return res, err
  end

  -- Errors are read BEFORE any status check: a GraphQL parse or validation
  -- failure comes back as HTTP 400 carrying the standard { errors = {...} }
  -- body, and the raw path represents a non-2xx as ok=false with no err — so
  -- returning early on status would discard the server's own diagnostics,
  -- which are the only useful part of that response.
  local errors = vs.getpath(res, "data.errors")

  if type(errors) == "table" and 0 < #errors then
    local msg = vs.getprop(errors[1], "message")
    if type(msg) ~= "string" or msg == "" then
      msg = "graphql error"
    end
    res.ok = false
    res.err = "CloudflareEmailSecuritySDK: graphql: " .. msg
    res.graphql = errors
  end

  return res, nil
end



-- Idiomatic facade: client:EmailAccountRulesPlanResponseSingle():list() / client:EmailAccountRulesPlanResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailAccountRulesPlanResponseSingle(data)
  local EntityMod = require("entity.email_account_rules_plan_response_single_entity")
  if data == nil then
    if self._email_account_rules_plan_response_single == nil then
      self._email_account_rules_plan_response_single = EntityMod.new(self, nil)
    end
    return self._email_account_rules_plan_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailAccountRulesResponseCollection():list() / client:EmailAccountRulesResponseCollection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailAccountRulesResponseCollection(data)
  local EntityMod = require("entity.email_account_rules_response_collection_entity")
  if data == nil then
    if self._email_account_rules_response_collection == nil then
      self._email_account_rules_response_collection = EntityMod.new(self, nil)
    end
    return self._email_account_rules_response_collection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailAuth():list() / client:EmailAuth():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailAuth(data)
  local EntityMod = require("entity.email_auth_entity")
  if data == nil then
    if self._email_auth == nil then
      self._email_auth = EntityMod.new(self, nil)
    end
    return self._email_auth
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailCatchAllRuleResponseSingle():list() / client:EmailCatchAllRuleResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailCatchAllRuleResponseSingle(data)
  local EntityMod = require("entity.email_catch_all_rule_response_single_entity")
  if data == nil then
    if self._email_catch_all_rule_response_single == nil then
      self._email_catch_all_rule_response_single = EntityMod.new(self, nil)
    end
    return self._email_catch_all_rule_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailDestinationAddressResponseSingle():list() / client:EmailDestinationAddressResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailDestinationAddressResponseSingle(data)
  local EntityMod = require("entity.email_destination_address_response_single_entity")
  if data == nil then
    if self._email_destination_address_response_single == nil then
      self._email_destination_address_response_single = EntityMod.new(self, nil)
    end
    return self._email_destination_address_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailDestinationAddressesResponseCollection():list() / client:EmailDestinationAddressesResponseCollection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailDestinationAddressesResponseCollection(data)
  local EntityMod = require("entity.email_destination_addresses_response_collection_entity")
  if data == nil then
    if self._email_destination_addresses_response_collection == nil then
      self._email_destination_addresses_response_collection = EntityMod.new(self, nil)
    end
    return self._email_destination_addresses_response_collection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailDnsSettingsResponseCollection():list() / client:EmailDnsSettingsResponseCollection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailDnsSettingsResponseCollection(data)
  local EntityMod = require("entity.email_dns_settings_response_collection_entity")
  if data == nil then
    if self._email_dns_settings_response_collection == nil then
      self._email_dns_settings_response_collection = EntityMod.new(self, nil)
    end
    return self._email_dns_settings_response_collection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailEmailSettingsResponseSingle():list() / client:EmailEmailSettingsResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailEmailSettingsResponseSingle(data)
  local EntityMod = require("entity.email_email_settings_response_single_entity")
  if data == nil then
    if self._email_email_settings_response_single == nil then
      self._email_email_settings_response_single = EntityMod.new(self, nil)
    end
    return self._email_email_settings_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailRoutingSuppression():list() / client:EmailRoutingSuppression():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailRoutingSuppression(data)
  local EntityMod = require("entity.email_routing_suppression_entity")
  if data == nil then
    if self._email_routing_suppression == nil then
      self._email_routing_suppression = EntityMod.new(self, nil)
    end
    return self._email_routing_suppression
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailRuleResponseSingle():list() / client:EmailRuleResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailRuleResponseSingle(data)
  local EntityMod = require("entity.email_rule_response_single_entity")
  if data == nil then
    if self._email_rule_response_single == nil then
      self._email_rule_response_single = EntityMod.new(self, nil)
    end
    return self._email_rule_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailRulesResponseCollection():list() / client:EmailRulesResponseCollection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailRulesResponseCollection(data)
  local EntityMod = require("entity.email_rules_response_collection_entity")
  if data == nil then
    if self._email_rules_response_collection == nil then
      self._email_rules_response_collection = EntityMod.new(self, nil)
    end
    return self._email_rules_response_collection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSecurity():list() / client:EmailSecurity():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSecurity(data)
  local EntityMod = require("entity.email_security_entity")
  if data == nil then
    if self._email_security == nil then
      self._email_security = EntityMod.new(self, nil)
    end
    return self._email_security
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSecurityAnalytics():list() / client:EmailSecurityAnalytics():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSecurityAnalytics(data)
  local EntityMod = require("entity.email_security_analytics_entity")
  if data == nil then
    if self._email_security_analytics == nil then
      self._email_security_analytics = EntityMod.new(self, nil)
    end
    return self._email_security_analytics
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSecuritySetting():list() / client:EmailSecuritySetting():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSecuritySetting(data)
  local EntityMod = require("entity.email_security_setting_entity")
  if data == nil then
    if self._email_security_setting == nil then
      self._email_security_setting = EntityMod.new(self, nil)
    end
    return self._email_security_setting
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSending():list() / client:EmailSending():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSending(data)
  local EntityMod = require("entity.email_sending_entity")
  if data == nil then
    if self._email_sending == nil then
      self._email_sending = EntityMod.new(self, nil)
    end
    return self._email_sending
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingLimitsResponseSingle():list() / client:EmailSendingLimitsResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingLimitsResponseSingle(data)
  local EntityMod = require("entity.email_sending_limits_response_single_entity")
  if data == nil then
    if self._email_sending_limits_response_single == nil then
      self._email_sending_limits_response_single = EntityMod.new(self, nil)
    end
    return self._email_sending_limits_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingReputationResponseSingle():list() / client:EmailSendingReputationResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingReputationResponseSingle(data)
  local EntityMod = require("entity.email_sending_reputation_response_single_entity")
  if data == nil then
    if self._email_sending_reputation_response_single == nil then
      self._email_sending_reputation_response_single = EntityMod.new(self, nil)
    end
    return self._email_sending_reputation_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSubdomainDnsStatus():list() / client:EmailSendingSubdomainDnsStatus():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSubdomainDnsStatus(data)
  local EntityMod = require("entity.email_sending_subdomain_dns_status_entity")
  if data == nil then
    if self._email_sending_subdomain_dns_status == nil then
      self._email_sending_subdomain_dns_status = EntityMod.new(self, nil)
    end
    return self._email_sending_subdomain_dns_status
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSubdomainPreview():list() / client:EmailSendingSubdomainPreview():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSubdomainPreview(data)
  local EntityMod = require("entity.email_sending_subdomain_preview_entity")
  if data == nil then
    if self._email_sending_subdomain_preview == nil then
      self._email_sending_subdomain_preview = EntityMod.new(self, nil)
    end
    return self._email_sending_subdomain_preview
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSubdomainReputationComplaint():list() / client:EmailSendingSubdomainReputationComplaint():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSubdomainReputationComplaint(data)
  local EntityMod = require("entity.email_sending_subdomain_reputation_complaint_entity")
  if data == nil then
    if self._email_sending_subdomain_reputation_complaint == nil then
      self._email_sending_subdomain_reputation_complaint = EntityMod.new(self, nil)
    end
    return self._email_sending_subdomain_reputation_complaint
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSubdomainResponseSingle():list() / client:EmailSendingSubdomainResponseSingle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSubdomainResponseSingle(data)
  local EntityMod = require("entity.email_sending_subdomain_response_single_entity")
  if data == nil then
    if self._email_sending_subdomain_response_single == nil then
      self._email_sending_subdomain_response_single = EntityMod.new(self, nil)
    end
    return self._email_sending_subdomain_response_single
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSubdomainsResponseCollection():list() / client:EmailSendingSubdomainsResponseCollection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSubdomainsResponseCollection(data)
  local EntityMod = require("entity.email_sending_subdomains_response_collection_entity")
  if data == nil then
    if self._email_sending_subdomains_response_collection == nil then
      self._email_sending_subdomains_response_collection = EntityMod.new(self, nil)
    end
    return self._email_sending_subdomains_response_collection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EmailSendingSuppression():list() / client:EmailSendingSuppression():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:EmailSendingSuppression(data)
  local EntityMod = require("entity.email_sending_suppression_entity")
  if data == nil then
    if self._email_sending_suppression == nil then
      self._email_sending_suppression = EntityMod.new(self, nil)
    end
    return self._email_sending_suppression
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Subdomain():list() / client:Subdomain():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function CloudflareEmailSecuritySDK:Subdomain(data)
  local EntityMod = require("entity.subdomain_entity")
  if data == nil then
    if self._subdomain == nil then
      self._subdomain = EntityMod.new(self, nil)
    end
    return self._subdomain
  end
  return EntityMod.new(self, data)
end




function CloudflareEmailSecuritySDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = CloudflareEmailSecuritySDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return CloudflareEmailSecuritySDK
