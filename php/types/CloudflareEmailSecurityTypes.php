<?php
declare(strict_types=1);

// Typed models for the CloudflareEmailSecurity SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** EmailAccountRulesPlanResponseSingle entity data model. */
class EmailAccountRulesPlanResponseSingle
{
    public ?array $catch_all_rules = null;
    public string $owner_worker_tag;
    public ?array $rules = null;
}

/** Request payload for EmailAccountRulesPlanResponseSingle#create. */
class EmailAccountRulesPlanResponseSingleCreateData
{
    public string $account_id;
    public ?array $catch_all_rules = null;
    public string $owner_worker_tag;
    public ?array $rules = null;
}

/** EmailAccountRulesResponseCollection entity data model. */
class EmailAccountRulesResponseCollection
{
    public ?array $result = null;
    public ?array $result_info = null;
}

/** Request payload for EmailAccountRulesResponseCollection#list. */
class EmailAccountRulesResponseCollectionListMatch
{
    public string $account_id;
    public ?bool $enabled = null;
    public ?float $page = null;
    public ?float $per_page = null;
}

/** EmailAuth entity data model. */
class EmailAuth
{
    public ?bool $enabled = null;
    public array $result;
    public ?bool $skip_wizard = null;
}

/** Request payload for EmailAuth#load. */
class EmailAuthLoadMatch
{
    public string $zone_id;
    public ?string $id = null;
}

/** Request payload for EmailAuth#update. */
class EmailAuthUpdateData
{
    public string $zone_id;
    public ?bool $enabled = null;
    public ?array $result = null;
    public ?bool $skip_wizard = null;
}

/** EmailCatchAllRuleResponseSingle entity data model. */
class EmailCatchAllRuleResponseSingle
{
    public array $actions;
    public ?bool $enabled = null;
    public array $matchers;
    public ?string $name = null;
    public ?string $owner_worker_tag = null;
    public ?array $result = null;
    public ?string $source = null;
}

/** Request payload for EmailCatchAllRuleResponseSingle#load. */
class EmailCatchAllRuleResponseSingleLoadMatch
{
    public string $zone_id;
}

/** Request payload for EmailCatchAllRuleResponseSingle#update. */
class EmailCatchAllRuleResponseSingleUpdateData
{
    public string $zone_id;
    public ?array $actions = null;
    public ?bool $enabled = null;
    public ?array $matchers = null;
    public ?string $name = null;
    public ?string $owner_worker_tag = null;
    public ?array $result = null;
    public ?string $source = null;
}

/** EmailDestinationAddressResponseSingle entity data model. */
class EmailDestinationAddressResponseSingle
{
    public string $email;
    public ?array $result = null;
    public string $status;
}

/** Request payload for EmailDestinationAddressResponseSingle#load. */
class EmailDestinationAddressResponseSingleLoadMatch
{
    public string $account_id;
    public string $destination_address_identifier;
}

/** Request payload for EmailDestinationAddressResponseSingle#create. */
class EmailDestinationAddressResponseSingleCreateData
{
    public string $account_id;
    public string $email;
    public ?array $result = null;
    public string $status;
}

/** Request payload for EmailDestinationAddressResponseSingle#update. */
class EmailDestinationAddressResponseSingleUpdateData
{
    public string $account_id;
    public string $destination_address_identifier;
    public ?string $email = null;
    public ?array $result = null;
    public ?string $status = null;
}

/** Request payload for EmailDestinationAddressResponseSingle#remove. */
class EmailDestinationAddressResponseSingleRemoveMatch
{
    public string $account_id;
    public string $destination_address_identifier;
}

/** EmailDestinationAddressesResponseCollection entity data model. */
class EmailDestinationAddressesResponseCollection
{
    public ?array $result = null;
    public ?array $result_info = null;
}

/** Request payload for EmailDestinationAddressesResponseCollection#list. */
class EmailDestinationAddressesResponseCollectionListMatch
{
    public string $account_id;
    public ?string $direction = null;
    public ?float $page = null;
    public ?float $per_page = null;
    public ?bool $verified = null;
}

/** EmailDnsSettingsResponseCollection entity data model. */
class EmailDnsSettingsResponseCollection
{
    public ?array $result = null;
}

