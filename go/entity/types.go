// Typed models for the CloudflareEmailSecurity SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/cloudflare-email-security-sdk/go/core"
)

// EmailAccountRulesPlanResponseSingle is the typed data model for the email_account_rules_plan_response_single entity.
type EmailAccountRulesPlanResponseSingle struct {
}

// EmailAccountRulesPlanResponseSingleCreateData is the typed request payload for EmailAccountRulesPlanResponseSingle.CreateTyped.
type EmailAccountRulesPlanResponseSingleCreateData struct {
	AccountId string `json:"account_id"`
	CatchAllRules *[]any `json:"catch_all_rules,omitempty"`
	OwnerWorkerTag string `json:"owner_worker_tag"`
	Rules *[]any `json:"rules,omitempty"`
}

// EmailAccountRulesResponseCollection is the typed data model for the email_account_rules_response_collection entity.
type EmailAccountRulesResponseCollection struct {
}

// EmailAccountRulesResponseCollectionListMatch is the typed request payload for EmailAccountRulesResponseCollection.ListTyped.
type EmailAccountRulesResponseCollectionListMatch struct {
	AccountId string `json:"account_id"`
	Enabled *bool `json:"enabled,omitempty"`
	Page *float64 `json:"page,omitempty"`
	PerPage *float64 `json:"per_page,omitempty"`
}

// EmailAuth is the typed data model for the email_auth entity.
type EmailAuth struct {
}

// EmailAuthLoadMatch is the typed request payload for EmailAuth.LoadTyped.
type EmailAuthLoadMatch struct {
	ZoneId string `json:"zone_id"`
	Id *string `json:"id,omitempty"`
}

// EmailAuthUpdateData is the typed request payload for EmailAuth.UpdateTyped.
type EmailAuthUpdateData struct {
	ZoneId string `json:"zone_id"`
	Enabled *bool `json:"enabled,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SkipWizard *bool `json:"skip_wizard,omitempty"`
}

// EmailCatchAllRuleResponseSingle is the typed data model for the email_catch_all_rule_response_single entity.
type EmailCatchAllRuleResponseSingle struct {
}

// EmailCatchAllRuleResponseSingleLoadMatch is the typed request payload for EmailCatchAllRuleResponseSingle.LoadTyped.
type EmailCatchAllRuleResponseSingleLoadMatch struct {
	ZoneId string `json:"zone_id"`
}

// EmailCatchAllRuleResponseSingleUpdateData is the typed request payload for EmailCatchAllRuleResponseSingle.UpdateTyped.
type EmailCatchAllRuleResponseSingleUpdateData struct {
	ZoneId string `json:"zone_id"`
	Actions *[]any `json:"actions,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Matchers *[]any `json:"matchers,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerWorkerTag *string `json:"owner_worker_tag,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Source *string `json:"source,omitempty"`
}

// EmailDestinationAddressResponseSingle is the typed data model for the email_destination_address_response_single entity.
type EmailDestinationAddressResponseSingle struct {
}

// EmailDestinationAddressResponseSingleLoadMatch is the typed request payload for EmailDestinationAddressResponseSingle.LoadTyped.
type EmailDestinationAddressResponseSingleLoadMatch struct {
	AccountId string `json:"account_id"`
	DestinationAddressIdentifier string `json:"destination_address_identifier"`
}

// EmailDestinationAddressResponseSingleCreateData is the typed request payload for EmailDestinationAddressResponseSingle.CreateTyped.
type EmailDestinationAddressResponseSingleCreateData struct {
	AccountId string `json:"account_id"`
	Email string `json:"email"`
	Result *map[string]any `json:"result,omitempty"`
	Status string `json:"status"`
}

// EmailDestinationAddressResponseSingleUpdateData is the typed request payload for EmailDestinationAddressResponseSingle.UpdateTyped.
type EmailDestinationAddressResponseSingleUpdateData struct {
	AccountId string `json:"account_id"`
	DestinationAddressIdentifier string `json:"destination_address_identifier"`
	Email *string `json:"email,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Status *string `json:"status,omitempty"`
}

// EmailDestinationAddressResponseSingleRemoveMatch is the typed request payload for EmailDestinationAddressResponseSingle.RemoveTyped.
type EmailDestinationAddressResponseSingleRemoveMatch struct {
	AccountId string `json:"account_id"`
	DestinationAddressIdentifier string `json:"destination_address_identifier"`
}

