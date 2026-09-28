# CloudflareEmailSecurity PHP SDK



The PHP SDK for the CloudflareEmailSecurity API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->EmailAccountRulesPlanResponseSingle()` — with named operations (`list`/`load`/`create`/`update`/`remove`/`patch`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'cloudflareemailsecurity_sdk.php';

$client = new CloudflareEmailSecuritySDK([
    "apikey" => getenv("CLOUDFLARE_EMAIL_SECURITY_APIKEY"),
]);
```

### 3. Load an emailsendingsubdomaindnsstatus

EmailSendingSubdomainDnsStatus is nested under subdomain, so provide the `subdomain_id`.

```php
try {
    // load() returns the ENTITY — call data_get() for the EmailSendingSubdomainDnsStatus record (throws on error).
    $emailsendingsubdomaindnsstatus = $client->EmailSendingSubdomainDnsStatus()->load(["subdomain_id" => "example_subdomain_id", "zone_id" => "example_zone_id"]);
    print_r($emailsendingsubdomaindnsstatus->data_get());
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created EmailAccountRulesPlanResponseSingle record.
$created = $client->EmailAccountRulesPlanResponseSingle()->create(["account_id" => "example_account_id", "owner_worker_tag" => "example_owner_worker_tag"]);

```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $emaildestinationaddressresponsesingle = $client->EmailDestinationAddressResponseSingle()->load(["account_id" => "example", "destination_address_identifier" => "example"]);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = CloudflareEmailSecuritySDK::test([
    "entity" => ["emailroutingsuppression" => ["test01" => ["id" => "test01"]]],
]);

// list() returns entity instances (throws on error);
// call data_get() for the mock record.
$emailroutingsuppression = $client->EmailRoutingSuppression()->list();
print_r(array_map(fn($item) => $item->data_get(), $emailroutingsuppression));
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new CloudflareEmailSecuritySDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE
CLOUDFLARE_EMAIL_SECURITY_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### CloudflareEmailSecuritySDK

