import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingReputationResponseSingle, EmailSendingReputationResponseSingleLoadMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingReputationResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailSendingReputationResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingReputationResponseSingleEntity): EmailSendingReputationResponseSingleEntity;
    load(this: any, reqmatch?: EmailSendingReputationResponseSingleLoadMatch, ctrl?: Control): Promise<EmailSendingReputationResponseSingleEntity>;
}
export { EmailSendingReputationResponseSingleEntity };
