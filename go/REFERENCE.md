# CloudflareEmailSecurity Golang SDK Reference

Complete API reference for the CloudflareEmailSecurity Golang SDK.


## CloudflareEmailSecuritySDK

### Constructor

```go
func NewCloudflareEmailSecuritySDK(options map[string]any) *CloudflareEmailSecuritySDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *CloudflareEmailSecuritySDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *CloudflareEmailSecuritySDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailAccountRulesPlanResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailAccountRulesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailAccountRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailAuth(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailAuth` entity instance. Pass `nil` for no initial data.

#### `EmailCatchAllRuleResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailCatchAllRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailDestinationAddressResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailDestinationAddressesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailDestinationAddressesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailDnsSettingsResponseCollection(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailDnsSettingsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailEmailSettingsResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailEmailSettingsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRoutingSuppression(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailRoutingSuppression` entity instance. Pass `nil` for no initial data.

#### `EmailRuleResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailRuleResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailRulesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailRulesResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSecurity(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSecurity` entity instance. Pass `nil` for no initial data.

#### `EmailSecurityAnalytics(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSecurityAnalytics` entity instance. Pass `nil` for no initial data.

#### `EmailSecuritySetting(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSecuritySetting` entity instance. Pass `nil` for no initial data.

#### `EmailSending(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSending` entity instance. Pass `nil` for no initial data.

#### `EmailSendingLimitsResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingLimitsResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingReputationResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingReputationResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainDnsStatus(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSubdomainDnsStatus` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainPreview(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSubdomainPreview` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainReputationComplaint(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSubdomainReputationComplaint` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainResponseSingle(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSubdomainResponseSingle` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSubdomainsResponseCollection(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSubdomainsResponseCollection` entity instance. Pass `nil` for no initial data.

#### `EmailSendingSuppression(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `EmailSendingSuppression` entity instance. Pass `nil` for no initial data.

#### `Subdomain(data map[string]any) CloudflareEmailSecurityEntity`

Create a new `Subdomain` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## EmailAccountRulesPlanResponseSingleEntity

```go
emailAccountRulesPlanResponseSingle := client.EmailAccountRulesPlanResponseSingle(nil)
fmt.Println(emailAccountRulesPlanResponseSingle.GetName()) // "email_account_rules_plan_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `[]any` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `[]any` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailAccountRulesPlanResponseSingle(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "owner_worker_tag": "example_owner_worker_tag",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailAccountRulesResponseCollectionEntity

```go
emailAccountRulesResponseCollection := client.EmailAccountRulesResponseCollection(nil)
fmt.Println(emailAccountRulesResponseCollection.GetName()) // "email_account_rules_response_collection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `[]any` | No |  |
| `result_info` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailAccountRulesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailAuthEntity

```go
emailAuth := client.EmailAuth(nil)
fmt.Println(emailAuth.GetName()) // "email_auth"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | Enable or disable DMARC reports for this zone |
| `result` | `map[string]any` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailAuth(nil).Load(map[string]any{"id": "email_auth_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailAuth(nil).Update(map[string]any{
    "zone_id": "zone_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailAuthEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailCatchAllRuleResponseSingleEntity

```go
emailCatchAllRuleResponseSingle := client.EmailCatchAllRuleResponseSingle(nil)
fmt.Println(emailCatchAllRuleResponseSingle.GetName()) // "email_catch_all_rule_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `[]any` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `[]any` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `map[string]any` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailCatchAllRuleResponseSingle(nil).Load(map[string]any{"zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailCatchAllRuleResponseSingle(nil).Update(map[string]any{
    "zone_id": "zone_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailDestinationAddressResponseSingleEntity

