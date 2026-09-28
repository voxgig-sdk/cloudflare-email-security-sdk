# Typed models for the CloudflareEmailSecurity SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class EmailAccountRulesPlanResponseSingleRequired(TypedDict):
    owner_worker_tag: str


class EmailAccountRulesPlanResponseSingle(EmailAccountRulesPlanResponseSingleRequired, total=False):
    catch_all_rules: list
    rules: list


class EmailAccountRulesPlanResponseSingleCreateDataRequired(TypedDict):
    account_id: str
    owner_worker_tag: str


class EmailAccountRulesPlanResponseSingleCreateData(EmailAccountRulesPlanResponseSingleCreateDataRequired, total=False):
    catch_all_rules: list
    rules: list


class EmailAccountRulesResponseCollection(TypedDict, total=False):
    result: list
    result_info: dict


class EmailAccountRulesResponseCollectionListMatchRequired(TypedDict):
    account_id: str


class EmailAccountRulesResponseCollectionListMatch(EmailAccountRulesResponseCollectionListMatchRequired, total=False):
    enabled: bool
    page: float
    per_page: float


class EmailAuthRequired(TypedDict):
    result: dict


class EmailAuth(EmailAuthRequired, total=False):
    enabled: bool
    skip_wizard: bool


class EmailAuthLoadMatchRequired(TypedDict):
    zone_id: str


class EmailAuthLoadMatch(EmailAuthLoadMatchRequired, total=False):
    id: str


class EmailAuthUpdateDataRequired(TypedDict):
    zone_id: str


class EmailAuthUpdateData(EmailAuthUpdateDataRequired, total=False):
    enabled: bool
    result: dict
    skip_wizard: bool


class EmailCatchAllRuleResponseSingleRequired(TypedDict):
    actions: list
    matchers: list


class EmailCatchAllRuleResponseSingle(EmailCatchAllRuleResponseSingleRequired, total=False):
    enabled: bool
    name: str
    owner_worker_tag: str
    result: dict
    source: str


class EmailCatchAllRuleResponseSingleLoadMatch(TypedDict):
    zone_id: str


class EmailCatchAllRuleResponseSingleUpdateDataRequired(TypedDict):
    zone_id: str


class EmailCatchAllRuleResponseSingleUpdateData(EmailCatchAllRuleResponseSingleUpdateDataRequired, total=False):
    actions: list
    enabled: bool
    matchers: list
    name: str
    owner_worker_tag: str
    result: dict
    source: str


class EmailDestinationAddressResponseSingleRequired(TypedDict):
    email: str
    status: str


class EmailDestinationAddressResponseSingle(EmailDestinationAddressResponseSingleRequired, total=False):
    result: dict


class EmailDestinationAddressResponseSingleLoadMatch(TypedDict):
    account_id: str
    destination_address_identifier: str


class EmailDestinationAddressResponseSingleCreateDataRequired(TypedDict):
    account_id: str
    email: str
    status: str


class EmailDestinationAddressResponseSingleCreateData(EmailDestinationAddressResponseSingleCreateDataRequired, total=False):
    result: dict


class EmailDestinationAddressResponseSingleUpdateDataRequired(TypedDict):
    account_id: str
    destination_address_identifier: str


class EmailDestinationAddressResponseSingleUpdateData(EmailDestinationAddressResponseSingleUpdateDataRequired, total=False):
    email: str
    result: dict
    status: str


class EmailDestinationAddressResponseSingleRemoveMatch(TypedDict):
    account_id: str
    destination_address_identifier: str


class EmailDestinationAddressesResponseCollection(TypedDict, total=False):
    result: list
    result_info: dict


class EmailDestinationAddressesResponseCollectionListMatchRequired(TypedDict):
    account_id: str


class EmailDestinationAddressesResponseCollectionListMatch(EmailDestinationAddressesResponseCollectionListMatchRequired, total=False):
    direction: str
    page: float
    per_page: float
    verified: bool


class EmailDnsSettingsResponseCollection(TypedDict, total=False):
    result: list


class EmailDnsSettingsResponseCollectionListMatchRequired(TypedDict):
    zone_id: str


class EmailDnsSettingsResponseCollectionListMatch(EmailDnsSettingsResponseCollectionListMatchRequired, total=False):
    subdomain: str
    subdomain_id: str


class EmailEmailSettingsResponseSingle(TypedDict, total=False):
    enabled: bool
    name: str
    result: dict
    skip_wizard: bool
    support_subaddress: bool


class EmailEmailSettingsResponseSingleLoadMatch(TypedDict):
    zone_id: str