// EmailDestinationAddressesResponseCollection is the typed data model for the email_destination_addresses_response_collection entity.
type EmailDestinationAddressesResponseCollection struct {
}

// EmailDestinationAddressesResponseCollectionListMatch is the typed request payload for EmailDestinationAddressesResponseCollection.ListTyped.
type EmailDestinationAddressesResponseCollectionListMatch struct {
	AccountId string `json:"account_id"`
	Direction *string `json:"direction,omitempty"`
	Page *float64 `json:"page,omitempty"`
	PerPage *float64 `json:"per_page,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

// EmailDnsSettingsResponseCollection is the typed data model for the email_dns_settings_response_collection entity.
type EmailDnsSettingsResponseCollection struct {
}

// EmailDnsSettingsResponseCollectionListMatch is the typed request payload for EmailDnsSettingsResponseCollection.ListTyped.
type EmailDnsSettingsResponseCollectionListMatch struct {
	ZoneId string `json:"zone_id"`
	Subdomain *string `json:"subdomain,omitempty"`
	SubdomainId *string `json:"subdomain_id,omitempty"`
}

// EmailEmailSettingsResponseSingle is the typed data model for the email_email_settings_response_single entity.
type EmailEmailSettingsResponseSingle struct {
}

// EmailEmailSettingsResponseSingleLoadMatch is the typed request payload for EmailEmailSettingsResponseSingle.LoadTyped.
type EmailEmailSettingsResponseSingleLoadMatch struct {
	ZoneId string `json:"zone_id"`
}

// EmailEmailSettingsResponseSingleCreateData is the typed request payload for EmailEmailSettingsResponseSingle.CreateTyped.
type EmailEmailSettingsResponseSingleCreateData struct {
	ZoneId string `json:"zone_id"`
	Enabled *bool `json:"enabled,omitempty"`
	Name *string `json:"name,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SkipWizard *bool `json:"skip_wizard,omitempty"`
	SupportSubaddress *bool `json:"support_subaddress,omitempty"`
}

// EmailEmailSettingsResponseSingleUpdateData is the typed request payload for EmailEmailSettingsResponseSingle.UpdateTyped.
type EmailEmailSettingsResponseSingleUpdateData struct {
	ZoneId string `json:"zone_id"`
	Enabled *bool `json:"enabled,omitempty"`
	Name *string `json:"name,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SkipWizard *bool `json:"skip_wizard,omitempty"`
	SupportSubaddress *bool `json:"support_subaddress,omitempty"`
}

// EmailEmailSettingsResponseSingleRemoveMatch is the typed request payload for EmailEmailSettingsResponseSingle.RemoveTyped.
type EmailEmailSettingsResponseSingleRemoveMatch struct {
	ZoneId string `json:"zone_id"`
}

// EmailRoutingSuppression is the typed data model for the email_routing_suppression entity.
type EmailRoutingSuppression struct {
}

// EmailRoutingSuppressionLoadMatch is the typed request payload for EmailRoutingSuppression.LoadTyped.
type EmailRoutingSuppressionLoadMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	ZoneId *string `json:"zone_id,omitempty"`
}

// EmailRoutingSuppressionListMatch is the typed request payload for EmailRoutingSuppression.ListTyped.
type EmailRoutingSuppressionListMatch struct {
	AccountId string `json:"account_id"`
	Direction *string `json:"direction,omitempty"`
	Order *string `json:"order,omitempty"`
	Page *int `json:"page,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
}

// EmailRoutingSuppressionCreateData is the typed request payload for EmailRoutingSuppression.CreateTyped.
type EmailRoutingSuppressionCreateData struct {
	AccountId string `json:"account_id"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email"`
	ExpiresAt string `json:"expires_at"`
	Id string `json:"id"`
	Reason string `json:"reason"`
	Zones *[]any `json:"zones,omitempty"`
}

// EmailRoutingSuppressionRemoveMatch is the typed request payload for EmailRoutingSuppression.RemoveTyped.
type EmailRoutingSuppressionRemoveMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	ZoneId *string `json:"zone_id,omitempty"`
}

// EmailRuleResponseSingle is the typed data model for the email_rule_response_single entity.
type EmailRuleResponseSingle struct {
}

// EmailRuleResponseSingleLoadMatch is the typed request payload for EmailRuleResponseSingle.LoadTyped.
type EmailRuleResponseSingleLoadMatch struct {
	RuleIdentifier string `json:"rule_identifier"`
	ZoneId string `json:"zone_id"`
}

