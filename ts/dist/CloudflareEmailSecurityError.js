"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CloudflareEmailSecurityError = void 0;
class CloudflareEmailSecurityError extends Error {
    isCloudflareEmailSecurityError = true;
    sdk = 'CloudflareEmailSecurity';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.CloudflareEmailSecurityError = CloudflareEmailSecurityError;
//# sourceMappingURL=CloudflareEmailSecurityError.js.map