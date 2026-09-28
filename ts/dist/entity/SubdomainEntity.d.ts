import { CloudflareEmailSecurityEntityBase } from '../CloudflareEmailSecurityEntityBase';
import type { CloudflareEmailSecuritySDK } from '../CloudflareEmailSecuritySDK';
import type { Control } from '../types';
import type { Subdomain, SubdomainRemoveMatch } from '../CloudflareEmailSecurityTypes';
declare class SubdomainEntity extends CloudflareEmailSecurityEntityBase<Subdomain> {
    constructor(client: CloudflareEmailSecuritySDK, entopts: any);
    make(this: SubdomainEntity): SubdomainEntity;
    remove(this: any, reqmatch?: SubdomainRemoveMatch, ctrl?: Control): Promise<SubdomainEntity>;
}
export { SubdomainEntity };
