import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailRuleResponseSingle, EmailRuleResponseSingleLoadMatch, EmailRuleResponseSingleCreateData, EmailRuleResponseSingleUpdateData, EmailRuleResponseSingleRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailRuleResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailRuleResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailRuleResponseSingleEntity): EmailRuleResponseSingleEntity;
    load(this: any, reqmatch?: EmailRuleResponseSingleLoadMatch, ctrl?: Control): Promise<EmailRuleResponseSingleEntity>;
    create(this: any, reqdata?: EmailRuleResponseSingleCreateData, ctrl?: Control): Promise<EmailRuleResponseSingleEntity>;
    update(this: any, reqdata?: EmailRuleResponseSingleUpdateData, ctrl?: Control): Promise<EmailRuleResponseSingleEntity>;
    remove(this: any, reqmatch?: EmailRuleResponseSingleRemoveMatch, ctrl?: Control): Promise<EmailRuleResponseSingleEntity>;
}
export { EmailRuleResponseSingleEntity };
