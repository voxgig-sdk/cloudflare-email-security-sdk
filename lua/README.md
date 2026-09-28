# CloudflareEmailSecurity Lua SDK



The Lua SDK for the CloudflareEmailSecurity API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:EmailAccountRulesPlanResponseSingle()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("cloudflare-email-security_sdk")

local client = sdk.new({
  apikey = os.getenv("CLOUDFLARE_EMAIL_SECURITY_APIKEY"),
})
```

### 3. Load an emailsendingsubdomaindnsstatus

EmailSendingSubdomainDnsStatus is nested under subdomain, so provide the `subdomain_id`.

```lua
local emailsendingsubdomaindnsstatus, err = client:EmailSendingSubdomainDnsStatus():load({ subdomain_id = "example_subdomain_id", zone_id = "example_zone_id" })
if err then error(err) end
print(emailsendingsubdomaindnsstatus)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:EmailAccountRulesPlanResponseSingle():create({ account_id = "example_account_id", owner_worker_tag = "example_owner_worker_tag" })
if err then error(err) end

```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local emaildestinationaddressresponsesingle, err = client:EmailDestinationAddressResponseSingle():load({ account_id = "example", destination_address_identifier = "example" })
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:EmailDestinationAddressResponseSingle():load({ account_id = "example", destination_address_identifier = "example" })
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
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
cd lua && busted test/
```


## Reference

### CloudflareEmailSecuritySDK

