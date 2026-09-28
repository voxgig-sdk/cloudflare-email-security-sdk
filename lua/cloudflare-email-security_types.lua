-- Typed models for the CloudflareEmailSecurity SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class EmailAccountRulesPlanResponseSingle
---@field catch_all_rules? table
---@field owner_worker_tag string
---@field rules? table

---@class EmailAccountRulesPlanResponseSingleCreateData
---@field account_id string
---@field catch_all_rules? table
---@field owner_worker_tag string
---@field rules? table

---@class EmailAccountRulesResponseCollection
---@field result? table
---@field result_info? table

---@class EmailAccountRulesResponseCollectionListMatch
---@field account_id string
---@field enabled? boolean
---@field page? number
---@field per_page? number

---@class EmailAuth
---@field enabled? boolean
---@field result table
---@field skip_wizard? boolean

---@class EmailAuthLoadMatch
---@field zone_id string
---@field id? string

---@class EmailAuthUpdateData
---@field zone_id string
---@field enabled? boolean
---@field result? table
---@field skip_wizard? boolean

---@class EmailCatchAllRuleResponseSingle
---@field actions table
---@field enabled? boolean
---@field matchers table
---@field name? string
---@field owner_worker_tag? string
---@field result? table
---@field source? string

---@class EmailCatchAllRuleResponseSingleLoadMatch
---@field zone_id string

---@class EmailCatchAllRuleResponseSingleUpdateData
---@field zone_id string
---@field actions? table
---@field enabled? boolean
---@field matchers? table
---@field name? string
---@field owner_worker_tag? string
---@field result? table
---@field source? string

---@class EmailDestinationAddressResponseSingle
---@field email string
---@field result? table
---@field status string

---@class EmailDestinationAddressResponseSingleLoadMatch
---@field account_id string
---@field destination_address_identifier string

---@class EmailDestinationAddressResponseSingleCreateData
---@field account_id string
---@field email string
---@field result? table
---@field status string

---@class EmailDestinationAddressResponseSingleUpdateData
---@field account_id string
---@field destination_address_identifier string
---@field email? string
---@field result? table
---@field status? string

---@class EmailDestinationAddressResponseSingleRemoveMatch
---@field account_id string
---@field destination_address_identifier string

---@class EmailDestinationAddressesResponseCollection
---@field result? table
---@field result_info? table

---@class EmailDestinationAddressesResponseCollectionListMatch
---@field account_id string
---@field direction? string
---@field page? number
---@field per_page? number
---@field verified? boolean

---@class EmailDnsSettingsResponseCollection
---@field result? table

---@class EmailDnsSettingsResponseCollectionListMatch
---@field zone_id string
---@field subdomain? string
---@field subdomain_id? string

---@class EmailEmailSettingsResponseSingle
---@field enabled? boolean
---@field name? string
---@field result? table
---@field skip_wizard? boolean
---@field support_subaddress? boolean

---@class EmailEmailSettingsResponseSingleLoadMatch
---@field zone_id string

---@class EmailEmailSettingsResponseSingleCreateData
---@field zone_id string
---@field enabled? boolean
---@field name? string
---@field result? table
---@field skip_wizard? boolean
---@field support_subaddress? boolean

---@class EmailEmailSettingsResponseSingleUpdateData
---@field zone_id string
---@field enabled? boolean
---@field name? string
---@field result? table
---@field skip_wizard? boolean
---@field support_subaddress? boolean

---@class EmailEmailSettingsResponseSingleRemoveMatch
---@field zone_id string

---@class EmailRoutingSuppression
---@field created_at string
---@field email string
---@field expires_at string
---@field id string
---@field reason string
---@field zones? table

---@class EmailRoutingSuppressionLoadMatch
---@field account_id? string
---@field id string
---@field zone_id? string

---@class EmailRoutingSuppressionListMatch
---@field account_id string
---@field direction? string
---@field order? string
---@field page? number
---@field per_page? number

---@class EmailRoutingSuppressionCreateData
---@field account_id string
---@field created_at string
---@field email string
---@field expires_at string
---@field id string
---@field reason string
---@field zones? table

---@class EmailRoutingSuppressionRemoveMatch
---@field account_id? string
---@field id string
---@field zone_id? string

---@class EmailRuleResponseSingle
---@field actions table
---@field enabled? boolean
---@field matchers table
---@field name? string
---@field owner_worker_tag? string
---@field priority? number
---@field result? table
---@field source? string

---@class EmailRuleResponseSingleLoadMatch
---@field rule_identifier string
---@field zone_id string

