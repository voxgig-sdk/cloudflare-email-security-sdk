# frozen_string_literal: true

# Typed models for the CloudflareEmailSecurity SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# EmailAccountRulesPlanResponseSingle entity data model.
#
# @!attribute [rw] catch_all_rules
#   @return [Array, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String]
#
# @!attribute [rw] rules
#   @return [Array, nil]
EmailAccountRulesPlanResponseSingle = Struct.new(
  :catch_all_rules,
  :owner_worker_tag,
  :rules,
  keyword_init: true
)

# Request payload for EmailAccountRulesPlanResponseSingle#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] catch_all_rules
#   @return [Array, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String]
#
# @!attribute [rw] rules
#   @return [Array, nil]
EmailAccountRulesPlanResponseSingleCreateData = Struct.new(
  :account_id,
  :catch_all_rules,
  :owner_worker_tag,
  :rules,
  keyword_init: true
)

# EmailAccountRulesResponseCollection entity data model.
#
# @!attribute [rw] result
#   @return [Array, nil]
#
# @!attribute [rw] result_info
#   @return [Hash, nil]
EmailAccountRulesResponseCollection = Struct.new(
  :result,
  :result_info,
  keyword_init: true
)

# Request payload for EmailAccountRulesResponseCollection#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] page
#   @return [Float, nil]
#
# @!attribute [rw] per_page
#   @return [Float, nil]
EmailAccountRulesResponseCollectionListMatch = Struct.new(
  :account_id,
  :enabled,
  :page,
  :per_page,
  keyword_init: true
)

# EmailAuth entity data model.
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] skip_wizard
#   @return [Boolean, nil]
EmailAuth = Struct.new(
  :enabled,
  :result,
  :skip_wizard,
  keyword_init: true
)

# Request payload for EmailAuth#load.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
EmailAuthLoadMatch = Struct.new(
  :zone_id,
  :id,
  keyword_init: true
)

# Request payload for EmailAuth#update.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] skip_wizard
#   @return [Boolean, nil]
EmailAuthUpdateData = Struct.new(
  :zone_id,
  :enabled,
  :result,
  :skip_wizard,
  keyword_init: true
)

# EmailCatchAllRuleResponseSingle entity data model.
#
# @!attribute [rw] actions
#   @return [Array]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] matchers
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
EmailCatchAllRuleResponseSingle = Struct.new(
  :actions,
  :enabled,
  :matchers,
  :name,
  :owner_worker_tag,
  :result,
  :source,
  keyword_init: true
)

# Request payload for EmailCatchAllRuleResponseSingle#load.
#
# @!attribute [rw] zone_id
#   @return [String]
EmailCatchAllRuleResponseSingleLoadMatch = Struct.new(
  :zone_id,
  keyword_init: true
)

# Request payload for EmailCatchAllRuleResponseSingle#update.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] actions
#   @return [Array, nil]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] matchers
#   @return [Array, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
EmailCatchAllRuleResponseSingleUpdateData = Struct.new(
  :zone_id,
  :actions,
  :enabled,
  :matchers,
  :name,
  :owner_worker_tag,
  :result,
  :source,
  keyword_init: true
)

# EmailDestinationAddressResponseSingle entity data model.
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String]
EmailDestinationAddressResponseSingle = Struct.new(
  :email,
  :result,
  :status,
  keyword_init: true
)

# Request payload for EmailDestinationAddressResponseSingle#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] destination_address_identifier
#   @return [String]
EmailDestinationAddressResponseSingleLoadMatch = Struct.new(
  :account_id,
  :destination_address_identifier,
  keyword_init: true
)

# Request payload for EmailDestinationAddressResponseSingle#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String]
EmailDestinationAddressResponseSingleCreateData = Struct.new(
  :account_id,
  :email,
  :result,
  :status,
  keyword_init: true
)

# Request payload for EmailDestinationAddressResponseSingle#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] destination_address_identifier
#   @return [String]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
EmailDestinationAddressResponseSingleUpdateData = Struct.new(
  :account_id,
  :destination_address_identifier,
  :email,
  :result,
  :status,
  keyword_init: true
)

