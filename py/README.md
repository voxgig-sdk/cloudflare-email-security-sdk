# CloudflareEmailSecurity Python SDK



The Python SDK for the CloudflareEmailSecurity API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.EmailAccountRulesPlanResponseSingle()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from cloudflareemailsecurity_sdk import CloudflareEmailSecuritySDK

client = CloudflareEmailSecuritySDK({
    "apikey": os.environ.get("CLOUDFLARE_EMAIL_SECURITY_APIKEY"),
})
```

### 3. Load an emailsendingsubdomaindnsstatus

EmailSendingSubdomainDnsStatus is nested under subdomain, so provide the `subdomain_id`.
`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    emailsendingsubdomaindnsstatus = client.EmailSendingSubdomainDnsStatus().load({"subdomain_id": "example_subdomain_id", "zone_id": "example_zone_id"})
    print(emailsendingsubdomaindnsstatus)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.EmailAccountRulesPlanResponseSingle().create({"account_id": "example_account_id", "owner_worker_tag": "example_owner_worker_tag"})

```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle().load({"account_id": "example", "destination_address_identifier": "example"})
    print(emaildestinationaddressresponsesingle)
except Exception as err:
    print(f"load failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = CloudflareEmailSecuritySDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle().load({"account_id": "example", "destination_address_identifier": "example"})
# emaildestinationaddressresponsesingle contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = CloudflareEmailSecuritySDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### CloudflareEmailSecuritySDK

```python
from cloudflareemailsecurity_sdk import CloudflareEmailSecuritySDK

client = CloudflareEmailSecuritySDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = CloudflareEmailSecuritySDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### CloudflareEmailSecuritySDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `list` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `str` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `list` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

```python
email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle().create({
    "account_id": "example_account_id",  # str
    "owner_worker_tag": "example_owner_worker_tag",  # str
})
```


### EmailAccountRulesResponseCollection

Create an instance: `email_account_rules_response_collection = client.EmailAccountRulesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `list` |  |
| `result_info` | `dict` |  |

#### Example: List

```python
email_account_rules_response_collections = client.EmailAccountRulesResponseCollection().list({"account_id": "example"})
```


### EmailAuth

Create an instance: `email_auth = client.EmailAuth()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `bool` | Enable or disable DMARC reports for this zone |
| `result` | `dict` | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | Skip the DMARC setup wizard |

#### Example: Load

```python
email_auth = client.EmailAuth().load({"id": "email_auth_id", "zone_id": "zone_id"})
```


### EmailCatchAllRuleResponseSingle

Create an instance: `email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `list` | List actions for the catch-all routing rule. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `list` | List of matchers for the catch-all routing rule. |
| `name` | `str` | Routing rule name. |
| `owner_worker_tag` | `str` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `dict` |  |
| `source` | `str` | Who manages the rule. |

#### Example: Load

```python
email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle().load({"zone_id": "zone_id"})
```


### EmailDestinationAddressResponseSingle

Create an instance: `email_destination_address_response_single = client.EmailDestinationAddressResponseSingle()`

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
| `email` | `str` | The contact email address of the user. |
| `result` | `dict` |  |
| `status` | `str` | Destination address status. |

#### Example: Load

```python
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle().load({"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"})
```

#### Example: Create

```python
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle().create({
    "account_id": "example_account_id",  # str
    "email": "example_email",  # str
    "status": "example_status",  # str
})
```


### EmailDestinationAddressesResponseCollection

Create an instance: `email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `list` |  |
| `result_info` | `dict` |  |

#### Example: List

```python
email_destination_addresses_response_collections = client.EmailDestinationAddressesResponseCollection().list({"account_id": "example"})
```


### EmailDnsSettingsResponseCollection

Create an instance: `email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `list` |  |

#### Example: List

```python
email_dns_settings_response_collections = client.EmailDnsSettingsResponseCollection().list({"zone_id": "example"})
```


### EmailEmailSettingsResponseSingle

Create an instance: `email_email_settings_response_single = client.EmailEmailSettingsResponseSingle()`

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
| `name` | `str` | Domain of your zone. |
| `result` | `dict` |  |
| `skip_wizard` | `bool` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```python
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle().load({"zone_id": "zone_id"})
```

#### Example: Create

```python
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle().create({
    "zone_id": "example_zone_id",  # str
})
```


### EmailRoutingSuppression

Create an instance: `email_routing_suppression = client.EmailRoutingSuppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` |  |
| `email` | `str` |  |
| `expires_at` | `str` |  |
| `id` | `str` |  |
| `reason` | `str` |  |
| `zones` | `list` |  |

#### Example: Load

```python
email_routing_suppression = client.EmailRoutingSuppression().load({"id": "email_routing_suppression_id"})
```

#### Example: List

```python
email_routing_suppressions = client.EmailRoutingSuppression().list({"account_id": "example"})
```

#### Example: Create

```python
email_routing_suppression = client.EmailRoutingSuppression().create({
    "account_id": "example_account_id",  # str
    "created_at": "example_created_at",  # str
    "email": "example_email",  # str
    "expires_at": "example_expires_at",  # str
    "id": "example_id",  # str
    "reason": "example_reason",  # str
})
```


### EmailRuleResponseSingle

Create an instance: `email_rule_response_single = client.EmailRuleResponseSingle()`

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
| `actions` | `list` | List actions patterns. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `list` | Matching patterns to forward to your actions. |
| `name` | `str` | Routing rule name. |
| `owner_worker_tag` | `str` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float` | Priority of the routing rule. |
| `result` | `dict` |  |
| `source` | `str` | Who manages the rule. |

