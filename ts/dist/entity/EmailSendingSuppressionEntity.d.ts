import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSendingSuppression, EmailSendingSuppressionLoadMatch, EmailSendingSuppressionListMatch, EmailSendingSuppressionCreateData, EmailSendingSuppressionUpdateData, EmailSendingSuppressionRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSendingSuppressionEntity extends CloudflareEmailSecurityEntityBase<EmailSendingSuppression> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSendingSuppressionEntity): EmailSendingSuppressionEntity;
    load(this: any, reqmatch?: EmailSendingSuppressionLoadMatch, ctrl?: Control): Promise<EmailSendingSuppressionEntity>;
    list(this: any, reqmatch?: EmailSendingSuppressionListMatch, ctrl?: Control): Promise<EmailSendingSuppressionEntity[]>;
    create(this: any, reqdata?: EmailSendingSuppressionCreateData, ctrl?: Control): Promise<EmailSendingSuppressionEntity>;
    update(this: any, reqdata?: EmailSendingSuppressionUpdateData, ctrl?: Control): Promise<EmailSendingSuppressionEntity>;
    remove(this: any, reqmatch?: EmailSendingSuppressionRemoveMatch, ctrl?: Control): Promise<EmailSendingSuppressionEntity>;
}
export { EmailSendingSuppressionEntity };