// EmailRuleResponseSingleCreateData is the typed request payload for EmailRuleResponseSingle.CreateTyped.
type EmailRuleResponseSingleCreateData struct {
	ZoneId string `json:"zone_id"`
	Actions []any `json:"actions"`
	Enabled *bool `json:"enabled,omitempty"`
	Matchers []any `json:"matchers"`
	Name *string `json:"name,omitempty"`
	OwnerWorkerTag *string `json:"owner_worker_tag,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Source *string `json:"source,omitempty"`
}

// EmailRuleResponseSingleUpdateData is the typed request payload for EmailRuleResponseSingle.UpdateTyped.
type EmailRuleResponseSingleUpdateData struct {
	RuleIdentifier string `json:"rule_identifier"`
	ZoneId string `json:"zone_id"`
	Actions *[]any `json:"actions,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Matchers *[]any `json:"matchers,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerWorkerTag *string `json:"owner_worker_tag,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Source *string `json:"source,omitempty"`
}

// EmailRuleResponseSingleRemoveMatch is the typed request payload for EmailRuleResponseSingle.RemoveTyped.
type EmailRuleResponseSingleRemoveMatch struct {
	RuleIdentifier string `json:"rule_identifier"`
	ZoneId string `json:"zone_id"`
}

// EmailRulesResponseCollection is the typed data model for the email_rules_response_collection entity.
type EmailRulesResponseCollection struct {
}

// EmailRulesResponseCollectionListMatch is the typed request payload for EmailRulesResponseCollection.ListTyped.
type EmailRulesResponseCollectionListMatch struct {
	ZoneId string `json:"zone_id"`
	Enabled *bool `json:"enabled,omitempty"`
	Page *float64 `json:"page,omitempty"`
	PerPage *float64 `json:"per_page,omitempty"`
}

// EmailSecurity is the typed data model for the email_security entity.
type EmailSecurity struct {
}

// EmailSecurityLoadMatch is the typed request payload for EmailSecurity.LoadTyped.
type EmailSecurityLoadMatch struct {
	AccountId string `json:"account_id"`
	InvestigateId *string `json:"investigate_id,omitempty"`
	Submission *bool `json:"submission,omitempty"`
	JobId *string `json:"job_id,omitempty"`
}

// EmailSecurityListMatch is the typed request payload for EmailSecurity.ListTyped.
type EmailSecurityListMatch struct {
	AccountId string `json:"account_id"`
	End *string `json:"end,omitempty"`
	FromDate *string `json:"from_date,omitempty"`
	Page *int `json:"page,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Start *string `json:"start,omitempty"`
	ToDate *string `json:"to_date,omitempty"`
	ActionType *string `json:"action_type,omitempty"`
	Status *string `json:"status,omitempty"`
	BulkId *string `json:"bulk_id,omitempty"`
	InvestigateId *string `json:"investigate_id,omitempty"`
}

// EmailSecurityCreateData is the typed request payload for EmailSecurity.CreateTyped.
type EmailSecurityCreateData struct {
	AccountId string `json:"account_id"`
	Action string `json:"action"`
	Comment *string `json:"comment,omitempty"`
	Destination *any `json:"destination,omitempty"`
	Errors []any `json:"errors"`
	ExpectedDisposition *string `json:"expected_disposition,omitempty"`
	Id string `json:"id"`
	Ids []any `json:"ids"`
	Messages []any `json:"messages"`
	PostfixIds *[]any `json:"postfix_ids,omitempty"`
	Result map[string]any `json:"result"`
	ResultInfo map[string]any `json:"result_info"`
	SearchParams map[string]any `json:"search_params"`
	Success bool `json:"success"`
}

// EmailSecurityRemoveMatch is the typed request payload for EmailSecurity.RemoveTyped.
type EmailSecurityRemoveMatch struct {
	AccountId string `json:"account_id"`
	JobId string `json:"job_id"`
}

// EmailSecurityAnalytics is the typed data model for the email_security_analytics entity.
type EmailSecurityAnalytics struct {
}

// EmailSecurityAnalyticsLoadMatch is the typed request payload for EmailSecurityAnalytics.LoadTyped.
type EmailSecurityAnalyticsLoadMatch struct {
	AccountId string `json:"account_id"`
}

// EmailSecuritySetting is the typed data model for the email_security_setting entity.
type EmailSecuritySetting struct {
}

