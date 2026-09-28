# CloudflareEmailSecurity Python SDK Reference

Complete API reference for the CloudflareEmailSecurity Python SDK.


## CloudflareEmailSecuritySDK

### Constructor

```python
from cloudflareemailsecurity_sdk import CloudflareEmailSecuritySDK

client = CloudflareEmailSecuritySDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CloudflareEmailSecuritySDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = CloudflareEmailSecuritySDK.test()
```


### Instance Methods

#### `EmailAccountRulesPlanResponseSingle(data=None)`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailAccountRulesResponseCollection(data=None)`

Create a new `EmailAccountRulesResponseCollectionEntity` instance. Pass `None` for no initial data.

#### `EmailAuth(data=None)`

Create a new `EmailAuthEntity` instance. Pass `None` for no initial data.

#### `EmailCatchAllRuleResponseSingle(data=None)`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailDestinationAddressResponseSingle(data=None)`

Create a new `EmailDestinationAddressResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailDestinationAddressesResponseCollection(data=None)`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance. Pass `None` for no initial data.

#### `EmailDnsSettingsResponseCollection(data=None)`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance. Pass `None` for no initial data.

#### `EmailEmailSettingsResponseSingle(data=None)`

Create a new `EmailEmailSettingsResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailRoutingSuppression(data=None)`

Create a new `EmailRoutingSuppressionEntity` instance. Pass `None` for no initial data.

#### `EmailRuleResponseSingle(data=None)`

Create a new `EmailRuleResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailRulesResponseCollection(data=None)`

Create a new `EmailRulesResponseCollectionEntity` instance. Pass `None` for no initial data.

#### `EmailSecurity(data=None)`

Create a new `EmailSecurityEntity` instance. Pass `None` for no initial data.

#### `EmailSecurityAnalytics(data=None)`

Create a new `EmailSecurityAnalyticsEntity` instance. Pass `None` for no initial data.

#### `EmailSecuritySetting(data=None)`

Create a new `EmailSecuritySettingEntity` instance. Pass `None` for no initial data.

#### `EmailSending(data=None)`

Create a new `EmailSendingEntity` instance. Pass `None` for no initial data.

#### `EmailSendingLimitsResponseSingle(data=None)`

Create a new `EmailSendingLimitsResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailSendingReputationResponseSingle(data=None)`

Create a new `EmailSendingReputationResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSubdomainDnsStatus(data=None)`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSubdomainPreview(data=None)`

Create a new `EmailSendingSubdomainPreviewEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSubdomainReputationComplaint(data=None)`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSubdomainResponseSingle(data=None)`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSubdomainsResponseCollection(data=None)`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance. Pass `None` for no initial data.

#### `EmailSendingSuppression(data=None)`

Create a new `EmailSendingSuppressionEntity` instance. Pass `None` for no initial data.

#### `Subdomain(data=None)`

Create a new `SubdomainEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## EmailAccountRulesPlanResponseSingleEntity

