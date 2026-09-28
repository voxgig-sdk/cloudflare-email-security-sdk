# CloudflareEmailSecurity TypeScript SDK Reference

Complete API reference for the CloudflareEmailSecurity TypeScript SDK.


## CloudflareEmailSecuritySDK

### Constructor

```ts
new CloudflareEmailSecuritySDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflareEmailSecuritySDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = CloudflareEmailSecuritySDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `CloudflareEmailSecuritySDK` instance in test mode.


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle(data?: object)`

Create a new `EmailAccountRulesPlanResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailAccountRulesPlanResponseSingleEntity` instance.

#### `EmailAccountRulesResponseCollection(data?: object)`

Create a new `EmailAccountRulesResponseCollection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailAccountRulesResponseCollectionEntity` instance.

#### `EmailAuth(data?: object)`

Create a new `EmailAuth` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailAuthEntity` instance.

#### `EmailCatchAllRuleResponseSingle(data?: object)`

Create a new `EmailCatchAllRuleResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailCatchAllRuleResponseSingleEntity` instance.

#### `EmailDestinationAddressResponseSingle(data?: object)`

Create a new `EmailDestinationAddressResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailDestinationAddressResponseSingleEntity` instance.

#### `EmailDestinationAddressesResponseCollection(data?: object)`

Create a new `EmailDestinationAddressesResponseCollection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailDestinationAddressesResponseCollectionEntity` instance.

#### `EmailDnsSettingsResponseCollection(data?: object)`

Create a new `EmailDnsSettingsResponseCollection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailDnsSettingsResponseCollectionEntity` instance.

#### `EmailEmailSettingsResponseSingle(data?: object)`

Create a new `EmailEmailSettingsResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailEmailSettingsResponseSingleEntity` instance.

#### `EmailRoutingSuppression(data?: object)`

Create a new `EmailRoutingSuppression` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailRoutingSuppressionEntity` instance.

#### `EmailRuleResponseSingle(data?: object)`

Create a new `EmailRuleResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailRuleResponseSingleEntity` instance.

#### `EmailRulesResponseCollection(data?: object)`

Create a new `EmailRulesResponseCollection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailRulesResponseCollectionEntity` instance.

#### `EmailSecurity(data?: object)`

Create a new `EmailSecurity` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSecurityEntity` instance.

#### `EmailSecurityAnalytics(data?: object)`

Create a new `EmailSecurityAnalytics` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSecurityAnalyticsEntity` instance.

#### `EmailSecuritySetting(data?: object)`

Create a new `EmailSecuritySetting` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSecuritySettingEntity` instance.

#### `EmailSending(data?: object)`

Create a new `EmailSending` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingEntity` instance.

#### `EmailSendingLimitsResponseSingle(data?: object)`

Create a new `EmailSendingLimitsResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingLimitsResponseSingleEntity` instance.

#### `EmailSendingReputationResponseSingle(data?: object)`

Create a new `EmailSendingReputationResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingReputationResponseSingleEntity` instance.

#### `EmailSendingSubdomainDnsStatus(data?: object)`

Create a new `EmailSendingSubdomainDnsStatus` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSubdomainDnsStatusEntity` instance.

#### `EmailSendingSubdomainPreview(data?: object)`

Create a new `EmailSendingSubdomainPreview` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSubdomainPreviewEntity` instance.

#### `EmailSendingSubdomainReputationComplaint(data?: object)`

Create a new `EmailSendingSubdomainReputationComplaint` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSubdomainReputationComplaintEntity` instance.

#### `EmailSendingSubdomainResponseSingle(data?: object)`

Create a new `EmailSendingSubdomainResponseSingle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSubdomainResponseSingleEntity` instance.

#### `EmailSendingSubdomainsResponseCollection(data?: object)`

Create a new `EmailSendingSubdomainsResponseCollection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSubdomainsResponseCollectionEntity` instance.

#### `EmailSendingSuppression(data?: object)`

Create a new `EmailSendingSuppression` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailSendingSuppressionEntity` instance.

#### `Subdomain(data?: object)`

Create a new `Subdomain` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SubdomainEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `CloudflareEmailSecuritySDK.test()`.

**Returns:** `CloudflareEmailSecuritySDK` instance in test mode.