// EmailSecuritySettingLoadMatch is the typed request payload for EmailSecuritySetting.LoadTyped.
type EmailSecuritySettingLoadMatch struct {
	AccountId string `json:"account_id"`
	DomainId *string `json:"domain_id,omitempty"`
	ImpersonationRegistryId *string `json:"impersonation_registry_id,omitempty"`
	PatternId *string `json:"pattern_id,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	SendingDomainRestrictionId *string `json:"sending_domain_restriction_id,omitempty"`
	TrustedDomainId *string `json:"trusted_domain_id,omitempty"`
}

// EmailSecuritySettingListMatch is the typed request payload for EmailSecuritySetting.ListTyped.
type EmailSecuritySettingListMatch struct {
	AccountId string `json:"account_id"`
	Direction *string `json:"direction,omitempty"`
	IsAcceptableSender *bool `json:"is_acceptable_sender,omitempty"`
	IsExemptRecipient *bool `json:"is_exempt_recipient,omitempty"`
	IsTrustedSender *bool `json:"is_trusted_sender,omitempty"`
	Order *string `json:"order,omitempty"`
	Page *int `json:"page,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	PatternType *any `json:"pattern_type,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Search *string `json:"search,omitempty"`
	VerifySender *bool `json:"verify_sender,omitempty"`
	ActiveDeliveryMode *string `json:"active_delivery_mode,omitempty"`
	AllowedDeliveryMode *string `json:"allowed_delivery_mode,omitempty"`
	Domain *[]any `json:"domain,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
	Status *string `json:"status,omitempty"`
	IsRecent *bool `json:"is_recent,omitempty"`
	IsSimilarity *bool `json:"is_similarity,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Name *string `json:"name,omitempty"`
	Provenance *string `json:"provenance,omitempty"`
}

// EmailSecuritySettingCreateData is the typed request payload for EmailSecuritySetting.CreateTyped.
type EmailSecuritySettingCreateData struct {
	AccountId string `json:"account_id"`
	AllowedDeliveryModes []any `json:"allowed_delivery_modes"`
	Comments *string `json:"comments,omitempty"`
	Deletes []any `json:"deletes"`
	Domain string `json:"domain"`
	DropDispositions []any `json:"drop_dispositions"`
	Folder *any `json:"folder,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
	IpRestrictions []any `json:"ip_restrictions"`
	LookbackHops *int `json:"lookback_hops,omitempty"`
	Patches []any `json:"patches"`
	Pattern string `json:"pattern"`
	Posts []any `json:"posts"`
	Puts []any `json:"puts"`
	Regions []any `json:"regions"`
	RequireTlsInbound *bool `json:"require_tls_inbound,omitempty"`
	RequireTlsOutbound *bool `json:"require_tls_outbound,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Transport *string `json:"transport,omitempty"`
}

// EmailSecuritySettingUpdateData is the typed request payload for EmailSecuritySetting.UpdateTyped.
type EmailSecuritySettingUpdateData struct {
	AccountId string `json:"account_id"`
	DomainId *string `json:"domain_id,omitempty"`
	PatternId *string `json:"pattern_id,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	TrustedDomainId *string `json:"trusted_domain_id,omitempty"`
	AllowedDeliveryModes *[]any `json:"allowed_delivery_modes,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Deletes *[]any `json:"deletes,omitempty"`
	Domain *string `json:"domain,omitempty"`
	DropDispositions *[]any `json:"drop_dispositions,omitempty"`
	Folder *any `json:"folder,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
	IpRestrictions *[]any `json:"ip_restrictions,omitempty"`
	LookbackHops *int `json:"lookback_hops,omitempty"`
	Patches *[]any `json:"patches,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	Posts *[]any `json:"posts,omitempty"`
	Puts *[]any `json:"puts,omitempty"`
	Regions *[]any `json:"regions,omitempty"`
	RequireTlsInbound *bool `json:"require_tls_inbound,omitempty"`
	RequireTlsOutbound *bool `json:"require_tls_outbound,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Transport *string `json:"transport,omitempty"`
}

// EmailSecuritySettingRemoveMatch is the typed request payload for EmailSecuritySetting.RemoveTyped.
type EmailSecuritySettingRemoveMatch struct {
	AccountId string `json:"account_id"`
	DomainId *string `json:"domain_id,omitempty"`
	ImpersonationRegistryId *string `json:"impersonation_registry_id,omitempty"`
	PatternId *string `json:"pattern_id,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	SendingDomainRestrictionId *string `json:"sending_domain_restriction_id,omitempty"`
	TrustedDomainId *string `json:"trusted_domain_id,omitempty"`
}