```php
require_once 'cloudflareemailsecurity_sdk.php';
$client = new CloudflareEmailSecuritySDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = CloudflareEmailSecuritySDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### CloudflareEmailSecuritySDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `EmailAccountRulesPlanResponseSingle` | `($data): EmailAccountRulesPlanResponseSingleEntity` | Create an EmailAccountRulesPlanResponseSingle entity instance. |
| `EmailAccountRulesResponseCollection` | `($data): EmailAccountRulesResponseCollectionEntity` | Create an EmailAccountRulesResponseCollection entity instance. |
| `EmailAuth` | `($data): EmailAuthEntity` | Create an EmailAuth entity instance. |
| `EmailCatchAllRuleResponseSingle` | `($data): EmailCatchAllRuleResponseSingleEntity` | Create an EmailCatchAllRuleResponseSingle entity instance. |
| `EmailDestinationAddressResponseSingle` | `($data): EmailDestinationAddressResponseSingleEntity` | Create an EmailDestinationAddressResponseSingle entity instance. |
| `EmailDestinationAddressesResponseCollection` | `($data): EmailDestinationAddressesResponseCollectionEntity` | Create an EmailDestinationAddressesResponseCollection entity instance. |
| `EmailDnsSettingsResponseCollection` | `($data): EmailDnsSettingsResponseCollectionEntity` | Create an EmailDnsSettingsResponseCollection entity instance. |
| `EmailEmailSettingsResponseSingle` | `($data): EmailEmailSettingsResponseSingleEntity` | Create an EmailEmailSettingsResponseSingle entity instance. |
| `EmailRoutingSuppression` | `($data): EmailRoutingSuppressionEntity` | Create an EmailRoutingSuppression entity instance. |
| `EmailRuleResponseSingle` | `($data): EmailRuleResponseSingleEntity` | Create an EmailRuleResponseSingle entity instance. |
| `EmailRulesResponseCollection` | `($data): EmailRulesResponseCollectionEntity` | Create an EmailRulesResponseCollection entity instance. |
| `EmailSecurity` | `($data): EmailSecurityEntity` | Create an EmailSecurity entity instance. |
| `EmailSecurityAnalytics` | `($data): EmailSecurityAnalyticsEntity` | Create an EmailSecurityAnalytics entity instance. |
| `EmailSecuritySetting` | `($data): EmailSecuritySettingEntity` | Create an EmailSecuritySetting entity instance. |
| `EmailSending` | `($data): EmailSendingEntity` | Create an EmailSending entity instance. |
| `EmailSendingLimitsResponseSingle` | `($data): EmailSendingLimitsResponseSingleEntity` | Create an EmailSendingLimitsResponseSingle entity instance. |
| `EmailSendingReputationResponseSingle` | `($data): EmailSendingReputationResponseSingleEntity` | Create an EmailSendingReputationResponseSingle entity instance. |
| `EmailSendingSubdomainDnsStatus` | `($data): EmailSendingSubdomainDnsStatusEntity` | Create an EmailSendingSubdomainDnsStatus entity instance. |
| `EmailSendingSubdomainPreview` | `($data): EmailSendingSubdomainPreviewEntity` | Create an EmailSendingSubdomainPreview entity instance. |
| `EmailSendingSubdomainReputationComplaint` | `($data): EmailSendingSubdomainReputationComplaintEntity` | Create an EmailSendingSubdomainReputationComplaint entity instance. |
| `EmailSendingSubdomainResponseSingle` | `($data): EmailSendingSubdomainResponseSingleEntity` | Create an EmailSendingSubdomainResponseSingle entity instance. |
| `EmailSendingSubdomainsResponseCollection` | `($data): EmailSendingSubdomainsResponseCollectionEntity` | Create an EmailSendingSubdomainsResponseCollection entity instance. |
| `EmailSendingSuppression` | `($data): EmailSendingSuppressionEntity` | Create an EmailSendingSuppression entity instance. |
| `Subdomain` | `($data): SubdomainEntity` | Create a Subdomain entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$email_account_rules_plan_response_single = $client->EmailAccountRulesPlanResponseSingle();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `array` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `array` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

```php
$email_account_rules_plan_response_single = $client->EmailAccountRulesPlanResponseSingle()->create([
    "account_id" => null, // string
    "owner_worker_tag" => null, // string
]);
```


### EmailAccountRulesResponseCollection

Create an instance: `$email_account_rules_response_collection = $client->EmailAccountRulesResponseCollection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |
| `result_info` | `array` |  |

#### Example: List

```php
// list() returns an array of EmailAccountRulesResponseCollection records (throws on error).
$email_account_rules_response_collections = $client->EmailAccountRulesResponseCollection()->list();
```


### EmailAuth

Create an instance: `$email_auth = $client->EmailAuth();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `bool` | Enable or disable DMARC reports for this zone |
| `result` | `array` | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | Skip the DMARC setup wizard |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailAuth record (throws on error).
$email_auth = $client->EmailAuth()->load(["id" => "email_auth_id", "zone_id" => "zone_id"]);
```


### EmailCatchAllRuleResponseSingle

Create an instance: `$email_catch_all_rule_response_single = $client->EmailCatchAllRuleResponseSingle();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `array` | List actions for the catch-all routing rule. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `array` | List of matchers for the catch-all routing rule. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `array` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailCatchAllRuleResponseSingle record (throws on error).
$email_catch_all_rule_response_single = $client->EmailCatchAllRuleResponseSingle()->load(["zone_id" => "zone_id"]);
```


### EmailDestinationAddressResponseSingle

