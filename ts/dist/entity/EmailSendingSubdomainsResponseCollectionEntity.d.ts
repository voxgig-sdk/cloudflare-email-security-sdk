import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSubdomainsResponseCollection, EmailSendingSubdomainsResponseCollectionListMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSubdomainsResponseCollectionEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSubdomainsResponseCollection> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSubdomainsResponseCollectionEntity): EmailSendingSubdomainsResponseCollectionEntity;
    list(this: any, reqmatch?: EmailSendingSubdomainsResponseCollectionListMatch, ctrl?: Control): Promise<EmailSendingSubdomainsResponseCollectionEntity[]>;
}
export { EmailSendingSubdomainsResponseCollectionEntity };
