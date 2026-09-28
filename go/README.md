# CloudflareEmailSecurity Golang SDK



The Golang SDK for the CloudflareEmailSecurity API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.EmailAccountRulesPlanResponseSingle(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`, `Patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/cloudflare-email-security-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/cloudflare-email-security-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/cloudflare-email-security-sdk/go=../cloudflare-email-security-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/cloudflare-email-security-sdk/go"
)

func main() {
    client := sdk.NewCloudflareEmailSecuritySDK(map[string]any{
        "apikey": os.Getenv("CLOUDFLARE_EMAIL_SECURITY_APIKEY"),
    })

    // Create a emailAccountRulesPlanResponseSingle.
    created, err := client.EmailAccountRulesPlanResponseSingle(nil).Create(map[string]any{"account_id": "example_account_id", "owner_worker_tag": "example_owner_worker_tag"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
emaildestinationaddressresponsesingle, err := client.EmailDestinationAddressResponseSingle(nil).Load(map[string]any{"account_id": "example", "destination_address_identifier": "example"}, nil)
if err != nil {
    // handle err
    return
}
_ = emaildestinationaddressresponsesingle
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

emailDestinationAddressResponseSingle, err := client.EmailDestinationAddressResponseSingle(nil).Load(
    map[string]any{"account_id": "example", "destination_address_identifier": "example"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(emailDestinationAddressResponseSingle) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewCloudflareEmailSecuritySDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewCloudflareEmailSecuritySDK

```go
func NewCloudflareEmailSecuritySDK(options map[string]any) *CloudflareEmailSecuritySDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *CloudflareEmailSecuritySDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CloudflareEmailSecuritySDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `EmailAccountRulesPlanResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailAccountRulesPlanResponseSingle entity instance. |
| `EmailAccountRulesResponseCollection` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailAccountRulesResponseCollection entity instance. |
| `EmailAuth` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailAuth entity instance. |
| `EmailCatchAllRuleResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailCatchAllRuleResponseSingle entity instance. |
| `EmailDestinationAddressResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailDestinationAddressResponseSingle entity instance. |
| `EmailDestinationAddressesResponseCollection` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailDestinationAddressesResponseCollection entity instance. |
| `EmailDnsSettingsResponseCollection` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailDnsSettingsResponseCollection entity instance. |
| `EmailEmailSettingsResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailEmailSettingsResponseSingle entity instance. |
| `EmailRoutingSuppression` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailRoutingSuppression entity instance. |
| `EmailRuleResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailRuleResponseSingle entity instance. |
| `EmailRulesResponseCollection` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailRulesResponseCollection entity instance. |
| `EmailSecurity` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSecurity entity instance. |
| `EmailSecurityAnalytics` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSecurityAnalytics entity instance. |
| `EmailSecuritySetting` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSecuritySetting entity instance. |
| `EmailSending` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSending entity instance. |
| `EmailSendingLimitsResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingLimitsResponseSingle entity instance. |
| `EmailSendingReputationResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingReputationResponseSingle entity instance. |
| `EmailSendingSubdomainDnsStatus` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSubdomainDnsStatus entity instance. |
| `EmailSendingSubdomainPreview` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSubdomainPreview entity instance. |
| `EmailSendingSubdomainReputationComplaint` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSubdomainReputationComplaint entity instance. |
| `EmailSendingSubdomainResponseSingle` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSubdomainResponseSingle entity instance. |
| `EmailSendingSubdomainsResponseCollection` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSubdomainsResponseCollection entity instance. |
| `EmailSendingSuppression` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create an EmailSendingSuppression entity instance. |
| `Subdomain` | `(data map[string]any) CloudflareEmailSecurityEntity` | Create a Subdomain entity instance. |

### Entity interface (CloudflareEmailSecurityEntity)

All entities implement the `CloudflareEmailSecurityEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    emailAccountRulesPlanResponseSingle, err := client.EmailAccountRulesPlanResponseSingle(nil).Create(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // emailAccountRulesPlanResponseSingle is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### EmailAccountRulesPlanResponseSingle

| Field | Description |
| --- | --- |
| `"catch_all_rules"` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `"owner_worker_tag"` | Public tag (script_tag) of the Worker that owns this rule. |
| `"rules"` | Desired normal Email Routing rules managed by the deploying Worker. |

Operations: Create.

API path: `/accounts/{account_id}/email/routing/rules/plan`

#### EmailAccountRulesResponseCollection

| Field | Description |
| --- | --- |
| `"result"` |  |
| `"result_info"` |  |

Operations: List.

API path: `/accounts/{account_id}/email/routing/rules`

#### EmailAuth

| Field | Description |
| --- | --- |
| `"enabled"` | Enable or disable DMARC reports for this zone |
| `"result"` | Recursive SPF inspection tree |
| `"skip_wizard"` | Skip the DMARC setup wizard |

Operations: Load, Update.

API path: `/zones/{zone_id}/email/auth/spf/inspect`

#### EmailCatchAllRuleResponseSingle

| Field | Description |
| --- | --- |
| `"actions"` | List actions for the catch-all routing rule. |
| `"enabled"` | Routing rule status. |
| `"matchers"` | List of matchers for the catch-all routing rule. |
| `"name"` | Routing rule name. |
| `"owner_worker_tag"` | Public tag (script_tag) of the Worker that owns this rule. |
| `"result"` |  |
| `"source"` | Who manages the rule. |

Operations: Load, Update.

API path: `/zones/{zone_id}/email/routing/rules/catch_all`

#### EmailDestinationAddressResponseSingle

| Field | Description |
| --- | --- |
| `"email"` | The contact email address of the user. |
| `"result"` |  |
| `"status"` | Destination address status. |

Operations: Create, Load, Remove, Update.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDestinationAddressesResponseCollection

| Field | Description |
| --- | --- |
| `"result"` |  |
| `"result_info"` |  |

Operations: List.

API path: `/accounts/{account_id}/email/routing/addresses`

#### EmailDnsSettingsResponseCollection

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: List.

API path: `/zones/{zone_id}/email/routing/dns`

#### EmailEmailSettingsResponseSingle

| Field | Description |
| --- | --- |
| `"enabled"` | State of your zone Email Routing settings. |
| `"name"` | Domain of your zone. |
| `"result"` |  |
| `"skip_wizard"` | Flag to check if the user skipped the configuration wizard. |
| `"support_subaddress"` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

Operations: Create, Load, Patch, Remove, Update.

API path: `/zones/{zone_id}/email/routing/disable`

#### EmailRoutingSuppression

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"email"` |  |
| `"expires_at"` |  |
| `"id"` |  |
| `"reason"` |  |
| `"zones"` |  |

Operations: Create, List, Load, Remove.

API path: `/accounts/{account_id}/email/routing/suppression`

#### EmailRuleResponseSingle

| Field | Description |
| --- | --- |
| `"actions"` | List actions patterns. |
| `"enabled"` | Routing rule status. |
| `"matchers"` | Matching patterns to forward to your actions. |
| `"name"` | Routing rule name. |
| `"owner_worker_tag"` | Public tag (script_tag) of the Worker that owns this rule. |
| `"priority"` | Priority of the routing rule. |
| `"result"` |  |
| `"source"` | Who manages the rule. |

Operations: Create, Load, Remove, Update.

API path: `/zones/{zone_id}/email/routing/rules`

#### EmailRulesResponseCollection

| Field | Description |
| --- | --- |
| `"result"` |  |
| `"result_info"` |  |

Operations: List.

API path: `/zones/{zone_id}/email/routing/rules`

#### EmailSecurity

| Field | Description |
| --- | --- |
| `"action"` | The action the job performs on every message matching the search parameters. |
| `"comment"` | Optional note describing the job. |
| `"destination"` | Required when action is 'MOVE'. |
| `"errors"` |  |
| `"expected_disposition"` | Nonfunctional field. |
| `"id"` | Unique identifier for a message retrieved from investigation. |
| `"ids"` | List of message IDs to move. |
| `"messages"` |  |
| `"postfix_ids"` | Deprecated, use `ids` instead. |
| `"result"` |  |
| `"result_info"` |  |
| `"search_params"` |  |
| `"success"` | Whether the API call was successful. |

Operations: Create, List, Load, Remove.

API path: `/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel`

#### EmailSecurityAnalytics

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: Load.

API path: `/accounts/{account_id}/email-security/analytics/monthly_report`

#### EmailSecuritySetting

| Field | Description |
| --- | --- |
| `"allowed_delivery_modes"` | Delivery modes to onboard the domain through. |
| `"comments"` | Optional note describing the reason for the ignore pattern. |
| `"deletes"` | IDs of the allow policies to delete. |
| `"domain"` | The email domain to protect. |
| `"drop_dispositions"` | Dispositions to drop instead of delivering, e.g. |
| `"folder"` | The mailbox folder to scan, for API-scanning domains. |
| `"integration_id"` | Identifier of the CASB integration that authorizes this domain. |
| `"ip_restrictions"` | Source IP ranges mail is accepted from. |
| `"lookback_hops"` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `"patches"` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `"pattern"` | Regular expression identifying URLs to exempt from rewriting. |
| `"posts"` | Allow policies to create. |
| `"puts"` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `"regions"` | Regions that process messages for this domain, e.g. |
| `"require_tls_inbound"` | Require TLS on inbound connections. |
| `"require_tls_outbound"` | Require TLS on outbound connections. |
| `"result"` | An impersonation registry entry. |
| `"transport"` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/accounts/{account_id}/email-security/settings/allow_policies`

#### EmailSending

| Field | Description |
| --- | --- |
| `"attachments"` | File attachments and inline images. |
| `"bcc"` | Blind carbon copy recipient(s). |
| `"cc"` | Carbon copy recipient(s). |
| `"errors"` |  |
| `"from"` | Sender email address. |
| `"headers"` | Custom email headers as key-value pairs. |
| `"html"` | HTML body of the email. |
| `"messages"` |  |
| `"mime_message"` | The full MIME-encoded email message. |
| `"recipients"` | List of recipient email addresses. |
| `"reply_to"` | Reply-to address. |
| `"result"` |  |
| `"result_info"` |  |
| `"subject"` | Email subject line. |
| `"success"` |  |
| `"text"` | Plain text body of the email. |
| `"to"` | Recipient(s). |

Operations: Create, Load.

API path: `/accounts/{account_id}/email/sending/send`

#### EmailSendingLimitsResponseSingle

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: Load.

API path: `/accounts/{account_id}/email/sending/limits`

#### EmailSendingReputationResponseSingle

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: Load.

API path: `/accounts/{account_id}/email/sending/reputation`

#### EmailSendingSubdomainDnsStatus

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: Create, Load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns`

#### EmailSendingSubdomainPreview

| Field | Description |
| --- | --- |
| `"name"` | The domain name within the zone. |

Operations: Create.

API path: `/zones/{zone_id}/email/sending/subdomains/preview`

#### EmailSendingSubdomainReputationComplaint

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: Load.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints`

#### EmailSendingSubdomainResponseSingle

| Field | Description |
| --- | --- |
| `"drop_suppressed_recipients"` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `"name"` | The domain name within the zone. |
| `"preview_enabled"` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `"result"` |  |

Operations: Create, Load, Update.

API path: `/zones/{zone_id}/email/sending/subdomains`

#### EmailSendingSubdomainsResponseCollection

| Field | Description |
| --- | --- |
| `"result"` |  |

Operations: List.

API path: `/zones/{zone_id}/email/sending/subdomains`

#### EmailSendingSuppression

| Field | Description |
| --- | --- |
| `"created_at"` | When the suppression was created. |
| `"email"` | The email address to suppress. |
| `"errors"` |  |
| `"expires_at"` | Expiration timestamp for the suppression. |
| `"id"` | Unique identifier for this suppression. |
| `"items"` | Suppressions to import. |
| `"messages"` |  |
| `"note"` | Advisory note for this suppression. |
| `"read_only"` | Whether clients may mutate this suppression. |
| `"reason"` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `"result"` |  |
| `"scope"` | Where the suppression applies. |
| `"success"` |  |
| `"zones"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/accounts/{account_id}/email/sending/suppression`

#### Subdomain

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/zones/{zone_id}/email/sending/subdomains/{subdomain_id}`



## Entities


### EmailAccountRulesPlanResponseSingle

Create an instance: `emailAccountRulesPlanResponseSingle := client.EmailAccountRulesPlanResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `catch_all_rules` | `[]any` | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `[]any` | Desired normal Email Routing rules managed by the deploying Worker. |

#### Example: Create

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


### EmailAccountRulesResponseCollection

Create an instance: `emailAccountRulesResponseCollection := client.EmailAccountRulesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `[]any` |  |
| `result_info` | `map[string]any` |  |

#### Example: List

```go
emailAccountRulesResponseCollections, err := client.EmailAccountRulesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailAccountRulesResponseCollections) // the array of records
```


### EmailAuth

Create an instance: `emailAuth := client.EmailAuth(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `bool` | Enable or disable DMARC reports for this zone |
| `result` | `map[string]any` | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | Skip the DMARC setup wizard |

#### Example: Load

```go
emailAuth, err := client.EmailAuth(nil).Load(map[string]any{"id": "email_auth_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailAuth) // the loaded record
```


### EmailCatchAllRuleResponseSingle

Create an instance: `emailCatchAllRuleResponseSingle := client.EmailCatchAllRuleResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `[]any` | List actions for the catch-all routing rule. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `[]any` | List of matchers for the catch-all routing rule. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `map[string]any` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```go
emailCatchAllRuleResponseSingle, err := client.EmailCatchAllRuleResponseSingle(nil).Load(map[string]any{"zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailCatchAllRuleResponseSingle) // the loaded record
```


### EmailDestinationAddressResponseSingle

Create an instance: `emailDestinationAddressResponseSingle := client.EmailDestinationAddressResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` | The contact email address of the user. |
| `result` | `map[string]any` |  |
| `status` | `string` | Destination address status. |

#### Example: Load

```go
emailDestinationAddressResponseSingle, err := client.EmailDestinationAddressResponseSingle(nil).Load(map[string]any{"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailDestinationAddressResponseSingle) // the loaded record
```

#### Example: Create

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


### EmailDestinationAddressesResponseCollection

Create an instance: `emailDestinationAddressesResponseCollection := client.EmailDestinationAddressesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `[]any` |  |
| `result_info` | `map[string]any` |  |

#### Example: List

```go
emailDestinationAddressesResponseCollections, err := client.EmailDestinationAddressesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailDestinationAddressesResponseCollections) // the array of records
```


### EmailDnsSettingsResponseCollection

Create an instance: `emailDnsSettingsResponseCollection := client.EmailDnsSettingsResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `[]any` |  |

#### Example: List

```go
emailDnsSettingsResponseCollections, err := client.EmailDnsSettingsResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailDnsSettingsResponseCollections) // the array of records
```


### EmailEmailSettingsResponseSingle

Create an instance: `emailEmailSettingsResponseSingle := client.EmailEmailSettingsResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enabled` | `bool` | State of your zone Email Routing settings. |
| `name` | `string` | Domain of your zone. |
| `result` | `map[string]any` |  |
| `skip_wizard` | `bool` | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

#### Example: Load

```go
emailEmailSettingsResponseSingle, err := client.EmailEmailSettingsResponseSingle(nil).Load(map[string]any{"zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailEmailSettingsResponseSingle) // the loaded record
```

#### Example: Create

```go
result, err := client.EmailEmailSettingsResponseSingle(nil).Create(map[string]any{
    "zone_id": "example_zone_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### EmailRoutingSuppression

Create an instance: `emailRoutingSuppression := client.EmailRoutingSuppression(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `expires_at` | `string` |  |
| `id` | `string` |  |
| `reason` | `string` |  |
| `zones` | `[]any` |  |

#### Example: Load

```go
emailRoutingSuppression, err := client.EmailRoutingSuppression(nil).Load(map[string]any{"id": "email_routing_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailRoutingSuppression) // the loaded record
```

#### Example: List

```go
emailRoutingSuppressions, err := client.EmailRoutingSuppression(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailRoutingSuppressions) // the array of records
```

#### Example: Create

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


### EmailRuleResponseSingle

Create an instance: `emailRuleResponseSingle := client.EmailRuleResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `[]any` | List actions patterns. |
| `enabled` | `bool` | Routing rule status. |
| `matchers` | `[]any` | Matching patterns to forward to your actions. |
| `name` | `string` | Routing rule name. |
| `owner_worker_tag` | `string` | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float64` | Priority of the routing rule. |
| `result` | `map[string]any` |  |
| `source` | `string` | Who manages the rule. |

#### Example: Load

```go
emailRuleResponseSingle, err := client.EmailRuleResponseSingle(nil).Load(map[string]any{"rule_identifier": "rule_identifier", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailRuleResponseSingle) // the loaded record
```

#### Example: Create

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


### EmailRulesResponseCollection

Create an instance: `emailRulesResponseCollection := client.EmailRulesResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `[]any` |  |
| `result_info` | `map[string]any` |  |

#### Example: List

```go
emailRulesResponseCollections, err := client.EmailRulesResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailRulesResponseCollections) // the array of records
```


### EmailSecurity

Create an instance: `emailSecurity := client.EmailSecurity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` | The action the job performs on every message matching the search parameters. |
| `comment` | `string` | Optional note describing the job. |
| `destination` | `any` | Required when action is 'MOVE'. |
| `errors` | `[]any` |  |
| `expected_disposition` | `string` | Nonfunctional field. |
| `id` | `string` | Unique identifier for a message retrieved from investigation. |
| `ids` | `[]any` | List of message IDs to move. |
| `messages` | `[]any` |  |
| `postfix_ids` | `[]any` | Deprecated, use `ids` instead. |
| `result` | `map[string]any` |  |
| `result_info` | `map[string]any` |  |
| `search_params` | `map[string]any` |  |
| `success` | `bool` | Whether the API call was successful. |

#### Example: Load

```go
emailSecurity, err := client.EmailSecurity(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSecurity) // the loaded record
```

#### Example: List

```go
emailSecuritys, err := client.EmailSecurity(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSecuritys) // the array of records
```

#### Example: Create

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


### EmailSecurityAnalytics

Create an instance: `emailSecurityAnalytics := client.EmailSecurityAnalytics(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `map[string]any` |  |

#### Example: Load

```go
emailSecurityAnalytics, err := client.EmailSecurityAnalytics(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSecurityAnalytics) // the loaded record
```


### EmailSecuritySetting

Create an instance: `emailSecuritySetting := client.EmailSecuritySetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `allowed_delivery_modes` | `[]any` | Delivery modes to onboard the domain through. |
| `comments` | `string` | Optional note describing the reason for the ignore pattern. |
| `deletes` | `[]any` | IDs of the allow policies to delete. |
| `domain` | `string` | The email domain to protect. |
| `drop_dispositions` | `[]any` | Dispositions to drop instead of delivering, e.g. |
| `folder` | `any` | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `string` | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `[]any` | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `[]any` | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `string` | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `[]any` | Allow policies to create. |
| `puts` | `[]any` | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `[]any` | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | Require TLS on outbound connections. |
| `result` | `map[string]any` | An impersonation registry entry. |
| `transport` | `string` | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

#### Example: Load

```go
emailSecuritySetting, err := client.EmailSecuritySetting(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSecuritySetting) // the loaded record
```

#### Example: List

```go
emailSecuritySettings, err := client.EmailSecuritySetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSecuritySettings) // the array of records
```

#### Example: Create

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


### EmailSending

Create an instance: `emailSending := client.EmailSending(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | File attachments and inline images. |
| `bcc` | `any` | Blind carbon copy recipient(s). |
| `cc` | `any` | Carbon copy recipient(s). |
| `errors` | `[]any` |  |
| `from` | `any` | Sender email address. |
| `headers` | `map[string]any` | Custom email headers as key-value pairs. |
| `html` | `string` | HTML body of the email. |
| `messages` | `[]any` |  |
| `mime_message` | `string` | The full MIME-encoded email message. |
| `recipients` | `[]any` | List of recipient email addresses. |
| `reply_to` | `any` | Reply-to address. |
| `result` | `map[string]any` |  |
| `result_info` | `map[string]any` |  |
| `subject` | `string` | Email subject line. |
| `success` | `bool` |  |
| `text` | `string` | Plain text body of the email. |
| `to` | `any` | Recipient(s). |

#### Example: Load

```go
emailSending, err := client.EmailSending(nil).Load(map[string]any{"account_id": "account_id", "message_id": "message_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSending) // the loaded record
```

#### Example: Create

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


### EmailSendingLimitsResponseSingle

Create an instance: `emailSendingLimitsResponseSingle := client.EmailSendingLimitsResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `map[string]any` |  |

#### Example: Load

```go
emailSendingLimitsResponseSingle, err := client.EmailSendingLimitsResponseSingle(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingLimitsResponseSingle) // the loaded record
```


### EmailSendingReputationResponseSingle

Create an instance: `emailSendingReputationResponseSingle := client.EmailSendingReputationResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `map[string]any` |  |

#### Example: Load

```go
emailSendingReputationResponseSingle, err := client.EmailSendingReputationResponseSingle(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingReputationResponseSingle) // the loaded record
```


### EmailSendingSubdomainDnsStatus

Create an instance: `emailSendingSubdomainDnsStatus := client.EmailSendingSubdomainDnsStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `map[string]any` |  |

#### Example: Load

```go
emailSendingSubdomainDnsStatus, err := client.EmailSendingSubdomainDnsStatus(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSubdomainDnsStatus) // the loaded record
```

#### Example: Create

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


### EmailSendingSubdomainPreview

Create an instance: `emailSendingSubdomainPreview := client.EmailSendingSubdomainPreview(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `name` | `string` | The domain name within the zone. |

#### Example: Create

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


### EmailSendingSubdomainReputationComplaint

Create an instance: `emailSendingSubdomainReputationComplaint := client.EmailSendingSubdomainReputationComplaint(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `map[string]any` |  |

#### Example: Load

```go
emailSendingSubdomainReputationComplaint, err := client.EmailSendingSubdomainReputationComplaint(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id", "end_at": "end_at", "start_at": "start_at"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSubdomainReputationComplaint) // the loaded record
```


### EmailSendingSubdomainResponseSingle

Create an instance: `emailSendingSubdomainResponseSingle := client.EmailSendingSubdomainResponseSingle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `drop_suppressed_recipients` | `bool` | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `string` | The domain name within the zone. |
| `preview_enabled` | `bool` | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `any` |  |

#### Example: Load

```go
emailSendingSubdomainResponseSingle, err := client.EmailSendingSubdomainResponseSingle(nil).Load(map[string]any{"subdomain_id": "subdomain_id", "zone_id": "zone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSubdomainResponseSingle) // the loaded record
```

#### Example: Create

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


### EmailSendingSubdomainsResponseCollection

Create an instance: `emailSendingSubdomainsResponseCollection := client.EmailSendingSubdomainsResponseCollection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `result` | `[]any` |  |

#### Example: List

```go
emailSendingSubdomainsResponseCollections, err := client.EmailSendingSubdomainsResponseCollection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSubdomainsResponseCollections) // the array of records
```


### EmailSendingSuppression

Create an instance: `emailSendingSuppression := client.EmailSendingSuppression(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | When the suppression was created. |
| `email` | `string` | The email address to suppress. |
| `errors` | `[]any` |  |
| `expires_at` | `string` | Expiration timestamp for the suppression. |
| `id` | `string` | Unique identifier for this suppression. |
| `items` | `[]any` | Suppressions to import. |
| `messages` | `[]any` |  |
| `note` | `string` | Advisory note for this suppression. |
| `read_only` | `bool` | Whether clients may mutate this suppression. |
| `reason` | `string` | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `map[string]any` |  |
| `scope` | `any` | Where the suppression applies. |
| `success` | `bool` |  |
| `zones` | `[]any` |  |

#### Example: Load

```go
emailSendingSuppression, err := client.EmailSendingSuppression(nil).Load(map[string]any{"id": "email_sending_suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSuppression) // the loaded record
```

#### Example: List

```go
emailSendingSuppressions, err := client.EmailSendingSuppression(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(emailSendingSuppressions) // the array of records
```

#### Example: Create

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


### Subdomain

Create an instance: `subdomain := client.Subdomain(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/cloudflare-email-security-sdk/go/
├── cloudflare-email-security.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/cloudflare-email-security-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
emaildestinationaddressresponsesingle := client.EmailDestinationAddressResponseSingle(nil)
emaildestinationaddressresponsesingle.Load(map[string]any{"account_id": "example", "destination_address_identifier": "example"}, nil)

// emaildestinationaddressresponsesingle.Data() now returns the emaildestinationaddressresponsesingle data from the last load
// emaildestinationaddressresponsesingle.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