---@class EmailRuleResponseSingleCreateData
---@field zone_id string
---@field actions table
---@field enabled? boolean
---@field matchers table
---@field name? string
---@field owner_worker_tag? string
---@field priority? number
---@field result? table
---@field source? string

---@class EmailRuleResponseSingleUpdateData
---@field rule_identifier string
---@field zone_id string
---@field actions? table
---@field enabled? boolean
---@field matchers? table
---@field name? string
---@field owner_worker_tag? string
---@field priority? number
---@field result? table
---@field source? string

---@class EmailRuleResponseSingleRemoveMatch
---@field rule_identifier string
---@field zone_id string

---@class EmailRulesResponseCollection
---@field result? table
---@field result_info? table

---@class EmailRulesResponseCollectionListMatch
---@field zone_id string
---@field enabled? boolean
---@field page? number
---@field per_page? number

---@class EmailSecurity
---@field action string
---@field comment? string
---@field destination? any
---@field errors table
---@field expected_disposition? string
---@field id string
---@field ids table
---@field messages table
---@field postfix_ids? table
---@field result table
---@field result_info table
---@field search_params table
---@field success boolean

---@class EmailSecurityLoadMatch
---@field account_id string
---@field investigate_id? string
---@field submission? boolean
---@field job_id? string

---@class EmailSecurityListMatch
---@field account_id string
---@field end? string
---@field from_date? string
---@field page? number
---@field per_page? number
---@field start? string
---@field to_date? string
---@field action_type? string
---@field status? string
---@field bulk_id? string
---@field investigate_id? string

---@class EmailSecurityCreateData
---@field account_id string
---@field action string
---@field comment? string
---@field destination? any
---@field errors table
---@field expected_disposition? string
---@field id string
---@field ids table
---@field messages table
---@field postfix_ids? table
---@field result table
---@field result_info table
---@field search_params table
---@field success boolean

---@class EmailSecurityRemoveMatch
---@field account_id string
---@field job_id string

---@class EmailSecurityAnalytics
---@field result? table

---@class EmailSecurityAnalyticsLoadMatch
---@field account_id string

---@class EmailSecuritySetting
---@field allowed_delivery_modes table
---@field comments? string
---@field deletes table
---@field domain string
---@field drop_dispositions table
---@field folder? any
---@field integration_id? string
---@field ip_restrictions table
---@field lookback_hops? number
---@field patches table
---@field pattern string
---@field posts table
---@field puts table
---@field regions table
---@field require_tls_inbound? boolean
---@field require_tls_outbound? boolean
---@field result? table
---@field transport? string

---@class EmailSecuritySettingLoadMatch
---@field account_id string
---@field domain_id? string
---@field impersonation_registry_id? string
---@field pattern_id? string
---@field policy_id? string
---@field sending_domain_restriction_id? string
---@field trusted_domain_id? string

---@class EmailSecuritySettingListMatch
---@field account_id string
---@field direction? string
---@field is_acceptable_sender? boolean
---@field is_exempt_recipient? boolean
---@field is_trusted_sender? boolean
---@field order? string
---@field page? number
---@field pattern? string
---@field pattern_type? any
---@field per_page? number
---@field search? string
---@field verify_sender? boolean
---@field active_delivery_mode? string
---@field allowed_delivery_mode? string
---@field domain? table
---@field integration_id? string
---@field status? string
---@field is_recent? boolean
---@field is_similarity? boolean
---@field enabled? boolean
---@field name? string
---@field provenance? string

---@class EmailSecuritySettingCreateData
---@field account_id string
---@field allowed_delivery_modes table
---@field comments? string
---@field deletes table
---@field domain string
---@field drop_dispositions table
---@field folder? any
---@field integration_id? string
---@field ip_restrictions table
---@field lookback_hops? number
---@field patches table
---@field pattern string
---@field posts table
---@field puts table
---@field regions table
---@field require_tls_inbound? boolean
---@field require_tls_outbound? boolean
---@field result? table
---@field transport? string

---@class EmailSecuritySettingUpdateData
---@field account_id string
---@field domain_id? string
---@field pattern_id? string
---@field policy_id? string
---@field trusted_domain_id? string
---@field allowed_delivery_modes? table
---@field comments? string
---@field deletes? table
---@field domain? string
---@field drop_dispositions? table
---@field folder? any
---@field integration_id? string
---@field ip_restrictions? table
---@field lookback_hops? number
---@field patches? table
---@field pattern? string
---@field posts? table
---@field puts? table
---@field regions? table
---@field require_tls_inbound? boolean
---@field require_tls_outbound? boolean
---@field result? table
---@field transport? string