/** Request payload for EmailDnsSettingsResponseCollection#list. */
class EmailDnsSettingsResponseCollectionListMatch
{
    public string $zone_id;
    public ?string $subdomain = null;
    public ?string $subdomain_id = null;
}

/** EmailEmailSettingsResponseSingle entity data model. */
class EmailEmailSettingsResponseSingle
{
    public ?bool $enabled = null;
    public ?string $name = null;
    public ?array $result = null;
    public ?bool $skip_wizard = null;
    public ?bool $support_subaddress = null;
}

/** Request payload for EmailEmailSettingsResponseSingle#load. */
class EmailEmailSettingsResponseSingleLoadMatch
{
    public string $zone_id;
}

/** Request payload for EmailEmailSettingsResponseSingle#create. */
class EmailEmailSettingsResponseSingleCreateData
{
    public string $zone_id;
    public ?bool $enabled = null;
    public ?string $name = null;
    public ?array $result = null;
    public ?bool $skip_wizard = null;
    public ?bool $support_subaddress = null;
}

/** Request payload for EmailEmailSettingsResponseSingle#update. */
class EmailEmailSettingsResponseSingleUpdateData
{
    public string $zone_id;
    public ?bool $enabled = null;
    public ?string $name = null;
    public ?array $result = null;
    public ?bool $skip_wizard = null;
    public ?bool $support_subaddress = null;
}

/** Request payload for EmailEmailSettingsResponseSingle#remove. */
class EmailEmailSettingsResponseSingleRemoveMatch
{
    public string $zone_id;
}

/** EmailRoutingSuppression entity data model. */
class EmailRoutingSuppression
{
    public string $created_at;
    public string $email;
    public string $expires_at;
    public string $id;
    public string $reason;
    public ?array $zones = null;
}

/** Request payload for EmailRoutingSuppression#load. */
class EmailRoutingSuppressionLoadMatch
{
    public ?string $account_id = null;
    public string $id;
    public ?string $zone_id = null;
}

/** Request payload for EmailRoutingSuppression#list. */
class EmailRoutingSuppressionListMatch
{
    public string $account_id;
    public ?string $direction = null;
    public ?string $order = null;
    public ?int $page = null;
    public ?int $per_page = null;
}

/** Request payload for EmailRoutingSuppression#create. */
class EmailRoutingSuppressionCreateData
{
    public string $account_id;
    public string $created_at;
    public string $email;
    public string $expires_at;
    public string $id;
    public string $reason;
    public ?array $zones = null;
}

/** Request payload for EmailRoutingSuppression#remove. */
class EmailRoutingSuppressionRemoveMatch
{
    public ?string $account_id = null;
    public string $id;
    public ?string $zone_id = null;
}

/** EmailRuleResponseSingle entity data model. */
class EmailRuleResponseSingle
{
    public array $actions;
    public ?bool $enabled = null;
    public array $matchers;
    public ?string $name = null;
    public ?string $owner_worker_tag = null;
    public ?float $priority = null;
    public ?array $result = null;
    public ?string $source = null;
}

/** Request payload for EmailRuleResponseSingle#load. */
class EmailRuleResponseSingleLoadMatch
{
    public string $rule_identifier;
    public string $zone_id;
}

/** Request payload for EmailRuleResponseSingle#create. */
class EmailRuleResponseSingleCreateData
{
    public string $zone_id;
    public array $actions;
    public ?bool $enabled = null;
    public array $matchers;
    public ?string $name = null;
    public ?string $owner_worker_tag = null;
    public ?float $priority = null;
    public ?array $result = null;
    public ?string $source = null;
}

/** Request payload for EmailRuleResponseSingle#update. */
class EmailRuleResponseSingleUpdateData
{
    public string $rule_identifier;
    public string $zone_id;
    public ?array $actions = null;
    public ?bool $enabled = null;
    public ?array $matchers = null;
    public ?string $name = null;
    public ?string $owner_worker_tag = null;
    public ?float $priority = null;
    public ?array $result = null;
    public ?string $source = null;
}

/** Request payload for EmailRuleResponseSingle#remove. */
class EmailRuleResponseSingleRemoveMatch
{
    public string $rule_identifier;
    public string $zone_id;
}

/** EmailRulesResponseCollection entity data model. */
class EmailRulesResponseCollection
{
    public ?array $result = null;
    public ?array $result_info = null;
}

