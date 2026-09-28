import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSending, EmailSendingLoadMatch, EmailSendingCreateData } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingEntity extends CloudflareEmailSecurityEntityBase<EmailSending> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingEntity): EmailSendingEntity;
    load(this: any, reqmatch?: EmailSendingLoadMatch, ctrl?: Control): Promise<EmailSendingEntity>;
    create(this: any, reqdata?: EmailSendingCreateData, ctrl?: Control): Promise<EmailSendingEntity>;
}
export { EmailSendingEntity };