Create an instance: `$email_destination_address_response_single = $client->EmailDestinationAddressResponseSingle();`

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
| `result` | `array` |  |
| `status` | `string` | Destination address status. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailDestinationAddressResponseSingle record (throws on error).
$email_destination_address_response_single = $client->EmailDestinationAddressResponseSingle()->load(["account_id" => "account_id", "destination_address_identifier" => "destination_address_identifier"]);
```

#### Example: Create

```php
$email_destination_address_response_single = $client->EmailDestinationAddressResponseSingle()->create([
    "account_id" => null, // string
    "email" => null, // string
    "status" => null, // string
]);
```


### EmailDestinationAddressesResponseCollection

Create an instance: `$email_destination_addresses_response_collection = $client->EmailDestinationAddressesResponseCollection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |
| `result_info` | `array` |  |

#### Example: List

```php
// list() returns an array of EmailDestinationAddressesResponseCollection records (throws on error).
$email_destination_addresses_response_collections = $client->EmailDestinationAddressesResponseCollection()->list();
```


### EmailDnsSettingsResponseCollection

Create an instance: `$email_dns_settings_response_collection = $client->EmailDnsSettingsResponseCollection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: List

```php
// list() returns an array of EmailDnsSettingsResponseCollection records (throws on error).
$email_dns_settings_response_collections = $client->EmailDnsSettingsResponseCollection()->list();
```


### EmailEmailSettingsResponseSingle

Create an instance: `$email_email_settings_response_single = $client->EmailEmailSettingsResponseSingle();`

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
| `enabled` | `bool` | State of your zone Email Routing settings. |
| `name` | `string` | Domain of your zone. |
| `result` | `array` |  |
| `skip_wizard` | `bool` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailEmailSettingsResponseSingle record (throws on error).
$email_email_settings_response_single = $client->EmailEmailSettingsResponseSingle()->load(["zone_id" => "zone_id"]);
```

#### Example: Create

```php
$email_email_settings_response_single = $client->EmailEmailSettingsResponseSingle()->create([
    "zone_id" => null, // string
]);
```


### EmailRoutingSuppression

Create an instance: `$email_routing_suppression = $client->EmailRoutingSuppression();`

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
| `zones` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailRoutingSuppression record (throws on error).
$email_routing_suppression = $client->EmailRoutingSuppression()->load(["id" => "email_routing_suppression_id"]);
```

#### Example: List

```php
// list() returns an array of EmailRoutingSuppression records (throws on error).
$email_routing_suppressions = $client->EmailRoutingSuppression()->list();
```

#### Example: Create

```php
$email_routing_suppression = $client->EmailRoutingSuppression()->create([
    "account_id" => null, // string
    "created_at" => null, // string
    "email" => null, // string
    "expires_at" => null, // string
    "id" => null, // string
    "reason" => null, // string
]);
```


### EmailRuleResponseSingle

Create an instance: `$email_rule_response_single = $client->EmailRuleResponseSingle();`

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
| `actions` | `array` | List actions patterns. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `array` | Matching patterns to forward to your actions. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float` | Priority of the routing rule. |
| `result` | `array` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailRuleResponseSingle record (throws on error).
$email_rule_response_single = $client->EmailRuleResponseSingle()->load(["rule_identifier" => "rule_identifier", "zone_id" => "zone_id"]);
```

#### Example: Create

```php
$email_rule_response_single = $client->EmailRuleResponseSingle()->create([
    "zone_id" => null, // string
    "actions" => null, // array
    "matchers" => null, // array
]);
```


### EmailRulesResponseCollection

Create an instance: `$email_rules_response_collection = $client->EmailRulesResponseCollection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |
| `result_info` | `array` |  |

#### Example: List

```php
// list() returns an array of EmailRulesResponseCollection records (throws on error).
$email_rules_response_collections = $client->EmailRulesResponseCollection()->list();
```


### EmailSecurity