class EmailEmailSettingsResponseSingleCreateDataRequired(TypedDict):
    zone_id: str


class EmailEmailSettingsResponseSingleCreateData(EmailEmailSettingsResponseSingleCreateDataRequired, total=False):
    enabled: bool
    name: str
    result: dict
    skip_wizard: bool
    support_subaddress: bool


class EmailEmailSettingsResponseSingleUpdateDataRequired(TypedDict):
    zone_id: str


class EmailEmailSettingsResponseSingleUpdateData(EmailEmailSettingsResponseSingleUpdateDataRequired, total=False):
    enabled: bool
    name: str
    result: dict
    skip_wizard: bool
    support_subaddress: bool


class EmailEmailSettingsResponseSingleRemoveMatch(TypedDict):
    zone_id: str


class EmailRoutingSuppressionRequired(TypedDict):
    created_at: str
    email: str
    expires_at: str
    id: str
    reason: str


class EmailRoutingSuppression(EmailRoutingSuppressionRequired, total=False):
    zones: list


class EmailRoutingSuppressionLoadMatchRequired(TypedDict):
    id: str


class EmailRoutingSuppressionLoadMatch(EmailRoutingSuppressionLoadMatchRequired, total=False):
    account_id: str
    zone_id: str


class EmailRoutingSuppressionListMatchRequired(TypedDict):
    account_id: str


class EmailRoutingSuppressionListMatch(EmailRoutingSuppressionListMatchRequired, total=False):
    direction: str
    order: str
    page: int
    per_page: int


class EmailRoutingSuppressionCreateDataRequired(TypedDict):
    account_id: str
    created_at: str
    email: str
    expires_at: str
    id: str
    reason: str


class EmailRoutingSuppressionCreateData(EmailRoutingSuppressionCreateDataRequired, total=False):
    zones: list


class EmailRoutingSuppressionRemoveMatchRequired(TypedDict):
    id: str


class EmailRoutingSuppressionRemoveMatch(EmailRoutingSuppressionRemoveMatchRequired, total=False):
    account_id: str
    zone_id: str


class EmailRuleResponseSingleRequired(TypedDict):
    actions: list
    matchers: list


class EmailRuleResponseSingle(EmailRuleResponseSingleRequired, total=False):
    enabled: bool
    name: str
    owner_worker_tag: str
    priority: float
    result: dict
    source: str


class EmailRuleResponseSingleLoadMatch(TypedDict):
    rule_identifier: str
    zone_id: str


class EmailRuleResponseSingleCreateDataRequired(TypedDict):
    zone_id: str
    actions: list
    matchers: list


class EmailRuleResponseSingleCreateData(EmailRuleResponseSingleCreateDataRequired, total=False):
    enabled: bool
    name: str
    owner_worker_tag: str
    priority: float
    result: dict
    source: str


class EmailRuleResponseSingleUpdateDataRequired(TypedDict):
    rule_identifier: str
    zone_id: str


class EmailRuleResponseSingleUpdateData(EmailRuleResponseSingleUpdateDataRequired, total=False):
    actions: list
    enabled: bool
    matchers: list
    name: str
    owner_worker_tag: str
    priority: float
    result: dict
    source: str


class EmailRuleResponseSingleRemoveMatch(TypedDict):
    rule_identifier: str
    zone_id: str


class EmailRulesResponseCollection(TypedDict, total=False):
    result: list
    result_info: dict


class EmailRulesResponseCollectionListMatchRequired(TypedDict):
    zone_id: str


class EmailRulesResponseCollectionListMatch(EmailRulesResponseCollectionListMatchRequired, total=False):
    enabled: bool
    page: float
    per_page: float


class EmailSecurityRequired(TypedDict):
    action: str
    errors: list
    id: str
    ids: list
    messages: list
    result: dict
    result_info: dict
    search_params: dict
    success: bool


class EmailSecurity(EmailSecurityRequired, total=False):
    comment: str
    destination: Any
    expected_disposition: str
    postfix_ids: list


class EmailSecurityLoadMatchRequired(TypedDict):
    account_id: str


class EmailSecurityLoadMatch(EmailSecurityLoadMatchRequired, total=False):
    investigate_id: str
    submission: bool
    job_id: str


class EmailSecurityListMatchRequired(TypedDict):
    account_id: str


class EmailSecurityListMatch(EmailSecurityListMatchRequired, total=False):
    end: str
    from_date: str
    page: int
    per_page: int
    start: str
    to_date: str
    action_type: str
    status: str
    bulk_id: str
    investigate_id: str