---

## EmailAccountRulesPlanResponseSingleEntity

```ts
const email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `any[]` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `any[]` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailAccountRulesPlanResponseSingle().create({
  account_id: 'example_account_id',
  owner_worker_tag: 'example_owner_worker_tag',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailAccountRulesResponseCollectionEntity

```ts
const email_account_rules_response_collection = client.EmailAccountRulesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `any[]` | No |  |
| `result_info` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailAccountRulesResponseCollection().list({ account_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailAuthEntity

```ts
const email_auth = client.EmailAuth()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | No | Enable or disable DMARC reports for this zone |
| `result` | `Record<string, any>` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `boolean` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailAuth().load({ id: 'email_auth_id', zone_id: 'zone_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailAuth().update({
  zone_id: 'zone_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailAuthEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailCatchAllRuleResponseSingleEntity

```ts
const email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `any[]` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `boolean` | No | Routing rule status. |
| `matchers` | `any[]` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `Record<string, any>` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailCatchAllRuleResponseSingle().load({ zone_id: 'zone_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailCatchAllRuleResponseSingle().update({
  zone_id: 'zone_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailDestinationAddressResponseSingleEntity

```ts
const email_destination_address_response_single = client.EmailDestinationAddressResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes | The contact email address of the user. |
| `result` | `Record<string, any>` | No |  |
| `status` | `string` | Yes | Destination address status. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailDestinationAddressResponseSingle().create({
  account_id: 'example_account_id',
  email: 'example_email',
  status: 'example_status',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailDestinationAddressResponseSingle().load({ account_id: 'account_id', destination_address_identifier: 'destination_address_identifier' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailDestinationAddressResponseSingle().remove({ account_id: 'account_id', destination_address_identifier: 'destination_address_identifier' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailDestinationAddressResponseSingle().update({
  account_id: 'account_id',
  destination_address_identifier: 'destination_address_identifier',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailDestinationAddressesResponseCollectionEntity

```ts
const email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `any[]` | No |  |
| `result_info` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailDestinationAddressesResponseCollection().list({ account_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailDnsSettingsResponseCollectionEntity

```ts
const email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailDnsSettingsResponseCollection().list({ zone_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailEmailSettingsResponseSingleEntity

```ts
const email_email_settings_response_single = client.EmailEmailSettingsResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | No | State of your zone Email Routing settings. |
| `name` | `string` | No | Domain of your zone. |
| `result` | `Record<string, any>` | No |  |
| `skip_wizard` | `boolean` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `boolean` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailEmailSettingsResponseSingle().create({
  zone_id: 'example_zone_id',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailEmailSettingsResponseSingle().load({ zone_id: 'zone_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailEmailSettingsResponseSingle().remove({ zone_id: 'zone_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailEmailSettingsResponseSingle().update({
  zone_id: 'zone_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailRoutingSuppressionEntity

```ts
const email_routing_suppression = client.EmailRoutingSuppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `expires_at` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `reason` | `string` | Yes |  |
| `zones` | `any[]` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailRoutingSuppression().create({
  account_id: 'example_account_id',
  created_at: 'example_created_at',
  email: 'example_email',
  expires_at: 'example_expires_at',
  id: 'example_id',
  reason: 'example_reason',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailRoutingSuppression().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailRoutingSuppression().load({ id: 'email_routing_suppression_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailRoutingSuppression().remove({ id: 'email_routing_suppression_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailRoutingSuppressionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailRuleResponseSingleEntity

```ts
const email_rule_response_single = client.EmailRuleResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `any[]` | Yes | List actions patterns. |
| `enabled` | `boolean` | No | Routing rule status. |
| `matchers` | `any[]` | Yes | Matching patterns to forward to your actions. |
| `name` | `string` | No | Routing rule name. |
| `owner_worker_tag` | `string` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `number` | No | Priority of the routing rule. |
| `result` | `Record<string, any>` | No |  |
| `source` | `string` | No | Who manages the rule. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailRuleResponseSingle().create({
  zone_id: 'example_zone_id',
  actions: [],
  matchers: [],
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailRuleResponseSingle().load({ rule_identifier: 'rule_identifier', zone_id: 'zone_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailRuleResponseSingle().remove({ rule_identifier: 'rule_identifier', zone_id: 'zone_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailRuleResponseSingle().update({
  rule_identifier: 'rule_identifier',
  zone_id: 'zone_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailRuleResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailRulesResponseCollectionEntity

```ts
const email_rules_response_collection = client.EmailRulesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `any[]` | No |  |
| `result_info` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailRulesResponseCollection().list({ zone_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailRulesResponseCollectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSecurityEntity

```ts
const email_security = client.EmailSecurity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `string` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | No | Optional note describing the job. |
| `destination` | `any` | No | Required when action is 'MOVE'. |
| `errors` | `any[]` | Yes |  |
| `expected_disposition` | `string` | No | Nonfunctional field. |
| `id` | `string` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `any[]` | Yes | List of message IDs to move. |
| `messages` | `any[]` | Yes |  |
| `postfix_ids` | `any[]` | No | Deprecated, use `ids` instead. |
| `result` | `Record<string, any>` | Yes |  |
| `result_info` | `Record<string, any>` | Yes |  |
| `search_params` | `Record<string, any>` | Yes |  |
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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `cancel` | `/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel` | `client.EmailSecurity().create({ $action: 'cancel', ... })` |
| `move` | `/accounts/{account_id}/email-security/investigate/{investigate_id}/move` | `client.EmailSecurity().create({ $action: 'move', ... })` |
| `reclassify` | `/accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify` | `client.EmailSecurity().create({ $action: 'reclassify', ... })` |
| `submission` | `/accounts/{account_id}/email-security/submissions` | `client.EmailSecurity().create({ $action: 'submission', ... })` |
| `investigate` | `/accounts/{account_id}/email-security/investigate` | `client.EmailSecurity().list({ $action: 'investigate', ... })` |
| `submission` | `/accounts/{account_id}/email-security/submissions` | `client.EmailSecurity().list({ $action: 'submission', ... })` |

An action returns that action's OWN response, which is not necessarily a
EmailSecurity record — check the API definition for its shape.

```ts
const result = await client.EmailSecurity().create({
  $action: 'cancel',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSecurity().create({
  account_id: 'example_account_id',
  action: 'example_action',
  errors: [],
  id: 'example_id',
  ids: [],
  messages: [],
  result: {},
  result_info: {},
  search_params: {},
  success: true,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailSecurity().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSecurity().load({ account_id: 'account_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailSecurity().remove({ account_id: 'account_id', job_id: 'job_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSecurityEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSecurityAnalyticsEntity

```ts
const email_security_analytics = client.EmailSecurityAnalytics()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSecurityAnalytics().load({ account_id: 'account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSecurityAnalyticsEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSecuritySettingEntity

```ts
const email_security_setting = client.EmailSecuritySetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `any[]` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `string` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `any[]` | Yes | IDs of the allow policies to delete. |
| `domain` | `string` | Yes | The email domain to protect. |
| `drop_dispositions` | `any[]` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `any[]` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `number` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `any[]` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `any[]` | Yes | Allow policies to create. |
| `puts` | `any[]` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `any[]` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `boolean` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `boolean` | No | Require TLS on outbound connections. |
| `result` | `Record<string, any>` | No | An impersonation registry entry. |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSecuritySetting().create({
  account_id: 'example_account_id',
  allowed_delivery_modes: [],
  deletes: [],
  domain: 'example_domain',
  drop_dispositions: [],
  ip_restrictions: [],
  patches: [],
  pattern: 'example_pattern',
  posts: [],
  puts: [],
  regions: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailSecuritySetting().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSecuritySetting().load({ account_id: 'account_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailSecuritySetting().remove({ account_id: 'account_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailSecuritySetting().update({
  account_id: 'account_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSecuritySettingEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingEntity

```ts
const email_sending = client.EmailSending()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any[]` | No | File attachments and inline images. |
| `bcc` | `any` | No | Blind carbon copy recipient(s). |
| `cc` | `any` | No | Carbon copy recipient(s). |
| `errors` | `any[]` | Yes |  |
| `from` | `any` | Yes | Sender email address. |
| `headers` | `Record<string, any>` | No | Custom email headers as key-value pairs. |
| `html` | `string` | No | HTML body of the email. |
| `messages` | `any[]` | Yes |  |
| `mime_message` | `string` | Yes | The full MIME-encoded email message. |
| `recipients` | `any[]` | Yes | List of recipient email addresses. |
| `reply_to` | `any` | No | Reply-to address. |
| `result` | `Record<string, any>` | Yes |  |
| `result_info` | `Record<string, any>` | Yes |  |
| `subject` | `string` | Yes | Email subject line. |
| `success` | `boolean` | Yes |  |
| `text` | `string` | No | Plain text body of the email. |
| `to` | `any` | No | Recipient(s). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSending().create({
  account_id: 'example_account_id',
  errors: [],
  from: 'example_from',
  messages: [],
  mime_message: 'example_mime_message',
  recipients: [],
  result: {},
  result_info: {},
  subject: 'example_subject',
  success: true,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSending().load({ account_id: 'account_id', message_id: 'message_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingLimitsResponseSingleEntity

```ts
const email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingLimitsResponseSingle().load({ account_id: 'account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingReputationResponseSingleEntity

```ts
const email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Record<string, any>` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingReputationResponseSingle().load({ account_id: 'account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSubdomainDnsStatusEntity

```ts
const email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Record<string, any>` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSendingSubdomainDnsStatus().create({
  subdomain_id: 'example_subdomain_id',
  zone_id: 'example_zone_id',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingSubdomainDnsStatus().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSubdomainPreviewEntity

```ts
const email_sending_subdomain_preview = client.EmailSendingSubdomainPreview()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The domain name within the zone. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSendingSubdomainPreview().create({
  zone_id: 'example_zone_id',
  name: 'example_name',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSubdomainReputationComplaintEntity

```ts
const email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `Record<string, any>` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingSubdomainReputationComplaint().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id', end_at: 'end_at', start_at: 'start_at' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSubdomainResponseSingleEntity

```ts
const email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `boolean` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | Yes | The domain name within the zone. |
| `preview_enabled` | `boolean` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `any` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSendingSubdomainResponseSingle().create({
  zone_id: 'example_zone_id',
  name: 'example_name',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingSubdomainResponseSingle().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailSendingSubdomainResponseSingle().update({
  subdomain_id: 'subdomain_id',
  zone_id: 'zone_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSubdomainsResponseCollectionEntity

```ts
const email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailSendingSubdomainsResponseCollection().list({ zone_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailSendingSuppressionEntity

```ts
const email_sending_suppression = client.EmailSendingSuppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | When the suppression was created. |
| `email` | `string` | Yes | The email address to suppress. |
| `errors` | `any[]` | Yes |  |
| `expires_at` | `string` | Yes | Expiration timestamp for the suppression. |
| `id` | `string` | Yes | Unique identifier for this suppression. |
| `items` | `any[]` | Yes | Suppressions to import. |
| `messages` | `any[]` | Yes |  |
| `note` | `string` | No | Advisory note for this suppression. |
| `read_only` | `boolean` | Yes | Whether clients may mutate this suppression. |
| `reason` | `string` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `Record<string, any>` | Yes |  |
| `scope` | `any` | No | Where the suppression applies. |
| `success` | `boolean` | Yes |  |
| `zones` | `any[]` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EmailSendingSuppression().create({
  account_id: 'example_account_id',
  created_at: 'example_created_at',
  email: 'example_email',
  errors: [],
  expires_at: 'example_expires_at',
  id: 'example_id',
  items: [],
  messages: [],
  read_only: true,
  reason: 'example_reason',
  result: {},
  success: true,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EmailSendingSuppression().list({ account_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EmailSendingSuppression().load({ id: 'email_sending_suppression_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EmailSendingSuppression().remove({ id: 'email_sending_suppression_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EmailSendingSuppression().update({
  id: 'email_sending_suppression_id',
  account_id: 'account_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailSendingSuppressionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SubdomainEntity

```ts
const subdomain = client.Subdomain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Subdomain().remove({ id: 'id', zone_id: 'zone_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SubdomainEntity` instance with the same client and
options.

#### `client()`

Return the parent `CloudflareEmailSecuritySDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```ts
const client = new CloudflareEmailSecuritySDK({
  feature: {
    test: { active: true },
  }
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