```python
email_account_rules_plan_response_single = client.EmailAccountRulesPlanResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `catch_all_rules` | `list` | No | Desired catch-all Email Routing rules managed by the deploying Worker. |
| `owner_worker_tag` | `str` | Yes | Public tag (script_tag) of the Worker that owns this rule. |
| `rules` | `list` | No | Desired normal Email Routing rules managed by the deploying Worker. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailAccountRulesPlanResponseSingle().create({
    "account_id": "example_account_id",  # str
    "owner_worker_tag": "example_owner_worker_tag",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAccountRulesPlanResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailAccountRulesResponseCollectionEntity

```python
email_account_rules_response_collection = client.EmailAccountRulesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `list` | No |  |
| `result_info` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailAccountRulesResponseCollection().list({"account_id": "example"})
for email_account_rules_response_collection in results:
    print(email_account_rules_response_collection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAccountRulesResponseCollectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailAuthEntity

```python
email_auth = client.EmailAuth()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | Enable or disable DMARC reports for this zone |
| `result` | `dict` | Yes | Recursive SPF inspection tree |
| `skip_wizard` | `bool` | No | Skip the DMARC setup wizard |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `enabled` | - | - |
| `result` | Yes | - |
| `skip_wizard` | - | - |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailAuth().load({"id": "email_auth_id", "zone_id": "zone_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailAuth().update({
    "zone_id": "zone_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailAuthEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailCatchAllRuleResponseSingleEntity

```python
email_catch_all_rule_response_single = client.EmailCatchAllRuleResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `list` | Yes | List actions for the catch-all routing rule. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `list` | Yes | List of matchers for the catch-all routing rule. |
| `name` | `str` | No | Routing rule name. |
| `owner_worker_tag` | `str` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `result` | `dict` | No |  |
| `source` | `str` | No | Who manages the rule. |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailCatchAllRuleResponseSingle().load({"zone_id": "zone_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailCatchAllRuleResponseSingle().update({
    "zone_id": "zone_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailCatchAllRuleResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailDestinationAddressResponseSingleEntity

```python
email_destination_address_response_single = client.EmailDestinationAddressResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `str` | Yes | The contact email address of the user. |
| `result` | `dict` | No |  |
| `status` | `str` | Yes | Destination address status. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailDestinationAddressResponseSingle().create({
    "account_id": "example_account_id",  # str
    "email": "example_email",  # str
    "status": "example_status",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailDestinationAddressResponseSingle().load({"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailDestinationAddressResponseSingle().remove({"account_id": "account_id", "destination_address_identifier": "destination_address_identifier"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailDestinationAddressResponseSingle().update({
    "account_id": "account_id",
    "destination_address_identifier": "destination_address_identifier",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDestinationAddressResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailDestinationAddressesResponseCollectionEntity

```python
email_destination_addresses_response_collection = client.EmailDestinationAddressesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `list` | No |  |
| `result_info` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailDestinationAddressesResponseCollection().list({"account_id": "example"})
for email_destination_addresses_response_collection in results:
    print(email_destination_addresses_response_collection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDestinationAddressesResponseCollectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailDnsSettingsResponseCollectionEntity

```python
email_dns_settings_response_collection = client.EmailDnsSettingsResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailDnsSettingsResponseCollection().list({"zone_id": "example"})
for email_dns_settings_response_collection in results:
    print(email_dns_settings_response_collection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailDnsSettingsResponseCollectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailEmailSettingsResponseSingleEntity

```python
email_email_settings_response_single = client.EmailEmailSettingsResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enabled` | `bool` | No | State of your zone Email Routing settings. |
| `name` | `str` | No | Domain of your zone. |
| `result` | `dict` | No |  |
| `skip_wizard` | `bool` | No | Flag to check if the user skipped the configuration wizard. |
| `support_subaddress` | `bool` | No | Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailEmailSettingsResponseSingle().create({
    "zone_id": "example_zone_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailEmailSettingsResponseSingle().load({"zone_id": "zone_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailEmailSettingsResponseSingle().remove({"zone_id": "zone_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailEmailSettingsResponseSingle().update({
    "zone_id": "zone_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailEmailSettingsResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailRoutingSuppressionEntity

```python
email_routing_suppression = client.EmailRoutingSuppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes |  |
| `email` | `str` | Yes |  |
| `expires_at` | `str` | Yes |  |
| `id` | `str` | Yes |  |
| `reason` | `str` | Yes |  |
| `zones` | `list` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailRoutingSuppression().create({
    "account_id": "example_account_id",  # str
    "created_at": "example_created_at",  # str
    "email": "example_email",  # str
    "expires_at": "example_expires_at",  # str
    "id": "example_id",  # str
    "reason": "example_reason",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailRoutingSuppression().list({"account_id": "example"})
for email_routing_suppression in results:
    print(email_routing_suppression)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailRoutingSuppression().load({"id": "email_routing_suppression_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailRoutingSuppression().remove({"id": "email_routing_suppression_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRoutingSuppressionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailRuleResponseSingleEntity

```python
email_rule_response_single = client.EmailRuleResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `list` | Yes | List actions patterns. |
| `enabled` | `bool` | No | Routing rule status. |
| `matchers` | `list` | Yes | Matching patterns to forward to your actions. |
| `name` | `str` | No | Routing rule name. |
| `owner_worker_tag` | `str` | No | Public tag (script_tag) of the Worker that owns this rule. |
| `priority` | `float` | No | Priority of the routing rule. |
| `result` | `dict` | No |  |
| `source` | `str` | No | Who manages the rule. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailRuleResponseSingle().create({
    "zone_id": "example_zone_id",  # str
    "actions": [],  # list
    "matchers": [],  # list
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailRuleResponseSingle().load({"rule_identifier": "rule_identifier", "zone_id": "zone_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailRuleResponseSingle().remove({"rule_identifier": "rule_identifier", "zone_id": "zone_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailRuleResponseSingle().update({
    "rule_identifier": "rule_identifier",
    "zone_id": "zone_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRuleResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailRulesResponseCollectionEntity

```python
email_rules_response_collection = client.EmailRulesResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `list` | No |  |
| `result_info` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailRulesResponseCollection().list({"zone_id": "example"})
for email_rules_response_collection in results:
    print(email_rules_response_collection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailRulesResponseCollectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSecurityEntity

```python
email_security = client.EmailSecurity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `action` | `str` | Yes | The action the job performs on every message matching the search parameters. |
| `comment` | `str` | No | Optional note describing the job. |
| `destination` | `Any` | No | Required when action is 'MOVE'. |
| `errors` | `list` | Yes |  |
| `expected_disposition` | `str` | No | Nonfunctional field. |
| `id` | `str` | Yes | Unique identifier for a message retrieved from investigation. |
| `ids` | `list` | Yes | List of message IDs to move. |
| `messages` | `list` | Yes |  |
| `postfix_ids` | `list` | No | Deprecated, use `ids` instead. |
| `result` | `dict` | Yes |  |
| `result_info` | `dict` | Yes |  |
| `search_params` | `dict` | Yes |  |
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSecurity().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailSecurity().list({"account_id": "example"})
for email_security in results:
    print(email_security)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSecurity().load({"account_id": "account_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailSecurity().remove({"account_id": "account_id", "job_id": "job_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecurityEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSecurityAnalyticsEntity

```python
email_security_analytics = client.EmailSecurityAnalytics()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSecurityAnalytics().load({"account_id": "account_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecurityAnalyticsEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSecuritySettingEntity

```python
email_security_setting = client.EmailSecuritySetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_delivery_modes` | `list` | Yes | Delivery modes to onboard the domain through. |
| `comments` | `str` | No | Optional note describing the reason for the ignore pattern. |
| `deletes` | `list` | Yes | IDs of the allow policies to delete. |
| `domain` | `str` | Yes | The email domain to protect. |
| `drop_dispositions` | `list` | Yes | Dispositions to drop instead of delivering, e.g. |
| `folder` | `Any` | No | The mailbox folder to scan, for API-scanning domains. |
| `integration_id` | `str` | No | Identifier of the CASB integration that authorizes this domain. |
| `ip_restrictions` | `list` | Yes | Source IP ranges mail is accepted from. |
| `lookback_hops` | `int` | No | Number of hops to trace back through received headers when reconstructing the original message (1-20). |
| `patches` | `list` | Yes | Partial updates to apply — each entry carries the policy's ID and only the fields to change. |
| `pattern` | `str` | Yes | Regular expression identifying URLs to exempt from rewriting. |
| `posts` | `list` | Yes | Allow policies to create. |
| `puts` | `list` | Yes | Full replacements to apply — each entry carries the policy's ID and every field of its new value. |
| `regions` | `list` | Yes | Regions that process messages for this domain, e.g. |
| `require_tls_inbound` | `bool` | No | Require TLS on inbound connections. |
| `require_tls_outbound` | `bool` | No | Require TLS on outbound connections. |
| `result` | `dict` | No | An impersonation registry entry. |
| `transport` | `str` | No | The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSecuritySetting().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailSecuritySetting().list({"account_id": "example"})
for email_security_setting in results:
    print(email_security_setting)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSecuritySetting().load({"account_id": "account_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailSecuritySetting().remove({"account_id": "account_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailSecuritySetting().update({
    "account_id": "account_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSecuritySettingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingEntity

```python
email_sending = client.EmailSending()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | No | File attachments and inline images. |
| `bcc` | `Any` | No | Blind carbon copy recipient(s). |
| `cc` | `Any` | No | Carbon copy recipient(s). |
| `errors` | `list` | Yes |  |
| `from` | `Any` | Yes | Sender email address. |
| `headers` | `dict` | No | Custom email headers as key-value pairs. |
| `html` | `str` | No | HTML body of the email. |
| `messages` | `list` | Yes |  |
| `mime_message` | `str` | Yes | The full MIME-encoded email message. |
| `recipients` | `list` | Yes | List of recipient email addresses. |
| `reply_to` | `Any` | No | Reply-to address. |
| `result` | `dict` | Yes |  |
| `result_info` | `dict` | Yes |  |
| `subject` | `str` | Yes | Email subject line. |
| `success` | `bool` | Yes |  |
| `text` | `str` | No | Plain text body of the email. |
| `to` | `Any` | No | Recipient(s). |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSending().create({
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSending().load({"account_id": "account_id", "message_id": "message_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingLimitsResponseSingleEntity

```python
email_sending_limits_response_single = client.EmailSendingLimitsResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingLimitsResponseSingle().load({"account_id": "account_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingLimitsResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingReputationResponseSingleEntity

```python
email_sending_reputation_response_single = client.EmailSendingReputationResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `dict` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingReputationResponseSingle().load({"account_id": "account_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingReputationResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSubdomainDnsStatusEntity

```python
email_sending_subdomain_dns_status = client.EmailSendingSubdomainDnsStatus()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSendingSubdomainDnsStatus().create({
    "subdomain_id": "example_subdomain_id",  # str
    "zone_id": "example_zone_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingSubdomainDnsStatus().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainDnsStatusEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSubdomainPreviewEntity

```python
email_sending_subdomain_preview = client.EmailSendingSubdomainPreview()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `str` | Yes | The domain name within the zone. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSendingSubdomainPreview().create({
    "zone_id": "example_zone_id",  # str
    "name": "example_name",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainPreviewEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSubdomainReputationComplaintEntity

```python
email_sending_subdomain_reputation_complaint = client.EmailSendingSubdomainReputationComplaint()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `dict` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingSubdomainReputationComplaint().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id", "end_at": "end_at", "start_at": "start_at"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainReputationComplaintEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSubdomainResponseSingleEntity

```python
email_sending_subdomain_response_single = client.EmailSendingSubdomainResponseSingle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `drop_suppressed_recipients` | `bool` | No | Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request. |
| `name` | `str` | Yes | The domain name within the zone. |
| `preview_enabled` | `bool` | No | Whether sent messages from this subdomain can be previewed in the activity log. |
| `result` | `Any` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSendingSubdomainResponseSingle().create({
    "zone_id": "example_zone_id",  # str
    "name": "example_name",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingSubdomainResponseSingle().load({"subdomain_id": "subdomain_id", "zone_id": "zone_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailSendingSubdomainResponseSingle().update({
    "subdomain_id": "subdomain_id",
    "zone_id": "zone_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainResponseSingleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSubdomainsResponseCollectionEntity

```python
email_sending_subdomains_response_collection = client.EmailSendingSubdomainsResponseCollection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `result` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailSendingSubdomainsResponseCollection().list({"zone_id": "example"})
for email_sending_subdomains_response_collection in results:
    print(email_sending_subdomains_response_collection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSubdomainsResponseCollectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailSendingSuppressionEntity

```python
email_sending_suppression = client.EmailSendingSuppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes | When the suppression was created. |
| `email` | `str` | Yes | The email address to suppress. |
| `errors` | `list` | Yes |  |
| `expires_at` | `str` | Yes | Expiration timestamp for the suppression. |
| `id` | `str` | Yes | Unique identifier for this suppression. |
| `items` | `list` | Yes | Suppressions to import. |
| `messages` | `list` | Yes |  |
| `note` | `str` | No | Advisory note for this suppression. |
| `read_only` | `bool` | Yes | Whether clients may mutate this suppression. |
| `reason` | `str` | Yes | Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`. |
| `result` | `dict` | Yes |  |
| `scope` | `Any` | No | Where the suppression applies. |
| `success` | `bool` | Yes |  |
| `zones` | `list` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EmailSendingSuppression().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EmailSendingSuppression().list({"account_id": "example"})
for email_sending_suppression in results:
    print(email_sending_suppression)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EmailSendingSuppression().load({"id": "email_sending_suppression_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EmailSendingSuppression().remove({"id": "email_sending_suppression_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EmailSendingSuppression().update({
    "id": "email_sending_suppression_id",
    "account_id": "account_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailSendingSuppressionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SubdomainEntity

```python
subdomain = client.Subdomain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Subdomain().remove({"id": "id", "zone_id": "zone_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubdomainEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```python
client = CloudflareEmailSecuritySDK({
    "feature": {
        "test": {"active": True},
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