#### Example: Load

```python
email_rule_response_single = client.EmailRuleResponseSingle().load({"rule_identifier": "rule_identifier", "zone_id": "zone_id"})
```

#### Example: Create

```python
email_rule_response_single = client.EmailRuleResponseSingle().create({
    "zone_id": "example_zone_id",  # str
    "actions": [],  # list
    "matchers": [],  # list
})
```


### EmailRulesResponseCollection

Create an instance: `email_rules_response_collection = client.EmailRulesResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `list` |  |
| `result_info` | `dict` |  |

#### Example: List

```python
email_rules_response_collections = client.EmailRulesResponseCollection().list({"zone_id": "example"})
```


### EmailSecurity

Create an instance: `email_security = client.EmailSecurity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `action` | `str` | The action the job performs on every message matching the search parameters. |
| `comment` | `str` | Optional note describing the job. |
| `destination` | `Any` | Required when action is 'MOVE'. |
| `errors` | `list` |  |
| `expected_disposition` | `str` | Nonfunctional field. |
| `id` | `str` | Unique identifier for a message retrieved from investigation. |
| `ids` | `list` | List of message IDs to move. |
| `messages` | `list` |  |
| `postfix_ids` | `list` | Deprecated, use `ids` instead. |
| `result` | `dict` |  |
| `result_info` | `dict` |  |
| `search_params` | `dict` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```python
email_security = client.EmailSecurity().load({"account_id": "account_id"})
```

#### Example: List

```python
email_securitys = client.EmailSecurity().list({"account_id": "example"})
```

#### Example: Create

```python
email_security = client.EmailSecurity().create({
    "account_id": "example_account_id",  # str
    "action": "example_action",  # str
    "errors": [],  # list
    "id": "example_id",  # str
    "ids": [],  # list
    "messages": [],  # list
    "result": {},  # dict
    "result_info": {},  # dict
    "search_params": {},  # dict
    "success": True,  # bool
})
```


### EmailSecurityAnalytics

Create an instance: `email_security_analytics = client.EmailSecurityAnalytics()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `dict` |  |

#### Example: Load

```python
email_security_analytics = client.EmailSecurityAnalytics().load({"account_id": "account_id"})
```


### EmailSecuritySetting

Create an instance: `email_security_setting = client.EmailSecuritySetting()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `allowed_delivery_modes` | `list` | Delivery modes to onboard the domain through. |
| `comments` | `str` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `list` | IDs of the allow policies to delete. |
| `domain` | `str` | The email domain to protect. |
| `drop_dispositions` | `list` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `Any` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `str` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `list` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `list` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `str` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `list` | Allow policies to create. |
| `puts` | `list` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `list` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | Require TLS on outbound connections. |
| `result` | `dict` | An impersonation registry entry. |
| `transport` | `str` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```python
email_security_setting = client.EmailSecuritySetting().load({"account_id": "account_id"})
```

#### Example: List

```python
email_security_settings = client.EmailSecuritySetting().list({"account_id": "example"})
```

#### Example: Create

```python
email_security_setting = client.EmailSecuritySetting().create({
    "account_id": "example_account_id",  # str
    "allowed_delivery_modes": [],  # list
    "deletes": [],  # list
    "domain": "example_domain",  # str
    "drop_dispositions": [],  # list
    "ip_restrictions": [],  # list
    "patches": [],  # list
    "pattern": "example_pattern",  # str
    "posts": [],  # list
    "puts": [],  # list
    "regions": [],  # list
})
```


### EmailSending

