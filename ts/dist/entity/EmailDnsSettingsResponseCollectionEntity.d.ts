import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailDnsSettingsResponseCollection, EmailDnsSettingsResponseCollectionListMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailDnsSettingsResponseCollectionEntity extends CloudflareEmailSecurityEntityBase<EmailDnsSettingsResponseCollection> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailDnsSettingsResponseCollectionEntity): EmailDnsSettingsResponseCollectionEntity;
    list(this: any, reqmatch?: EmailDnsSettingsResponseCollectionListMatch, ctrl?: Control): Promise<EmailDnsSettingsResponseCollectionEntity[]>;
}
export { EmailDnsSettingsResponseCollectionEntity };