# Request payload for EmailDestinationAddressResponseSingle#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] destination_address_identifier
#   @return [String]
EmailDestinationAddressResponseSingleRemoveMatch = Struct.new(
  :account_id,
  :destination_address_identifier,
  keyword_init: true
)

# EmailDestinationAddressesResponseCollection entity data model.
#
# @!attribute [rw] result
#   @return [Array, nil]
#
# @!attribute [rw] result_info
#   @return [Hash, nil]
EmailDestinationAddressesResponseCollection = Struct.new(
  :result,
  :result_info,
  keyword_init: true
)

# Request payload for EmailDestinationAddressesResponseCollection#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] direction
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Float, nil]
#
# @!attribute [rw] per_page
#   @return [Float, nil]
#
# @!attribute [rw] verified
#   @return [Boolean, nil]
EmailDestinationAddressesResponseCollectionListMatch = Struct.new(
  :account_id,
  :direction,
  :page,
  :per_page,
  :verified,
  keyword_init: true
)

# EmailDnsSettingsResponseCollection entity data model.
#
# @!attribute [rw] result
#   @return [Array, nil]
EmailDnsSettingsResponseCollection = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailDnsSettingsResponseCollection#list.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] subdomain
#   @return [String, nil]
#
# @!attribute [rw] subdomain_id
#   @return [String, nil]
EmailDnsSettingsResponseCollectionListMatch = Struct.new(
  :zone_id,
  :subdomain,
  :subdomain_id,
  keyword_init: true
)

# EmailEmailSettingsResponseSingle entity data model.
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] skip_wizard
#   @return [Boolean, nil]
#
# @!attribute [rw] support_subaddress
#   @return [Boolean, nil]
EmailEmailSettingsResponseSingle = Struct.new(
  :enabled,
  :name,
  :result,
  :skip_wizard,
  :support_subaddress,
  keyword_init: true
)

# Request payload for EmailEmailSettingsResponseSingle#load.
#
# @!attribute [rw] zone_id
#   @return [String]
EmailEmailSettingsResponseSingleLoadMatch = Struct.new(
  :zone_id,
  keyword_init: true
)

# Request payload for EmailEmailSettingsResponseSingle#create.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] skip_wizard
#   @return [Boolean, nil]
#
# @!attribute [rw] support_subaddress
#   @return [Boolean, nil]
EmailEmailSettingsResponseSingleCreateData = Struct.new(
  :zone_id,
  :enabled,
  :name,
  :result,
  :skip_wizard,
  :support_subaddress,
  keyword_init: true
)

# Request payload for EmailEmailSettingsResponseSingle#update.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] skip_wizard
#   @return [Boolean, nil]
#
# @!attribute [rw] support_subaddress
#   @return [Boolean, nil]
EmailEmailSettingsResponseSingleUpdateData = Struct.new(
  :zone_id,
  :enabled,
  :name,
  :result,
  :skip_wizard,
  :support_subaddress,
  keyword_init: true
)

# Request payload for EmailEmailSettingsResponseSingle#remove.
#
# @!attribute [rw] zone_id
#   @return [String]
EmailEmailSettingsResponseSingleRemoveMatch = Struct.new(
  :zone_id,
  keyword_init: true
)

# EmailRoutingSuppression entity data model.
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] expires_at
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] reason
#   @return [String]
#
# @!attribute [rw] zones
#   @return [Array, nil]
EmailRoutingSuppression = Struct.new(
  :created_at,
  :email,
  :expires_at,
  :id,
  :reason,
  :zones,
  keyword_init: true
)

# Request payload for EmailRoutingSuppression#load.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String, nil]
EmailRoutingSuppressionLoadMatch = Struct.new(
  :account_id,
  :id,
  :zone_id,
  keyword_init: true
)

