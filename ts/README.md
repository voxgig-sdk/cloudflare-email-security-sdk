# CloudflareEmailSecurity TypeScript SDK



The TypeScript SDK for the CloudflareEmailSecurity API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.EmailAccountRulesPlanResponseSingle()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { CloudflareEmailSecuritySDK } from '@voxgig-sdk/cloudflare-email-security-sdk'

const client = new CloudflareEmailSecuritySDK({
  apikey: process.env.CLOUDFLARE_EMAIL_SECURITY_APIKEY,
})
```

### 3. Load an emailsendingsubdomaindnsstatus

EmailSendingSubdomainDnsStatus is nested under subdomain, so provide the `subdomain_id`.
`load()` returns the entity directly and throws on failure:

```ts
try {
  const emailsendingsubdomaindnsstatus = await client.EmailSendingSubdomainDnsStatus().load({
    subdomain_id: 'example_subdomain_id',
    zone_id: 'example_zone_id',
  })
  console.log(emailsendingsubdomaindnsstatus)
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created EmailAccountRulesPlanResponseSingle ENTITY (.data() for the record)
const created = await client.EmailAccountRulesPlanResponseSingle().create({
  account_id: 'example_account_id',
  owner_worker_tag: 'example_owner_worker_tag',
})

```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const emaildestinationaddressresponsesingle = await client.EmailDestinationAddressResponseSingle().load({ account_id: "example", destination_address_identifier: "example" })
  console.log(emaildestinationaddressresponsesingle)
} catch (err) {
  console.error('load failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = CloudflareEmailSecuritySDK.test()

const emaildestinationaddressresponsesingle = await client.EmailDestinationAddressResponseSingle().load({ account_id: 'example_account_id', destination_address_identifier: 'example_destination_address_identifier' })
// emaildestinationaddressresponsesingle is the entity, populated with mock response data
// — call emaildestinationaddressresponsesingle.data() for the record itself
console.log(emaildestinationaddressresponsesingle)
```

You can also use the instance method:

```ts
const client = new CloudflareEmailSecuritySDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.EmailDestinationAddressResponseSingle()

// First call runs the operation and stores its result
await entity.load({ account_id: 'example_account_id', destination_address_identifier: 'example_destination_address_identifier' })

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new CloudflareEmailSecuritySDK({
  apikey: '...',
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE
CLOUDFLARE_EMAIL_SECURITY_APIKEY=<your-key>
```

Then run:

```bash
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### CloudflareEmailSecuritySDK

#### Constructor

```ts
new CloudflareEmailSecuritySDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `EmailAccountRulesPlanResponseSingle(data?)` | `EmailAccountRulesPlanResponseSingleEntity` | Create an EmailAccountRulesPlanResponseSingle entity instance. |
| `EmailAccountRulesResponseCollection(data?)` | `EmailAccountRulesResponseCollectionEntity` | Create an EmailAccountRulesResponseCollection entity instance. |
| `EmailAuth(data?)` | `EmailAuthEntity` | Create an EmailAuth entity instance. |
| `EmailCatchAllRuleResponseSingle(data?)` | `EmailCatchAllRuleResponseSingleEntity` | Create an EmailCatchAllRuleResponseSingle entity instance. |
| `EmailDestinationAddressResponseSingle(data?)` | `EmailDestinationAddressResponseSingleEntity` | Create an EmailDestinationAddressResponseSingle entity instance. |
| `EmailDestinationAddressesResponseCollection(data?)` | `EmailDestinationAddressesResponseCollectionEntity` | Create an EmailDestinationAddressesResponseCollection entity instance. |
| `EmailDnsSettingsResponseCollection(data?)` | `EmailDnsSettingsResponseCollectionEntity` | Create an EmailDnsSettingsResponseCollection entity instance. |
| `EmailEmailSettingsResponseSingle(data?)` | `EmailEmailSettingsResponseSingleEntity` | Create an EmailEmailSettingsResponseSingle entity instance. |
| `EmailRoutingSuppression(data?)` | `EmailRoutingSuppressionEntity` | Create an EmailRoutingSuppression entity instance. |
| `EmailRuleResponseSingle(data?)` | `EmailRuleResponseSingleEntity` | Create an EmailRuleResponseSingle entity instance. |
| `EmailRulesResponseCollection(data?)` | `EmailRulesResponseCollectionEntity` | Create an EmailRulesResponseCollection entity instance. |
| `EmailSecurity(data?)` | `EmailSecurityEntity` | Create an EmailSecurity entity instance. |
| `EmailSecurityAnalytics(data?)` | `EmailSecurityAnalyticsEntity` | Create an EmailSecurityAnalytics entity instance. |
| `EmailSecuritySetting(data?)` | `EmailSecuritySettingEntity` | Create an EmailSecuritySetting entity instance. |
| `EmailSending(data?)` | `EmailSendingEntity` | Create an EmailSending entity instance. |
| `EmailSendingLimitsResponseSingle(data?)` | `EmailSendingLimitsResponseSingleEntity` | Create an EmailSendingLimitsResponseSingle entity instance. |
| `EmailSendingReputationResponseSingle(data?)` | `EmailSendingReputationResponseSingleEntity` | Create an EmailSendingReputationResponseSingle entity instance. |
| `EmailSendingSubdomainDnsStatus(data?)` | `EmailSendingSubdomainDnsStatusEntity` | Create an EmailSendingSubdomainDnsStatus entity instance. |
| `EmailSendingSubdomainPreview(data?)` | `EmailSendingSubdomainPreviewEntity` | Create an EmailSendingSubdomainPreview entity instance. |
| `EmailSendingSubdomainReputationComplaint(data?)` | `EmailSendingSubdomainReputationComplaintEntity` | Create an EmailSendingSubdomainReputationComplaint entity instance. |
| `EmailSendingSubdomainResponseSingle(data?)` | `EmailSendingSubdomainResponseSingleEntity` | Create an EmailSendingSubdomainResponseSingle entity instance. |
| `EmailSendingSubdomainsResponseCollection(data?)` | `EmailSendingSubdomainsResponseCollectionEntity` | Create an EmailSendingSubdomainsResponseCollection entity instance. |
| `EmailSendingSuppression(data?)` | `EmailSendingSuppressionEntity` | Create an EmailSendingSuppression entity instance. |
| `Subdomain(data?)` | `SubdomainEntity` | Create a Subdomain entity instance. |
| `tester(testopts?, sdkopts?)` | `CloudflareEmailSecuritySDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `CloudflareEmailSecuritySDK.test(testopts?, sdkopts?)` | `CloudflareEmailSecuritySDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): CloudflareEmailSecuritySDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### EmailAccountRulesPlanResponseSingle