// EmailSending is the typed data model for the email_sending entity.
type EmailSending struct {
}

// EmailSendingLoadMatch is the typed request payload for EmailSending.LoadTyped.
type EmailSendingLoadMatch struct {
	AccountId string `json:"account_id"`
	MessageId string `json:"message_id"`
}

// EmailSendingCreateData is the typed request payload for EmailSending.CreateTyped.
type EmailSendingCreateData struct {
	AccountId string `json:"account_id"`
	Attachments *[]any `json:"attachments,omitempty"`
	Bcc *any `json:"bcc,omitempty"`
	Cc *any `json:"cc,omitempty"`
	Errors []any `json:"errors"`
	From any `json:"from"`
	Headers *map[string]any `json:"headers,omitempty"`
	Html *string `json:"html,omitempty"`
	Messages []any `json:"messages"`
	MimeMessage string `json:"mime_message"`
	Recipients []any `json:"recipients"`
	ReplyTo *any `json:"reply_to,omitempty"`
	Result map[string]any `json:"result"`
	ResultInfo map[string]any `json:"result_info"`
	Subject string `json:"subject"`
	Success bool `json:"success"`
	Text *string `json:"text,omitempty"`
	To *any `json:"to,omitempty"`
}

// EmailSendingLimitsResponseSingle is the typed data model for the email_sending_limits_response_single entity.
type EmailSendingLimitsResponseSingle struct {
}

// EmailSendingLimitsResponseSingleLoadMatch is the typed request payload for EmailSendingLimitsResponseSingle.LoadTyped.
type EmailSendingLimitsResponseSingleLoadMatch struct {
	AccountId string `json:"account_id"`
}

// EmailSendingReputationResponseSingle is the typed data model for the email_sending_reputation_response_single entity.
type EmailSendingReputationResponseSingle struct {
}

// EmailSendingReputationResponseSingleLoadMatch is the typed request payload for EmailSendingReputationResponseSingle.LoadTyped.
type EmailSendingReputationResponseSingleLoadMatch struct {
	AccountId string `json:"account_id"`
}

// EmailSendingSubdomainDnsStatus is the typed data model for the email_sending_subdomain_dns_status entity.
type EmailSendingSubdomainDnsStatus struct {
}

// EmailSendingSubdomainDnsStatusLoadMatch is the typed request payload for EmailSendingSubdomainDnsStatus.LoadTyped.
type EmailSendingSubdomainDnsStatusLoadMatch struct {
	SubdomainId string `json:"subdomain_id"`
	ZoneId string `json:"zone_id"`
}

// EmailSendingSubdomainDnsStatusCreateData is the typed request payload for EmailSendingSubdomainDnsStatus.CreateTyped.
type EmailSendingSubdomainDnsStatusCreateData struct {
	SubdomainId string `json:"subdomain_id"`
	ZoneId string `json:"zone_id"`
	Result *map[string]any `json:"result,omitempty"`
}

// EmailSendingSubdomainPreview is the typed data model for the email_sending_subdomain_preview entity.
type EmailSendingSubdomainPreview struct {
}

// EmailSendingSubdomainPreviewCreateData is the typed request payload for EmailSendingSubdomainPreview.CreateTyped.
type EmailSendingSubdomainPreviewCreateData struct {
	ZoneId string `json:"zone_id"`
	Name string `json:"name"`
}

// EmailSendingSubdomainReputationComplaint is the typed data model for the email_sending_subdomain_reputation_complaint entity.
type EmailSendingSubdomainReputationComplaint struct {
}

// EmailSendingSubdomainReputationComplaintLoadMatch is the typed request payload for EmailSendingSubdomainReputationComplaint.LoadTyped.
type EmailSendingSubdomainReputationComplaintLoadMatch struct {
	SubdomainId string `json:"subdomain_id"`
	ZoneId string `json:"zone_id"`
	EndAt string `json:"end_at"`
	StartAt string `json:"start_at"`
}

// EmailSendingSubdomainResponseSingle is the typed data model for the email_sending_subdomain_response_single entity.
type EmailSendingSubdomainResponseSingle struct {
}

