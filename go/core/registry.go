package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewEmailAccountRulesPlanResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailAccountRulesResponseCollectionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailAuthEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailCatchAllRuleResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailDestinationAddressResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailDestinationAddressesResponseCollectionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailDnsSettingsResponseCollectionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailEmailSettingsResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailRoutingSuppressionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailRuleResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailRulesResponseCollectionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSecurityEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSecurityAnalyticsEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSecuritySettingEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingLimitsResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingReputationResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSubdomainDnsStatusEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSubdomainPreviewEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSubdomainReputationComplaintEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSubdomainResponseSingleEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSubdomainsResponseCollectionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewEmailSendingSuppressionEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

var NewSubdomainEntityFunc func(client *CloudflareEmailSecuritySDK, entopts map[string]any) CloudflareEmailSecurityEntity

