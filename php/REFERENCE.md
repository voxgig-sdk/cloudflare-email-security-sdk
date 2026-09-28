# CloudflareEmailSecurity PHP SDK Reference

Complete API reference for the CloudflareEmailSecurity PHP SDK.


## CloudflareEmailSecuritySDK

### Constructor

```php
require_once __DIR__ . '/cloudflareemailsecurity_sdk.php';

$client = new CloudflareEmailSecuritySDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflareEmailSecuritySDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = CloudflareEmailSecuritySDK::test();
```


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle($data = null)`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailAccountRulesResponseCollection($data = null)`

Create a new `EmailAccountRulesResponseCollectionEntity` instance. Pass `null` for no initial data.

#### `EmailAuth($data = null)`

Create a new `EmailAuthEntity` instance. Pass `null` for no initial data.

#### `EmailCatchAllRuleResponseSingle($data = null)`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailDestinationAddressResponseSingle($data = null)`

Create a new `EmailDestinationAddressResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailDestinationAddressesResponseCollection($data = null)`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance. Pass `null` for no initial data.

#### `EmailDnsSettingsResponseCollection($data = null)`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance. Pass `null` for no initial data.

#### `EmailEmailSettingsResponseSingle($data = null)`

Create a new `EmailEmailSettingsResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailRoutingSuppression($data = null)`

Create a new `EmailRoutingSuppressionEntity` instance. Pass `null` for no initial data.

#### `EmailRuleResponseSingle($data = null)`

Create a new `EmailRuleResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailRulesResponseCollection($data = null)`

Create a new `EmailRulesResponseCollectionEntity` instance. Pass `null` for no initial data.

#### `EmailSecurity($data = null)`

Create a new `EmailSecurityEntity` instance. Pass `null` for no initial data.

#### `EmailSecurityAnalytics($data = null)`

Create a new `EmailSecurityAnalyticsEntity` instance. Pass `null` for no initial data.

#### `EmailSecuritySetting($data = null)`

Create a new `EmailSecuritySettingEntity` instance. Pass `null` for no initial data.

#### `EmailSending($data = null)`

Create a new `EmailSendingEntity` instance. Pass `null` for no initial data.

#### `EmailSendingLimitsResponseSingle($data = null)`

Create a new `EmailSendingLimitsResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailSendingReputationResponseSingle($data = null)`

Create a new `EmailSendingReputationResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSubdomainDnsStatus($data = null)`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSubdomainPreview($data = null)`

Create a new `EmailSendingSubdomainPreviewEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSubdomainReputationComplaint($data = null)`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSubdomainResponseSingle($data = null)`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSubdomainsResponseCollection($data = null)`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance. Pass `null` for no initial data.

#### `EmailSendingSuppression($data = null)`

Create a new `EmailSendingSuppressionEntity` instance. Pass `null` for no initial data.

#### `Subdomain($data = null)`

Create a new `SubdomainEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): CloudflareEmailSecurityUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## EmailAccountRulesPlanResponseSingleEntity

```php
$email_account_rules_plan_response_single = $client->EmailAccountRulesPlanResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `array` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `array` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailAccountRulesPlanResponseSingle()->create([
  "account_id" => null, // string
  "owner_worker_tag" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailAccountRulesPlanResponseSingleEntity`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailAccountRulesResponseCollectionEntity

```php
$email_account_rules_response_collection = $client->EmailAccountRulesResponseCollection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |
| `result_info` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailAccountRulesResponseCollection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailAccountRulesResponseCollectionEntity`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailAuthEntity

```php
$email_auth = $client->EmailAuth();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | Enable or disable DMARC reports for this zone |
| `result` | `array` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailAuth()->load(["id" => "email_auth_id", "zone_id" => "zone_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailAuth()->update([
  "zone_id" => "zone_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailAuthEntity`

