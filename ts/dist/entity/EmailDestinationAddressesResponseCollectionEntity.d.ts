import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailDestinationAddressesResponseCollection, EmailDestinationAddressesResponseCollectionListMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailDestinationAddressesResponseCollectionEntity extends CloudflareEmailSecurityEntityBase<EmailDestinationAddressesResponseCollection> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailDestinationAddressesResponseCollectionEntity): EmailDestinationAddressesResponseCollectionEntity;
    list(this: any, reqmatch?: EmailDestinationAddressesResponseCollectionListMatch, ctrl?: Control): Promise<EmailDestinationAddressesResponseCollectionEntity[]>;
}
export { EmailDestinationAddressesResponseCollectionEntity };
