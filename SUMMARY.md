# Cloudflare Email Security

Cloudflare Email Security: email routing rules, addresses, and email security settings.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 24 entities and 139 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [EmailAccountRulesPlanResponseSingle](docs/api/email_account_rules_plan_response_single.html)

Results: Account routing rules plan response.

SDK operations: `create`.

Key fields to recognise:

- `catch_all_rules`: Desired catch-all Email Routing rules managed by the deploying Worker.
- `owner_worker_tag`: Public tag (script_tag) of the Worker that owns this rule.
- `rules`: Desired normal Email Routing rules managed by the deploying Worker.

### [EmailAccountRulesResponseCollection](docs/api/email_account_rules_response_collection.html)

Results: List account routing rules response.

SDK operations: `list`.

### [EmailAuth](docs/api/email_auth.html)

Results: SPF record inspected successfully; DMARC report status retrieved successfully; DMARC report configuration updated successfully.

SDK operations: `load`, `update`.

Key fields to recognise:

- `enabled`: Whether DMARC reports are enabled
- `result`: Recursive SPF inspection tree
- `skip_wizard`: Whether to skip the setup wizard

### [EmailCatchAllRuleResponseSingle](docs/api/email_catch_all_rule_response_single.html)

Results: Get catch-all rule response; Update catch-all rule response.

SDK operations: `load`, `update`.

Key fields to recognise:

- `actions`: List actions for the catch-all routing rule.
- `enabled`: Routing rule status.
- `matchers`: List of matchers for the catch-all routing rule.
- `name`: Routing rule name.
- `owner_worker_tag`: Public tag (script_tag) of the Worker that owns this rule.

### [EmailDestinationAddressResponseSingle](docs/api/email_destination_address_response_single.html)

Results: Create a destination address response; Get a destination address response; Delete destination address response; Update destination address response.

SDK operations: `create`, `load`, `remove`, `update`.

Key fields to recognise:

- `email`: The contact email address of the user.
- `status`: Destination address status.

### [EmailDestinationAddressesResponseCollection](docs/api/email_destination_addresses_response_collection.html)

Results: List destination addresses response.

SDK operations: `list`.

### [EmailDnsSettingsResponseCollection](docs/api/email_dns_settings_response_collection.html)

Results: Email Routing - DNS settings response; Get sending subdomain DNS records response.

SDK operations: `list`.

### [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html)

Results: Disable Email Routing response; Enable Email Routing response; Unlock Email Routing response; Get Email Routing settings response; Update Email Routing settings response; Unlock Email Routing MX records; Apply Email Routing settings response.

SDK operations: `create`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `enabled`: State of the zone settings for Email Routing.
- `name`: Domain of your zone.
- `skip_wizard`: Flag to check if the user skipped the configuration wizard.
- `support_subaddress`: Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.

### [EmailRoutingSuppression](docs/api/email_routing_suppression.html)

Results: Returns suppression.; Returns suppressed emails.; Deletes suppression.

SDK operations: `create`, `list`, `load`, `remove`.

### [EmailRuleResponseSingle](docs/api/email_rule_response_single.html)

Results: Create routing rule response; Get routing rule response; Delete routing rule response; Update routing rule response.

SDK operations: `create`, `load`, `remove`, `update`.

Key fields to recognise:

- `actions`: List actions patterns.
- `enabled`: Routing rule status.
- `matchers`: Matching patterns to forward to your actions.
- `name`: Routing rule name.
- `owner_worker_tag`: Public tag (script_tag) of the Worker that owns this rule.

### [EmailRulesResponseCollection](docs/api/email_rules_response_collection.html)

Results: List routing rules response.

SDK operations: `list`.

### [EmailSecurity](docs/api/email_security.html)

