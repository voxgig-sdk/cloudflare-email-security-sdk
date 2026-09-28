# CloudflareEmailSecurity Ruby SDK



The Ruby SDK for the CloudflareEmailSecurity API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.EmailAccountRulesPlanResponseSingle` — with named operations (`list`/`load`/`create`/`update`/`remove`/`patch`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "CloudflareEmailSecurity_sdk"

client = CloudflareEmailSecuritySDK.new({
  "apikey" => ENV["CLOUDFLARE_EMAIL_SECURITY_APIKEY"],
})
```

### 3. Load an emailsendingsubdomaindnsstatus

EmailSendingSubdomainDnsStatus is nested under subdomain, so provide the `subdomain_id`.

```ruby
begin
  # load returns the ENTITY — call data_get for the EmailSendingSubdomainDnsStatus record (raises on error).
  emailsendingsubdomaindnsstatus = client.EmailSendingSubdomainDnsStatus.load({ "subdomain_id" => "example_subdomain_id", "zone_id" => "example_zone_id" })
  puts emailsendingsubdomaindnsstatus
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the ENTITY — call data_get for the created EmailAccountRulesPlanResponseSingle record.
created = client.EmailAccountRulesPlanResponseSingle.create({ "account_id" => "example_account_id", "owner_worker_tag" => "example_owner_worker_tag" })

```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle.load({ "account_id" => "example", "destination_address_identifier" => "example" })
rescue => err
  warn "load failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = CloudflareEmailSecuritySDK.test

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle.load({ "account_id" => "example", "destination_address_identifier" => "example" })
puts emaildestinationaddressresponsesingle
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = CloudflareEmailSecuritySDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### CloudflareEmailSecuritySDK

```ruby
require_relative "CloudflareEmailSecurity_sdk"
client = CloudflareEmailSecuritySDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = CloudflareEmailSecuritySDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CloudflareEmailSecuritySDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `CloudflareEmailSecurityError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Create an instance: `email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `Array` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `String` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `Array` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

```ruby
email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle.create({
  "account_id" => "example_account_id", # String
  "owner_worker_tag" => "example_owner_worker_tag", # String
})
```


### EmailAccountRulesResponseCollection

Create an instance: `email_account_rules_response_collection = client.EmailAccountRulesResponseCollection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Array` |  |
| `result_info` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of EmailAccountRulesResponseCollection records (raises on error).
email_account_rules_response_collections = client.EmailAccountRulesResponseCollection.list
```


### EmailAuth

Create an instance: `email_auth = client.EmailAuth`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `Boolean` | Enable or disable DMARC reports for this zone |
| `result` | `Hash` | Recursive SPF inspection tree |
| `skip_wizard` | `Boolean` | Skip the DMARC setup wizard |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailAuth record (raises on error).
email_auth = client.EmailAuth.load({ "id" => "email_auth_id", "zone_id" => "zone_id" })
```


### EmailCatchAllRuleResponseSingle

Create an instance: `email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `Array` | List actions for the catch-all routing rule. |
| `enabled` | `Boolean` | Routing rule status. |
| `matchers` | `Array` | List of matchers for the catch-all routing rule. |
| `name` | `String` | Routing rule name. |
| `owner_worker_tag` | `String` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `Hash` |  |
| `source` | `String` | Who manages the rule. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailCatchAllRuleResponseSingle record (raises on error).
email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle.load({ "zone_id" => "zone_id" })
```


### EmailDestinationAddressResponseSingle

Create an instance: `email_destination_address_response_single = client.EmailDestinationAddressResponseSingle`

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
| `email` | `String` | The contact email address of the user. |
| `result` | `Hash` |  |
| `status` | `String` | Destination address status. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailDestinationAddressResponseSingle record (raises on error).
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle.load({ "account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier" })
```

#### Example: Create

```ruby
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle.create({
  "account_id" => "example_account_id", # String
  "email" => "example_email", # String
  "status" => "example_status", # String
})
```


### EmailDestinationAddressesResponseCollection

Create an instance: `email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Array` |  |
| `result_info` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of EmailDestinationAddressesResponseCollection records (raises on error).
email_destination_addresses_response_collections = client.EmailDestinationAddressesResponseCollection.list
```


### EmailDnsSettingsResponseCollection

Create an instance: `email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Array` |  |

#### Example: List

```ruby
# list returns an Array of EmailDnsSettingsResponseCollection records (raises on error).
email_dns_settings_response_collections = client.EmailDnsSettingsResponseCollection.list
```


### EmailEmailSettingsResponseSingle

Create an instance: `email_email_settings_response_single = client.EmailEmailSettingsResponseSingle`

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
| `enabled` | `Boolean` | State of your zone Email Routing settings. |
| `name` | `String` | Domain of your zone. |
| `result` | `Hash` |  |
| `skip_wizard` | `Boolean` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `Boolean` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailEmailSettingsResponseSingle record (raises on error).
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle.load({ "zone_id" => "zone_id" })
```