# Request payload for EmailRoutingSuppression#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] direction
#   @return [String, nil]
#
# @!attribute [rw] order
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
EmailRoutingSuppressionListMatch = Struct.new(
  :account_id,
  :direction,
  :order,
  :page,
  :per_page,
  keyword_init: true
)

# Request payload for EmailRoutingSuppression#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] expires_at
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] reason
#   @return [String]
#
# @!attribute [rw] zones
#   @return [Array, nil]
EmailRoutingSuppressionCreateData = Struct.new(
  :account_id,
  :created_at,
  :email,
  :expires_at,
  :id,
  :reason,
  :zones,
  keyword_init: true
)

# Request payload for EmailRoutingSuppression#remove.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String, nil]
EmailRoutingSuppressionRemoveMatch = Struct.new(
  :account_id,
  :id,
  :zone_id,
  keyword_init: true
)

# EmailRuleResponseSingle entity data model.
#
# @!attribute [rw] actions
#   @return [Array]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] matchers
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String, nil]
#
# @!attribute [rw] priority
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
EmailRuleResponseSingle = Struct.new(
  :actions,
  :enabled,
  :matchers,
  :name,
  :owner_worker_tag,
  :priority,
  :result,
  :source,
  keyword_init: true
)

# Request payload for EmailRuleResponseSingle#load.
#
# @!attribute [rw] rule_identifier
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
EmailRuleResponseSingleLoadMatch = Struct.new(
  :rule_identifier,
  :zone_id,
  keyword_init: true
)

# Request payload for EmailRuleResponseSingle#create.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] actions
#   @return [Array]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] matchers
#   @return [Array]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String, nil]
#
# @!attribute [rw] priority
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
EmailRuleResponseSingleCreateData = Struct.new(
  :zone_id,
  :actions,
  :enabled,
  :matchers,
  :name,
  :owner_worker_tag,
  :priority,
  :result,
  :source,
  keyword_init: true
)

# Request payload for EmailRuleResponseSingle#update.
#
# @!attribute [rw] rule_identifier
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] actions
#   @return [Array, nil]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] matchers
#   @return [Array, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] owner_worker_tag
#   @return [String, nil]
#
# @!attribute [rw] priority
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
EmailRuleResponseSingleUpdateData = Struct.new(
  :rule_identifier,
  :zone_id,
  :actions,
  :enabled,
  :matchers,
  :name,
  :owner_worker_tag,
  :priority,
  :result,
  :source,
  keyword_init: true
)

# Request payload for EmailRuleResponseSingle#remove.
#
# @!attribute [rw] rule_identifier
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
EmailRuleResponseSingleRemoveMatch = Struct.new(
  :rule_identifier,
  :zone_id,
  keyword_init: true
)

# EmailRulesResponseCollection entity data model.
#
# @!attribute [rw] result
#   @return [Array, nil]
#
# @!attribute [rw] result_info
#   @return [Hash, nil]
EmailRulesResponseCollection = Struct.new(
  :result,
  :result_info,
  keyword_init: true
)

# Request payload for EmailRulesResponseCollection#list.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] page
#   @return [Float, nil]
#
# @!attribute [rw] per_page
#   @return [Float, nil]
EmailRulesResponseCollectionListMatch = Struct.new(
  :zone_id,
  :enabled,
  :page,
  :per_page,
  keyword_init: true
)

# EmailSecurity entity data model.
#
# @!attribute [rw] action
#   @return [String]
#
# @!attribute [rw] comment
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Object, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] expected_disposition
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] ids
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] postfix_ids
#   @return [Array, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] result_info
#   @return [Hash]
#
# @!attribute [rw] search_params
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
EmailSecurity = Struct.new(
  :action,
  :comment,
  :destination,
  :errors,
  :expected_disposition,
  :id,
  :ids,
  :messages,
  :postfix_ids,
  :result,
  :result_info,
  :search_params,
  :success,
  keyword_init: true
)

# Request payload for EmailSecurity#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] investigate_id
#   @return [String, nil]
#
# @!attribute [rw] submission
#   @return [Boolean, nil]
#
# @!attribute [rw] job_id
#   @return [String, nil]
EmailSecurityLoadMatch = Struct.new(
  :account_id,
  :investigate_id,
  :submission,
  :job_id,
  keyword_init: true
)

