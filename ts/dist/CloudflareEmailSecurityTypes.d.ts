export interface EmailAccountRulesPlanResponseSingle {
    catch_all_rules?: any[];
    owner_worker_tag: string;
    rules?: any[];
}
export interface EmailAccountRulesPlanResponseSingleCreateData {
    account_id: string;
    catch_all_rules?: any[];
    owner_worker_tag: string;
    rules?: any[];
}
export interface EmailAccountRulesResponseCollection {
    result?: any[];
    result_info?: Record<string, any>;
}
export interface EmailAccountRulesResponseCollectionListMatch {
    account_id: string;
    enabled?: boolean;
    page?: number;
    per_page?: number;
}
export interface EmailAuth {
    enabled?: boolean;
    result: Record<string, any>;
    skip_wizard?: boolean;
}
export interface EmailAuthLoadMatch {
    zone_id: string;
    id?: string;
}
export interface EmailAuthUpdateData {
    zone_id: string;
    enabled?: boolean;
    result?: Record<string, any>;
    skip_wizard?: boolean;
}
export interface EmailCatchAllRuleResponseSingle {
    actions: any[];
    enabled?: boolean;
    matchers: any[];
    name?: string;
    owner_worker_tag?: string;
    result?: Record<string, any>;
    source?: string;
}
export interface EmailCatchAllRuleResponseSingleLoadMatch {
    zone_id: string;
}
export interface EmailCatchAllRuleResponseSingleUpdateData {
    zone_id: string;
    actions?: any[];
    enabled?: boolean;
    matchers?: any[];
    name?: string;
    owner_worker_tag?: string;
    result?: Record<string, any>;
    source?: string;
}
export interface EmailDestinationAddressResponseSingle {
    email: string;
    result?: Record<string, any>;
    status: string;
}
export interface EmailDestinationAddressResponseSingleLoadMatch {
    account_id: string;
    destination_address_identifier: string;
}
export interface EmailDestinationAddressResponseSingleCreateData {
    account_id: string;
    email: string;
    result?: Record<string, any>;
    status: string;
}
export interface EmailDestinationAddressResponseSingleUpdateData {
    account_id: string;
    destination_address_identifier: string;
    email?: string;
    result?: Record<string, any>;
    status?: string;
}
export interface EmailDestinationAddressResponseSingleRemoveMatch {
    account_id: string;
    destination_address_identifier: string;
}
export interface EmailDestinationAddressesResponseCollection {
    result?: any[];
    result_info?: Record<string, any>;
}
export interface EmailDestinationAddressesResponseCollectionListMatch {
    account_id: string;
    direction?: string;
    page?: number;
    per_page?: number;
    verified?: boolean;
}
export interface EmailDnsSettingsResponseCollection {
    result?: any[];
}
export interface EmailDnsSettingsResponseCollectionListMatch {
    zone_id: string;
    subdomain?: string;
    subdomain_id?: string;
}
export interface EmailEmailSettingsResponseSingle {
    enabled?: boolean;
    name?: string;
    result?: Record<string, any>;
    skip_wizard?: boolean;
    support_subaddress?: boolean;
}
export interface EmailEmailSettingsResponseSingleLoadMatch {
    zone_id: string;
}
export interface EmailEmailSettingsResponseSingleCreateData {
    zone_id: string;
    enabled?: boolean;
    name?: string;
    result?: Record<string, any>;
    skip_wizard?: boolean;
    support_subaddress?: boolean;
}
export interface EmailEmailSettingsResponseSingleUpdateData {
    zone_id: string;
    enabled?: boolean;
    name?: string;
    result?: Record<string, any>;
    skip_wizard?: boolean;
    support_subaddress?: boolean;
}
export interface EmailEmailSettingsResponseSingleRemoveMatch {
    zone_id: string;
}
export interface EmailRoutingSuppression {
    created_at: string;
    email: string;
    expires_at: string;
    id: string;
    reason: string;
    zones?: any[];
}
export interface EmailRoutingSuppressionLoadMatch {
    account_id?: string;
    id: string;
    zone_id?: string;
}
export interface EmailRoutingSuppressionListMatch {
    account_id: string;
    direction?: string;
    order?: string;
    page?: number;
    per_page?: number;
}
export interface EmailRoutingSuppressionCreateData {
    account_id: string;
    created_at: string;
    email: string;
    expires_at: string;
    id: string;
    reason: string;
    zones?: any[];
}
export interface EmailRoutingSuppressionRemoveMatch {
    account_id?: string;
    id: string;
    zone_id?: string;
}
export interface EmailRuleResponseSingle {
    actions: any[];
    enabled?: boolean;
    matchers: any[];
    name?: string;
    owner_worker_tag?: string;
    priority?: number;
    result?: Record<string, any>;
    source?: string;
}
export interface EmailRuleResponseSingleLoadMatch {
    rule_identifier: string;
    zone_id: string;
}
export interface EmailRuleResponseSingleCreateData {
    zone_id: string;
    actions: any[];
    enabled?: boolean;
    matchers: any[];
    name?: string;
    owner_worker_tag?: string;
    priority?: number;
    result?: Record<string, any>;
    source?: string;
}
export interface EmailRuleResponseSingleUpdateData {
    rule_identifier: string;
    zone_id: string;
    actions?: any[];
    enabled?: boolean;
    matchers?: any[];
    name?: string;
    owner_worker_tag?: string;
    priority?: number;
    result?: Record<string, any>;
    source?: string;
}
export interface EmailRuleResponseSingleRemoveMatch {
    rule_identifier: string;
    zone_id: string;
}
export interface EmailRulesResponseCollection {
    result?: any[];
    result_info?: Record<string, any>;
}
export interface EmailRulesResponseCollectionListMatch {
    zone_id: string;
    enabled?: boolean;
    page?: number;
    per_page?: number;
}
export interface EmailSecurity {
    action: string;
    comment?: string;
    destination?: any;
    errors: any[];
    expected_disposition?: string;
    id: string;
    ids: any[];
    messages: any[];
    postfix_ids?: any[];
    result: Record<string, any>;
    result_info: Record<string, any>;
    search_params: Record<string, any>;
    success: boolean;
}
export interface EmailSecurityLoadMatch {
    account_id: string;
    investigate_id?: string;
    submission?: boolean;
    job_id?: string;
}
export interface EmailSecurityListMatch {
    account_id: string;
    end?: string;
    from_date?: string;
    page?: number;
    per_page?: number;
    start?: string;
    to_date?: string;
    action_type?: string;
    status?: string;
    bulk_id?: string;
    investigate_id?: string;
    $action?: string;
    [action: string]: any;
}
export interface EmailSecurityCreateData {
    account_id: string;
    action: string;
    comment?: string;
    destination?: any;
    errors: any[];
    expected_disposition?: string;
    id: string;
    ids: any[];
    messages: any[];
    postfix_ids?: any[];
    result: Record<string, any>;
    result_info: Record<string, any>;
    search_params: Record<string, any>;
    success: boolean;
    $action?: string;
    [action: string]: any;
}
export interface EmailSecurityRemoveMatch {
    account_id: string;
    job_id: string;
}
export interface EmailSecurityAnalytics {
    result?: Record<string, any>;
}
export interface EmailSecurityAnalyticsLoadMatch {
    account_id: string;
}
export interface EmailSecuritySetting {
    allowed_delivery_modes: any[];
    comments?: string;
    deletes: any[];
    domain: string;
    drop_dispositions: any[];
    folder?: any;
    integration_id?: string;
    ip_restrictions: any[];
    lookback_hops?: number;
    patches: any[];
    pattern: string;
    posts: any[];
    puts: any[];
    regions: any[];
    require_tls_inbound?: boolean;
    require_tls_outbound?: boolean;
    result?: Record<string, any>;
    transport?: string;
}
export interface EmailSecuritySettingLoadMatch {
    account_id: string;
    domain_id?: string;
    impersonation_registry_id?: string;
    pattern_id?: string;
    policy_id?: string;
    sending_domain_restriction_id?: string;
    trusted_domain_id?: string;
}
export interface EmailSecuritySettingListMatch {
    account_id: string;
    direction?: string;
    is_acceptable_sender?: boolean;
    is_exempt_recipient?: boolean;
    is_trusted_sender?: boolean;
    order?: string;
    page?: number;
    pattern?: string;
    pattern_type?: any;
    per_page?: number;
    search?: string;
    verify_sender?: boolean;
    active_delivery_mode?: string;
    allowed_delivery_mode?: string;
    domain?: any[];
    integration_id?: string;
    status?: string;
    is_recent?: boolean;
    is_similarity?: boolean;
    enabled?: boolean;
    name?: string;
    provenance?: string;
}
export interface EmailSecuritySettingCreateData {
    account_id: string;
    allowed_delivery_modes: any[];
    comments?: string;
    deletes: any[];
    domain: string;
    drop_dispositions: any[];
    folder?: any;
    integration_id?: string;
    ip_restrictions: any[];
    lookback_hops?: number;
    patches: any[];
    pattern: string;
    posts: any[];
    puts: any[];
    regions: any[];
    require_tls_inbound?: boolean;
    require_tls_outbound?: boolean;
    result?: Record<string, any>;
    transport?: string;
}
export interface EmailSecuritySettingUpdateData {
    account_id: string;
    domain_id?: string;
    pattern_id?: string;
    policy_id?: string;
    trusted_domain_id?: string;
    allowed_delivery_modes?: any[];
    comments?: string;
    deletes?: any[];
    domain?: string;
    drop_dispositions?: any[];
    folder?: any;
    integration_id?: string;
    ip_restrictions?: any[];
    lookback_hops?: number;
    patches?: any[];
    pattern?: string;
    posts?: any[];
    puts?: any[];
    regions?: any[];
    require_tls_inbound?: boolean;
    require_tls_outbound?: boolean;
    result?: Record<string, any>;
    transport?: string;
}
export interface EmailSecuritySettingRemoveMatch {
    account_id: string;
    domain_id?: string;
    impersonation_registry_id?: string;
    pattern_id?: string;
    policy_id?: string;
    sending_domain_restriction_id?: string;
    trusted_domain_id?: string;
}
export interface EmailSending {
    attachments?: any[];
    bcc?: any;
    cc?: any;
    errors: any[];
    from: any;
    headers?: Record<string, any>;
    html?: string;
    messages: any[];
    mime_message: string;
    recipients: any[];
    reply_to?: any;
    result: Record<string, any>;
    result_info: Record<string, any>;
    subject: string;
    success: boolean;
    text?: string;
    to?: any;
}
export interface EmailSendingLoadMatch {
    account_id: string;
    message_id: string;
}
export interface EmailSendingCreateData {
    account_id: string;
    attachments?: any[];
    bcc?: any;
    cc?: any;
    errors: any[];
    from: any;
    headers?: Record<string, any>;
    html?: string;
    messages: any[];
    mime_message: string;
    recipients: any[];
    reply_to?: any;
    result: Record<string, any>;
    result_info: Record<string, any>;
    subject: string;
    success: boolean;
    text?: string;
    to?: any;
}
export interface EmailSendingLimitsResponseSingle {
    result?: Record<string, any>;
}
export interface EmailSendingLimitsResponseSingleLoadMatch {
    account_id: string;
}
export interface EmailSendingReputationResponseSingle {
    result: Record<string, any>;
}
export interface EmailSendingReputationResponseSingleLoadMatch {
    account_id: string;
}
export interface EmailSendingSubdomainDnsStatus {
    result?: Record<string, any>;
}
export interface EmailSendingSubdomainDnsStatusLoadMatch {
    subdomain_id: string;
    zone_id: string;
}
export interface EmailSendingSubdomainDnsStatusCreateData {
    subdomain_id: string;
    zone_id: string;
    result?: Record<string, any>;
}
export interface EmailSendingSubdomainPreview {
    name: string;
}
export interface EmailSendingSubdomainPreviewCreateData {
    zone_id: string;
    name: string;
}
export interface EmailSendingSubdomainReputationComplaint {
    result: Record<string, any>;
}
export interface EmailSendingSubdomainReputationComplaintLoadMatch {
    subdomain_id: string;
    zone_id: string;
    end_at: string;
    start_at: string;
}
export interface EmailSendingSubdomainResponseSingle {
    drop_suppressed_recipients?: boolean;
    name: string;
    preview_enabled?: boolean;
    result?: any;
}
export interface EmailSendingSubdomainResponseSingleLoadMatch {
    subdomain_id: string;
    zone_id: string;
}
export interface EmailSendingSubdomainResponseSingleCreateData {
    zone_id: string;
    drop_suppressed_recipients?: boolean;
    name: string;
    preview_enabled?: boolean;
    result?: any;
}
export interface EmailSendingSubdomainResponseSingleUpdateData {
    subdomain_id: string;
    zone_id: string;
    drop_suppressed_recipients?: boolean;
    name?: string;
    preview_enabled?: boolean;
    result?: any;
}
export interface EmailSendingSubdomainsResponseCollection {
    result?: any[];
}
export interface EmailSendingSubdomainsResponseCollectionListMatch {
    zone_id: string;
}
export interface EmailSendingSuppression {
    created_at: string;
    email: string;
    errors: any[];
    expires_at: string;
    id: string;
    items: any[];
    messages: any[];
    note?: string;
    read_only: boolean;
    reason: string;
    result: Record<string, any>;
    scope?: any;
    success: boolean;
    zones?: any[];
}
export interface EmailSendingSuppressionLoadMatch {
    account_id?: string;
    id: string;
    zone_id?: string;
}
export interface EmailSendingSuppressionListMatch {
    account_id: string;
    cursor?: string;
    email?: string;
    per_page?: number;
    reason?: string;
    scope_type?: string;
    scope_value?: string;
    search?: string;
}
export interface EmailSendingSuppressionCreateData {
    account_id: string;
    created_at: string;
    email: string;
    errors: any[];
    expires_at: string;
    id: string;
    items: any[];
    messages: any[];
    note?: string;
    read_only: boolean;
    reason: string;
    result: Record<string, any>;
    scope?: any;
    success: boolean;
    zones?: any[];
}
export interface EmailSendingSuppressionUpdateData {
    account_id: string;
    id: string;
    created_at?: string;
    email?: string;
    errors?: any[];
    expires_at?: string;
    items?: any[];
    messages?: any[];
    note?: string;
    read_only?: boolean;
    reason?: string;
    result?: Record<string, any>;
    scope?: any;
    success?: boolean;
    zones?: any[];
}
export interface EmailSendingSuppressionRemoveMatch {
    account_id?: string;
    id: string;
    zone_id?: string;
}
export interface Subdomain {
    id?: string;
}
export interface SubdomainRemoveMatch {
    id: string;
    zone_id: string;
}