#### Example: Create

```ruby
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle.create({
  "zone_id" => "example_zone_id", # String
})
```


### EmailRoutingSuppression

Create an instance: `email_routing_suppression = client.EmailRoutingSuppression`

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
| `created_at` | `String` |  |
| `email` | `String` |  |
| `expires_at` | `String` |  |
| `id` | `String` |  |
| `reason` | `String` |  |
| `zones` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailRoutingSuppression record (raises on error).
email_routing_suppression = client.EmailRoutingSuppression.load({ "id" => "email_routing_suppression_id" })
```

#### Example: List

```ruby
# list returns an Array of EmailRoutingSuppression records (raises on error).
email_routing_suppressions = client.EmailRoutingSuppression.list
```

#### Example: Create

```ruby
email_routing_suppression = client.EmailRoutingSuppression.create({
  "account_id" => "example_account_id", # String
  "created_at" => "example_created_at", # String
  "email" => "example_email", # String
  "expires_at" => "example_expires_at", # String
  "id" => "example_id", # String
  "reason" => "example_reason", # String
})
```


### EmailRuleResponseSingle

Create an instance: `email_rule_response_single = client.EmailRuleResponseSingle`

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
| `actions` | `Array` | List actions patterns. |
| `enabled` | `Boolean` | Routing rule status. |
| `matchers` | `Array` | Matching patterns to forward to your actions. |
| `name` | `String` | Routing rule name. |
| `owner_worker_tag` | `String` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `Float` | Priority of the routing rule. |
| `result` | `Hash` |  |
| `source` | `String` | Who manages the rule. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailRuleResponseSingle record (raises on error).
email_rule_response_single = client.EmailRuleResponseSingle.load({ "rule_identifier" => "rule_identifier", "zone_id" => "zone_id" })
```

#### Example: Create

```ruby
email_rule_response_single = client.EmailRuleResponseSingle.create({
  "zone_id" => "example_zone_id", # String
  "actions" => [], # Array
  "matchers" => [], # Array
})
```


### EmailRulesResponseCollection

Create an instance: `email_rules_response_collection = client.EmailRulesResponseCollection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Array` |  |
| `result_info` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of EmailRulesResponseCollection records (raises on error).
email_rules_response_collections = client.EmailRulesResponseCollection.list
```


### EmailSecurity

Create an instance: `email_security = client.EmailSecurity`

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
| `action` | `String` | The action the job performs on every message matching the search parameters. |
| `comment` | `String` | Optional note describing the job. |
| `destination` | `Object` | Required when action is 'MOVE'. |
| `errors` | `Array` |  |
| `expected_disposition` | `String` | Nonfunctional field. |
| `id` | `String` | Unique identifier for a message retrieved from investigation. |
| `ids` | `Array` | List of message IDs to move. |
| `messages` | `Array` |  |
| `postfix_ids` | `Array` | Deprecated, use `ids` instead. |
| `result` | `Hash` |  |
| `result_info` | `Hash` |  |
| `search_params` | `Hash` |  |
| `success` | `Boolean` | Whether the API call was successful. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSecurity record (raises on error).
email_security = client.EmailSecurity.load({ "account_id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of EmailSecurity records (raises on error).
email_securitys = client.EmailSecurity.list
```

#### Example: Create

```ruby
email_security = client.EmailSecurity.create({
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


### EmailSecurityAnalytics

Create an instance: `email_security_analytics = client.EmailSecurityAnalytics`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSecurityAnalytics record (raises on error).
email_security_analytics = client.EmailSecurityAnalytics.load({ "account_id" => "account_id" })
```


### EmailSecuritySetting