| Field | Description |
| --- | --- |
| `catch_all_rules` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | Desired normal Email Routing rules managed by the deploying Worker. |

Operations: create.

API path: `/accounts/{account_id}/email/routing/rules/plan`

#### EmailAccountRulesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: list.

API path: `/accounts/{account_id}/email/routing/rules`

#### EmailAuth

| Field | Description |
| --- | --- |
| `enabled` | Enable or disable DMARC reports for this zone |
| `result` | Recursive SPF inspection tree |
| `skip_wizard` | Skip the DMARC setup wizard |

Operations: load, update.

API path: `/zones/{zone_id}/email/auth/spf/inspect`

#### EmailCatchAllRuleResponseSingle

| Field | Description |
| --- | --- |
| `actions` | List actions for the catch-all routing rule. |
| `enabled` | Routing rule status. |
| `matchers` | List of matchers for the catch-all routing rule. |
| `name` | Routing rule name. |
| `owner_worker_tag` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` |  |
| `source` | Who manages the rule. |

Operations: load, update.

API path: `/zones/{zone_id}/email/routing/rules/catch_all`

#### EmailDestinationAddressResponseSingle

| Field | Description |
| --- | --- |
| `email` | The contact email address of the user. |
| `result` |  |
| `status` | Destination address status. |

Operations: create, load, remove, update.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDestinationAddressesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: list.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDnsSettingsResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |

Operations: list.

API path: `/zones/{zone_id}/email/routing/dns`

#### EmailEmailSettingsResponseSingle

| Field | Description |
| --- | --- |
| `enabled` | State of your zone Email Routing settings. |
| `name` | Domain of your zone. |
| `result` |  |
| `skip_wizard` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

Operations: create, load, patch, remove, update.

API path: `/zones/{zone_id}/email/routing/disable`

#### EmailRoutingSuppression

| Field | Description |
| --- | --- |
| `created_at` |  |
| `email` |  |
| `expires_at` |  |
| `id` |  |
| `reason` |  |
| `zones` |  |

Operations: create, list, load, remove.

API path: `/accounts/{account_id}/email/routing/suppression`

#### EmailRuleResponseSingle

| Field | Description |
| --- | --- |
| `actions` | List actions patterns. |
| `enabled` | Routing rule status. |
| `matchers` | Matching patterns to forward to your actions. |
| `name` | Routing rule name. |
| `owner_worker_tag` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | Priority of the routing rule. |
| `result` |  |
| `source` | Who manages the rule. |

Operations: create, load, remove, update.

API path: `/zones/{zone_id}/email/routing/rules`

#### EmailRulesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: list.

API path: `/zones/{zone_id}/email/routing/rules`

#### EmailSecurity

| Field | Description |
| --- | --- |
| `action` | The action the job performs on every message matching the search parameters. |
| `comment` | Optional note describing the job. |
| `destination` | Required when action is 'MOVE'. |
| `errors` |  |
| `expected_disposition` | Nonfunctional field. |
| `id` | Unique identifier for a message retrieved from investigation. |
| `ids` | List of message IDs to move. |
| `messages` |  |
| `postfix_ids` | Deprecated, use `ids` instead. |
| `result` |  |
| `result_info` |  |
| `search_params` |  |
| `success` | Whether the API call was successful. |

Operations: create, list, load, remove.

API path: `/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel`

#### EmailSecurityAnalytics

| Field | Description |
| --- | --- |
| `result` |  |

Operations: load.

API path: `/accounts/{account_id}/email-security/analytics/monthly_report`

#### EmailSecuritySetting

| Field | Description |
| --- | --- |
| `allowed_delivery_modes` | Delivery modes to onboard the domain through. |
| `comments` | Optional note describing the reason for the ignore pattern. |
| `deletes` | IDs of the allow policies to delete. |
| `domain` | The email domain to protect. |
| `drop_dispositions` | Dispositions to drop instead of delivering, e.g. |
| `folder` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | Source IP ranges mail is accepted from. |
| `lookback_hops` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | Allow policies to create. |
| `puts` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | Require TLS on inbound connections. |
| `require_tls_outbound` | Require TLS on outbound connections. |
| `result` | An impersonation registry entry. |
| `transport` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

Operations: create, list, load, patch, remove, update.

API path: `/accounts/{account_id}/email-security/settings/allow_policies`

#### EmailSending

| Field | Description |
| --- | --- |
| `attachments` | File attachments and inline images. |
| `bcc` | Blind carbon copy recipient(s). |
| `cc` | Carbon copy recipient(s). |
| `errors` |  |
| `from` | Sender email address. |
| `headers` | Custom email headers as key-value pairs. |
| `html` | HTML body of the email. |
| `messages` |  |
| `mime_message` | The full MIME-encoded email message. |
| `recipients` | List of recipient email addresses. |
| `reply_to` | Reply-to address. |
| `result` |  |
| `result_info` |  |
| `subject` | Email subject line. |
| `success` |  |
| `text` | Plain text body of the email. |
| `to` | Recipient(s). |

Operations: create, load.

API path: `/accounts/{account_id}/email/sending/send`

#### EmailSendingLimitsResponseSingle

| Field | Description |
| --- | --- |
| `result` |  |

Operations: load.

API path: `/accounts/{account_id}/email/sending/limits`

#### EmailSendingReputationResponseSingle

| Field | Description |
| --- | --- |
| `result` |  |

Operations: load.

API path: `/accounts/{account_id}/email/sending/reputation`

#### EmailSendingSubdomainDnsStatus

| Field | Description |
| --- | --- |
| `result` |  |

Operations: create, load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns`