# Request payload for EmailSecurity#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] end
#   @return [String, nil]
#
# @!attribute [rw] from_date
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] start
#   @return [String, nil]
#
# @!attribute [rw] to_date
#   @return [String, nil]
#
# @!attribute [rw] action_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] bulk_id
#   @return [String, nil]
#
# @!attribute [rw] investigate_id
#   @return [String, nil]
EmailSecurityListMatch = Struct.new(
  :account_id,
  :end,
  :from_date,
  :page,
  :per_page,
  :start,
  :to_date,
  :action_type,
  :status,
  :bulk_id,
  :investigate_id,
  keyword_init: true
)

# Request payload for EmailSecurity#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] action
#   @return [String]
#
# @!attribute [rw] comment
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Object, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] expected_disposition
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] ids
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] postfix_ids
#   @return [Array, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] result_info
#   @return [Hash]
#
# @!attribute [rw] search_params
#   @return [Hash]
#
# @!attribute [rw] success
#   @return [Boolean]
EmailSecurityCreateData = Struct.new(
  :account_id,
  :action,
  :comment,
  :destination,
  :errors,
  :expected_disposition,
  :id,
  :ids,
  :messages,
  :postfix_ids,
  :result,
  :result_info,
  :search_params,
  :success,
  keyword_init: true
)

# Request payload for EmailSecurity#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] job_id
#   @return [String]
EmailSecurityRemoveMatch = Struct.new(
  :account_id,
  :job_id,
  keyword_init: true
)

# EmailSecurityAnalytics entity data model.
#
# @!attribute [rw] result
#   @return [Hash, nil]
EmailSecurityAnalytics = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSecurityAnalytics#load.
#
# @!attribute [rw] account_id
#   @return [String]
EmailSecurityAnalyticsLoadMatch = Struct.new(
  :account_id,
  keyword_init: true
)

# EmailSecuritySetting entity data model.
#
# @!attribute [rw] allowed_delivery_modes
#   @return [Array]
#
# @!attribute [rw] comments
#   @return [String, nil]
#
# @!attribute [rw] deletes
#   @return [Array]
#
# @!attribute [rw] domain
#   @return [String]
#
# @!attribute [rw] drop_dispositions
#   @return [Array]
#
# @!attribute [rw] folder
#   @return [Object, nil]
#
# @!attribute [rw] integration_id
#   @return [String, nil]
#
# @!attribute [rw] ip_restrictions
#   @return [Array]
#
# @!attribute [rw] lookback_hops
#   @return [Integer, nil]
#
# @!attribute [rw] patches
#   @return [Array]
#
# @!attribute [rw] pattern
#   @return [String]
#
# @!attribute [rw] posts
#   @return [Array]
#
# @!attribute [rw] puts
#   @return [Array]
#
# @!attribute [rw] regions
#   @return [Array]
#
# @!attribute [rw] require_tls_inbound
#   @return [Boolean, nil]
#
# @!attribute [rw] require_tls_outbound
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] transport
#   @return [String, nil]
EmailSecuritySetting = Struct.new(
  :allowed_delivery_modes,
  :comments,
  :deletes,
  :domain,
  :drop_dispositions,
  :folder,
  :integration_id,
  :ip_restrictions,
  :lookback_hops,
  :patches,
  :pattern,
  :posts,
  :puts,
  :regions,
  :require_tls_inbound,
  :require_tls_outbound,
  :result,
  :transport,
  keyword_init: true
)

# Request payload for EmailSecuritySetting#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] domain_id
#   @return [String, nil]
#
# @!attribute [rw] impersonation_registry_id
#   @return [String, nil]
#
# @!attribute [rw] pattern_id
#   @return [String, nil]
#
# @!attribute [rw] policy_id
#   @return [String, nil]
#
# @!attribute [rw] sending_domain_restriction_id
#   @return [String, nil]
#
# @!attribute [rw] trusted_domain_id
#   @return [String, nil]
EmailSecuritySettingLoadMatch = Struct.new(
  :account_id,
  :domain_id,
  :impersonation_registry_id,
  :pattern_id,
  :policy_id,
  :sending_domain_restriction_id,
  :trusted_domain_id,
  keyword_init: true
)