/** Request payload for EmailRulesResponseCollection#list. */
class EmailRulesResponseCollectionListMatch
{
    public string $zone_id;
    public ?bool $enabled = null;
    public ?float $page = null;
    public ?float $per_page = null;
}

/** EmailSecurity entity data model. */
class EmailSecurity
{
    public string $action;
    public ?string $comment = null;
    public mixed $destination = null;
    public array $errors;
    public ?string $expected_disposition = null;
    public string $id;
    public array $ids;
    public array $messages;
    public ?array $postfix_ids = null;
    public array $result;
    public array $result_info;
    public array $search_params;
    public bool $success;
}

/** Request payload for EmailSecurity#load. */
class EmailSecurityLoadMatch
{
    public string $account_id;
    public ?string $investigate_id = null;
    public ?bool $submission = null;
    public ?string $job_id = null;
}

/** Request payload for EmailSecurity#list. */
class EmailSecurityListMatch
{
    public string $account_id;
    public ?string $end = null;
    public ?string $from_date = null;
    public ?int $page = null;
    public ?int $per_page = null;
    public ?string $start = null;
    public ?string $to_date = null;
    public ?string $action_type = null;
    public ?string $status = null;
    public ?string $bulk_id = null;
    public ?string $investigate_id = null;
}

/** Request payload for EmailSecurity#create. */
class EmailSecurityCreateData
{
    public string $account_id;
    public string $action;
    public ?string $comment = null;
    public mixed $destination = null;
    public array $errors;
    public ?string $expected_disposition = null;
    public string $id;
    public array $ids;
    public array $messages;
    public ?array $postfix_ids = null;
    public array $result;
    public array $result_info;
    public array $search_params;
    public bool $success;
}

/** Request payload for EmailSecurity#remove. */
class EmailSecurityRemoveMatch
{
    public string $account_id;
    public string $job_id;
}

/** EmailSecurityAnalytics entity data model. */
class EmailSecurityAnalytics
{
    public ?array $result = null;
}

/** Request payload for EmailSecurityAnalytics#load. */
class EmailSecurityAnalyticsLoadMatch
{
    public string $account_id;
}

/** EmailSecuritySetting entity data model. */
class EmailSecuritySetting
{
    public array $allowed_delivery_modes;
    public ?string $comments = null;
    public array $deletes;
    public string $domain;
    public array $drop_dispositions;
    public mixed $folder = null;
    public ?string $integration_id = null;
    public array $ip_restrictions;
    public ?int $lookback_hops = null;
    public array $patches;
    public string $pattern;
    public array $posts;
    public array $puts;
    public array $regions;
    public ?bool $require_tls_inbound = null;
    public ?bool $require_tls_outbound = null;
    public ?array $result = null;
    public ?string $transport = null;
}

/** Request payload for EmailSecuritySetting#load. */
class EmailSecuritySettingLoadMatch
{
    public string $account_id;
    public ?string $domain_id = null;
    public ?string $impersonation_registry_id = null;
    public ?string $pattern_id = null;
    public ?string $policy_id = null;
    public ?string $sending_domain_restriction_id = null;
    public ?string $trusted_domain_id = null;
}

/** Request payload for EmailSecuritySetting#list. */
class EmailSecuritySettingListMatch
{
    public string $account_id;
    public ?string $direction = null;
    public ?bool $is_acceptable_sender = null;
    public ?bool $is_exempt_recipient = null;
    public ?bool $is_trusted_sender = null;
    public ?string $order = null;
    public ?int $page = null;
    public ?string $pattern = null;
    public mixed $pattern_type = null;
    public ?int $per_page = null;
    public ?string $search = null;
    public ?bool $verify_sender = null;
    public ?string $active_delivery_mode = null;
    public ?string $allowed_delivery_mode = null;
    public ?array $domain = null;
    public ?string $integration_id = null;
    public ?string $status = null;
    public ?bool $is_recent = null;
    public ?bool $is_similarity = null;
    public ?bool $enabled = null;
    public ?string $name = null;
    public ?string $provenance = null;
}