#### EmailSendingSubdomainPreview

| Field | Description |
| --- | --- |
| `name` | The domain name within the zone. |

Operations: create.

API path: `/zones/{zone_id}/email/sending/subdomains/preview`

#### EmailSendingSubdomainReputationComplaint

| Field | Description |
| --- | --- |
| `result` |  |

Operations: load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints`

#### EmailSendingSubdomainResponseSingle

| Field | Description |
| --- | --- |
| `drop_suppressed_recipients` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | The domain name within the zone. |
| `preview_enabled` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` |  |

Operations: create, load, update.

API path: `/zones/{zone_id}/email/sending/subdomains`

#### EmailSendingSubdomainsResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |

Operations: list.

API path: `/zones/{zone_id}/email/sending/subdomains`

#### EmailSendingSuppression

| Field | Description |
| --- | --- |
| `created_at` | When the suppression was created. |
| `email` | The email address to suppress. |
| `errors` |  |
| `expires_at` | Expiration timestamp for the suppression. |
| `id` | Unique identifier for this suppression. |
| `items` | Suppressions to import. |
| `messages` |  |
| `note` | Advisory note for this suppression. |
| `read_only` | Whether clients may mutate this suppression. |
| `reason` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` |  |
| `scope` | Where the suppression applies. |
| `success` |  |
| `zones` |  |

Operations: create, list, load, remove, update.

API path: `/accounts/{account_id}/email/sending/suppression`

#### Subdomain

| Field | Description |
| --- | --- |
| `id` |  |

Operations: remove.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}`