# Request payload for EmailSecuritySetting#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] direction
#   @return [String, nil]
#
# @!attribute [rw] is_acceptable_sender
#   @return [Boolean, nil]
#
# @!attribute [rw] is_exempt_recipient
#   @return [Boolean, nil]
#
# @!attribute [rw] is_trusted_sender
#   @return [Boolean, nil]
#
# @!attribute [rw] order
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] pattern
#   @return [String, nil]
#
# @!attribute [rw] pattern_type
#   @return [Object, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] search
#   @return [String, nil]
#
# @!attribute [rw] verify_sender
#   @return [Boolean, nil]
#
# @!attribute [rw] active_delivery_mode
#   @return [String, nil]
#
# @!attribute [rw] allowed_delivery_mode
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [Array, nil]
#
# @!attribute [rw] integration_id
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] is_recent
#   @return [Boolean, nil]
#
# @!attribute [rw] is_similarity
#   @return [Boolean, nil]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [String, nil]
EmailSecuritySettingListMatch = Struct.new(
  :account_id,
  :direction,
  :is_acceptable_sender,
  :is_exempt_recipient,
  :is_trusted_sender,
  :order,
  :page,
  :pattern,
  :pattern_type,
  :per_page,
  :search,
  :verify_sender,
  :active_delivery_mode,
  :allowed_delivery_mode,
  :domain,
  :integration_id,
  :status,
  :is_recent,
  :is_similarity,
  :enabled,
  :name,
  :provenance,
  keyword_init: true
)

# Request payload for EmailSecuritySetting#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] allowed_delivery_modes
#   @return [Array]
#
# @!attribute [rw] comments
#   @return [String, nil]
#
# @!attribute [rw] deletes
#   @return [Array]
#
# @!attribute [rw] domain
#   @return [String]
#
# @!attribute [rw] drop_dispositions
#   @return [Array]
#
# @!attribute [rw] folder
#   @return [Object, nil]
#
# @!attribute [rw] integration_id
#   @return [String, nil]
#
# @!attribute [rw] ip_restrictions
#   @return [Array]
#
# @!attribute [rw] lookback_hops
#   @return [Integer, nil]
#
# @!attribute [rw] patches
#   @return [Array]
#
# @!attribute [rw] pattern
#   @return [String]
#
# @!attribute [rw] posts
#   @return [Array]
#
# @!attribute [rw] puts
#   @return [Array]
#
# @!attribute [rw] regions
#   @return [Array]
#
# @!attribute [rw] require_tls_inbound
#   @return [Boolean, nil]
#
# @!attribute [rw] require_tls_outbound
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] transport
#   @return [String, nil]
EmailSecuritySettingCreateData = Struct.new(
  :account_id,
  :allowed_delivery_modes,
  :comments,
  :deletes,
  :domain,
  :drop_dispositions,
  :folder,
  :integration_id,
  :ip_restrictions,
  :lookback_hops,
  :patches,
  :pattern,
  :posts,
  :puts,
  :regions,
  :require_tls_inbound,
  :require_tls_outbound,
  :result,
  :transport,
  keyword_init: true
)

