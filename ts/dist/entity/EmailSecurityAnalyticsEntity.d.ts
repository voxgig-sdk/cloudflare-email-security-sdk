import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSecurityAnalytics, EmailSecurityAnalyticsLoadMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSecurityAnalyticsEntity extends CloudflareEmailSecurityEntityBase<EmailSecurityAnalytics> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSecurityAnalyticsEntity): EmailSecurityAnalyticsEntity;
    load(this: any, reqmatch?: EmailSecurityAnalyticsLoadMatch, ctrl?: Control): Promise<EmailSecurityAnalyticsEntity>;
}
export { EmailSecurityAnalyticsEntity };