## Entities


### EmailAccountRulesPlanResponseSingle

Create an instance: `const email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `any[]` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `any[]` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

```ts
const email_account_rules_plan_response_single = await client.EmailAccountRulesPlanResponseSingle().create({
  account_id: 'example_account_id',
  owner_worker_tag: 'example_owner_worker_tag',
})
```


### EmailAccountRulesResponseCollection

Create an instance: `const email_account_rules_response_collection = client.EmailAccountRulesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `any[]` |  |
| `result_info` | `Record<string, any>` |  |

#### Example: List

```ts
const email_account_rules_response_collections = await client.EmailAccountRulesResponseCollection().list({ account_id: "example" })
```


### EmailAuth

Create an instance: `const email_auth = client.EmailAuth()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `boolean` | Enable or disable DMARC reports for this zone |
| `result` | `Record<string, any>` | Recursive SPF inspection tree |
| `skip_wizard` | `boolean` | Skip the DMARC setup wizard |

#### Example: Load

```ts
const email_auth = await client.EmailAuth().load({ id: 'email_auth_id', zone_id: 'zone_id' })
```


### EmailCatchAllRuleResponseSingle

Create an instance: `const email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `any[]` | List actions for the catch-all routing rule. |
| `enabled` | `boolean` | Routing rule status. |
| `matchers` | `any[]` | List of matchers for the catch-all routing rule. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `Record<string, any>` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```ts
const email_catch_all_rule_response_single = await client.EmailCatchAllRuleResponseSingle().load({ zone_id: 'zone_id' })
```


### EmailDestinationAddressResponseSingle

Create an instance: `const email_destination_address_response_single = client.EmailDestinationAddressResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` | The contact email address of the user. |
| `result` | `Record<string, any>` |  |
| `status` | `string` | Destination address status. |

#### Example: Load

