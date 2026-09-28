# CloudflareEmailSecurity Lua SDK Reference

Complete API reference for the CloudflareEmailSecurity Lua SDK.


## CloudflareEmailSecuritySDK

### Constructor

```lua
local sdk = require("cloudflare-email-security_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle(data)`

Create a new `EmailAccountRulesPlanResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailAccountRulesResponseCollection(data)`

Create a new `EmailAccountRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailAuth(data)`

Create a new `EmailAuth` entity instance. Pass `nil` for no initial data.

#### `EmailCatchAllRuleResponseSingle(data)`

Create a new `EmailCatchAllRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressResponseSingle(data)`

Create a new `EmailDestinationAddressResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressesResponseCollection(data)`

Create a new `EmailDestinationAddressesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailDnsSettingsResponseCollection(data)`

Create a new `EmailDnsSettingsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailEmailSettingsResponseSingle(data)`

Create a new `EmailEmailSettingsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRoutingSuppression(data)`

Create a new `EmailRoutingSuppression` entity instance. Pass `nil` for no initial data.

#### `EmailRuleResponseSingle(data)`

Create a new `EmailRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRulesResponseCollection(data)`

Create a new `EmailRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSecurity(data)`

Create a new `EmailSecurity` entity instance. Pass `nil` for no initial data.

#### `EmailSecurityAnalytics(data)`

Create a new `EmailSecurityAnalytics` entity instance. Pass `nil` for no initial data.

#### `EmailSecuritySetting(data)`

Create a new `EmailSecuritySetting` entity instance. Pass `nil` for no initial data.

#### `EmailSending(data)`

Create a new `EmailSending` entity instance. Pass `nil` for no initial data.

#### `EmailSendingLimitsResponseSingle(data)`

Create a new `EmailSendingLimitsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingReputationResponseSingle(data)`

Create a new `EmailSendingReputationResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainDnsStatus(data)`

Create a new `EmailSendingSubdomainDnsStatus` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainPreview(data)`

Create a new `EmailSendingSubdomainPreview` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainReputationComplaint(data)`