# Request payload for EmailSecuritySetting#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] domain_id
#   @return [String, nil]
#
# @!attribute [rw] pattern_id
#   @return [String, nil]
#
# @!attribute [rw] policy_id
#   @return [String, nil]
#
# @!attribute [rw] trusted_domain_id
#   @return [String, nil]
#
# @!attribute [rw] allowed_delivery_modes
#   @return [Array, nil]
#
# @!attribute [rw] comments
#   @return [String, nil]
#
# @!attribute [rw] deletes
#   @return [Array, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] drop_dispositions
#   @return [Array, nil]
#
# @!attribute [rw] folder
#   @return [Object, nil]
#
# @!attribute [rw] integration_id
#   @return [String, nil]
#
# @!attribute [rw] ip_restrictions
#   @return [Array, nil]
#
# @!attribute [rw] lookback_hops
#   @return [Integer, nil]
#
# @!attribute [rw] patches
#   @return [Array, nil]
#
# @!attribute [rw] pattern
#   @return [String, nil]
#
# @!attribute [rw] posts
#   @return [Array, nil]
#
# @!attribute [rw] puts
#   @return [Array, nil]
#
# @!attribute [rw] regions
#   @return [Array, nil]
#
# @!attribute [rw] require_tls_inbound
#   @return [Boolean, nil]
#
# @!attribute [rw] require_tls_outbound
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] transport
#   @return [String, nil]
EmailSecuritySettingUpdateData = Struct.new(
  :account_id,
  :domain_id,
  :pattern_id,
  :policy_id,
  :trusted_domain_id,
  :allowed_delivery_modes,
  :comments,
  :deletes,
  :domain,
  :drop_dispositions,
  :folder,
  :integration_id,
  :ip_restrictions,
  :lookback_hops,
  :patches,
  :pattern,
  :posts,
  :puts,
  :regions,
  :require_tls_inbound,
  :require_tls_outbound,
  :result,
  :transport,
  keyword_init: true
)

# Request payload for EmailSecuritySetting#remove.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] domain_id
#   @return [String, nil]
#
# @!attribute [rw] impersonation_registry_id
#   @return [String, nil]
#
# @!attribute [rw] pattern_id
#   @return [String, nil]
#
# @!attribute [rw] policy_id
#   @return [String, nil]
#
# @!attribute [rw] sending_domain_restriction_id
#   @return [String, nil]
#
# @!attribute [rw] trusted_domain_id
#   @return [String, nil]
EmailSecuritySettingRemoveMatch = Struct.new(
  :account_id,
  :domain_id,
  :impersonation_registry_id,
  :pattern_id,
  :policy_id,
  :sending_domain_restriction_id,
  :trusted_domain_id,
  keyword_init: true
)

# EmailSending entity data model.
#
# @!attribute [rw] attachments
#   @return [Array, nil]
#
# @!attribute [rw] bcc
#   @return [Object, nil]
#
# @!attribute [rw] cc
#   @return [Object, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] from
#   @return [Object]
#
# @!attribute [rw] headers
#   @return [Hash, nil]
#
# @!attribute [rw] html
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] mime_message
#   @return [String]
#
# @!attribute [rw] recipients
#   @return [Array]
#
# @!attribute [rw] reply_to
#   @return [Object, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] result_info
#   @return [Hash]
#
# @!attribute [rw] subject
#   @return [String]
#
# @!attribute [rw] success
#   @return [Boolean]
#
# @!attribute [rw] text
#   @return [String, nil]
#
# @!attribute [rw] to
#   @return [Object, nil]
EmailSending = Struct.new(
  :attachments,
  :bcc,
  :cc,
  :errors,
  :from,
  :headers,
  :html,
  :messages,
  :mime_message,
  :recipients,
  :reply_to,
  :result,
  :result_info,
  :subject,
  :success,
  :text,
  :to,
  keyword_init: true
)

# Request payload for EmailSending#load.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] message_id
#   @return [String]
EmailSendingLoadMatch = Struct.new(
  :account_id,
  :message_id,
  keyword_init: true
)