Create an instance: `$email_security = $client->EmailSecurity();`

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
| `destination` | `mixed` | Required when action is 'MOVE'. |
| `errors` | `array` |  |
| `expected_disposition` | `string` | Nonfunctional field. |
| `id` | `string` | Unique identifier for a message retrieved from investigation. |
| `ids` | `array` | List of message IDs to move. |
| `messages` | `array` |  |
| `postfix_ids` | `array` | Deprecated, use `ids` instead. |
| `result` | `array` |  |
| `result_info` | `array` |  |
| `search_params` | `array` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSecurity record (throws on error).
$email_security = $client->EmailSecurity()->load(["account_id" => "account_id"]);
```

#### Example: List

```php
// list() returns an array of EmailSecurity records (throws on error).
$email_securitys = $client->EmailSecurity()->list();
```

#### Example: Create

```php
$email_security = $client->EmailSecurity()->create([
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


### EmailSecurityAnalytics

Create an instance: `$email_security_analytics = $client->EmailSecurityAnalytics();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSecurityAnalytics record (throws on error).
$email_security_analytics = $client->EmailSecurityAnalytics()->load(["account_id" => "account_id"]);
```


### EmailSecuritySetting

Create an instance: `$email_security_setting = $client->EmailSecuritySetting();`

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
| `allowed_delivery_modes` | `array` | Delivery modes to onboard the domain through. |
| `comments` | `string` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `array` | IDs of the allow policies to delete. |
| `domain` | `string` | The email domain to protect. |
| `drop_dispositions` | `array` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `mixed` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `array` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `array` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `array` | Allow policies to create. |
| `puts` | `array` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `array` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | Require TLS on outbound connections. |
| `result` | `array` | An impersonation registry entry. |
| `transport` | `string` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSecuritySetting record (throws on error).
$email_security_setting = $client->EmailSecuritySetting()->load(["account_id" => "account_id"]);
```

#### Example: List

```php
// list() returns an array of EmailSecuritySetting records (throws on error).
$email_security_settings = $client->EmailSecuritySetting()->list();
```

#### Example: Create

```php
$email_security_setting = $client->EmailSecuritySetting()->create([
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


### EmailSending

Create an instance: `$email_sending = $client->EmailSending();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `array` | File attachments and inline images. |
| `bcc` | `mixed` | Blind carbon copy recipient(s). |
| `cc` | `mixed` | Carbon copy recipient(s). |
| `errors` | `array` |  |
| `from` | `mixed` | Sender email address. |
| `headers` | `array` | Custom email headers as key-value pairs. |
| `html` | `string` | HTML body of the email. |
| `messages` | `array` |  |
| `mime_message` | `string` | The full MIME-encoded email message. |
| `recipients` | `array` | List of recipient email addresses. |
| `reply_to` | `mixed` | Reply-to address. |
| `result` | `array` |  |
| `result_info` | `array` |  |
| `subject` | `string` | Email subject line. |
| `success` | `bool` |  |
| `text` | `string` | Plain text body of the email. |
| `to` | `mixed` | Recipient(s). |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSending record (throws on error).
$email_sending = $client->EmailSending()->load(["account_id" => "account_id", "message_id" => "message_id"]);
```

#### Example: Create

```php
$email_sending = $client->EmailSending()->create([
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


### EmailSendingLimitsResponseSingle

Create an instance: `$email_sending_limits_response_single = $client->EmailSendingLimitsResponseSingle();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingLimitsResponseSingle record (throws on error).
$email_sending_limits_response_single = $client->EmailSendingLimitsResponseSingle()->load(["account_id" => "account_id"]);
```


### EmailSendingReputationResponseSingle

Create an instance: `$email_sending_reputation_response_single = $client->EmailSendingReputationResponseSingle();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingReputationResponseSingle record (throws on error).
$email_sending_reputation_response_single = $client->EmailSendingReputationResponseSingle()->load(["account_id" => "account_id"]);
```


### EmailSendingSubdomainDnsStatus

Create an instance: `$email_sending_subdomain_dns_status = $client->EmailSendingSubdomainDnsStatus();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingSubdomainDnsStatus record (throws on error).
$email_sending_subdomain_dns_status = $client->EmailSendingSubdomainDnsStatus()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id"]);
```

#### Example: Create

```php
$email_sending_subdomain_dns_status = $client->EmailSendingSubdomainDnsStatus()->create([
    "subdomain_id" => null, // string
    "zone_id" => null, // string
]);
```


### EmailSendingSubdomainPreview

Create an instance: `$email_sending_subdomain_preview = $client->EmailSendingSubdomainPreview();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `string` | The domain name within the zone. |

#### Example: Create

```php
$email_sending_subdomain_preview = $client->EmailSendingSubdomainPreview()->create([
    "zone_id" => null, // string
    "name" => null, // string
]);
```


### EmailSendingSubdomainReputationComplaint

Create an instance: `$email_sending_subdomain_reputation_complaint = $client->EmailSendingSubdomainReputationComplaint();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingSubdomainReputationComplaint record (throws on error).
$email_sending_subdomain_reputation_complaint = $client->EmailSendingSubdomainReputationComplaint()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id", "end_at" => "end_at", "start_at" => "start_at"]);
```


### EmailSendingSubdomainResponseSingle

Create an instance: `$email_sending_subdomain_response_single = $client->EmailSendingSubdomainResponseSingle();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `drop_suppressed_recipients` | `bool` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | The domain name within the zone. |
| `preview_enabled` | `bool` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `mixed` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingSubdomainResponseSingle record (throws on error).
$email_sending_subdomain_response_single = $client->EmailSendingSubdomainResponseSingle()->load(["subdomain_id" => "subdomain_id", "zone_id" => "zone_id"]);
```

#### Example: Create

```php
$email_sending_subdomain_response_single = $client->EmailSendingSubdomainResponseSingle()->create([
    "zone_id" => null, // string
    "name" => null, // string
]);
```


### EmailSendingSubdomainsResponseCollection

Create an instance: `$email_sending_subdomains_response_collection = $client->EmailSendingSubdomainsResponseCollection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `array` |  |

#### Example: List

```php
// list() returns an array of EmailSendingSubdomainsResponseCollection records (throws on error).
$email_sending_subdomains_response_collections = $client->EmailSendingSubdomainsResponseCollection()->list();
```


### EmailSendingSuppression

Create an instance: `$email_sending_suppression = $client->EmailSendingSuppression();`

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
| `errors` | `array` |  |
| `expires_at` | `string` | Expiration timestamp for the suppression. |
| `id` | `string` | Unique identifier for this suppression. |
| `items` | `array` | Suppressions to import. |
| `messages` | `array` |  |
| `note` | `string` | Advisory note for this suppression. |
| `read_only` | `bool` | Whether clients may mutate this suppression. |
| `reason` | `string` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `array` |  |
| `scope` | `mixed` | Where the suppression applies. |
| `success` | `bool` |  |
| `zones` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EmailSendingSuppression record (throws on error).
$email_sending_suppression = $client->EmailSendingSuppression()->load(["id" => "email_sending_suppression_id"]);
```

#### Example: List

```php
// list() returns an array of EmailSendingSuppression records (throws on error).
$email_sending_suppressions = $client->EmailSendingSuppression()->list();
```

#### Example: Create

```php
$email_sending_suppression = $client->EmailSendingSuppression()->create([
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


### Subdomain

Create an instance: `$subdomain = $client->Subdomain();`

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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── cloudflareemailsecurity_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── schema.php                     -- Generated option + entity specs
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`cloudflareemailsecurity_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$emaildestinationaddressresponsesingle = $client->EmailDestinationAddressResponseSingle();
$emaildestinationaddressresponsesingle->load(["account_id" => "example", "destination_address_identifier" => "example"]);

// $emaildestinationaddressresponsesingle->data_get() now returns the emaildestinationaddressresponsesingle data from the last load
// $emaildestinationaddressresponsesingle->match_get() returns the last match criteria
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
