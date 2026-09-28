import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSubdomainPreview, EmailSendingSubdomainPreviewCreateData } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSubdomainPreviewEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSubdomainPreview> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSubdomainPreviewEntity): EmailSendingSubdomainPreviewEntity;
    create(this: any, reqdata?: EmailSendingSubdomainPreviewCreateData, ctrl?: Control): Promise<EmailSendingSubdomainPreviewEntity>;
}
export { EmailSendingSubdomainPreviewEntity };