```lua
local sdk = require("cloudflare-email-security_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CloudflareEmailSecuritySDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `EmailAccountRulesPlanResponseSingle` | `(data) -> EmailAccountRulesPlanResponseSingleEntity` | Create an EmailAccountRulesPlanResponseSingle entity instance. |
| `EmailAccountRulesResponseCollection` | `(data) -> EmailAccountRulesResponseCollectionEntity` | Create an EmailAccountRulesResponseCollection entity instance. |
| `EmailAuth` | `(data) -> EmailAuthEntity` | Create an EmailAuth entity instance. |
| `EmailCatchAllRuleResponseSingle` | `(data) -> EmailCatchAllRuleResponseSingleEntity` | Create an EmailCatchAllRuleResponseSingle entity instance. |
| `EmailDestinationAddressResponseSingle` | `(data) -> EmailDestinationAddressResponseSingleEntity` | Create an EmailDestinationAddressResponseSingle entity instance. |
| `EmailDestinationAddressesResponseCollection` | `(data) -> EmailDestinationAddressesResponseCollectionEntity` | Create an EmailDestinationAddressesResponseCollection entity instance. |
| `EmailDnsSettingsResponseCollection` | `(data) -> EmailDnsSettingsResponseCollectionEntity` | Create an EmailDnsSettingsResponseCollection entity instance. |
| `EmailEmailSettingsResponseSingle` | `(data) -> EmailEmailSettingsResponseSingleEntity` | Create an EmailEmailSettingsResponseSingle entity instance. |
| `EmailRoutingSuppression` | `(data) -> EmailRoutingSuppressionEntity` | Create an EmailRoutingSuppression entity instance. |
| `EmailRuleResponseSingle` | `(data) -> EmailRuleResponseSingleEntity` | Create an EmailRuleResponseSingle entity instance. |
| `EmailRulesResponseCollection` | `(data) -> EmailRulesResponseCollectionEntity` | Create an EmailRulesResponseCollection entity instance. |
| `EmailSecurity` | `(data) -> EmailSecurityEntity` | Create an EmailSecurity entity instance. |
| `EmailSecurityAnalytics` | `(data) -> EmailSecurityAnalyticsEntity` | Create an EmailSecurityAnalytics entity instance. |
| `EmailSecuritySetting` | `(data) -> EmailSecuritySettingEntity` | Create an EmailSecuritySetting entity instance. |
| `EmailSending` | `(data) -> EmailSendingEntity` | Create an EmailSending entity instance. |
| `EmailSendingLimitsResponseSingle` | `(data) -> EmailSendingLimitsResponseSingleEntity` | Create an EmailSendingLimitsResponseSingle entity instance. |
| `EmailSendingReputationResponseSingle` | `(data) -> EmailSendingReputationResponseSingleEntity` | Create an EmailSendingReputationResponseSingle entity instance. |
| `EmailSendingSubdomainDnsStatus` | `(data) -> EmailSendingSubdomainDnsStatusEntity` | Create an EmailSendingSubdomainDnsStatus entity instance. |
| `EmailSendingSubdomainPreview` | `(data) -> EmailSendingSubdomainPreviewEntity` | Create an EmailSendingSubdomainPreview entity instance. |
| `EmailSendingSubdomainReputationComplaint` | `(data) -> EmailSendingSubdomainReputationComplaintEntity` | Create an EmailSendingSubdomainReputationComplaint entity instance. |
| `EmailSendingSubdomainResponseSingle` | `(data) -> EmailSendingSubdomainResponseSingleEntity` | Create an EmailSendingSubdomainResponseSingle entity instance. |
| `EmailSendingSubdomainsResponseCollection` | `(data) -> EmailSendingSubdomainsResponseCollectionEntity` | Create an EmailSendingSubdomainsResponseCollection entity instance. |
| `EmailSendingSuppression` | `(data) -> EmailSendingSuppressionEntity` | Create an EmailSendingSuppression entity instance. |
| `Subdomain` | `(data) -> SubdomainEntity` | Create a Subdomain entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local email_auth, err = client:EmailAuth():load({ id = "example_id" })
    if err then error(err) end
    -- email_auth is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### EmailAccountRulesPlanResponseSingle

| Field | Description |
| --- | --- |
| `catch_all_rules` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | Desired normal Email Routing rules managed by the deploying Worker. |

Operations: Create.

API path: `/accounts/{account_id}/email/routing/rules/plan`

#### EmailAccountRulesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: List.

API path: `/accounts/{account_id}/email/routing/rules`

#### EmailAuth

| Field | Description |
| --- | --- |
| `enabled` | Enable or disable DMARC reports for this zone |
| `result` | Recursive SPF inspection tree |
| `skip_wizard` | Skip the DMARC setup wizard |

Operations: Load, Update.

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

Operations: Load, Update.

API path: `/zones/{zone_id}/email/routing/rules/catch_all`

#### EmailDestinationAddressResponseSingle

| Field | Description |
| --- | --- |
| `email` | The contact email address of the user. |
| `result` |  |
| `status` | Destination address status. |

Operations: Create, Load, Remove, Update.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDestinationAddressesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: List.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDnsSettingsResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |

Operations: List.

API path: `/zones/{zone_id}/email/routing/dns`

#### EmailEmailSettingsResponseSingle

| Field | Description |
| --- | --- |
| `enabled` | State of your zone Email Routing settings. |
| `name` | Domain of your zone. |
| `result` |  |
| `skip_wizard` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

Operations: Create, Load, Patch, Remove, Update.

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

Operations: Create, List, Load, Remove.

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

Operations: Create, Load, Remove, Update.

API path: `/zones/{zone_id}/email/routing/rules`

#### EmailRulesResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |
| `result_info` |  |

Operations: List.

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

Operations: Create, List, Load, Remove.

API path: `/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel`

#### EmailSecurityAnalytics

| Field | Description |
| --- | --- |
| `result` |  |

Operations: Load.

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

Operations: Create, List, Load, Patch, Remove, Update.

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

Operations: Create, Load.

API path: `/accounts/{account_id}/email/sending/send`

#### EmailSendingLimitsResponseSingle

| Field | Description |
| --- | --- |
| `result` |  |

Operations: Load.

API path: `/accounts/{account_id}/email/sending/limits`

#### EmailSendingReputationResponseSingle

| Field | Description |
| --- | --- |
| `result` |  |

Operations: Load.

API path: `/accounts/{account_id}/email/sending/reputation`

#### EmailSendingSubdomainDnsStatus

| Field | Description |
| --- | --- |
| `result` |  |

Operations: Create, Load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns`

#### EmailSendingSubdomainPreview

| Field | Description |
| --- | --- |
| `name` | The domain name within the zone. |

Operations: Create.

API path: `/zones/{zone_id}/email/sending/subdomains/preview`

#### EmailSendingSubdomainReputationComplaint

| Field | Description |
| --- | --- |
| `result` |  |

Operations: Load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints`

#### EmailSendingSubdomainResponseSingle

| Field | Description |
| --- | --- |
| `drop_suppressed_recipients` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | The domain name within the zone. |
| `preview_enabled` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` |  |

Operations: Create, Load, Update.

API path: `/zones/{zone_id}/email/sending/subdomains`

#### EmailSendingSubdomainsResponseCollection

| Field | Description |
| --- | --- |
| `result` |  |

Operations: List.

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

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/email/sending/suppression`

#### Subdomain

| Field | Description |
| --- | --- |
| `id` |  |

Operations: Remove.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}`



## Entities


### EmailAccountRulesPlanResponseSingle

