import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailAccountRulesResponseCollection, EmailAccountRulesResponseCollectionListMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailAccountRulesResponseCollectionEntity extends CloudflareEmailSecurityEntityBase<EmailAccountRulesResponseCollection> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailAccountRulesResponseCollectionEntity): EmailAccountRulesResponseCollectionEntity;
    list(this: any, reqmatch?: EmailAccountRulesResponseCollectionListMatch, ctrl?: Control): Promise<EmailAccountRulesResponseCollectionEntity[]>;
}
export { EmailAccountRulesResponseCollectionEntity };