class EmailSecurityCreateDataRequired(TypedDict):
    account_id: str
    action: str
    errors: list
    id: str
    ids: list
    messages: list
    result: dict
    result_info: dict
    search_params: dict
    success: bool


class EmailSecurityCreateData(EmailSecurityCreateDataRequired, total=False):
    comment: str
    destination: Any
    expected_disposition: str
    postfix_ids: list


class EmailSecurityRemoveMatch(TypedDict):
    account_id: str
    job_id: str


class EmailSecurityAnalytics(TypedDict, total=False):
    result: dict


class EmailSecurityAnalyticsLoadMatch(TypedDict):
    account_id: str


class EmailSecuritySettingRequired(TypedDict):
    allowed_delivery_modes: list
    deletes: list
    domain: str
    drop_dispositions: list
    ip_restrictions: list
    patches: list
    pattern: str
    posts: list
    puts: list
    regions: list


class EmailSecuritySetting(EmailSecuritySettingRequired, total=False):
    comments: str
    folder: Any
    integration_id: str
    lookback_hops: int
    require_tls_inbound: bool
    require_tls_outbound: bool
    result: dict
    transport: str


class EmailSecuritySettingLoadMatchRequired(TypedDict):
    account_id: str


class EmailSecuritySettingLoadMatch(EmailSecuritySettingLoadMatchRequired, total=False):
    domain_id: str
    impersonation_registry_id: str
    pattern_id: str
    policy_id: str
    sending_domain_restriction_id: str
    trusted_domain_id: str


class EmailSecuritySettingListMatchRequired(TypedDict):
    account_id: str


class EmailSecuritySettingListMatch(EmailSecuritySettingListMatchRequired, total=False):
    direction: str
    is_acceptable_sender: bool
    is_exempt_recipient: bool
    is_trusted_sender: bool
    order: str
    page: int
    pattern: str
    pattern_type: Any
    per_page: int
    search: str
    verify_sender: bool
    active_delivery_mode: str
    allowed_delivery_mode: str
    domain: list
    integration_id: str
    status: str
    is_recent: bool
    is_similarity: bool
    enabled: bool
    name: str
    provenance: str


class EmailSecuritySettingCreateDataRequired(TypedDict):
    account_id: str
    allowed_delivery_modes: list
    deletes: list
    domain: str
    drop_dispositions: list
    ip_restrictions: list
    patches: list
    pattern: str
    posts: list
    puts: list
    regions: list


class EmailSecuritySettingCreateData(EmailSecuritySettingCreateDataRequired, total=False):
    comments: str
    folder: Any
    integration_id: str
    lookback_hops: int
    require_tls_inbound: bool
    require_tls_outbound: bool
    result: dict
    transport: str


class EmailSecuritySettingUpdateDataRequired(TypedDict):
    account_id: str


class EmailSecuritySettingUpdateData(EmailSecuritySettingUpdateDataRequired, total=False):
    domain_id: str
    pattern_id: str
    policy_id: str
    trusted_domain_id: str
    allowed_delivery_modes: list
    comments: str
    deletes: list
    domain: str
    drop_dispositions: list
    folder: Any
    integration_id: str
    ip_restrictions: list
    lookback_hops: int
    patches: list
    pattern: str
    posts: list
    puts: list
    regions: list
    require_tls_inbound: bool
    require_tls_outbound: bool
    result: dict
    transport: str


class EmailSecuritySettingRemoveMatchRequired(TypedDict):
    account_id: str


class EmailSecuritySettingRemoveMatch(EmailSecuritySettingRemoveMatchRequired, total=False):
    domain_id: str
    impersonation_registry_id: str
    pattern_id: str
    policy_id: str
    sending_domain_restriction_id: str
    trusted_domain_id: str


class EmailSendingRequired(TypedDict):
    errors: list
    messages: list
    mime_message: str
    recipients: list
    result: dict
    result_info: dict
    subject: str
    success: bool


class EmailSending(EmailSendingRequired, total=False):
    attachments: list
    bcc: Any
    cc: Any
    headers: dict
    html: str
    reply_to: Any
    text: str
    to: Any


class EmailSendingLoadMatch(TypedDict):
    account_id: str
    message_id: str


class EmailSendingCreateDataRequired(TypedDict):
    account_id: str
    errors: list
    messages: list
    mime_message: str
    recipients: list
    result: dict
    result_info: dict
    subject: str
    success: bool


class EmailSendingCreateData(EmailSendingCreateDataRequired, total=False):
    attachments: list
    bcc: Any
    cc: Any
    headers: dict
    html: str
    reply_to: Any
    text: str
    to: Any


class EmailSendingLimitsResponseSingle(TypedDict, total=False):
    result: dict