Create an instance: `email_sending = client.EmailSending()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `list` | File attachments and inline images. |
| `bcc` | `Any` | Blind carbon copy recipient(s). |
| `cc` | `Any` | Carbon copy recipient(s). |
| `errors` | `list` |  |
| `from` | `Any` | Sender email address. |
| `headers` | `dict` | Custom email headers as key-value pairs. |
| `html` | `str` | HTML body of the email. |
| `messages` | `list` |  |
| `mime_message` | `str` | The full MIME-encoded email message. |
| `recipients` | `list` | List of recipient email addresses. |
| `reply_to` | `Any` | Reply-to address. |
| `result` | `dict` |  |
| `result_info` | `dict` |  |
| `subject` | `str` | Email subject line. |
| `success` | `bool` |  |
| `text` | `str` | Plain text body of the email. |
| `to` | `Any` | Recipient(s). |

#### Example: Load

```python
email_sending = client.EmailSending().load({"account_id": "account_id", "message_id": "message_id"})
```

#### Example: Create

```python
email_sending = client.EmailSending().create({
    "account_id": "example_account_id",  # str
    "errors": [],  # list
    "from": "example_from",  # Any
    "messages": [],  # list
    "mime_message": "example_mime_message",  # str
    "recipients": [],  # list
    "result": {},  # dict
    "result_info": {},  # dict
    "subject": "example_subject",  # str
    "success": True,  # bool
})
```


### EmailSendingLimitsResponseSingle

Create an instance: `email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `dict` |  |

#### Example: Load

```python
email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle().load({"account_id": "account_id"})
```


### EmailSendingReputationResponseSingle

Create an instance: `email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `dict` |  |

#### Example: Load

```python
email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle().load({"account_id": "account_id"})
```


### EmailSendingSubdomainDnsStatus

Create an instance: `email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `dict` |  |

#### Example: Load

```python
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id"})
```

#### Example: Create

```python
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus().create({
    "subdomain_id": "example_subdomain_id",  # str
    "zone_id": "example_zone_id",  # str
})
```


### EmailSendingSubdomainPreview

Create an instance: `email_sending_subdomain_preview = client.EmailSendingSubdomainPreview()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `str` | The domain name within the zone. |

#### Example: Create

```python
email_sending_subdomain_preview = client.EmailSendingSubdomainPreview().create({
    "zone_id": "example_zone_id",  # str
    "name": "example_name",  # str
})
```


### EmailSendingSubdomainReputationComplaint

Create an instance: `email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `dict` |  |

#### Example: Load

```python
email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id", "end_at": "end_at", "start_at": "start_at"})
```


### EmailSendingSubdomainResponseSingle

Create an instance: `email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle()`

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
| `name` | `str` | The domain name within the zone. |
| `preview_enabled` | `bool` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `Any` |  |

#### Example: Load

```python
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id"})
```

#### Example: Create

```python
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle().create({
    "zone_id": "example_zone_id",  # str
    "name": "example_name",  # str
})
```


### EmailSendingSubdomainsResponseCollection

Create an instance: `email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `list` |  |

#### Example: List

```python
email_sending_subdomains_response_collections = client.EmailSendingSubdomainsResponseCollection().list({"zone_id": "example"})
```


### EmailSendingSuppression

Create an instance: `email_sending_suppression = client.EmailSendingSuppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | When the suppression was created. |
| `email` | `str` | The email address to suppress. |
| `errors` | `list` |  |
| `expires_at` | `str` | Expiration timestamp for the suppression. |
| `id` | `str` | Unique identifier for this suppression. |
| `items` | `list` | Suppressions to import. |
| `messages` | `list` |  |
| `note` | `str` | Advisory note for this suppression. |
| `read_only` | `bool` | Whether clients may mutate this suppression. |
| `reason` | `str` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `dict` |  |
| `scope` | `Any` | Where the suppression applies. |
| `success` | `bool` |  |
| `zones` | `list` |  |

#### Example: Load

```python
email_sending_suppression = client.EmailSendingSuppression().load({"id": "email_sending_suppression_id"})
```

#### Example: List

```python
email_sending_suppressions = client.EmailSendingSuppression().list({"account_id": "example"})
```

#### Example: Create

```python
email_sending_suppression = client.EmailSendingSuppression().create({
    "account_id": "example_account_id",  # str
    "created_at": "example_created_at",  # str
    "email": "example_email",  # str
    "errors": [],  # list
    "expires_at": "example_expires_at",  # str
    "id": "example_id",  # str
    "items": [],  # list
    "messages": [],  # list
    "read_only": True,  # bool
    "reason": "example_reason",  # str
    "result": {},  # dict
    "success": True,  # bool
})
```


### Subdomain

Create an instance: `subdomain = client.Subdomain()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `str` |  |

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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── cloudflareemailsecurity_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`cloudflareemailsecurity_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
emaildestinationaddressresponsesingle = client.EmailDestinationAddressResponseSingle()
emaildestinationaddressresponsesingle.load({"account_id": "example", "destination_address_identifier": "example"})

# emaildestinationaddressresponsesingle.data_get() now returns the emaildestinationaddressresponsesingle data from the last load
# emaildestinationaddressresponsesingle.match_get() returns the last match criteria
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