```ts
const email_destination_address_response_single = await client.EmailDestinationAddressResponseSingle().load({ account_id: 'account_id', destination_address_identifier: 'destination_address_identifier' })
```

#### Example: Create

```ts
const email_destination_address_response_single = await client.EmailDestinationAddressResponseSingle().create({
  account_id: 'example_account_id',
  email: 'example_email',
  status: 'example_status',
})
```


### EmailDestinationAddressesResponseCollection

Create an instance: `const email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `any[]` |  |
| `result_info` | `Record<string, any>` |  |

#### Example: List

```ts
const email_destination_addresses_response_collections = await client.EmailDestinationAddressesResponseCollection().list({ account_id: "example" })
```


### EmailDnsSettingsResponseCollection

Create an instance: `const email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `any[]` |  |

#### Example: List

```ts
const email_dns_settings_response_collections = await client.EmailDnsSettingsResponseCollection().list({ zone_id: "example" })
```


### EmailEmailSettingsResponseSingle

Create an instance: `const email_email_settings_response_single = client.EmailEmailSettingsResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `boolean` | State of your zone Email Routing settings. |
| `name` | `string` | Domain of your zone. |
| `result` | `Record<string, any>` |  |
| `skip_wizard` | `boolean` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `boolean` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```ts
const email_email_settings_response_single = await client.EmailEmailSettingsResponseSingle().load({ zone_id: 'zone_id' })
```

#### Example: Create

```ts
const email_email_settings_response_single = await client.EmailEmailSettingsResponseSingle().create({
  zone_id: 'example_zone_id',
})
```


### EmailRoutingSuppression

Create an instance: `const email_routing_suppression = client.EmailRoutingSuppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `expires_at` | `string` |  |
| `id` | `string` |  |
| `reason` | `string` |  |
| `zones` | `any[]` |  |

#### Example: Load

```ts
const email_routing_suppression = await client.EmailRoutingSuppression().load({ id: 'email_routing_suppression_id' })
```

#### Example: List

```ts
const email_routing_suppressions = await client.EmailRoutingSuppression().list({ account_id: "example" })
```

#### Example: Create

```ts
const email_routing_suppression = await client.EmailRoutingSuppression().create({
  account_id: 'example_account_id',
  created_at: 'example_created_at',
  email: 'example_email',
  expires_at: 'example_expires_at',
  id: 'example_id',
  reason: 'example_reason',
})
```


### EmailRuleResponseSingle

Create an instance: `const email_rule_response_single = client.EmailRuleResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `any[]` | List actions patterns. |
| `enabled` | `boolean` | Routing rule status. |
| `matchers` | `any[]` | Matching patterns to forward to your actions. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `number` | Priority of the routing rule. |
| `result` | `Record<string, any>` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```ts
const email_rule_response_single = await client.EmailRuleResponseSingle().load({ rule_identifier: 'rule_identifier', zone_id: 'zone_id' })
```

#### Example: Create

```ts
const email_rule_response_single = await client.EmailRuleResponseSingle().create({
  zone_id: 'example_zone_id',
  actions: [],
  matchers: [],
})
```


### EmailRulesResponseCollection

Create an instance: `const email_rules_response_collection = client.EmailRulesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `any[]` |  |
| `result_info` | `Record<string, any>` |  |

#### Example: List

```ts
const email_rules_response_collections = await client.EmailRulesResponseCollection().list({ zone_id: "example" })
```


### EmailSecurity

