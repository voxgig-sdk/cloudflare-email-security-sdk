import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailRulesResponseCollection, EmailRulesResponseCollectionListMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailRulesResponseCollectionEntity extends CloudflareEmailSecurityEntityBase<EmailRulesResponseCollection> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailRulesResponseCollectionEntity): EmailRulesResponseCollectionEntity;
    list(this: any, reqmatch?: EmailRulesResponseCollectionListMatch, ctrl?: Control): Promise<EmailRulesResponseCollectionEntity[]>;
}
export { EmailRulesResponseCollectionEntity };