# Request payload for EmailSending#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] attachments
#   @return [Array, nil]
#
# @!attribute [rw] bcc
#   @return [Object, nil]
#
# @!attribute [rw] cc
#   @return [Object, nil]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] from
#   @return [Object]
#
# @!attribute [rw] headers
#   @return [Hash, nil]
#
# @!attribute [rw] html
#   @return [String, nil]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] mime_message
#   @return [String]
#
# @!attribute [rw] recipients
#   @return [Array]
#
# @!attribute [rw] reply_to
#   @return [Object, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] result_info
#   @return [Hash]
#
# @!attribute [rw] subject
#   @return [String]
#
# @!attribute [rw] success
#   @return [Boolean]
#
# @!attribute [rw] text
#   @return [String, nil]
#
# @!attribute [rw] to
#   @return [Object, nil]
EmailSendingCreateData = Struct.new(
  :account_id,
  :attachments,
  :bcc,
  :cc,
  :errors,
  :from,
  :headers,
  :html,
  :messages,
  :mime_message,
  :recipients,
  :reply_to,
  :result,
  :result_info,
  :subject,
  :success,
  :text,
  :to,
  keyword_init: true
)

# EmailSendingLimitsResponseSingle entity data model.
#
# @!attribute [rw] result
#   @return [Hash, nil]
EmailSendingLimitsResponseSingle = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSendingLimitsResponseSingle#load.
#
# @!attribute [rw] account_id
#   @return [String]
EmailSendingLimitsResponseSingleLoadMatch = Struct.new(
  :account_id,
  keyword_init: true
)

# EmailSendingReputationResponseSingle entity data model.
#
# @!attribute [rw] result
#   @return [Hash]
EmailSendingReputationResponseSingle = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSendingReputationResponseSingle#load.
#
# @!attribute [rw] account_id
#   @return [String]
EmailSendingReputationResponseSingleLoadMatch = Struct.new(
  :account_id,
  keyword_init: true
)

# EmailSendingSubdomainDnsStatus entity data model.
#
# @!attribute [rw] result
#   @return [Hash, nil]
EmailSendingSubdomainDnsStatus = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainDnsStatus#load.
#
# @!attribute [rw] subdomain_id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
EmailSendingSubdomainDnsStatusLoadMatch = Struct.new(
  :subdomain_id,
  :zone_id,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainDnsStatus#create.
#
# @!attribute [rw] subdomain_id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash, nil]
EmailSendingSubdomainDnsStatusCreateData = Struct.new(
  :subdomain_id,
  :zone_id,
  :result,
  keyword_init: true
)

# EmailSendingSubdomainPreview entity data model.
#
# @!attribute [rw] name
#   @return [String]
EmailSendingSubdomainPreview = Struct.new(
  :name,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainPreview#create.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] name
#   @return [String]
EmailSendingSubdomainPreviewCreateData = Struct.new(
  :zone_id,
  :name,
  keyword_init: true
)

# EmailSendingSubdomainReputationComplaint entity data model.
#
# @!attribute [rw] result
#   @return [Hash]
EmailSendingSubdomainReputationComplaint = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainReputationComplaint#load.
#
# @!attribute [rw] subdomain_id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] end_at
#   @return [String]
#
# @!attribute [rw] start_at
#   @return [String]
EmailSendingSubdomainReputationComplaintLoadMatch = Struct.new(
  :subdomain_id,
  :zone_id,
  :end_at,
  :start_at,
  keyword_init: true
)

# EmailSendingSubdomainResponseSingle entity data model.
#
# @!attribute [rw] drop_suppressed_recipients
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] preview_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Object, nil]
EmailSendingSubdomainResponseSingle = Struct.new(
  :drop_suppressed_recipients,
  :name,
  :preview_enabled,
  :result,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainResponseSingle#load.
#
# @!attribute [rw] subdomain_id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
EmailSendingSubdomainResponseSingleLoadMatch = Struct.new(
  :subdomain_id,
  :zone_id,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainResponseSingle#create.
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] drop_suppressed_recipients
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] preview_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Object, nil]
EmailSendingSubdomainResponseSingleCreateData = Struct.new(
  :zone_id,
  :drop_suppressed_recipients,
  :name,
  :preview_enabled,
  :result,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainResponseSingle#update.
#
# @!attribute [rw] subdomain_id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
#
# @!attribute [rw] drop_suppressed_recipients
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] preview_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Object, nil]
EmailSendingSubdomainResponseSingleUpdateData = Struct.new(
  :subdomain_id,
  :zone_id,
  :drop_suppressed_recipients,
  :name,
  :preview_enabled,
  :result,
  keyword_init: true
)