Create an instance: `email_security_setting = client.EmailSecuritySetting`

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
| `allowed_delivery_modes` | `Array` | Delivery modes to onboard the domain through. |
| `comments` | `String` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `Array` | IDs of the allow policies to delete. |
| `domain` | `String` | The email domain to protect. |
| `drop_dispositions` | `Array` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `Object` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `String` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `Array` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `Integer` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `Array` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `String` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `Array` | Allow policies to create. |
| `puts` | `Array` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `Array` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `Boolean` | Require TLS on inbound connections. |
| `require_tls_outbound` | `Boolean` | Require TLS on outbound connections. |
| `result` | `Hash` | An impersonation registry entry. |
| `transport` | `String` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSecuritySetting record (raises on error).
email_security_setting = client.EmailSecuritySetting.load({ "account_id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of EmailSecuritySetting records (raises on error).
email_security_settings = client.EmailSecuritySetting.list
```

#### Example: Create

```ruby
email_security_setting = client.EmailSecuritySetting.create({
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


### EmailSending

Create an instance: `email_sending = client.EmailSending`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `Array` | File attachments and inline images. |
| `bcc` | `Object` | Blind carbon copy recipient(s). |
| `cc` | `Object` | Carbon copy recipient(s). |
| `errors` | `Array` |  |
| `from` | `Object` | Sender email address. |
| `headers` | `Hash` | Custom email headers as key-value pairs. |
| `html` | `String` | HTML body of the email. |
| `messages` | `Array` |  |
| `mime_message` | `String` | The full MIME-encoded email message. |
| `recipients` | `Array` | List of recipient email addresses. |
| `reply_to` | `Object` | Reply-to address. |
| `result` | `Hash` |  |
| `result_info` | `Hash` |  |
| `subject` | `String` | Email subject line. |
| `success` | `Boolean` |  |
| `text` | `String` | Plain text body of the email. |
| `to` | `Object` | Recipient(s). |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSending record (raises on error).
email_sending = client.EmailSending.load({ "account_id" => "account_id", "message_id" => "message_id" })
```

#### Example: Create

```ruby
email_sending = client.EmailSending.create({
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


### EmailSendingLimitsResponseSingle

Create an instance: `email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingLimitsResponseSingle record (raises on error).
email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle.load({ "account_id" => "account_id" })
```


### EmailSendingReputationResponseSingle

Create an instance: `email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingReputationResponseSingle record (raises on error).
email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle.load({ "account_id" => "account_id" })
```


### EmailSendingSubdomainDnsStatus

Create an instance: `email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingSubdomainDnsStatus record (raises on error).
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id" })
```

#### Example: Create

```ruby
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus.create({
  "subdomain_id" => "example_subdomain_id", # String
  "zone_id" => "example_zone_id", # String
})
```


### EmailSendingSubdomainPreview

Create an instance: `email_sending_subdomain_preview = client.EmailSendingSubdomainPreview`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `String` | The domain name within the zone. |

#### Example: Create

```ruby
email_sending_subdomain_preview = client.EmailSendingSubdomainPreview.create({
  "zone_id" => "example_zone_id", # String
  "name" => "example_name", # String
})
```


### EmailSendingSubdomainReputationComplaint

Create an instance: `email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingSubdomainReputationComplaint record (raises on error).
email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id", "end_at" => "end_at", "start_at" => "start_at" })
```


### EmailSendingSubdomainResponseSingle

Create an instance: `email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `drop_suppressed_recipients` | `Boolean` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `String` | The domain name within the zone. |
| `preview_enabled` | `Boolean` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `Object` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingSubdomainResponseSingle record (raises on error).
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle.load({ "subdomain_id" => "subdomain_id", "zone_id" => "zone_id" })
```

#### Example: Create

```ruby
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle.create({
  "zone_id" => "example_zone_id", # String
  "name" => "example_name", # String
})
```


### EmailSendingSubdomainsResponseCollection

Create an instance: `email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `Array` |  |

#### Example: List

```ruby
# list returns an Array of EmailSendingSubdomainsResponseCollection records (raises on error).
email_sending_subdomains_response_collections = client.EmailSendingSubdomainsResponseCollection.list
```


### EmailSendingSuppression

Create an instance: `email_sending_suppression = client.EmailSendingSuppression`

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
| `created_at` | `String` | When the suppression was created. |
| `email` | `String` | The email address to suppress. |
| `errors` | `Array` |  |
| `expires_at` | `String` | Expiration timestamp for the suppression. |
| `id` | `String` | Unique identifier for this suppression. |
| `items` | `Array` | Suppressions to import. |
| `messages` | `Array` |  |
| `note` | `String` | Advisory note for this suppression. |
| `read_only` | `Boolean` | Whether clients may mutate this suppression. |
| `reason` | `String` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `Hash` |  |
| `scope` | `Object` | Where the suppression applies. |
| `success` | `Boolean` |  |
| `zones` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EmailSendingSuppression record (raises on error).
email_sending_suppression = client.EmailSendingSuppression.load({ "id" => "email_sending_suppression_id" })
```

#### Example: List

```ruby
# list returns an Array of EmailSendingSuppression records (raises on error).
email_sending_suppressions = client.EmailSendingSuppression.list
```

#### Example: Create

```ruby
email_sending_suppression = client.EmailSendingSuppression.create({
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


### Subdomain

Create an instance: `subdomain = client.Subdomain`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `String` |  |

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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── CloudflareEmailSecurity_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── schema.rb                  -- Generated option + entity specs
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`CloudflareEmailSecurity_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle
emaildestinationaddressresponsesingle.load({ "account_id" => "example", "destination_address_identifier" => "example" })

# emaildestinationaddressresponsesingle.data_get now returns the emaildestinationaddressresponsesingle data from the last load
# emaildestinationaddressresponsesingle.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
