import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingLimitsResponseSingle, EmailSendingLimitsResponseSingleLoadMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingLimitsResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailSendingLimitsResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingLimitsResponseSingleEntity): EmailSendingLimitsResponseSingleEntity;
    load(this: any, reqmatch?: EmailSendingLimitsResponseSingleLoadMatch, ctrl?: Control): Promise<EmailSendingLimitsResponseSingleEntity>;
}
export { EmailSendingLimitsResponseSingleEntity };