Create an instance: `local email_account_rules_plan_response_single = client:EmailAccountRulesPlanResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `table` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `table` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

```lua
local email_account_rules_plan_response_single, err = client:EmailAccountRulesPlanResponseSingle():create({
  account_id = "example_account_id", -- string
  owner_worker_tag = "example_owner_worker_tag", -- string
})
```


### EmailAccountRulesResponseCollection

Create an instance: `local email_account_rules_response_collection = client:EmailAccountRulesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |
| `result_info` | `table` |  |

#### Example: List

```lua
local email_account_rules_response_collections, err = client:EmailAccountRulesResponseCollection():list()
```


### EmailAuth

Create an instance: `local email_auth = client:EmailAuth(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `boolean` | Enable or disable DMARC reports for this zone |
| `result` | `table` | Recursive SPF inspection tree |
| `skip_wizard` | `boolean` | Skip the DMARC setup wizard |

#### Example: Load

```lua
local email_auth, err = client:EmailAuth():load({ id = "email_auth_id", zone_id = "zone_id" })
```


### EmailCatchAllRuleResponseSingle

Create an instance: `local email_catch_all_rule_response_single = client:EmailCatchAllRuleResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `table` | List actions for the catch-all routing rule. |
| `enabled` | `boolean` | Routing rule status. |
| `matchers` | `table` | List of matchers for the catch-all routing rule. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `table` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```lua
local email_catch_all_rule_response_single, err = client:EmailCatchAllRuleResponseSingle():load({ zone_id = "zone_id" })
```


### EmailDestinationAddressResponseSingle

Create an instance: `local email_destination_address_response_single = client:EmailDestinationAddressResponseSingle(nil)`

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
| `result` | `table` |  |
| `status` | `string` | Destination address status. |

#### Example: Load

```lua
local email_destination_address_response_single, err = client:EmailDestinationAddressResponseSingle():load({ account_id = "account_id", destination_address_identifier = "destination_address_identifier" })
```

#### Example: Create

```lua
local email_destination_address_response_single, err = client:EmailDestinationAddressResponseSingle():create({
  account_id = "example_account_id", -- string
  email = "example_email", -- string
  status = "example_status", -- string
})
```


### EmailDestinationAddressesResponseCollection

Create an instance: `local email_destination_addresses_response_collection = client:EmailDestinationAddressesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |
| `result_info` | `table` |  |

#### Example: List

```lua
local email_destination_addresses_response_collections, err = client:EmailDestinationAddressesResponseCollection():list()
```


### EmailDnsSettingsResponseCollection

Create an instance: `local email_dns_settings_response_collection = client:EmailDnsSettingsResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: List

```lua
local email_dns_settings_response_collections, err = client:EmailDnsSettingsResponseCollection():list()
```


### EmailEmailSettingsResponseSingle

Create an instance: `local email_email_settings_response_single = client:EmailEmailSettingsResponseSingle(nil)`

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
| `result` | `table` |  |
| `skip_wizard` | `boolean` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `boolean` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```lua
local email_email_settings_response_single, err = client:EmailEmailSettingsResponseSingle():load({ zone_id = "zone_id" })
```

#### Example: Create

```lua
local email_email_settings_response_single, err = client:EmailEmailSettingsResponseSingle():create({
  zone_id = "example_zone_id", -- string
})
```


### EmailRoutingSuppression

Create an instance: `local email_routing_suppression = client:EmailRoutingSuppression(nil)`

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
| `zones` | `table` |  |

#### Example: Load

```lua
local email_routing_suppression, err = client:EmailRoutingSuppression():load({ id = "email_routing_suppression_id" })
```

#### Example: List

```lua
local email_routing_suppressions, err = client:EmailRoutingSuppression():list()
```

#### Example: Create

```lua
local email_routing_suppression, err = client:EmailRoutingSuppression():create({
  account_id = "example_account_id", -- string
  created_at = "example_created_at", -- string
  email = "example_email", -- string
  expires_at = "example_expires_at", -- string
  id = "example_id", -- string
  reason = "example_reason", -- string
})
```


### EmailRuleResponseSingle

Create an instance: `local email_rule_response_single = client:EmailRuleResponseSingle(nil)`

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
| `actions` | `table` | List actions patterns. |
| `enabled` | `boolean` | Routing rule status. |
| `matchers` | `table` | Matching patterns to forward to your actions. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `number` | Priority of the routing rule. |
| `result` | `table` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```lua
local email_rule_response_single, err = client:EmailRuleResponseSingle():load({ rule_identifier = "rule_identifier", zone_id = "zone_id" })
```

#### Example: Create

```lua
local email_rule_response_single, err = client:EmailRuleResponseSingle():create({
  zone_id = "example_zone_id", -- string
  actions = {}, -- table
  matchers = {}, -- table
})
```


### EmailRulesResponseCollection

