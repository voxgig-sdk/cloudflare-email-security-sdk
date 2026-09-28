import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailCatchAllRuleResponseSingle, EmailCatchAllRuleResponseSingleLoadMatch, EmailCatchAllRuleResponseSingleUpdateData } from '../CloudflareEmailSecurityTypes';
declare class EmailCatchAllRuleResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailCatchAllRuleResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailCatchAllRuleResponseSingleEntity): EmailCatchAllRuleResponseSingleEntity;
    load(this: any, reqmatch?: EmailCatchAllRuleResponseSingleLoadMatch, ctrl?: Control): Promise<EmailCatchAllRuleResponseSingleEntity>;
    update(this: any, reqdata?: EmailCatchAllRuleResponseSingleUpdateData, ctrl?: Control): Promise<EmailCatchAllRuleResponseSingleEntity>;
}
export { EmailCatchAllRuleResponseSingleEntity };