Create a new `EmailSendingSubdomainReputationComplaint` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainResponseSingle(data)`

Create a new `EmailSendingSubdomainResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainsResponseCollection(data)`

Create a new `EmailSendingSubdomainsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSuppression(data)`

Create a new `EmailSendingSuppression` entity instance. Pass `nil` for no initial data.

#### `Subdomain(data)`

Create a new `Subdomain` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## EmailAccountRulesPlanResponseSingleEntity

```lua
local email_account_rules_plan_response_single = client:EmailAccountRulesPlanResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `table` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `table` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailAccountRulesPlanResponseSingle():create({
  account_id = --[[ string ]],
  owner_worker_tag = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailAccountRulesResponseCollectionEntity

```lua
local email_account_rules_response_collection = client:EmailAccountRulesResponseCollection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |
| `result_info` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailAccountRulesResponseCollection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailAuthEntity

```lua
local email_auth = client:EmailAuth(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | No | Enable or disable DMARC reports for this zone |
| `result` | `table` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `boolean` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailAuth():load({ id = "email_auth_id", zone_id = "zone_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailAuth():update({
  zone_id = "zone_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAuthEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailCatchAllRuleResponseSingleEntity

```lua
local email_catch_all_rule_response_single = client:EmailCatchAllRuleResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `table` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `boolean` | No | Routing rule status. |
| `matchers` | `table` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `table` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailCatchAllRuleResponseSingle():load({ zone_id = "zone_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailCatchAllRuleResponseSingle():update({
  zone_id = "zone_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailDestinationAddressResponseSingleEntity

```lua
local email_destination_address_response_single = client:EmailDestinationAddressResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes | The contact email address of the user. |
| `result` | `table` | No |  |
| `status` | `string` | Yes | Destination address status. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailDestinationAddressResponseSingle():create({
  account_id = --[[ string ]],
  email = --[[ string ]],
  status = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailDestinationAddressResponseSingle():load({ account_id = "account_id", destination_address_identifier = "destination_address_identifier" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailDestinationAddressResponseSingle():remove({ account_id = "account_id", destination_address_identifier = "destination_address_identifier" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailDestinationAddressResponseSingle():update({
  account_id = "account_id",
  destination_address_identifier = "destination_address_identifier",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailDestinationAddressesResponseCollectionEntity

```lua
local email_destination_addresses_response_collection = client:EmailDestinationAddressesResponseCollection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |
| `result_info` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailDestinationAddressesResponseCollection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailDnsSettingsResponseCollectionEntity

```lua
local email_dns_settings_response_collection = client:EmailDnsSettingsResponseCollection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailDnsSettingsResponseCollection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailEmailSettingsResponseSingleEntity

```lua
local email_email_settings_response_single = client:EmailEmailSettingsResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | No | State of your zone Email Routing settings. |
| `name` | `string` | No | Domain of your zone. |
| `result` | `table` | No |  |
| `skip_wizard` | `boolean` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `boolean` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailEmailSettingsResponseSingle():create({
  zone_id = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailEmailSettingsResponseSingle():load({ zone_id = "zone_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailEmailSettingsResponseSingle():remove({ zone_id = "zone_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailEmailSettingsResponseSingle():update({
  zone_id = "zone_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailRoutingSuppressionEntity

```lua
local email_routing_suppression = client:EmailRoutingSuppression(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `expires_at` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `reason` | `string` | Yes |  |
| `zones` | `table` | No |  |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `email` | - | - | - | - |
| `expires_at` | - | - | Yes | - |
| `id` | - | - | - | - |
| `reason` | - | - | - | - |
| `zones` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailRoutingSuppression():create({
  account_id = --[[ string ]],
  created_at = --[[ string ]],
  email = --[[ string ]],
  expires_at = --[[ string ]],
  id = --[[ string ]],
  reason = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailRoutingSuppression():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailRoutingSuppression():load({ id = "email_routing_suppression_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailRoutingSuppression():remove({ id = "email_routing_suppression_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRoutingSuppressionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailRuleResponseSingleEntity

```lua
local email_rule_response_single = client:EmailRuleResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `table` | Yes | List actions patterns. |
| `enabled` | `boolean` | No | Routing rule status. |
| `matchers` | `table` | Yes | Matching patterns to forward to your actions. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `number` | No | Priority of the routing rule. |
| `result` | `table` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailRuleResponseSingle():create({
  zone_id = --[[ string ]],
  actions = --[[ table ]],
  matchers = --[[ table ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailRuleResponseSingle():load({ rule_identifier = "rule_identifier", zone_id = "zone_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailRuleResponseSingle():remove({ rule_identifier = "rule_identifier", zone_id = "zone_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailRuleResponseSingle():update({
  rule_identifier = "rule_identifier",
  zone_id = "zone_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailRulesResponseCollectionEntity

```lua
local email_rules_response_collection = client:EmailRulesResponseCollection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |
| `result_info` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailRulesResponseCollection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSecurityEntity

```lua
local email_security = client:EmailSecurity(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `string` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | No | Optional note describing the job. |
| `destination` | `any` | No | Required when action is 'MOVE'. |
| `errors` | `table` | Yes |  |
| `expected_disposition` | `string` | No | Nonfunctional field. |
| `id` | `string` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `table` | Yes | List of message IDs to move. |
| `messages` | `table` | Yes |  |
| `postfix_ids` | `table` | No | Deprecated, use `ids` instead. |
| `result` | `table` | Yes |  |
| `result_info` | `table` | Yes |  |
| `search_params` | `table` | Yes |  |
| `success` | `boolean` | Yes | Whether the API call was successful. |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `action` | - | - | - | - |
| `comment` | - | - | - | - |
| `destination` | - | - | Yes | - |
| `errors` | - | - | - | - |
| `expected_disposition` | - | - | - | - |
| `id` | - | - | - | - |
| `ids` | - | - | - | - |
| `messages` | - | - | - | - |
| `postfix_ids` | - | - | - | - |
| `result` | - | - | - | - |
| `result_info` | - | - | - | - |
| `search_params` | - | - | - | - |
| `success` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSecurity():create({
  account_id = --[[ string ]],
  action = --[[ string ]],
  errors = --[[ table ]],
  id = --[[ string ]],
  ids = --[[ table ]],
  messages = --[[ table ]],
  result = --[[ table ]],
  result_info = --[[ table ]],
  search_params = --[[ table ]],
  success = --[[ boolean ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailSecurity():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSecurity():load({ account_id = "account_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailSecurity():remove({ account_id = "account_id", job_id = "job_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecurityEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSecurityAnalyticsEntity

```lua
local email_security_analytics = client:EmailSecurityAnalytics(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSecurityAnalytics():load({ account_id = "account_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecurityAnalyticsEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSecuritySettingEntity

```lua
local email_security_setting = client:EmailSecuritySetting(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `table` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `string` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `table` | Yes | IDs of the allow policies to delete. |
| `domain` | `string` | Yes | The email domain to protect. |
| `drop_dispositions` | `table` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `table` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `number` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `table` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `table` | Yes | Allow policies to create. |
| `puts` | `table` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `table` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `boolean` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `boolean` | No | Require TLS on outbound connections. |
| `result` | `table` | No | An impersonation registry entry. |
| `transport` | `string` | No | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `allowed_delivery_modes` | - | - | - | - | - |
| `comments` | - | - | - | - | - |
| `deletes` | - | - | - | - | - |
| `domain` | - | - | - | - | - |
| `drop_dispositions` | - | - | - | - | - |
| `folder` | - | - | - | - | - |
| `integration_id` | - | - | - | - | - |
| `ip_restrictions` | - | - | - | - | - |
| `lookback_hops` | - | - | - | - | - |
| `patches` | - | - | - | - | - |
| `pattern` | - | - | - | - | - |
| `posts` | - | - | - | - | - |
| `puts` | - | - | - | - | - |
| `regions` | - | - | - | - | - |
| `require_tls_inbound` | - | - | - | - | - |
| `require_tls_outbound` | - | - | - | - | - |
| `result` | Yes | - | - | - | - |
| `transport` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSecuritySetting():create({
  account_id = --[[ string ]],
  allowed_delivery_modes = --[[ table ]],
  deletes = --[[ table ]],
  domain = --[[ string ]],
  drop_dispositions = --[[ table ]],
  ip_restrictions = --[[ table ]],
  patches = --[[ table ]],
  pattern = --[[ string ]],
  posts = --[[ table ]],
  puts = --[[ table ]],
  regions = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailSecuritySetting():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSecuritySetting():load({ account_id = "account_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailSecuritySetting():remove({ account_id = "account_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailSecuritySetting():update({
  account_id = "account_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecuritySettingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingEntity

```lua
local email_sending = client:EmailSending(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `table` | No | File attachments and inline images. |
| `bcc` | `any` | No | Blind carbon copy recipient(s). |
| `cc` | `any` | No | Carbon copy recipient(s). |
| `errors` | `table` | Yes |  |
| `from` | `any` | Yes | Sender email address. |
| `headers` | `table` | No | Custom email headers as key-value pairs. |
| `html` | `string` | No | HTML body of the email. |
| `messages` | `table` | Yes |  |
| `mime_message` | `string` | Yes | The full MIME-encoded email message. |
| `recipients` | `table` | Yes | List of recipient email addresses. |
| `reply_to` | `any` | No | Reply-to address. |
| `result` | `table` | Yes |  |
| `result_info` | `table` | Yes |  |
| `subject` | `string` | Yes | Email subject line. |
| `success` | `boolean` | Yes |  |
| `text` | `string` | No | Plain text body of the email. |
| `to` | `any` | No | Recipient(s). |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSending():create({
  account_id = --[[ string ]],
  errors = --[[ table ]],
  from = --[[ any ]],
  messages = --[[ table ]],
  mime_message = --[[ string ]],
  recipients = --[[ table ]],
  result = --[[ table ]],
  result_info = --[[ table ]],
  subject = --[[ string ]],
  success = --[[ boolean ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSending():load({ account_id = "account_id", message_id = "message_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingLimitsResponseSingleEntity

```lua
local email_sending_limits_response_single = client:EmailSendingLimitsResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingLimitsResponseSingle():load({ account_id = "account_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingReputationResponseSingleEntity

```lua
local email_sending_reputation_response_single = client:EmailSendingReputationResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingReputationResponseSingle():load({ account_id = "account_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSubdomainDnsStatusEntity

```lua
local email_sending_subdomain_dns_status = client:EmailSendingSubdomainDnsStatus(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSendingSubdomainDnsStatus():create({
  subdomain_id = --[[ string ]],
  zone_id = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingSubdomainDnsStatus():load({ subdomain_id = "subdomain_id", zone_id = "zone_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSubdomainPreviewEntity

```lua
local email_sending_subdomain_preview = client:EmailSendingSubdomainPreview(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The domain name within the zone. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSendingSubdomainPreview():create({
  zone_id = --[[ string ]],
  name = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSubdomainReputationComplaintEntity

```lua
local email_sending_subdomain_reputation_complaint = client:EmailSendingSubdomainReputationComplaint(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingSubdomainReputationComplaint():load({ subdomain_id = "subdomain_id", zone_id = "zone_id", end_at = "end_at", start_at = "start_at" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSubdomainResponseSingleEntity

```lua
local email_sending_subdomain_response_single = client:EmailSendingSubdomainResponseSingle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `boolean` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | Yes | The domain name within the zone. |
| `preview_enabled` | `boolean` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `any` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSendingSubdomainResponseSingle():create({
  zone_id = --[[ string ]],
  name = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingSubdomainResponseSingle():load({ subdomain_id = "subdomain_id", zone_id = "zone_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailSendingSubdomainResponseSingle():update({
  subdomain_id = "subdomain_id",
  zone_id = "zone_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSubdomainsResponseCollectionEntity

```lua
local email_sending_subdomains_response_collection = client:EmailSendingSubdomainsResponseCollection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailSendingSubdomainsResponseCollection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EmailSendingSuppressionEntity

```lua
local email_sending_suppression = client:EmailSendingSuppression(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | When the suppression was created. |
| `email` | `string` | Yes | The email address to suppress. |
| `errors` | `table` | Yes |  |
| `expires_at` | `string` | Yes | Expiration timestamp for the suppression. |
| `id` | `string` | Yes | Unique identifier for this suppression. |
| `items` | `table` | Yes | Suppressions to import. |
| `messages` | `table` | Yes |  |
| `note` | `string` | No | Advisory note for this suppression. |
| `read_only` | `boolean` | Yes | Whether clients may mutate this suppression. |
| `reason` | `string` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `table` | Yes |  |
| `scope` | `any` | No | Where the suppression applies. |
| `success` | `boolean` | Yes |  |
| `zones` | `table` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `email` | - | - | - | - | - |
| `errors` | - | - | - | - | - |
| `expires_at` | - | - | Yes | Yes | - |
| `id` | - | - | - | - | - |
| `items` | - | - | - | - | - |
| `messages` | - | - | - | - | - |
| `note` | - | - | - | - | - |
| `read_only` | - | - | - | - | - |
| `reason` | - | - | - | - | - |
| `result` | - | - | - | - | - |
| `scope` | - | - | - | - | - |
| `success` | - | - | - | - | - |
| `zones` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EmailSendingSuppression():create({
  account_id = --[[ string ]],
  created_at = --[[ string ]],
  email = --[[ string ]],
  errors = --[[ table ]],
  expires_at = --[[ string ]],
  id = --[[ string ]],
  items = --[[ table ]],
  messages = --[[ table ]],
  read_only = --[[ boolean ]],
  reason = --[[ string ]],
  result = --[[ table ]],
  success = --[[ boolean ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EmailSendingSuppression():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EmailSendingSuppression():load({ id = "email_sending_suppression_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EmailSendingSuppression():remove({ id = "email_sending_suppression_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EmailSendingSuppression():update({
  id = "email_sending_suppression_id",
  account_id = "account_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSuppressionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SubdomainEntity

```lua
local subdomain = client:Subdomain(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Subdomain():remove({ id = "id", zone_id = "zone_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubdomainEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

