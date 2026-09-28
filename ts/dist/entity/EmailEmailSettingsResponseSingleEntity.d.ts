import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailEmailSettingsResponseSingle, EmailEmailSettingsResponseSingleLoadMatch, EmailEmailSettingsResponseSingleCreateData, EmailEmailSettingsResponseSingleUpdateData, EmailEmailSettingsResponseSingleRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailEmailSettingsResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailEmailSettingsResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailEmailSettingsResponseSingleEntity): EmailEmailSettingsResponseSingleEntity;
    load(this: any, reqmatch?: EmailEmailSettingsResponseSingleLoadMatch, ctrl?: Control): Promise<EmailEmailSettingsResponseSingleEntity>;
    create(this: any, reqdata?: EmailEmailSettingsResponseSingleCreateData, ctrl?: Control): Promise<EmailEmailSettingsResponseSingleEntity>;
    update(this: any, reqdata?: EmailEmailSettingsResponseSingleUpdateData, ctrl?: Control): Promise<EmailEmailSettingsResponseSingleEntity>;
    remove(this: any, reqmatch?: EmailEmailSettingsResponseSingleRemoveMatch, ctrl?: Control): Promise<EmailEmailSettingsResponseSingleEntity>;
}
export { EmailEmailSettingsResponseSingleEntity };
