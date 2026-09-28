# CloudflareEmailSecurity Ruby SDK Reference

Complete API reference for the CloudflareEmailSecurity Ruby SDK.


## CloudflareEmailSecuritySDK

### Constructor

```ruby
require_relative 'CloudflareEmailSecurity_sdk'

client = CloudflareEmailSecuritySDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflareEmailSecuritySDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = CloudflareEmailSecuritySDK.test
```


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle(data = nil)`

Create a new `EmailAccountRulesPlanResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailAccountRulesResponseCollection(data = nil)`

Create a new `EmailAccountRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailAuth(data = nil)`

Create a new `EmailAuth` entity instance. Pass `nil` for no initial data.

#### `EmailCatchAllRuleResponseSingle(data = nil)`

Create a new `EmailCatchAllRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressResponseSingle(data = nil)`

Create a new `EmailDestinationAddressResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressesResponseCollection(data = nil)`

Create a new `EmailDestinationAddressesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailDnsSettingsResponseCollection(data = nil)`

Create a new `EmailDnsSettingsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailEmailSettingsResponseSingle(data = nil)`

Create a new `EmailEmailSettingsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRoutingSuppression(data = nil)`

Create a new `EmailRoutingSuppression` entity instance. Pass `nil` for no initial data.

#### `EmailRuleResponseSingle(data = nil)`

Create a new `EmailRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRulesResponseCollection(data = nil)`

Create a new `EmailRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSecurity(data = nil)`

Create a new `EmailSecurity` entity instance. Pass `nil` for no initial data.

#### `EmailSecurityAnalytics(data = nil)`

Create a new `EmailSecurityAnalytics` entity instance. Pass `nil` for no initial data.

#### `EmailSecuritySetting(data = nil)`

Create a new `EmailSecuritySetting` entity instance. Pass `nil` for no initial data.

#### `EmailSending(data = nil)`

Create a new `EmailSending` entity instance. Pass `nil` for no initial data.

#### `EmailSendingLimitsResponseSingle(data = nil)`

Create a new `EmailSendingLimitsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingReputationResponseSingle(data = nil)`

Create a new `EmailSendingReputationResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainDnsStatus(data = nil)`

Create a new `EmailSendingSubdomainDnsStatus` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainPreview(data = nil)`

Create a new `EmailSendingSubdomainPreview` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainReputationComplaint(data = nil)`

Create a new `EmailSendingSubdomainReputationComplaint` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainResponseSingle(data = nil)`

Create a new `EmailSendingSubdomainResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainsResponseCollection(data = nil)`

Create a new `EmailSendingSubdomainsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSuppression(data = nil)`

Create a new `EmailSendingSuppression` entity instance. Pass `nil` for no initial data.

#### `Subdomain(data = nil)`

Create a new `Subdomain` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## EmailAccountRulesPlanResponseSingleEntity