// EmailSendingSubdomainResponseSingleLoadMatch is the typed request payload for EmailSendingSubdomainResponseSingle.LoadTyped.
type EmailSendingSubdomainResponseSingleLoadMatch struct {
	SubdomainId string `json:"subdomain_id"`
	ZoneId string `json:"zone_id"`
}

// EmailSendingSubdomainResponseSingleCreateData is the typed request payload for EmailSendingSubdomainResponseSingle.CreateTyped.
type EmailSendingSubdomainResponseSingleCreateData struct {
	ZoneId string `json:"zone_id"`
	DropSuppressedRecipients *bool `json:"drop_suppressed_recipients,omitempty"`
	Name string `json:"name"`
	PreviewEnabled *bool `json:"preview_enabled,omitempty"`
	Result *any `json:"result,omitempty"`
}

// EmailSendingSubdomainResponseSingleUpdateData is the typed request payload for EmailSendingSubdomainResponseSingle.UpdateTyped.
type EmailSendingSubdomainResponseSingleUpdateData struct {
	SubdomainId string `json:"subdomain_id"`
	ZoneId string `json:"zone_id"`
	DropSuppressedRecipients *bool `json:"drop_suppressed_recipients,omitempty"`
	Name *string `json:"name,omitempty"`
	PreviewEnabled *bool `json:"preview_enabled,omitempty"`
	Result *any `json:"result,omitempty"`
}

// EmailSendingSubdomainsResponseCollection is the typed data model for the email_sending_subdomains_response_collection entity.
type EmailSendingSubdomainsResponseCollection struct {
}

// EmailSendingSubdomainsResponseCollectionListMatch is the typed request payload for EmailSendingSubdomainsResponseCollection.ListTyped.
type EmailSendingSubdomainsResponseCollectionListMatch struct {
	ZoneId string `json:"zone_id"`
}

// EmailSendingSuppression is the typed data model for the email_sending_suppression entity.
type EmailSendingSuppression struct {
}

// EmailSendingSuppressionLoadMatch is the typed request payload for EmailSendingSuppression.LoadTyped.
type EmailSendingSuppressionLoadMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	ZoneId *string `json:"zone_id,omitempty"`
}

// EmailSendingSuppressionListMatch is the typed request payload for EmailSendingSuppression.ListTyped.
type EmailSendingSuppressionListMatch struct {
	AccountId string `json:"account_id"`
	Cursor *string `json:"cursor,omitempty"`
	Email *string `json:"email,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScopeType *string `json:"scope_type,omitempty"`
	ScopeValue *string `json:"scope_value,omitempty"`
	Search *string `json:"search,omitempty"`
}

// EmailSendingSuppressionCreateData is the typed request payload for EmailSendingSuppression.CreateTyped.
type EmailSendingSuppressionCreateData struct {
	AccountId string `json:"account_id"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email"`
	Errors []any `json:"errors"`
	ExpiresAt string `json:"expires_at"`
	Id string `json:"id"`
	Items []any `json:"items"`
	Messages []any `json:"messages"`
	Note *string `json:"note,omitempty"`
	ReadOnly bool `json:"read_only"`
	Reason string `json:"reason"`
	Result map[string]any `json:"result"`
	Scope *any `json:"scope,omitempty"`
	Success bool `json:"success"`
	Zones *[]any `json:"zones,omitempty"`
}

// EmailSendingSuppressionUpdateData is the typed request payload for EmailSendingSuppression.UpdateTyped.
type EmailSendingSuppressionUpdateData struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Errors *[]any `json:"errors,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	Items *[]any `json:"items,omitempty"`
	Messages *[]any `json:"messages,omitempty"`
	Note *string `json:"note,omitempty"`
	ReadOnly *bool `json:"read_only,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Scope *any `json:"scope,omitempty"`
	Success *bool `json:"success,omitempty"`
	Zones *[]any `json:"zones,omitempty"`
}

// EmailSendingSuppressionRemoveMatch is the typed request payload for EmailSendingSuppression.RemoveTyped.
type EmailSendingSuppressionRemoveMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	ZoneId *string `json:"zone_id,omitempty"`
}

// Subdomain is the typed data model for the subdomain entity.
type Subdomain struct {
}

// SubdomainRemoveMatch is the typed request payload for Subdomain.RemoveTyped.
type SubdomainRemoveMatch struct {
	Id string `json:"id"`
	ZoneId string `json:"zone_id"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