Create an instance: `local email_rules_response_collection = client:EmailRulesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |
| `result_info` | `table` |  |

#### Example: List

```lua
local email_rules_response_collections, err = client:EmailRulesResponseCollection():list()
```


### EmailSecurity

Create an instance: `local email_security = client:EmailSecurity(nil)`

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
| `errors` | `table` |  |
| `expected_disposition` | `string` | Nonfunctional field. |
| `id` | `string` | Unique identifier for a message retrieved from investigation. |
| `ids` | `table` | List of message IDs to move. |
| `messages` | `table` |  |
| `postfix_ids` | `table` | Deprecated, use `ids` instead. |
| `result` | `table` |  |
| `result_info` | `table` |  |
| `search_params` | `table` |  |
| `success` | `boolean` | Whether the API call was successful. |

#### Example: Load

```lua
local email_security, err = client:EmailSecurity():load({ account_id = "account_id" })
```

#### Example: List

```lua
local email_securitys, err = client:EmailSecurity():list()
```

#### Example: Create

```lua
local email_security, err = client:EmailSecurity():create({
  account_id = "example_account_id", -- string
  action = "example_action", -- string
  errors = {}, -- table
  id = "example_id", -- string
  ids = {}, -- table
  messages = {}, -- table
  result = {}, -- table
  result_info = {}, -- table
  search_params = {}, -- table
  success = true, -- boolean
})
```


### EmailSecurityAnalytics

Create an instance: `local email_security_analytics = client:EmailSecurityAnalytics(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: Load

```lua
local email_security_analytics, err = client:EmailSecurityAnalytics():load({ account_id = "account_id" })
```


### EmailSecuritySetting

Create an instance: `local email_security_setting = client:EmailSecuritySetting(nil)`

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
| `allowed_delivery_modes` | `table` | Delivery modes to onboard the domain through. |
| `comments` | `string` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `table` | IDs of the allow policies to delete. |
| `domain` | `string` | The email domain to protect. |
| `drop_dispositions` | `table` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `table` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `number` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `table` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `table` | Allow policies to create. |
| `puts` | `table` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `table` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `boolean` | Require TLS on inbound connections. |
| `require_tls_outbound` | `boolean` | Require TLS on outbound connections. |
| `result` | `table` | An impersonation registry entry. |
| `transport` | `string` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```lua
local email_security_setting, err = client:EmailSecuritySetting():load({ account_id = "account_id" })
```

#### Example: List

```lua
local email_security_settings, err = client:EmailSecuritySetting():list()
```

#### Example: Create

```lua
local email_security_setting, err = client:EmailSecuritySetting():create({
  account_id = "example_account_id", -- string
  allowed_delivery_modes = {}, -- table
  deletes = {}, -- table
  domain = "example_domain", -- string
  drop_dispositions = {}, -- table
  ip_restrictions = {}, -- table
  patches = {}, -- table
  pattern = "example_pattern", -- string
  posts = {}, -- table
  puts = {}, -- table
  regions = {}, -- table
})
```


### EmailSending

Create an instance: `local email_sending = client:EmailSending(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `table` | File attachments and inline images. |
| `bcc` | `any` | Blind carbon copy recipient(s). |
| `cc` | `any` | Carbon copy recipient(s). |
| `errors` | `table` |  |
| `from` | `any` | Sender email address. |
| `headers` | `table` | Custom email headers as key-value pairs. |
| `html` | `string` | HTML body of the email. |
| `messages` | `table` |  |
| `mime_message` | `string` | The full MIME-encoded email message. |
| `recipients` | `table` | List of recipient email addresses. |
| `reply_to` | `any` | Reply-to address. |
| `result` | `table` |  |
| `result_info` | `table` |  |
| `subject` | `string` | Email subject line. |
| `success` | `boolean` |  |
| `text` | `string` | Plain text body of the email. |
| `to` | `any` | Recipient(s). |

#### Example: Load

```lua
local email_sending, err = client:EmailSending():load({ account_id = "account_id", message_id = "message_id" })
```

#### Example: Create

```lua
local email_sending, err = client:EmailSending():create({
  account_id = "example_account_id", -- string
  errors = {}, -- table
  from = "example_from", -- any
  messages = {}, -- table
  mime_message = "example_mime_message", -- string
  recipients = {}, -- table
  result = {}, -- table
  result_info = {}, -- table
  subject = "example_subject", -- string
  success = true, -- boolean
})
```


### EmailSendingLimitsResponseSingle

Create an instance: `local email_sending_limits_response_single = client:EmailSendingLimitsResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: Load

```lua
local email_sending_limits_response_single, err = client:EmailSendingLimitsResponseSingle():load({ account_id = "account_id" })
```


### EmailSendingReputationResponseSingle

Create an instance: `local email_sending_reputation_response_single = client:EmailSendingReputationResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: Load

```lua
local email_sending_reputation_response_single, err = client:EmailSendingReputationResponseSingle():load({ account_id = "account_id" })
```


### EmailSendingSubdomainDnsStatus

Create an instance: `local email_sending_subdomain_dns_status = client:EmailSendingSubdomainDnsStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: Load

```lua
local email_sending_subdomain_dns_status, err = client:EmailSendingSubdomainDnsStatus():load({ subdomain_id = "subdomain_id", zone_id = "zone_id" })
```

#### Example: Create

```lua
local email_sending_subdomain_dns_status, err = client:EmailSendingSubdomainDnsStatus():create({
  subdomain_id = "example_subdomain_id", -- string
  zone_id = "example_zone_id", -- string
})
```


### EmailSendingSubdomainPreview

Create an instance: `local email_sending_subdomain_preview = client:EmailSendingSubdomainPreview(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `string` | The domain name within the zone. |

#### Example: Create

```lua
local email_sending_subdomain_preview, err = client:EmailSendingSubdomainPreview():create({
  zone_id = "example_zone_id", -- string
  name = "example_name", -- string
})
```


### EmailSendingSubdomainReputationComplaint

Create an instance: `local email_sending_subdomain_reputation_complaint = client:EmailSendingSubdomainReputationComplaint(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: Load

```lua
local email_sending_subdomain_reputation_complaint, err = client:EmailSendingSubdomainReputationComplaint():load({ subdomain_id = "subdomain_id", zone_id = "zone_id", end_at = "end_at", start_at = "start_at" })
```


### EmailSendingSubdomainResponseSingle

Create an instance: `local email_sending_subdomain_response_single = client:EmailSendingSubdomainResponseSingle(nil)`

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

```lua
local email_sending_subdomain_response_single, err = client:EmailSendingSubdomainResponseSingle():load({ subdomain_id = "subdomain_id", zone_id = "zone_id" })
```

#### Example: Create

```lua
local email_sending_subdomain_response_single, err = client:EmailSendingSubdomainResponseSingle():create({
  zone_id = "example_zone_id", -- string
  name = "example_name", -- string
})
```


### EmailSendingSubdomainsResponseCollection

Create an instance: `local email_sending_subdomains_response_collection = client:EmailSendingSubdomainsResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `table` |  |

#### Example: List

```lua
local email_sending_subdomains_response_collections, err = client:EmailSendingSubdomainsResponseCollection():list()
```


### EmailSendingSuppression

Create an instance: `local email_sending_suppression = client:EmailSendingSuppression(nil)`

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
| `errors` | `table` |  |
| `expires_at` | `string` | Expiration timestamp for the suppression. |
| `id` | `string` | Unique identifier for this suppression. |
| `items` | `table` | Suppressions to import. |
| `messages` | `table` |  |
| `note` | `string` | Advisory note for this suppression. |
| `read_only` | `boolean` | Whether clients may mutate this suppression. |
| `reason` | `string` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `table` |  |
| `scope` | `any` | Where the suppression applies. |
| `success` | `boolean` |  |
| `zones` | `table` |  |

#### Example: Load

```lua
local email_sending_suppression, err = client:EmailSendingSuppression():load({ id = "email_sending_suppression_id" })
```

#### Example: List

```lua
local email_sending_suppressions, err = client:EmailSendingSuppression():list()
```

#### Example: Create

```lua
local email_sending_suppression, err = client:EmailSendingSuppression():create({
  account_id = "example_account_id", -- string
  created_at = "example_created_at", -- string
  email = "example_email", -- string
  errors = {}, -- table
  expires_at = "example_expires_at", -- string
  id = "example_id", -- string
  items = {}, -- table
  messages = {}, -- table
  read_only = true, -- boolean
  reason = "example_reason", -- string
  result = {}, -- table
  success = true, -- boolean
})
```


### Subdomain

Create an instance: `local subdomain = client:Subdomain(nil)`

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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── cloudflare-email-security_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── schema.lua               -- Generated option + entity specs
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`cloudflare-email-security_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```lua
local emaildestinationaddressresponsesingle = client:EmailDestinationAddressResponseSingle()
emaildestinationaddressresponsesingle:load({ account_id = "example", destination_address_identifier = "example" })

-- emaildestinationaddressresponsesingle:data_get() now returns the emaildestinationaddressresponsesingle data from the last load
-- emaildestinationaddressresponsesingle:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
