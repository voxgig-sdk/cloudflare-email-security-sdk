import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailAuth, EmailAuthLoadMatch, EmailAuthUpdateData } from '../CloudflareEmailSecurityTypes';
declare class EmailAuthEntity extends CloudflareEmailSecurityEntityBase<EmailAuth> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailAuthEntity): EmailAuthEntity;
    load(this: any, reqmatch?: EmailAuthLoadMatch, ctrl?: Control): Promise<EmailAuthEntity>;
    update(this: any, reqdata?: EmailAuthUpdateData, ctrl?: Control): Promise<EmailAuthEntity>;
}
export { EmailAuthEntity };
