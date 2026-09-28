import { Context } from './Context';
declare class CloudflareEmailSecurityError extends Error {
    isCloudflareEmailSecurityError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { CloudflareEmailSecurityError };
