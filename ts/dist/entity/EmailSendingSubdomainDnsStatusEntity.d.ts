import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSubdomainDnsStatus, EmailSendingSubdomainDnsStatusLoadMatch, EmailSendingSubdomainDnsStatusCreateData } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSubdomainDnsStatusEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSubdomainDnsStatus> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSubdomainDnsStatusEntity): EmailSendingSubdomainDnsStatusEntity;
    load(this: any, reqmatch?: EmailSendingSubdomainDnsStatusLoadMatch, ctrl?: Control): Promise<EmailSendingSubdomainDnsStatusEntity>;
    create(this: any, reqdata?: EmailSendingSubdomainDnsStatusCreateData, ctrl?: Control): Promise<EmailSendingSubdomainDnsStatusEntity>;
}
export { EmailSendingSubdomainDnsStatusEntity };