Results: Job cancelled.; Move operation results.; Reclassification request accepted.; Bulk job created.; Bulk move operation results.; Email preview.; Release operation results.; Submission accepted.; Search results for the provided query.; The search is taking longer than expected. Use the Location header to poll for results.; List of submissions.; List of PhishGuard reports.; Jobs retrieved successfully.; Messages for the bulk job.; Action log for the email.; Email message detection details.; Raw email content.; Email trace.; Email message details.; Job details retrieved.; Job deleted.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `action`: The action the job performs on every message matching the search parameters.
- `comment`: Optional note describing the job.
- `destination`: The mailbox folder to move messages to.
- `expected_disposition`: Nonfunctional field. End of life: December 1, 2026.
- `id`: Unique identifier for a message retrieved from investigation.

### [EmailSecurityAnalytics](docs/api/email_security_analytics.html)

Results: Monthly report for the trailing 12 months.

SDK operations: `load`.

### [EmailSecuritySetting](docs/api/email_security_setting.html)

Results: Created allow policy.; Batch operation results.; Created blocked sender.; Created content policy.; Created domain.; Created impersonation registry entry.; Created sending domain restriction.; Created trusted domain.; Created URL ignore pattern.; List of allow policies.; List of domains.; List of trusted domains.; List of content policies.; List of blocked senders.; List of impersonation registry entries.; List of sending domain restrictions.; List of URL ignore patterns for the provided account.; Domain details.; DNS TXT verification details.; Impersonation registry entry details.; Blocked sender details.; URL ignore pattern.; Allow policy details.; Content policy details.; Sending domain restriction details.; Trusted domain details.; Updated domain.; Updated impersonation registry entry.; Updated blocked sender.; Updated URL ignore pattern.; Updated allow policy.; Updated content policy.; Updated sending domain restriction.; Updated trusted domain.; Deleted domain.; Deleted impersonation registry entry.; Deleted blocked sender.; Deleted URL ignore pattern.; Deleted allow policy.; Deleted content policy.; Deleted sending domain restriction.; Deleted trusted domain.; Deleted domains.; Replaced domain.; Replaced blocked sender.; Replaced allow policy.; Replaced trusted domain.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `allowed_delivery_modes`: Delivery modes to onboard the domain through.
- `comments`: Optional note describing the entry.
- `deletes`: IDs of the allow policies to delete.
- `domain`: Domain that requires TLS enforcement.
- `drop_dispositions`: Dispositions to drop instead of delivering, for example

### [EmailSending](docs/api/email_sending.html)

Results: Email sending results.; The raw RFC 5322 MIME message.

SDK operations: `create`, `load`.

Key fields to recognise:

- `attachments`: File attachments and inline images.
- `bcc`: Blind carbon copy recipient(s).
- `cc`: Carbon copy recipient(s).
- `from`: Sender email address.
- `headers`: Custom email headers as key-value pairs.

### [EmailSendingLimitsResponseSingle](docs/api/email_sending_limits_response_single.html)

Results: Get sending limits response.

SDK operations: `load`.

### [EmailSendingReputationResponseSingle](docs/api/email_sending_reputation_response_single.html)

Results: Get account reputation response.

SDK operations: `load`.

### [EmailSendingSubdomainDnsStatus](docs/api/email_sending_subdomain_dns_status.html)

Results: Fix sending subdomain DNS records response; Get sending subdomain DNS status response.

SDK operations: `create`, `load`.

### [EmailSendingSubdomainPreview](docs/api/email_sending_subdomain_preview.html)

Results: Preview sending subdomain DNS response.

SDK operations: `create`.

Key fields to recognise:

- `name`: DNS record name (or @ for the zone apex).

### [EmailSendingSubdomainReputationComplaint](docs/api/email_sending_subdomain_reputation_complaint.html)

Results: Sending subdomain complaint count response.

SDK operations: `load`.

### [EmailSendingSubdomainResponseSingle](docs/api/email_sending_subdomain_response_single.html)

Results: Create a sending subdomain response; Get a sending subdomain response; Update a sending subdomain response.

SDK operations: `create`, `load`, `update`.

Key fields to recognise:

- `drop_suppressed_recipients`: Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request.
- `name`: The exact domain name or a leftmost wildcard such as `*.example.com`.
- `preview_enabled`: Whether sent messages from this subdomain can be previewed in the activity log.

### [EmailSendingSubdomainsResponseCollection](docs/api/email_sending_subdomains_response_collection.html)

