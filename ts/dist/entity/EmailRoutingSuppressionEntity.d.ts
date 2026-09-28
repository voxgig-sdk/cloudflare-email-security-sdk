import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailRoutingSuppression, EmailRoutingSuppressionLoadMatch, EmailRoutingSuppressionListMatch, EmailRoutingSuppressionCreateData, EmailRoutingSuppressionRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailRoutingSuppressionEntity extends CloudflareEmailSecurityEntityBase<EmailRoutingSuppression> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailRoutingSuppressionEntity): EmailRoutingSuppressionEntity;
    load(this: any, reqmatch?: EmailRoutingSuppressionLoadMatch, ctrl?: Control): Promise<EmailRoutingSuppressionEntity>;
    list(this: any, reqmatch?: EmailRoutingSuppressionListMatch, ctrl?: Control): Promise<EmailRoutingSuppressionEntity[]>;
    create(this: any, reqdata?: EmailRoutingSuppressionCreateData, ctrl?: Control): Promise<EmailRoutingSuppressionEntity>;
    remove(this: any, reqmatch?: EmailRoutingSuppressionRemoveMatch, ctrl?: Control): Promise<EmailRoutingSuppressionEntity>;
}
export { EmailRoutingSuppressionEntity };
