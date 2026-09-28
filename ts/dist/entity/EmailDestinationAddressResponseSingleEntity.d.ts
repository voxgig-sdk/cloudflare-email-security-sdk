import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { EmailDestinationAddressResponseSingle, EmailDestinationAddressResponseSingleLoadMatch, EmailDestinationAddressResponseSingleCreateData, EmailDestinationAddressResponseSingleUpdateData, EmailDestinationAddressResponseSingleRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class EmailDestinationAddressResponseSingleEntity extends CloudflareEmailSecurityEntityBase<EmailDestinationAddressResponseSingle> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: EmailDestinationAddressResponseSingleEntity): EmailDestinationAddressResponseSingleEntity;
    load(this: any, reqmatch?: EmailDestinationAddressResponseSingleLoadMatch, ctrl?: Control): Promise<EmailDestinationAddressResponseSingleEntity>;
    create(this: any, reqdata?: EmailDestinationAddressResponseSingleCreateData, ctrl?: Control): Promise<EmailDestinationAddressResponseSingleEntity>;
    update(this: any, reqdata?: EmailDestinationAddressResponseSingleUpdateData, ctrl?: Control): Promise<EmailDestinationAddressResponseSingleEntity>;
    remove(this: any, reqmatch?: EmailDestinationAddressResponseSingleRemoveMatch, ctrl?: Control): Promise<EmailDestinationAddressResponseSingleEntity>;
}
export { EmailDestinationAddressResponseSingleEntity };