/** Request payload for EmailSecuritySetting#create. */
class EmailSecuritySettingCreateData
{
    public string $account_id;
    public array $allowed_delivery_modes;
    public ?string $comments = null;
    public array $deletes;
    public string $domain;
    public array $drop_dispositions;
    public mixed $folder = null;
    public ?string $integration_id = null;
    public array $ip_restrictions;
    public ?int $lookback_hops = null;
    public array $patches;
    public string $pattern;
    public array $posts;
    public array $puts;
    public array $regions;
    public ?bool $require_tls_inbound = null;
    public ?bool $require_tls_outbound = null;
    public ?array $result = null;
    public ?string $transport = null;
}

/** Request payload for EmailSecuritySetting#update. */
class EmailSecuritySettingUpdateData
{
    public string $account_id;
    public ?string $domain_id = null;
    public ?string $pattern_id = null;
    public ?string $policy_id = null;
    public ?string $trusted_domain_id = null;
    public ?array $allowed_delivery_modes = null;
    public ?string $comments = null;
    public ?array $deletes = null;
    public ?string $domain = null;
    public ?array $drop_dispositions = null;
    public mixed $folder = null;
    public ?string $integration_id = null;
    public ?array $ip_restrictions = null;
    public ?int $lookback_hops = null;
    public ?array $patches = null;
    public ?string $pattern = null;
    public ?array $posts = null;
    public ?array $puts = null;
    public ?array $regions = null;
    public ?bool $require_tls_inbound = null;
    public ?bool $require_tls_outbound = null;
    public ?array $result = null;
    public ?string $transport = null;
}

/** Request payload for EmailSecuritySetting#remove. */
class EmailSecuritySettingRemoveMatch
{
    public string $account_id;
    public ?string $domain_id = null;
    public ?string $impersonation_registry_id = null;
    public ?string $pattern_id = null;
    public ?string $policy_id = null;
    public ?string $sending_domain_restriction_id = null;
    public ?string $trusted_domain_id = null;
}

/** EmailSending entity data model. */
class EmailSending
{
    public ?array $attachments = null;
    public mixed $bcc = null;
    public mixed $cc = null;
    public array $errors;
    public mixed $from;
    public ?array $headers = null;
    public ?string $html = null;
    public array $messages;
    public string $mime_message;
    public array $recipients;
    public mixed $reply_to = null;
    public array $result;
    public array $result_info;
    public string $subject;
    public bool $success;
    public ?string $text = null;
    public mixed $to = null;
}

/** Request payload for EmailSending#load. */
class EmailSendingLoadMatch
{
    public string $account_id;
    public string $message_id;
}

/** Request payload for EmailSending#create. */
class EmailSendingCreateData
{
    public string $account_id;
    public ?array $attachments = null;
    public mixed $bcc = null;
    public mixed $cc = null;
    public array $errors;
    public mixed $from;
    public ?array $headers = null;
    public ?string $html = null;
    public array $messages;
    public string $mime_message;
    public array $recipients;
    public mixed $reply_to = null;
    public array $result;
    public array $result_info;
    public string $subject;
    public bool $success;
    public ?string $text = null;
    public mixed $to = null;
}

/** EmailSendingLimitsResponseSingle entity data model. */
class EmailSendingLimitsResponseSingle
{
    public ?array $result = null;
}

/** Request payload for EmailSendingLimitsResponseSingle#load. */
class EmailSendingLimitsResponseSingleLoadMatch
{
    public string $account_id;
}

/** EmailSendingReputationResponseSingle entity data model. */
class EmailSendingReputationResponseSingle
{
    public array $result;
}

/** Request payload for EmailSendingReputationResponseSingle#load. */
class EmailSendingReputationResponseSingleLoadMatch
{
    public string $account_id;
}

/** EmailSendingSubdomainDnsStatus entity data model. */
class EmailSendingSubdomainDnsStatus
{
    public ?array $result = null;
}

/** Request payload for EmailSendingSubdomainDnsStatus#load. */
class EmailSendingSubdomainDnsStatusLoadMatch
{
    public string $subdomain_id;
    public string $zone_id;
}

/** Request payload for EmailSendingSubdomainDnsStatus#create. */
class EmailSendingSubdomainDnsStatusCreateData
{
    public string $subdomain_id;
    public string $zone_id;
    public ?array $result = null;
}

/** EmailSendingSubdomainPreview entity data model. */
class EmailSendingSubdomainPreview
{
    public string $name;
}

/** Request payload for EmailSendingSubdomainPreview#create. */
class EmailSendingSubdomainPreviewCreateData
{
    public string $zone_id;
    public string $name;
}

