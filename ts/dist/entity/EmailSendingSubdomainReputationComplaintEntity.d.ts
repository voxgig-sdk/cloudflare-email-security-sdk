import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSubdomainReputationComplaint, EmailSendingSubdomainReputationComplaintLoadMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSubdomainReputationComplaintEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSubdomainReputationComplaint> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSubdomainReputationComplaintEntity): EmailSendingSubdomainReputationComplaintEntity;
    load(this: any, reqmatch?: EmailSendingSubdomainReputationComplaintLoadMatch, ctrl?: Control): Promise<EmailSendingSubdomainReputationComplaintEntity>;
}
export { EmailSendingSubdomainReputationComplaintEntity };