```ruby
email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `Array` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `String` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `Array` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailAccountRulesPlanResponseSingle.create({
  "account_id" => "example_account_id", # String
  "owner_worker_tag" => "example_owner_worker_tag", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailAccountRulesResponseCollectionEntity

```ruby
email_account_rules_response_collection = client.EmailAccountRulesResponseCollection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Array` | No |  |
| `result_info` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailAccountRulesResponseCollection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailAuthEntity

```ruby
email_auth = client.EmailAuth
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `Boolean` | No | Enable or disable DMARC reports for this zone |
| `result` | `Hash` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `Boolean` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailAuth.load({ "id" => "email_auth_id", "zone_id" => "zone_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailAuth.update({
  "zone_id" => "zone_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailAuthEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailCatchAllRuleResponseSingleEntity

```ruby
email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `Array` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `Boolean` | No | Routing rule status. |
| `matchers` | `Array` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `String` | No | Routing rule name. |
| `owner_worker_tag` | `String` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `Hash` | No |  |
| `source` | `String` | No | Who manages the rule. |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailCatchAllRuleResponseSingle.load({ "zone_id" => "zone_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailCatchAllRuleResponseSingle.update({
  "zone_id" => "zone_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailDestinationAddressResponseSingleEntity

```ruby
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `String` | Yes | The contact email address of the user. |
| `result` | `Hash` | No |  |
| `status` | `String` | Yes | Destination address status. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailDestinationAddressResponseSingle.create({
  "account_id" => "example_account_id", # String
  "email" => "example_email", # String
  "status" => "example_status", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailDestinationAddressResponseSingle.load({ "account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailDestinationAddressResponseSingle.remove({ "account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailDestinationAddressResponseSingle.update({
  "account_id" => "account_id",
  "destination_address_identifier" => "destination_address_identifier",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailDestinationAddressesResponseCollectionEntity

```ruby
email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Array` | No |  |
| `result_info` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailDestinationAddressesResponseCollection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailDnsSettingsResponseCollectionEntity

```ruby
email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Array` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailDnsSettingsResponseCollection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailEmailSettingsResponseSingleEntity

```ruby
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `Boolean` | No | State of your zone Email Routing settings. |
| `name` | `String` | No | Domain of your zone. |
| `result` | `Hash` | No |  |
| `skip_wizard` | `Boolean` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `Boolean` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailEmailSettingsResponseSingle.create({
  "zone_id" => "example_zone_id", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailEmailSettingsResponseSingle.load({ "zone_id" => "zone_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailEmailSettingsResponseSingle.remove({ "zone_id" => "zone_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailEmailSettingsResponseSingle.update({
  "zone_id" => "zone_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailRoutingSuppressionEntity

```ruby
email_routing_suppression = client.EmailRoutingSuppression
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `String` | Yes |  |
| `email` | `String` | Yes |  |
| `expires_at` | `String` | Yes |  |
| `id` | `String` | Yes |  |
| `reason` | `String` | Yes |  |
| `zones` | `Array` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailRoutingSuppression.create({
  "account_id" => "example_account_id", # String
  "created_at" => "example_created_at", # String
  "email" => "example_email", # String
  "expires_at" => "example_expires_at", # String
  "id" => "example_id", # String
  "reason" => "example_reason", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailRoutingSuppression.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailRoutingSuppression.load({ "id" => "email_routing_suppression_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailRoutingSuppression.remove({ "id" => "email_routing_suppression_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailRoutingSuppressionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailRuleResponseSingleEntity

```ruby
email_rule_response_single = client.EmailRuleResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `Array` | Yes | List actions patterns. |
| `enabled` | `Boolean` | No | Routing rule status. |
| `matchers` | `Array` | Yes | Matching patterns to forward to your actions. |
| `name` | `String` | No | Routing rule name. |
| `owner_worker_tag` | `String` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `Float` | No | Priority of the routing rule. |
| `result` | `Hash` | No |  |
| `source` | `String` | No | Who manages the rule. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailRuleResponseSingle.create({
  "zone_id" => "example_zone_id", # String
  "actions" => [], # Array
  "matchers" => [], # Array
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailRuleResponseSingle.load({ "rule_identifier" => "rule_identifier", "zone_id" => "zone_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailRuleResponseSingle.remove({ "rule_identifier" => "rule_identifier", "zone_id" => "zone_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailRuleResponseSingle.update({
  "rule_identifier" => "rule_identifier",
  "zone_id" => "zone_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailRulesResponseCollectionEntity

```ruby
email_rules_response_collection = client.EmailRulesResponseCollection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Array` | No |  |
| `result_info` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailRulesResponseCollection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSecurityEntity

```ruby
email_security = client.EmailSecurity
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `String` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `String` | No | Optional note describing the job. |
| `destination` | `Object` | No | Required when action is 'MOVE'. |
| `errors` | `Array` | Yes |  |
| `expected_disposition` | `String` | No | Nonfunctional field. |
| `id` | `String` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `Array` | Yes | List of message IDs to move. |
| `messages` | `Array` | Yes |  |
| `postfix_ids` | `Array` | No | Deprecated, use `ids` instead. |
| `result` | `Hash` | Yes |  |
| `result_info` | `Hash` | Yes |  |
| `search_params` | `Hash` | Yes |  |
| `success` | `Boolean` | Yes | Whether the API call was successful. |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSecurity.create({
  "account_id" => "example_account_id", # String
  "action" => "example_action", # String
  "errors" => [], # Array
  "id" => "example_id", # String
  "ids" => [], # Array
  "messages" => [], # Array
  "result" => {}, # Hash
  "result_info" => {}, # Hash
  "search_params" => {}, # Hash
  "success" => true, # Boolean
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailSecurity.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSecurity.load({ "account_id" => "account_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSecurity.remove({ "account_id" => "account_id", "job_id" => "job_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSecurityEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSecurityAnalyticsEntity

```ruby
email_security_analytics = client.EmailSecurityAnalytics
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSecurityAnalytics.load({ "account_id" => "account_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSecurityAnalyticsEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSecuritySettingEntity

```ruby
email_security_setting = client.EmailSecuritySetting
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `Array` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `String` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `Array` | Yes | IDs of the allow policies to delete. |
| `domain` | `String` | Yes | The email domain to protect. |
| `drop_dispositions` | `Array` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `Object` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `String` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `Array` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `Integer` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `Array` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `String` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `Array` | Yes | Allow policies to create. |
| `puts` | `Array` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `Array` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `Boolean` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `Boolean` | No | Require TLS on outbound connections. |
| `result` | `Hash` | No | An impersonation registry entry. |
| `transport` | `String` | No | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSecuritySetting.create({
  "account_id" => "example_account_id", # String
  "allowed_delivery_modes" => [], # Array
  "deletes" => [], # Array
  "domain" => "example_domain", # String
  "drop_dispositions" => [], # Array
  "ip_restrictions" => [], # Array
  "patches" => [], # Array
  "pattern" => "example_pattern", # String
  "posts" => [], # Array
  "puts" => [], # Array
  "regions" => [], # Array
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailSecuritySetting.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSecuritySetting.load({ "account_id" => "account_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSecuritySetting.remove({ "account_id" => "account_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailSecuritySetting.update({
  "account_id" => "account_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSecuritySettingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingEntity

```ruby
email_sending = client.EmailSending
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `Array` | No | File attachments and inline images. |
| `bcc` | `Object` | No | Blind carbon copy recipient(s). |
| `cc` | `Object` | No | Carbon copy recipient(s). |
| `errors` | `Array` | Yes |  |
| `from` | `Object` | Yes | Sender email address. |
| `headers` | `Hash` | No | Custom email headers as key-value pairs. |
| `html` | `String` | No | HTML body of the email. |
| `messages` | `Array` | Yes |  |
| `mime_message` | `String` | Yes | The full MIME-encoded email message. |
| `recipients` | `Array` | Yes | List of recipient email addresses. |
| `reply_to` | `Object` | No | Reply-to address. |
| `result` | `Hash` | Yes |  |
| `result_info` | `Hash` | Yes |  |
| `subject` | `String` | Yes | Email subject line. |
| `success` | `Boolean` | Yes |  |
| `text` | `String` | No | Plain text body of the email. |
| `to` | `Object` | No | Recipient(s). |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSending.create({
  "account_id" => "example_account_id", # String
  "errors" => [], # Array
  "from" => "example_from", # Object
  "messages" => [], # Array
  "mime_message" => "example_mime_message", # String
  "recipients" => [], # Array
  "result" => {}, # Hash
  "result_info" => {}, # Hash
  "subject" => "example_subject", # String
  "success" => true, # Boolean
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSending.load({ "account_id" => "account_id", "message_id" => "message_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingLimitsResponseSingleEntity

```ruby
email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingLimitsResponseSingle.load({ "account_id" => "account_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingReputationResponseSingleEntity

```ruby
email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Hash` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingReputationResponseSingle.load({ "account_id" => "account_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSubdomainDnsStatusEntity

```ruby
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSendingSubdomainDnsStatus.create({
  "subdomain_id" => "example_subdomain_id", # String
  "zone_id" => "example_zone_id", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingSubdomainDnsStatus.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSubdomainPreviewEntity

```ruby
email_sending_subdomain_preview = client.EmailSendingSubdomainPreview
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `String` | Yes | The domain name within the zone. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSendingSubdomainPreview.create({
  "zone_id" => "example_zone_id", # String
  "name" => "example_name", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSubdomainReputationComplaintEntity

```ruby
email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Hash` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingSubdomainReputationComplaint.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id", "end_at" => "end_at", "start_at" => "start_at" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSubdomainResponseSingleEntity

```ruby
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `Boolean` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `String` | Yes | The domain name within the zone. |
| `preview_enabled` | `Boolean` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `Object` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSendingSubdomainResponseSingle.create({
  "zone_id" => "example_zone_id", # String
  "name" => "example_name", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingSubdomainResponseSingle.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailSendingSubdomainResponseSingle.update({
  "subdomain_id" => "subdomain_id",
  "zone_id" => "zone_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSubdomainsResponseCollectionEntity

```ruby
email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Array` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailSendingSubdomainsResponseCollection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EmailSendingSuppressionEntity

```ruby
email_sending_suppression = client.EmailSendingSuppression
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `String` | Yes | When the suppression was created. |
| `email` | `String` | Yes | The email address to suppress. |
| `errors` | `Array` | Yes |  |
| `expires_at` | `String` | Yes | Expiration timestamp for the suppression. |
| `id` | `String` | Yes | Unique identifier for this suppression. |
| `items` | `Array` | Yes | Suppressions to import. |
| `messages` | `Array` | Yes |  |
| `note` | `String` | No | Advisory note for this suppression. |
| `read_only` | `Boolean` | Yes | Whether clients may mutate this suppression. |
| `reason` | `String` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `Hash` | Yes |  |
| `scope` | `Object` | No | Where the suppression applies. |
| `success` | `Boolean` | Yes |  |
| `zones` | `Array` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.EmailSendingSuppression.create({
  "account_id" => "example_account_id", # String
  "created_at" => "example_created_at", # String
  "email" => "example_email", # String
  "errors" => [], # Array
  "expires_at" => "example_expires_at", # String
  "id" => "example_id", # String
  "items" => [], # Array
  "messages" => [], # Array
  "read_only" => true, # Boolean
  "reason" => "example_reason", # String
  "result" => {}, # Hash
  "success" => true, # Boolean
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EmailSendingSuppression.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingSuppression.load({ "id" => "email_sending_suppression_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.EmailSendingSuppression.remove({ "id" => "email_sending_suppression_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.EmailSendingSuppression.update({
  "id" => "email_sending_suppression_id",
  "account_id" => "account_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EmailSendingSuppressionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SubdomainEntity

```ruby
subdomain = client.Subdomain
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `String` | No |  |

### Operations

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Subdomain.remove({ "id" => "id", "zone_id" => "zone_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SubdomainEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```ruby
client = CloudflareEmailSecuritySDK.new({
  "feature" => {
    "test" => { "active" => true },
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

