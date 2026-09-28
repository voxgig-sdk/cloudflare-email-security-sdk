import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailSecuritySetting, EmailSecuritySettingLoadMatch, EmailSecuritySettingListMatch, EmailSecuritySettingCreateData, EmailSecuritySettingUpdateData, EmailSecuritySettingRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailSecuritySettingEntity extends CloudflareEmailSecurityEntityBase<EmailSecuritySetting> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailSecuritySettingEntity): EmailSecuritySettingEntity;
    load(this: any, reqmatch?: EmailSecuritySettingLoadMatch, ctrl?: Control): Promise<EmailSecuritySettingEntity>;
    list(this: any, reqmatch?: EmailSecuritySettingListMatch, ctrl?: Control): Promise<EmailSecuritySettingEntity[]>;
    create(this: any, reqdata?: EmailSecuritySettingCreateData, ctrl?: Control): Promise<EmailSecuritySettingEntity>;
    update(this: any, reqdata?: EmailSecuritySettingUpdateData, ctrl?: Control): Promise<EmailSecuritySettingEntity>;
    remove(this: any, reqmatch?: EmailSecuritySettingRemoveMatch, ctrl?: Control): Promise<EmailSecuritySettingEntity>;
}
export { EmailSecuritySettingEntity };
