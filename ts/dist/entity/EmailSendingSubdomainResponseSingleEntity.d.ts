import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSubdomainResponseSingle, EmailSendingSubdomainResponseSingleLoadMatch, EmailSendingSubdomainResponseSingleCreateData, EmailSendingSubdomainResponseSingleUpdateData } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSubdomainResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSubdomainResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSubdomainResponseSingleEntity): EmailSendingSubdomainResponseSingleEntity;
    load(this: any, reqmatch?: EmailSendingSubdomainResponseSingleLoadMatch, ctrl?: Control): Promise<EmailSendingSubdomainResponseSingleEntity>;
    create(this: any, reqdata?: EmailSendingSubdomainResponseSingleCreateData, ctrl?: Control): Promise<EmailSendingSubdomainResponseSingleEntity>;
    update(this: any, reqdata?: EmailSendingSubdomainResponseSingleUpdateData, ctrl?: Control): Promise<EmailSendingSubdomainResponseSingleEntity>;
}
export { EmailSendingSubdomainResponseSingleEntity };