Create a new `EmailAuthEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailCatchAllRuleResponseSingleEntity

```php
$email_catch_all_rule_response_single = $client->EmailCatchAllRuleResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `array` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `array` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `array` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailCatchAllRuleResponseSingle()->load(["zone_id" => "zone_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailCatchAllRuleResponseSingle()->update([
  "zone_id" => "zone_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailCatchAllRuleResponseSingleEntity`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailDestinationAddressResponseSingleEntity

```php
$email_destination_address_response_single = $client->EmailDestinationAddressResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes | The contact email address of the user. |
| `result` | `array` | No |  |
| `status` | `string` | Yes | Destination address status. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailDestinationAddressResponseSingle()->create([
  "account_id" => null, // string
  "email" => null, // string
  "status" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailDestinationAddressResponseSingle()->load(["account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailDestinationAddressResponseSingle()->remove(["account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailDestinationAddressResponseSingle()->update([
  "account_id" => "account_id",
  "destination_address_identifier" => "destination_address_identifier",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailDestinationAddressResponseSingleEntity`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailDestinationAddressesResponseCollectionEntity

```php
$email_destination_addresses_response_collection = $client->EmailDestinationAddressesResponseCollection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |
| `result_info` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailDestinationAddressesResponseCollection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailDestinationAddressesResponseCollectionEntity`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailDnsSettingsResponseCollectionEntity

```php
$email_dns_settings_response_collection = $client->EmailDnsSettingsResponseCollection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailDnsSettingsResponseCollection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailDnsSettingsResponseCollectionEntity`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailEmailSettingsResponseSingleEntity

```php
$email_email_settings_response_single = $client->EmailEmailSettingsResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | State of your zone Email Routing settings. |
| `name` | `string` | No | Domain of your zone. |
| `result` | `array` | No |  |
| `skip_wizard` | `bool` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailEmailSettingsResponseSingle()->create([
  "zone_id" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailEmailSettingsResponseSingle()->load(["zone_id" => "zone_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailEmailSettingsResponseSingle()->remove(["zone_id" => "zone_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailEmailSettingsResponseSingle()->update([
  "zone_id" => "zone_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailEmailSettingsResponseSingleEntity`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailRoutingSuppressionEntity

```php
$email_routing_suppression = $client->EmailRoutingSuppression();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `expires_at` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `reason` | `string` | Yes |  |
| `zones` | `array` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailRoutingSuppression()->create([
  "account_id" => null, // string
  "created_at" => null, // string
  "email" => null, // string
  "expires_at" => null, // string
  "id" => null, // string
  "reason" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailRoutingSuppression()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailRoutingSuppression()->load(["id" => "email_routing_suppression_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailRoutingSuppression()->remove(["id" => "email_routing_suppression_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailRoutingSuppressionEntity`

Create a new `EmailRoutingSuppressionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailRuleResponseSingleEntity

```php
$email_rule_response_single = $client->EmailRuleResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `array` | Yes | List actions patterns. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `array` | Yes | Matching patterns to forward to your actions. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float` | No | Priority of the routing rule. |
| `result` | `array` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailRuleResponseSingle()->create([
  "zone_id" => null, // string
  "actions" => null, // array
  "matchers" => null, // array
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailRuleResponseSingle()->load(["rule_identifier" => "rule_identifier", "zone_id" => "zone_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailRuleResponseSingle()->remove(["rule_identifier" => "rule_identifier", "zone_id" => "zone_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailRuleResponseSingle()->update([
  "rule_identifier" => "rule_identifier",
  "zone_id" => "zone_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailRuleResponseSingleEntity`

Create a new `EmailRuleResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailRulesResponseCollectionEntity

```php
$email_rules_response_collection = $client->EmailRulesResponseCollection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |
| `result_info` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailRulesResponseCollection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailRulesResponseCollectionEntity`

Create a new `EmailRulesResponseCollectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSecurityEntity

```php
$email_security = $client->EmailSecurity();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `string` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | No | Optional note describing the job. |
| `destination` | `mixed` | No | Required when action is 'MOVE'. |
| `errors` | `array` | Yes |  |
| `expected_disposition` | `string` | No | Nonfunctional field. |
| `id` | `string` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `array` | Yes | List of message IDs to move. |
| `messages` | `array` | Yes |  |
| `postfix_ids` | `array` | No | Deprecated, use `ids` instead. |
| `result` | `array` | Yes |  |
| `result_info` | `array` | Yes |  |
| `search_params` | `array` | Yes |  |
| `success` | `bool` | Yes | Whether the API call was successful. |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSecurity()->create([
  "account_id" => null, // string
  "action" => null, // string
  "errors" => null, // array
  "id" => null, // string
  "ids" => null, // array
  "messages" => null, // array
  "result" => null, // array
  "result_info" => null, // array
  "search_params" => null, // array
  "success" => null, // bool
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailSecurity()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSecurity()->load(["account_id" => "account_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSecurity()->remove(["account_id" => "account_id", "job_id" => "job_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSecurityEntity`

Create a new `EmailSecurityEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSecurityAnalyticsEntity

```php
$email_security_analytics = $client->EmailSecurityAnalytics();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSecurityAnalytics()->load(["account_id" => "account_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSecurityAnalyticsEntity`

Create a new `EmailSecurityAnalyticsEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSecuritySettingEntity

```php
$email_security_setting = $client->EmailSecuritySetting();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `array` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `string` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `array` | Yes | IDs of the allow policies to delete. |
| `domain` | `string` | Yes | The email domain to protect. |
| `drop_dispositions` | `array` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `mixed` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `array` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `array` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `array` | Yes | Allow policies to create. |
| `puts` | `array` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `array` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | No | Require TLS on outbound connections. |
| `result` | `array` | No | An impersonation registry entry. |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSecuritySetting()->create([
  "account_id" => null, // string
  "allowed_delivery_modes" => null, // array
  "deletes" => null, // array
  "domain" => null, // string
  "drop_dispositions" => null, // array
  "ip_restrictions" => null, // array
  "patches" => null, // array
  "pattern" => null, // string
  "posts" => null, // array
  "puts" => null, // array
  "regions" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailSecuritySetting()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSecuritySetting()->load(["account_id" => "account_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSecuritySetting()->remove(["account_id" => "account_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailSecuritySetting()->update([
  "account_id" => "account_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSecuritySettingEntity`

Create a new `EmailSecuritySettingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingEntity

```php
$email_sending = $client->EmailSending();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `array` | No | File attachments and inline images. |
| `bcc` | `mixed` | No | Blind carbon copy recipient(s). |
| `cc` | `mixed` | No | Carbon copy recipient(s). |
| `errors` | `array` | Yes |  |
| `from` | `mixed` | Yes | Sender email address. |
| `headers` | `array` | No | Custom email headers as key-value pairs. |
| `html` | `string` | No | HTML body of the email. |
| `messages` | `array` | Yes |  |
| `mime_message` | `string` | Yes | The full MIME-encoded email message. |
| `recipients` | `array` | Yes | List of recipient email addresses. |
| `reply_to` | `mixed` | No | Reply-to address. |
| `result` | `array` | Yes |  |
| `result_info` | `array` | Yes |  |
| `subject` | `string` | Yes | Email subject line. |
| `success` | `bool` | Yes |  |
| `text` | `string` | No | Plain text body of the email. |
| `to` | `mixed` | No | Recipient(s). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSending()->create([
  "account_id" => null, // string
  "errors" => null, // array
  "from" => null, // mixed
  "messages" => null, // array
  "mime_message" => null, // string
  "recipients" => null, // array
  "result" => null, // array
  "result_info" => null, // array
  "subject" => null, // string
  "success" => null, // bool
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSending()->load(["account_id" => "account_id", "message_id" => "message_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingEntity`

Create a new `EmailSendingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingLimitsResponseSingleEntity

```php
$email_sending_limits_response_single = $client->EmailSendingLimitsResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingLimitsResponseSingle()->load(["account_id" => "account_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingLimitsResponseSingleEntity`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingReputationResponseSingleEntity

```php
$email_sending_reputation_response_single = $client->EmailSendingReputationResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingReputationResponseSingle()->load(["account_id" => "account_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingReputationResponseSingleEntity`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSubdomainDnsStatusEntity

```php
$email_sending_subdomain_dns_status = $client->EmailSendingSubdomainDnsStatus();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSendingSubdomainDnsStatus()->create([
  "subdomain_id" => null, // string
  "zone_id" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingSubdomainDnsStatus()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSubdomainDnsStatusEntity`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSubdomainPreviewEntity

```php
$email_sending_subdomain_preview = $client->EmailSendingSubdomainPreview();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The domain name within the zone. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSendingSubdomainPreview()->create([
  "zone_id" => null, // string
  "name" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSubdomainPreviewEntity`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSubdomainReputationComplaintEntity

```php
$email_sending_subdomain_reputation_complaint = $client->EmailSendingSubdomainReputationComplaint();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingSubdomainReputationComplaint()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id", "end_at" => "end_at", "start_at" => "start_at"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSubdomainReputationComplaintEntity`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSubdomainResponseSingleEntity

```php
$email_sending_subdomain_response_single = $client->EmailSendingSubdomainResponseSingle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `bool` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | Yes | The domain name within the zone. |
| `preview_enabled` | `bool` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `mixed` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSendingSubdomainResponseSingle()->create([
  "zone_id" => null, // string
  "name" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingSubdomainResponseSingle()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailSendingSubdomainResponseSingle()->update([
  "subdomain_id" => "subdomain_id",
  "zone_id" => "zone_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSubdomainResponseSingleEntity`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSubdomainsResponseCollectionEntity

```php
$email_sending_subdomains_response_collection = $client->EmailSendingSubdomainsResponseCollection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailSendingSubdomainsResponseCollection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSubdomainsResponseCollectionEntity`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EmailSendingSuppressionEntity

```php
$email_sending_suppression = $client->EmailSendingSuppression();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | When the suppression was created. |
| `email` | `string` | Yes | The email address to suppress. |
| `errors` | `array` | Yes |  |
| `expires_at` | `string` | Yes | Expiration timestamp for the suppression. |
| `id` | `string` | Yes | Unique identifier for this suppression. |
| `items` | `array` | Yes | Suppressions to import. |
| `messages` | `array` | Yes |  |
| `note` | `string` | No | Advisory note for this suppression. |
| `read_only` | `bool` | Yes | Whether clients may mutate this suppression. |
| `reason` | `string` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `array` | Yes |  |
| `scope` | `mixed` | No | Where the suppression applies. |
| `success` | `bool` | Yes |  |
| `zones` | `array` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EmailSendingSuppression()->create([
  "account_id" => null, // string
  "created_at" => null, // string
  "email" => null, // string
  "errors" => null, // array
  "expires_at" => null, // string
  "id" => null, // string
  "items" => null, // array
  "messages" => null, // array
  "read_only" => null, // bool
  "reason" => null, // string
  "result" => null, // array
  "success" => null, // bool
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EmailSendingSuppression()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingSuppression()->load(["id" => "email_sending_suppression_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EmailSendingSuppression()->remove(["id" => "email_sending_suppression_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EmailSendingSuppression()->update([
  "id" => "email_sending_suppression_id",
  "account_id" => "account_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EmailSendingSuppressionEntity`

Create a new `EmailSendingSuppressionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SubdomainEntity

```php
$subdomain = $client->Subdomain();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Subdomain()->remove(["id" => "id", "zone_id" => "zone_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SubdomainEntity`

Create a new `SubdomainEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```php
$client = new CloudflareEmailSecuritySDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
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