# EmailSendingSubdomainsResponseCollection entity data model.
#
# @!attribute [rw] result
#   @return [Array, nil]
EmailSendingSubdomainsResponseCollection = Struct.new(
  :result,
  keyword_init: true
)

# Request payload for EmailSendingSubdomainsResponseCollection#list.
#
# @!attribute [rw] zone_id
#   @return [String]
EmailSendingSubdomainsResponseCollectionListMatch = Struct.new(
  :zone_id,
  keyword_init: true
)

# EmailSendingSuppression entity data model.
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] expires_at
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] items
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] read_only
#   @return [Boolean]
#
# @!attribute [rw] reason
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] scope
#   @return [Object, nil]
#
# @!attribute [rw] success
#   @return [Boolean]
#
# @!attribute [rw] zones
#   @return [Array, nil]
EmailSendingSuppression = Struct.new(
  :created_at,
  :email,
  :errors,
  :expires_at,
  :id,
  :items,
  :messages,
  :note,
  :read_only,
  :reason,
  :result,
  :scope,
  :success,
  :zones,
  keyword_init: true
)

# Request payload for EmailSendingSuppression#load.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String, nil]
EmailSendingSuppressionLoadMatch = Struct.new(
  :account_id,
  :id,
  :zone_id,
  keyword_init: true
)

# Request payload for EmailSendingSuppression#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] cursor
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] scope_type
#   @return [String, nil]
#
# @!attribute [rw] scope_value
#   @return [String, nil]
#
# @!attribute [rw] search
#   @return [String, nil]
EmailSendingSuppressionListMatch = Struct.new(
  :account_id,
  :cursor,
  :email,
  :per_page,
  :reason,
  :scope_type,
  :scope_value,
  :search,
  keyword_init: true
)

# Request payload for EmailSendingSuppression#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] errors
#   @return [Array]
#
# @!attribute [rw] expires_at
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] items
#   @return [Array]
#
# @!attribute [rw] messages
#   @return [Array]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] read_only
#   @return [Boolean]
#
# @!attribute [rw] reason
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] scope
#   @return [Object, nil]
#
# @!attribute [rw] success
#   @return [Boolean]
#
# @!attribute [rw] zones
#   @return [Array, nil]
EmailSendingSuppressionCreateData = Struct.new(
  :account_id,
  :created_at,
  :email,
  :errors,
  :expires_at,
  :id,
  :items,
  :messages,
  :note,
  :read_only,
  :reason,
  :result,
  :scope,
  :success,
  :zones,
  keyword_init: true
)

# Request payload for EmailSendingSuppression#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] errors
#   @return [Array, nil]
#
# @!attribute [rw] expires_at
#   @return [String, nil]
#
# @!attribute [rw] items
#   @return [Array, nil]
#
# @!attribute [rw] messages
#   @return [Array, nil]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] read_only
#   @return [Boolean, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] scope
#   @return [Object, nil]
#
# @!attribute [rw] success
#   @return [Boolean, nil]
#
# @!attribute [rw] zones
#   @return [Array, nil]
EmailSendingSuppressionUpdateData = Struct.new(
  :account_id,
  :id,
  :created_at,
  :email,
  :errors,
  :expires_at,
  :items,
  :messages,
  :note,
  :read_only,
  :reason,
  :result,
  :scope,
  :success,
  :zones,
  keyword_init: true
)

# Request payload for EmailSendingSuppression#remove.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String, nil]
EmailSendingSuppressionRemoveMatch = Struct.new(
  :account_id,
  :id,
  :zone_id,
  keyword_init: true
)

# Subdomain entity data model.
#
# @!attribute [rw] id
#   @return [String, nil]
Subdomain = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Subdomain#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] zone_id
#   @return [String]
SubdomainRemoveMatch = Struct.new(
  :id,
  :zone_id,
  keyword_init: true
)

