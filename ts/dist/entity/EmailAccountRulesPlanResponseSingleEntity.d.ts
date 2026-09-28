import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailAccountRulesPlanResponseSingle, EmailAccountRulesPlanResponseSingleCreateData } from '../CloudflareEmailSecurityTypes';
declare class EmailAccountRulesPlanResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailAccountRulesPlanResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailAccountRulesPlanResponseSingleEntity): EmailAccountRulesPlanResponseSingleEntity;
    create(this: any, reqdata?: EmailAccountRulesPlanResponseSingleCreateData, ctrl?: Control): Promise<EmailAccountRulesPlanResponseSingleEntity>;
}
export { EmailAccountRulesPlanResponseSingleEntity };