class EmailSendingLimitsResponseSingleLoadMatch(TypedDict):
    account_id: str


class EmailSendingReputationResponseSingle(TypedDict):
    result: dict


class EmailSendingReputationResponseSingleLoadMatch(TypedDict):
    account_id: str


class EmailSendingSubdomainDnsStatus(TypedDict, total=False):
    result: dict


class EmailSendingSubdomainDnsStatusLoadMatch(TypedDict):
    subdomain_id: str
    zone_id: str


class EmailSendingSubdomainDnsStatusCreateDataRequired(TypedDict):
    subdomain_id: str
    zone_id: str


class EmailSendingSubdomainDnsStatusCreateData(EmailSendingSubdomainDnsStatusCreateDataRequired, total=False):
    result: dict


class EmailSendingSubdomainPreview(TypedDict):
    name: str


class EmailSendingSubdomainPreviewCreateData(TypedDict):
    zone_id: str
    name: str


class EmailSendingSubdomainReputationComplaint(TypedDict):
    result: dict


class EmailSendingSubdomainReputationComplaintLoadMatch(TypedDict):
    subdomain_id: str
    zone_id: str
    end_at: str
    start_at: str


class EmailSendingSubdomainResponseSingleRequired(TypedDict):
    name: str


class EmailSendingSubdomainResponseSingle(EmailSendingSubdomainResponseSingleRequired, total=False):
    drop_suppressed_recipients: bool
    preview_enabled: bool
    result: Any


class EmailSendingSubdomainResponseSingleLoadMatch(TypedDict):
    subdomain_id: str
    zone_id: str


class EmailSendingSubdomainResponseSingleCreateDataRequired(TypedDict):
    zone_id: str
    name: str


class EmailSendingSubdomainResponseSingleCreateData(EmailSendingSubdomainResponseSingleCreateDataRequired, total=False):
    drop_suppressed_recipients: bool
    preview_enabled: bool
    result: Any


class EmailSendingSubdomainResponseSingleUpdateDataRequired(TypedDict):
    subdomain_id: str
    zone_id: str


class EmailSendingSubdomainResponseSingleUpdateData(EmailSendingSubdomainResponseSingleUpdateDataRequired, total=False):
    drop_suppressed_recipients: bool
    name: str
    preview_enabled: bool
    result: Any


class EmailSendingSubdomainsResponseCollection(TypedDict, total=False):
    result: list


class EmailSendingSubdomainsResponseCollectionListMatch(TypedDict):
    zone_id: str


class EmailSendingSuppressionRequired(TypedDict):
    created_at: str
    email: str
    errors: list
    expires_at: str
    id: str
    items: list
    messages: list
    read_only: bool
    reason: str
    result: dict
    success: bool


class EmailSendingSuppression(EmailSendingSuppressionRequired, total=False):
    note: str
    scope: Any
    zones: list


class EmailSendingSuppressionLoadMatchRequired(TypedDict):
    id: str


class EmailSendingSuppressionLoadMatch(EmailSendingSuppressionLoadMatchRequired, total=False):
    account_id: str
    zone_id: str


class EmailSendingSuppressionListMatchRequired(TypedDict):
    account_id: str


class EmailSendingSuppressionListMatch(EmailSendingSuppressionListMatchRequired, total=False):
    cursor: str
    email: str
    per_page: int
    reason: str
    scope_type: str
    scope_value: str
    search: str


class EmailSendingSuppressionCreateDataRequired(TypedDict):
    account_id: str
    created_at: str
    email: str
    errors: list
    expires_at: str
    id: str
    items: list
    messages: list
    read_only: bool
    reason: str
    result: dict
    success: bool


class EmailSendingSuppressionCreateData(EmailSendingSuppressionCreateDataRequired, total=False):
    note: str
    scope: Any
    zones: list


class EmailSendingSuppressionUpdateDataRequired(TypedDict):
    account_id: str
    id: str


class EmailSendingSuppressionUpdateData(EmailSendingSuppressionUpdateDataRequired, total=False):
    created_at: str
    email: str
    errors: list
    expires_at: str
    items: list
    messages: list
    note: str
    read_only: bool
    reason: str
    result: dict
    scope: Any
    success: bool
    zones: list


class EmailSendingSuppressionRemoveMatchRequired(TypedDict):
    id: str


class EmailSendingSuppressionRemoveMatch(EmailSendingSuppressionRemoveMatchRequired, total=False):
    account_id: str
    zone_id: str


class Subdomain(TypedDict, total=False):
    id: str


class SubdomainRemoveMatch(TypedDict):
    id: str
    zone_id: str
