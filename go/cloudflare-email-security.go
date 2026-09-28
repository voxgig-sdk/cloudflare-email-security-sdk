package voxgigcloudflareemailsecuritysdk

import (
	"github.com/voxgig-sdk/cloudflare-email-security-sdk/go/core"
	"github.com/voxgig-sdk/cloudflare-email-security-sdk/go/entity"
	"github.com/voxgig-sdk/cloudflare-email-security-sdk/go/feature"
	_ "github.com/voxgig-sdk/cloudflare-email-security-sdk/go/utility"
)

// Type aliases preserve external API.
type CloudflareEmailSecuritySDK = core.CloudflareEmailSecuritySDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type CloudflareEmailSecurityEntity = core.CloudflareEmailSecurityEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type CloudflareEmailSecurityError = core.CloudflareEmailSecurityError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewEmailAccountRulesPlanResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailAccountRulesPlanResponseSingleEntity(client, entopts)
	}
	core.NewEmailAccountRulesResponseCollectionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailAccountRulesResponseCollectionEntity(client, entopts)
	}
	core.NewEmailAuthEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailAuthEntity(client, entopts)
	}
	core.NewEmailCatchAllRuleResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailCatchAllRuleResponseSingleEntity(client, entopts)
	}
	core.NewEmailDestinationAddressResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailDestinationAddressResponseSingleEntity(client, entopts)
	}
	core.NewEmailDestinationAddressesResponseCollectionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailDestinationAddressesResponseCollectionEntity(client, entopts)
	}
	core.NewEmailDnsSettingsResponseCollectionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailDnsSettingsResponseCollectionEntity(client, entopts)
	}
	core.NewEmailEmailSettingsResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailEmailSettingsResponseSingleEntity(client, entopts)
	}
	core.NewEmailRoutingSuppressionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailRoutingSuppressionEntity(client, entopts)
	}
	core.NewEmailRuleResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailRuleResponseSingleEntity(client, entopts)
	}
	core.NewEmailRulesResponseCollectionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailRulesResponseCollectionEntity(client, entopts)
	}
	core.NewEmailSecurityEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSecurityEntity(client, entopts)
	}
	core.NewEmailSecurityAnalyticsEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSecurityAnalyticsEntity(client, entopts)
	}
	core.NewEmailSecuritySettingEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSecuritySettingEntity(client, entopts)
	}
	core.NewEmailSendingEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingEntity(client, entopts)
	}
	core.NewEmailSendingLimitsResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingLimitsResponseSingleEntity(client, entopts)
	}
	core.NewEmailSendingReputationResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingReputationResponseSingleEntity(client, entopts)
	}
	core.NewEmailSendingSubdomainDnsStatusEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSubdomainDnsStatusEntity(client, entopts)
	}
	core.NewEmailSendingSubdomainPreviewEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSubdomainPreviewEntity(client, entopts)
	}
	core.NewEmailSendingSubdomainReputationComplaintEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSubdomainReputationComplaintEntity(client, entopts)
	}
	core.NewEmailSendingSubdomainResponseSingleEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSubdomainResponseSingleEntity(client, entopts)
	}
	core.NewEmailSendingSubdomainsResponseCollectionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSubdomainsResponseCollectionEntity(client, entopts)
	}
	core.NewEmailSendingSuppressionEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewEmailSendingSuppressionEntity(client, entopts)
	}
	core.NewSubdomainEntityFunc = func(client *core.CloudflareEmailSecuritySDK, entopts map[string]any) core.CloudflareEmailSecurityEntity {
		return entity.NewSubdomainEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewCloudflareEmailSecuritySDK = core.NewCloudflareEmailSecuritySDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewCloudflareEmailSecuritySDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *CloudflareEmailSecuritySDK  { return NewCloudflareEmailSecuritySDK(nil) }
func Test() *CloudflareEmailSecuritySDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
