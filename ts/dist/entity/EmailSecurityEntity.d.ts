import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSecurity, EmailSecurityLoadMatch, EmailSecurityListMatch, EmailSecurityCreateData, EmailSecurityRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSecurityEntity extends CloudflareEmailSecurityEntityBase<EmailSecurity> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSecurityEntity): EmailSecurityEntity;
    load(this: any, reqmatch?: EmailSecurityLoadMatch, ctrl?: Control): Promise<EmailSecurityEntity>;
    list(this: any, reqmatch?: EmailSecurityListMatch, ctrl?: Control): Promise<EmailSecurityEntity[]>;
    create(this: any, reqdata?: EmailSecurityCreateData, ctrl?: Control): Promise<EmailSecurityEntity>;
    remove(this: any, reqmatch?: EmailSecurityRemoveMatch, ctrl?: Control): Promise<EmailSecurityEntity>;
}
export { EmailSecurityEntity };