---@class EmailSecuritySettingRemoveMatch
---@field account_id string
---@field domain_id? string
---@field impersonation_registry_id? string
---@field pattern_id? string
---@field policy_id? string
---@field sending_domain_restriction_id? string
---@field trusted_domain_id? string

---@class EmailSending
---@field attachments? table
---@field bcc? any
---@field cc? any
---@field errors table
---@field from any
---@field headers? table
---@field html? string
---@field messages table
---@field mime_message string
---@field recipients table
---@field reply_to? any
---@field result table
---@field result_info table
---@field subject string
---@field success boolean
---@field text? string
---@field to? any

---@class EmailSendingLoadMatch
---@field account_id string
---@field message_id string

---@class EmailSendingCreateData
---@field account_id string
---@field attachments? table
---@field bcc? any
---@field cc? any
---@field errors table
---@field from any
---@field headers? table
---@field html? string
---@field messages table
---@field mime_message string
---@field recipients table
---@field reply_to? any
---@field result table
---@field result_info table
---@field subject string
---@field success boolean
---@field text? string
---@field to? any

---@class EmailSendingLimitsResponseSingle
---@field result? table

---@class EmailSendingLimitsResponseSingleLoadMatch
---@field account_id string

---@class EmailSendingReputationResponseSingle
---@field result table

---@class EmailSendingReputationResponseSingleLoadMatch
---@field account_id string

---@class EmailSendingSubdomainDnsStatus
---@field result? table

---@class EmailSendingSubdomainDnsStatusLoadMatch
---@field subdomain_id string
---@field zone_id string

---@class EmailSendingSubdomainDnsStatusCreateData
---@field subdomain_id string
---@field zone_id string
---@field result? table

---@class EmailSendingSubdomainPreview
---@field name string

---@class EmailSendingSubdomainPreviewCreateData
---@field zone_id string
---@field name string

---@class EmailSendingSubdomainReputationComplaint
---@field result table

---@class EmailSendingSubdomainReputationComplaintLoadMatch
---@field subdomain_id string
---@field zone_id string
---@field end_at string
---@field start_at string

---@class EmailSendingSubdomainResponseSingle
---@field drop_suppressed_recipients? boolean
---@field name string
---@field preview_enabled? boolean
---@field result? any

---@class EmailSendingSubdomainResponseSingleLoadMatch
---@field subdomain_id string
---@field zone_id string

---@class EmailSendingSubdomainResponseSingleCreateData
---@field zone_id string
---@field drop_suppressed_recipients? boolean
---@field name string
---@field preview_enabled? boolean
---@field result? any

---@class EmailSendingSubdomainResponseSingleUpdateData
---@field subdomain_id string
---@field zone_id string
---@field drop_suppressed_recipients? boolean
---@field name? string
---@field preview_enabled? boolean
---@field result? any

---@class EmailSendingSubdomainsResponseCollection
---@field result? table

---@class EmailSendingSubdomainsResponseCollectionListMatch
---@field zone_id string

---@class EmailSendingSuppression
---@field created_at string
---@field email string
---@field errors table
---@field expires_at string
---@field id string
---@field items table
---@field messages table
---@field note? string
---@field read_only boolean
---@field reason string
---@field result table
---@field scope? any
---@field success boolean
---@field zones? table

---@class EmailSendingSuppressionLoadMatch
---@field account_id? string
---@field id string
---@field zone_id? string

---@class EmailSendingSuppressionListMatch
---@field account_id string
---@field cursor? string
---@field email? string
---@field per_page? number
---@field reason? string
---@field scope_type? string
---@field scope_value? string
---@field search? string

---@class EmailSendingSuppressionCreateData
---@field account_id string
---@field created_at string
---@field email string
---@field errors table
---@field expires_at string
---@field id string
---@field items table
---@field messages table
---@field note? string
---@field read_only boolean
---@field reason string
---@field result table
---@field scope? any
---@field success boolean
---@field zones? table

---@class EmailSendingSuppressionUpdateData
---@field account_id string
---@field id string
---@field created_at? string
---@field email? string
---@field errors? table
---@field expires_at? string
---@field items? table
---@field messages? table
---@field note? string
---@field read_only? boolean
---@field reason? string
---@field result? table
---@field scope? any
---@field success? boolean
---@field zones? table

---@class EmailSendingSuppressionRemoveMatch
---@field account_id? string
---@field id string
---@field zone_id? string

---@class Subdomain
---@field id? string

---@class SubdomainRemoveMatch
---@field id string
---@field zone_id string

local M = {}

return M