Create an instance: `const email_security = client.EmailSecurity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | Optional note describing the job. |
| `destination` | `any` | Required when action is 'MOVE'. |
| `errors` | `any[]` |  |
| `expected_disposition` | `string` | Nonfunctional field. |
| `id` | `string` | Unique identifier for a message retrieved from investigation. |
| `ids` | `any[]` | List of message IDs to move. |
| `messages` | `any[]` |  |
| `postfix_ids` | `any[]` | Deprecated, use `ids` instead. |
| `result` | `Record<string, any>` |  |
| `result_info` | `Record<string, any>` |  |
| `search_params` | `Record<string, any>` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Load

```ts
const email_security = await client.EmailSecurity().load({ account_id: 'account_id' })
```

#### Example: List

```ts
const email_securitys = await client.EmailSecurity().list({ account_id: "example" })
```

#### Example: Create

```ts
const email_security = await client.EmailSecurity().create({
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


### EmailSecurityAnalytics

Create an instance: `const email_security_analytics = client.EmailSecurityAnalytics()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Record<string, any>` |  |

#### Example: Load

```ts
const email_security_analytics = await client.EmailSecurityAnalytics().load({ account_id: 'account_id' })
```


### EmailSecuritySetting

Create an instance: `const email_security_setting = client.EmailSecuritySetting()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `allowed_delivery_modes` | `any[]` | Delivery modes to onboard the domain through. |
| `comments` | `string` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `any[]` | IDs of the allow policies to delete. |
| `domain` | `string` | The email domain to protect. |
| `drop_dispositions` | `any[]` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `any[]` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `number` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `any[]` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `any[]` | Allow policies to create. |
| `puts` | `any[]` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `any[]` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `boolean` | Require TLS on inbound connections. |
| `require_tls_outbound` | `boolean` | Require TLS on outbound connections. |
| `result` | `Record<string, any>` | An impersonation registry entry. |
| `transport` | `string` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```ts
const email_security_setting = await client.EmailSecuritySetting().load({ account_id: 'account_id' })
```

#### Example: List

```ts
const email_security_settings = await client.EmailSecuritySetting().list({ account_id: "example" })
```

#### Example: Create

```ts
const email_security_setting = await client.EmailSecuritySetting().create({
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


### EmailSending

Create an instance: `const email_sending = client.EmailSending()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any[]` | File attachments and inline images. |
| `bcc` | `any` | Blind carbon copy recipient(s). |
| `cc` | `any` | Carbon copy recipient(s). |
| `errors` | `any[]` |  |
| `from` | `any` | Sender email address. |
| `headers` | `Record<string, any>` | Custom email headers as key-value pairs. |
| `html` | `string` | HTML body of the email. |
| `messages` | `any[]` |  |
| `mime_message` | `string` | The full MIME-encoded email message. |
| `recipients` | `any[]` | List of recipient email addresses. |
| `reply_to` | `any` | Reply-to address. |
| `result` | `Record<string, any>` |  |
| `result_info` | `Record<string, any>` |  |
| `subject` | `string` | Email subject line. |
| `success` | `boolean` |  |
| `text` | `string` | Plain text body of the email. |
| `to` | `any` | Recipient(s). |

#### Example: Load

```ts
const email_sending = await client.EmailSending().load({ account_id: 'account_id', message_id: 'message_id' })
```

#### Example: Create

```ts
const email_sending = await client.EmailSending().create({
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


### EmailSendingLimitsResponseSingle

Create an instance: `const email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Record<string, any>` |  |

#### Example: Load

```ts
const email_sending_limits_response_single = await client.EmailSendingLimitsResponseSingle().load({ account_id: 'account_id' })
```


### EmailSendingReputationResponseSingle

Create an instance: `const email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Record<string, any>` |  |

#### Example: Load

```ts
const email_sending_reputation_response_single = await client.EmailSendingReputationResponseSingle().load({ account_id: 'account_id' })
```


### EmailSendingSubdomainDnsStatus

Create an instance: `const email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Record<string, any>` |  |

#### Example: Load

```ts
const email_sending_subdomain_dns_status = await client.EmailSendingSubdomainDnsStatus().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id' })
```

#### Example: Create

```ts
const email_sending_subdomain_dns_status = await client.EmailSendingSubdomainDnsStatus().create({
  subdomain_id: 'example_subdomain_id',
  zone_id: 'example_zone_id',
})
```


### EmailSendingSubdomainPreview

Create an instance: `const email_sending_subdomain_preview = client.EmailSendingSubdomainPreview()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `string` | The domain name within the zone. |

#### Example: Create

```ts
const email_sending_subdomain_preview = await client.EmailSendingSubdomainPreview().create({
  zone_id: 'example_zone_id',
  name: 'example_name',
})
```


### EmailSendingSubdomainReputationComplaint

Create an instance: `const email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Record<string, any>` |  |

#### Example: Load

```ts
const email_sending_subdomain_reputation_complaint = await client.EmailSendingSubdomainReputationComplaint().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id', end_at: 'end_at', start_at: 'start_at' })
```


### EmailSendingSubdomainResponseSingle

Create an instance: `const email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `drop_suppressed_recipients` | `boolean` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | The domain name within the zone. |
| `preview_enabled` | `boolean` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `any` |  |

#### Example: Load

```ts
const email_sending_subdomain_response_single = await client.EmailSendingSubdomainResponseSingle().load({ subdomain_id: 'subdomain_id', zone_id: 'zone_id' })
```

#### Example: Create

```ts
const email_sending_subdomain_response_single = await client.EmailSendingSubdomainResponseSingle().create({
  zone_id: 'example_zone_id',
  name: 'example_name',
})
```


### EmailSendingSubdomainsResponseCollection

Create an instance: `const email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `any[]` |  |

#### Example: List

```ts
const email_sending_subdomains_response_collections = await client.EmailSendingSubdomainsResponseCollection().list({ zone_id: "example" })
```


### EmailSendingSuppression

Create an instance: `const email_sending_suppression = client.EmailSendingSuppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | When the suppression was created. |
| `email` | `string` | The email address to suppress. |
| `errors` | `any[]` |  |
| `expires_at` | `string` | Expiration timestamp for the suppression. |
| `id` | `string` | Unique identifier for this suppression. |
| `items` | `any[]` | Suppressions to import. |
| `messages` | `any[]` |  |
| `note` | `string` | Advisory note for this suppression. |
| `read_only` | `boolean` | Whether clients may mutate this suppression. |
| `reason` | `string` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `Record<string, any>` |  |
| `scope` | `any` | Where the suppression applies. |
| `success` | `boolean` |  |
| `zones` | `any[]` |  |

#### Example: Load

```ts
const email_sending_suppression = await client.EmailSendingSuppression().load({ id: 'email_sending_suppression_id' })
```

#### Example: List

```ts
const email_sending_suppressions = await client.EmailSendingSuppression().list({ account_id: "example" })
```

#### Example: Create

```ts
const email_sending_suppression = await client.EmailSendingSuppression().create({
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


### Subdomain

Create an instance: `const subdomain = client.Subdomain()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | Test transport |

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Open types

3 fields are carried as open values rather than typed structures.
This follows from the API definition, not from a gap in this SDK: the
definition describes them with untagged unions —
`oneOf`/`anyOf` branches with no `discriminator` — so it never states which
variant a given value is. Nothing can select a branch reliably, so the SDK
passes the value through unchanged rather than assert a shape the API does not
guarantee.

| Entity | Field | Variants | Nesting |
| --- | --- | --- | --- |
| `email_sending` | `bcc` | 3 | 3 levels |
| `email_sending` | `cc` | 3 | 3 levels |
| `email_sending` | `to` | 3 | 3 levels |

These values round-trip unchanged — read them, modify them, send them back. If
the API adds a `discriminator` to the definition, regenerating will type them.
Every other field is typed normally.

## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
cloudflare-email-security/
├── src/
│   ├── CloudflareEmailSecuritySDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { CloudflareEmailSecuritySDK } from '@voxgig-sdk/cloudflare-email-security-sdk'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle()
await emaildestinationaddressresponsesingle.load({ account_id: "example", destination_address_identifier: "example" })

// emaildestinationaddressresponsesingle.data() now returns the emaildestinationaddressresponsesingle data from the last `load`
// emaildestinationaddressresponsesingle.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