Results: List sending subdomains response.

SDK operations: `list`.

### [EmailSendingSuppression](docs/api/email_sending_suppression.html)

Results: Returns suppression.; Suppression mutation completed.; Bulk import result.; Returns Email Sending suppressions.; Returns suppressed emails.; Returns an Email Sending suppression owned by the account.; Deletes suppression.; Returns the updated Email Sending suppression.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: When the suppression was created.
- `email`: The submitted email address for this item.
- `errors`: Number of items that failed to import due to an unexpected error.
- `expires_at`: When the suppression expires. Null for a permanent suppression.
- `id`: The suppression&#39;s identifier.

### [Subdomain](docs/api/subdomain.html)

Results: Delete a sending subdomain response.

SDK operations: `remove`.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [EmailAccountRulesPlanResponseSingle](docs/api/email_account_rules_plan_response_single.html) | `create` | `POST /accounts/{account_id}/email/routing/rules/plan` | Required |
| [EmailAccountRulesResponseCollection](docs/api/email_account_rules_response_collection.html) | `list` | `GET /accounts/{account_id}/email/routing/rules` | Required |
| [EmailAuth](docs/api/email_auth.html) | `load` | `GET /zones/{zone_id}/email/auth/spf/inspect` | Required |
| [EmailAuth](docs/api/email_auth.html) | `load` | `GET /zones/{zone_id}/email/auth/dmarc-reports` | Required |
| [EmailAuth](docs/api/email_auth.html) | `update` | `PATCH /zones/{zone_id}/email/auth/dmarc-reports` | Required |
| [EmailCatchAllRuleResponseSingle](docs/api/email_catch_all_rule_response_single.html) | `load` | `GET /zones/{zone_id}/email/routing/rules/catch_all` | Required |
| [EmailCatchAllRuleResponseSingle](docs/api/email_catch_all_rule_response_single.html) | `update` | `PUT /zones/{zone_id}/email/routing/rules/catch_all` | Required |
| [EmailDestinationAddressResponseSingle](docs/api/email_destination_address_response_single.html) | `create` | `POST /accounts/{account_id}/email/routing/addresses` | Required |
| [EmailDestinationAddressResponseSingle](docs/api/email_destination_address_response_single.html) | `load` | `GET /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}` | Required |
| [EmailDestinationAddressResponseSingle](docs/api/email_destination_address_response_single.html) | `remove` | `DELETE /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}` | Required |
| [EmailDestinationAddressResponseSingle](docs/api/email_destination_address_response_single.html) | `update` | `PATCH /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}` | Required |
| [EmailDestinationAddressesResponseCollection](docs/api/email_destination_addresses_response_collection.html) | `list` | `GET /accounts/{account_id}/email/routing/addresses` | Required |
| [EmailDnsSettingsResponseCollection](docs/api/email_dns_settings_response_collection.html) | `list` | `GET /zones/{zone_id}/email/routing/dns` | Required |
| [EmailDnsSettingsResponseCollection](docs/api/email_dns_settings_response_collection.html) | `list` | `GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `create` | `POST /zones/{zone_id}/email/routing/disable` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `create` | `POST /zones/{zone_id}/email/routing/dns` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `create` | `POST /zones/{zone_id}/email/routing/enable` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `create` | `POST /zones/{zone_id}/email/routing/unlock` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `load` | `GET /zones/{zone_id}/email/routing` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `patch` | `PATCH /zones/{zone_id}/email/routing` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `patch` | `PATCH /zones/{zone_id}/email/routing/dns` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `remove` | `DELETE /zones/{zone_id}/email/routing/dns` | Required |
| [EmailEmailSettingsResponseSingle](docs/api/email_email_settings_response_single.html) | `update` | `PUT /zones/{zone_id}/email/routing` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `create` | `POST /accounts/{account_id}/email/routing/suppression` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `create` | `POST /zones/{zone_id}/email/routing/suppression` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `list` | `GET /accounts/{account_id}/email/routing/suppression` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `list` | `GET /zones/{zone_id}/email/routing/suppression` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `load` | `GET /accounts/{account_id}/email/routing/suppression/{suppression_id}` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `load` | `GET /zones/{zone_id}/email/routing/suppression/{suppression_id}` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `remove` | `DELETE /accounts/{account_id}/email/routing/suppression/{suppression_id}` | Required |
| [EmailRoutingSuppression](docs/api/email_routing_suppression.html) | `remove` | `DELETE /zones/{zone_id}/email/routing/suppression/{suppression_id}` | Required |
| [EmailRuleResponseSingle](docs/api/email_rule_response_single.html) | `create` | `POST /zones/{zone_id}/email/routing/rules` | Required |
| [EmailRuleResponseSingle](docs/api/email_rule_response_single.html) | `load` | `GET /zones/{zone_id}/email/routing/rules/{rule_identifier}` | Required |
| [EmailRuleResponseSingle](docs/api/email_rule_response_single.html) | `remove` | `DELETE /zones/{zone_id}/email/routing/rules/{rule_identifier}` | Required |
| [EmailRuleResponseSingle](docs/api/email_rule_response_single.html) | `update` | `PUT /zones/{zone_id}/email/routing/rules/{rule_identifier}` | Required |
| [EmailRulesResponseCollection](docs/api/email_rules_response_collection.html) | `list` | `GET /zones/{zone_id}/email/routing/rules` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/{investigate_id}/move` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/bulk` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/move` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/preview` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/investigate/release` | Required |
| [EmailSecurity](docs/api/email_security.html) | `create` | `POST /accounts/{account_id}/email-security/submissions` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/submissions` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/phishguard/reports` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/bulk` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/bulk/{job_id}/messages` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}/action_log` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}/detections` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}/preview` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}/raw` | Required |
| [EmailSecurity](docs/api/email_security.html) | `list` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}/trace` | Required |
| [EmailSecurity](docs/api/email_security.html) | `load` | `GET /accounts/{account_id}/email-security/investigate/{investigate_id}` | Required |
| [EmailSecurity](docs/api/email_security.html) | `load` | `GET /accounts/{account_id}/email-security/investigate/bulk/{job_id}` | Required |
| [EmailSecurity](docs/api/email_security.html) | `remove` | `DELETE /accounts/{account_id}/email-security/investigate/bulk/{job_id}` | Required |
| [EmailSecurityAnalytics](docs/api/email_security_analytics.html) | `load` | `GET /accounts/{account_id}/email-security/analytics/monthly_report` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/allow_policies` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/allow_policies/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/block_senders` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/block_senders/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/content_policies` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/content_policies/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/domains` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/domains/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/impersonation_registry` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/sending_domain_restrictions` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/sending_domain_restrictions/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/trusted_domains` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/trusted_domains/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/url_ignore_patterns` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `create` | `POST /accounts/{account_id}/email-security/settings/url_ignore_patterns/batch` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/allow_policies` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/domains` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/trusted_domains` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/content_policies` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/block_senders` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/impersonation_registry` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/sending_domain_restrictions` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `list` | `GET /accounts/{account_id}/email-security/settings/url_ignore_patterns` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/domains/{domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/domains/{domain_id}/verification` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/content_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `load` | `GET /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/domains/{domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/content_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `patch` | `PATCH /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/domains/{domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/content_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `remove` | `DELETE /accounts/{account_id}/email-security/settings/domains` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `update` | `PUT /accounts/{account_id}/email-security/settings/domains/{domain_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `update` | `PUT /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `update` | `PUT /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}` | Required |
| [EmailSecuritySetting](docs/api/email_security_setting.html) | `update` | `PUT /accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}` | Required |
| [EmailSending](docs/api/email_sending.html) | `create` | `POST /accounts/{account_id}/email/sending/send` | Required |
| [EmailSending](docs/api/email_sending.html) | `create` | `POST /accounts/{account_id}/email/sending/send_raw` | Required |
| [EmailSending](docs/api/email_sending.html) | `load` | `GET /accounts/{account_id}/email/sending/messages/{message_id}` | Required |
| [EmailSendingLimitsResponseSingle](docs/api/email_sending_limits_response_single.html) | `load` | `GET /accounts/{account_id}/email/sending/limits` | Required |
| [EmailSendingReputationResponseSingle](docs/api/email_sending_reputation_response_single.html) | `load` | `GET /accounts/{account_id}/email/sending/reputation` | Required |
| [EmailSendingSubdomainDnsStatus](docs/api/email_sending_subdomain_dns_status.html) | `create` | `POST /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns` | Required |
| [EmailSendingSubdomainDnsStatus](docs/api/email_sending_subdomain_dns_status.html) | `load` | `GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns/status` | Required |
| [EmailSendingSubdomainPreview](docs/api/email_sending_subdomain_preview.html) | `create` | `POST /zones/{zone_id}/email/sending/subdomains/preview` | Required |
| [EmailSendingSubdomainReputationComplaint](docs/api/email_sending_subdomain_reputation_complaint.html) | `load` | `GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints` | Required |
| [EmailSendingSubdomainResponseSingle](docs/api/email_sending_subdomain_response_single.html) | `create` | `POST /zones/{zone_id}/email/sending/subdomains` | Required |
| [EmailSendingSubdomainResponseSingle](docs/api/email_sending_subdomain_response_single.html) | `load` | `GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}` | Required |
| [EmailSendingSubdomainResponseSingle](docs/api/email_sending_subdomain_response_single.html) | `update` | `PATCH /zones/{zone_id}/email/sending/subdomains/{subdomain_id}` | Required |
| [EmailSendingSubdomainsResponseCollection](docs/api/email_sending_subdomains_response_collection.html) | `list` | `GET /zones/{zone_id}/email/sending/subdomains` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `create` | `POST /accounts/{account_id}/email/sending/suppression` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `create` | `POST /accounts/{account_id}/email/sending/suppressions` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `create` | `POST /accounts/{account_id}/email/sending/suppressions/bulk` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `create` | `POST /zones/{zone_id}/email/sending/suppression` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `list` | `GET /accounts/{account_id}/email/sending/suppressions` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `list` | `GET /accounts/{account_id}/email/sending/suppression` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `list` | `GET /zones/{zone_id}/email/sending/suppression` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `load` | `GET /accounts/{account_id}/email/sending/suppression/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `load` | `GET /accounts/{account_id}/email/sending/suppressions/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `load` | `GET /zones/{zone_id}/email/sending/suppression/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `remove` | `DELETE /accounts/{account_id}/email/sending/suppression/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `remove` | `DELETE /accounts/{account_id}/email/sending/suppressions/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `remove` | `DELETE /zones/{zone_id}/email/sending/suppression/{suppression_id}` | Required |
| [EmailSendingSuppression](docs/api/email_sending_suppression.html) | `update` | `PATCH /accounts/{account_id}/email/sending/suppressions/{suppression_id}` | Required |
| [Subdomain](docs/api/subdomain.html) | `remove` | `DELETE /zones/{zone_id}/email/sending/subdomains/{subdomain_id}` | Required |

## Connect to the API

- Client API: `https://api.cloudflare.com/client/v4`

The default credential is sent in the `X-Auth-Email` header.

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [Ruby](docs/sdks/rb.html) | `rb/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `cloudflare-email-security_list`: List records for an entity. Supported entities: `email_account_rules_response_collection`, `email_destination_addresses_response_collection`, `email_dns_settings_response_collection`, `email_routing_suppression`, `email_rules_response_collection`, `email_security`, `email_security_setting`, `email_sending_subdomains_response_collection`, `email_sending_suppression`.
- `cloudflare-email-security_load`: Load one record for an entity. Supported entities: `email_auth`, `email_catch_all_rule_response_single`, `email_destination_address_response_single`, `email_email_settings_response_single`, `email_routing_suppression`, `email_rule_response_single`, `email_security`, `email_security_analytics`, `email_security_setting`, `email_sending`, `email_sending_limits_response_single`, `email_sending_reputation_response_single`, `email_sending_subdomain_dns_status`, `email_sending_subdomain_reputation_complaint`, `email_sending_subdomain_response_single`, `email_sending_suppression`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