```go
emailDestinationAddressResponseSingle := client.EmailDestinationAddressResponseSingle(nil)
fmt.Println(emailDestinationAddressResponseSingle.GetName()) // "email_destination_address_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes | The contact email address of the user. |
| `result` | `map[string]any` | No |  |
| `status` | `string` | Yes | Destination address status. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailDestinationAddressResponseSingle(nil).Load(map[string]any{"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailDestinationAddressResponseSingle(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "email": "example_email",
    "status": "example_status",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailDestinationAddressResponseSingle(nil).Update(map[string]any{
    "account_id": "account_id",
    "destination_address_identifier": "destination_address_identifier",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailDestinationAddressResponseSingle(nil).Remove(map[string]any{"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailDestinationAddressesResponseCollectionEntity

```go
emailDestinationAddressesResponseCollection := client.EmailDestinationAddressesResponseCollection(nil)
fmt.Println(emailDestinationAddressesResponseCollection.GetName()) // "email_destination_addresses_response_collection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `[]any` | No |  |
| `result_info` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailDestinationAddressesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailDnsSettingsResponseCollectionEntity

```go
emailDnsSettingsResponseCollection := client.EmailDnsSettingsResponseCollection(nil)
fmt.Println(emailDnsSettingsResponseCollection.GetName()) // "email_dns_settings_response_collection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailDnsSettingsResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailEmailSettingsResponseSingleEntity

```go
emailEmailSettingsResponseSingle := client.EmailEmailSettingsResponseSingle(nil)
fmt.Println(emailEmailSettingsResponseSingle.GetName()) // "email_email_settings_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | State of your zone Email Routing settings. |
| `name` | `string` | No | Domain of your zone. |
| `result` | `map[string]any` | No |  |
| `skip_wizard` | `bool` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailEmailSettingsResponseSingle(nil).Load(map[string]any{"zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailEmailSettingsResponseSingle(nil).Create(map[string]any{
    "zone_id": "example_zone_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailEmailSettingsResponseSingle(nil).Update(map[string]any{
    "zone_id": "zone_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailEmailSettingsResponseSingle(nil).Remove(map[string]any{"zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailRoutingSuppressionEntity

```go
emailRoutingSuppression := client.EmailRoutingSuppression(nil)
fmt.Println(emailRoutingSuppression.GetName()) // "email_routing_suppression"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `expires_at` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `reason` | `string` | Yes |  |
| `zones` | `[]any` | No |  |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailRoutingSuppression(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailRoutingSuppression(nil).Load(map[string]any{"id": "email_routing_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailRoutingSuppression(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "created_at": "example_created_at",
    "email": "example_email",
    "expires_at": "example_expires_at",
    "id": "example_id",
    "reason": "example_reason",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailRoutingSuppression(nil).Remove(map[string]any{"id": "email_routing_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailRoutingSuppressionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailRuleResponseSingleEntity

```go
emailRuleResponseSingle := client.EmailRuleResponseSingle(nil)
fmt.Println(emailRuleResponseSingle.GetName()) // "email_rule_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `[]any` | Yes | List actions patterns. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `[]any` | Yes | Matching patterns to forward to your actions. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float64` | No | Priority of the routing rule. |
| `result` | `map[string]any` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailRuleResponseSingle(nil).Load(map[string]any{"rule_identifier": "rule_identifier", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailRuleResponseSingle(nil).Create(map[string]any{
    "zone_id": "example_zone_id",
    "actions": []any{},
    "matchers": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailRuleResponseSingle(nil).Update(map[string]any{
    "rule_identifier": "rule_identifier",
    "zone_id": "zone_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailRuleResponseSingle(nil).Remove(map[string]any{"rule_identifier": "rule_identifier", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailRuleResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailRulesResponseCollectionEntity

```go
emailRulesResponseCollection := client.EmailRulesResponseCollection(nil)
fmt.Println(emailRulesResponseCollection.GetName()) // "email_rules_response_collection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `[]any` | No |  |
| `result_info` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailRulesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailRulesResponseCollectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSecurityEntity

```go
emailSecurity := client.EmailSecurity(nil)
fmt.Println(emailSecurity.GetName()) // "email_security"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `string` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | No | Optional note describing the job. |
| `destination` | `any` | No | Required when action is 'MOVE'. |
| `errors` | `[]any` | Yes |  |
| `expected_disposition` | `string` | No | Nonfunctional field. |
| `id` | `string` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `[]any` | Yes | List of message IDs to move. |
| `messages` | `[]any` | Yes |  |
| `postfix_ids` | `[]any` | No | Deprecated, use `ids` instead. |
| `result` | `map[string]any` | Yes |  |
| `result_info` | `map[string]any` | Yes |  |
| `search_params` | `map[string]any` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailSecurity(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSecurity(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSecurity(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "action": "example_action",
    "errors": []any{},
    "id": "example_id",
    "ids": []any{},
    "messages": []any{},
    "result": map[string]any{},
    "result_info": map[string]any{},
    "search_params": map[string]any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailSecurity(nil).Remove(map[string]any{"account_id": "account_id", "job_id": "job_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSecurityEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSecurityAnalyticsEntity

```go
emailSecurityAnalytics := client.EmailSecurityAnalytics(nil)
fmt.Println(emailSecurityAnalytics.GetName()) // "email_security_analytics"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSecurityAnalytics(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSecurityAnalyticsEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSecuritySettingEntity

```go
emailSecuritySetting := client.EmailSecuritySetting(nil)
fmt.Println(emailSecuritySetting.GetName()) // "email_security_setting"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `[]any` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `string` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `[]any` | Yes | IDs of the allow policies to delete. |
| `domain` | `string` | Yes | The email domain to protect. |
| `drop_dispositions` | `[]any` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `[]any` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `[]any` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `[]any` | Yes | Allow policies to create. |
| `puts` | `[]any` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `[]any` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | No | Require TLS on outbound connections. |
| `result` | `map[string]any` | No | An impersonation registry entry. |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailSecuritySetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSecuritySetting(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSecuritySetting(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "allowed_delivery_modes": []any{},
    "deletes": []any{},
    "domain": "example_domain",
    "drop_dispositions": []any{},
    "ip_restrictions": []any{},
    "patches": []any{},
    "pattern": "example_pattern",
    "posts": []any{},
    "puts": []any{},
    "regions": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailSecuritySetting(nil).Update(map[string]any{
    "account_id": "account_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailSecuritySetting(nil).Remove(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSecuritySettingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingEntity

```go
emailSending := client.EmailSending(nil)
fmt.Println(emailSending.GetName()) // "email_sending"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | No | File attachments and inline images. |
| `bcc` | `any` | No | Blind carbon copy recipient(s). |
| `cc` | `any` | No | Carbon copy recipient(s). |
| `errors` | `[]any` | Yes |  |
| `from` | `any` | Yes | Sender email address. |
| `headers` | `map[string]any` | No | Custom email headers as key-value pairs. |
| `html` | `string` | No | HTML body of the email. |
| `messages` | `[]any` | Yes |  |
| `mime_message` | `string` | Yes | The full MIME-encoded email message. |
| `recipients` | `[]any` | Yes | List of recipient email addresses. |
| `reply_to` | `any` | No | Reply-to address. |
| `result` | `map[string]any` | Yes |  |
| `result_info` | `map[string]any` | Yes |  |
| `subject` | `string` | Yes | Email subject line. |
| `success` | `bool` | Yes |  |
| `text` | `string` | No | Plain text body of the email. |
| `to` | `any` | No | Recipient(s). |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSending(nil).Load(map[string]any{"account_id": "account_id", "message_id": "message_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSending(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "errors": []any{},
    "from": "example_from",
    "messages": []any{},
    "mime_message": "example_mime_message",
    "recipients": []any{},
    "result": map[string]any{},
    "result_info": map[string]any{},
    "subject": "example_subject",
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingLimitsResponseSingleEntity

```go
emailSendingLimitsResponseSingle := client.EmailSendingLimitsResponseSingle(nil)
fmt.Println(emailSendingLimitsResponseSingle.GetName()) // "email_sending_limits_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingLimitsResponseSingle(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingReputationResponseSingleEntity

```go
emailSendingReputationResponseSingle := client.EmailSendingReputationResponseSingle(nil)
fmt.Println(emailSendingReputationResponseSingle.GetName()) // "email_sending_reputation_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `map[string]any` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingReputationResponseSingle(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSubdomainDnsStatusEntity

```go
emailSendingSubdomainDnsStatus := client.EmailSendingSubdomainDnsStatus(nil)
fmt.Println(emailSendingSubdomainDnsStatus.GetName()) // "email_sending_subdomain_dns_status"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingSubdomainDnsStatus(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSendingSubdomainDnsStatus(nil).Create(map[string]any{
    "subdomain_id": "example_subdomain_id",
    "zone_id": "example_zone_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSubdomainPreviewEntity

```go
emailSendingSubdomainPreview := client.EmailSendingSubdomainPreview(nil)
fmt.Println(emailSendingSubdomainPreview.GetName()) // "email_sending_subdomain_preview"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The domain name within the zone. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSendingSubdomainPreview(nil).Create(map[string]any{
    "zone_id": "example_zone_id",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSubdomainReputationComplaintEntity

```go
emailSendingSubdomainReputationComplaint := client.EmailSendingSubdomainReputationComplaint(nil)
fmt.Println(emailSendingSubdomainReputationComplaint.GetName()) // "email_sending_subdomain_reputation_complaint"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `map[string]any` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingSubdomainReputationComplaint(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id", "end_at": "end_at", "start_at": "start_at"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSubdomainResponseSingleEntity

```go
emailSendingSubdomainResponseSingle := client.EmailSendingSubdomainResponseSingle(nil)
fmt.Println(emailSendingSubdomainResponseSingle.GetName()) // "email_sending_subdomain_response_single"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `bool` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | Yes | The domain name within the zone. |
| `preview_enabled` | `bool` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingSubdomainResponseSingle(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSendingSubdomainResponseSingle(nil).Create(map[string]any{
    "zone_id": "example_zone_id",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailSendingSubdomainResponseSingle(nil).Update(map[string]any{
    "subdomain_id": "subdomain_id",
    "zone_id": "zone_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSubdomainsResponseCollectionEntity

```go
emailSendingSubdomainsResponseCollection := client.EmailSendingSubdomainsResponseCollection(nil)
fmt.Println(emailSendingSubdomainsResponseCollection.GetName()) // "email_sending_subdomains_response_collection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailSendingSubdomainsResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EmailSendingSuppressionEntity

```go
emailSendingSuppression := client.EmailSendingSuppression(nil)
fmt.Println(emailSendingSuppression.GetName()) // "email_sending_suppression"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | When the suppression was created. |
| `email` | `string` | Yes | The email address to suppress. |
| `errors` | `[]any` | Yes |  |
| `expires_at` | `string` | Yes | Expiration timestamp for the suppression. |
| `id` | `string` | Yes | Unique identifier for this suppression. |
| `items` | `[]any` | Yes | Suppressions to import. |
| `messages` | `[]any` | Yes |  |
| `note` | `string` | No | Advisory note for this suppression. |
| `read_only` | `bool` | Yes | Whether clients may mutate this suppression. |
| `reason` | `string` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `map[string]any` | Yes |  |
| `scope` | `any` | No | Where the suppression applies. |
| `success` | `bool` | Yes |  |
| `zones` | `[]any` | No |  |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EmailSendingSuppression(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EmailSendingSuppression(nil).Load(map[string]any{"id": "email_sending_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EmailSendingSuppression(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "created_at": "example_created_at",
    "email": "example_email",
    "errors": []any{},
    "expires_at": "example_expires_at",
    "id": "example_id",
    "items": []any{},
    "messages": []any{},
    "read_only": true,
    "reason": "example_reason",
    "result": map[string]any{},
    "success": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EmailSendingSuppression(nil).Update(map[string]any{
    "id": "email_sending_suppression_id",
    "account_id": "account_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EmailSendingSuppression(nil).Remove(map[string]any{"id": "email_sending_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailSendingSuppressionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SubdomainEntity

```go
subdomain := client.Subdomain(nil)
fmt.Println(subdomain.GetName()) // "subdomain"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Subdomain(nil).Remove(map[string]any{"id": "id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SubdomainEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```go
client := sdk.NewCloudflareEmailSecuritySDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
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