/** EmailSendingSubdomainReputationComplaint entity data model. */
class EmailSendingSubdomainReputationComplaint
{
    public array $result;
}

/** Request payload for EmailSendingSubdomainReputationComplaint#load. */
class EmailSendingSubdomainReputationComplaintLoadMatch
{
    public string $subdomain_id;
    public string $zone_id;
    public string $end_at;
    public string $start_at;
}

/** EmailSendingSubdomainResponseSingle entity data model. */
class EmailSendingSubdomainResponseSingle
{
    public ?bool $drop_suppressed_recipients = null;
    public string $name;
    public ?bool $preview_enabled = null;
    public mixed $result = null;
}

/** Request payload for EmailSendingSubdomainResponseSingle#load. */
class EmailSendingSubdomainResponseSingleLoadMatch
{
    public string $subdomain_id;
    public string $zone_id;
}

/** Request payload for EmailSendingSubdomainResponseSingle#create. */
class EmailSendingSubdomainResponseSingleCreateData
{
    public string $zone_id;
    public ?bool $drop_suppressed_recipients = null;
    public string $name;
    public ?bool $preview_enabled = null;
    public mixed $result = null;
}

/** Request payload for EmailSendingSubdomainResponseSingle#update. */
class EmailSendingSubdomainResponseSingleUpdateData
{
    public string $subdomain_id;
    public string $zone_id;
    public ?bool $drop_suppressed_recipients = null;
    public ?string $name = null;
    public ?bool $preview_enabled = null;
    public mixed $result = null;
}

/** EmailSendingSubdomainsResponseCollection entity data model. */
class EmailSendingSubdomainsResponseCollection
{
    public ?array $result = null;
}

/** Request payload for EmailSendingSubdomainsResponseCollection#list. */
class EmailSendingSubdomainsResponseCollectionListMatch
{
    public string $zone_id;
}

/** EmailSendingSuppression entity data model. */
class EmailSendingSuppression
{
    public string $created_at;
    public string $email;
    public array $errors;
    public string $expires_at;
    public string $id;
    public array $items;
    public array $messages;
    public ?string $note = null;
    public bool $read_only;
    public string $reason;
    public array $result;
    public mixed $scope = null;
    public bool $success;
    public ?array $zones = null;
}

/** Request payload for EmailSendingSuppression#load. */
class EmailSendingSuppressionLoadMatch
{
    public ?string $account_id = null;
    public string $id;
    public ?string $zone_id = null;
}

/** Request payload for EmailSendingSuppression#list. */
class EmailSendingSuppressionListMatch
{
    public string $account_id;
    public ?string $cursor = null;
    public ?string $email = null;
    public ?int $per_page = null;
    public ?string $reason = null;
    public ?string $scope_type = null;
    public ?string $scope_value = null;
    public ?string $search = null;
}

/** Request payload for EmailSendingSuppression#create. */
class EmailSendingSuppressionCreateData
{
    public string $account_id;
    public string $created_at;
    public string $email;
    public array $errors;
    public string $expires_at;
    public string $id;
    public array $items;
    public array $messages;
    public ?string $note = null;
    public bool $read_only;
    public string $reason;
    public array $result;
    public mixed $scope = null;
    public bool $success;
    public ?array $zones = null;
}

/** Request payload for EmailSendingSuppression#update. */
class EmailSendingSuppressionUpdateData
{
    public string $account_id;
    public string $id;
    public ?string $created_at = null;
    public ?string $email = null;
    public ?array $errors = null;
    public ?string $expires_at = null;
    public ?array $items = null;
    public ?array $messages = null;
    public ?string $note = null;
    public ?bool $read_only = null;
    public ?string $reason = null;
    public ?array $result = null;
    public mixed $scope = null;
    public ?bool $success = null;
    public ?array $zones = null;
}

/** Request payload for EmailSendingSuppression#remove. */
class EmailSendingSuppressionRemoveMatch
{
    public ?string $account_id = null;
    public string $id;
    public ?string $zone_id = null;
}

/** Subdomain entity data model. */
class Subdomain
{
    public ?string $id = null;
}

/** Request payload for Subdomain#remove. */
class SubdomainRemoveMatch
{
    public string $id;
    public string $zone_id;
}

