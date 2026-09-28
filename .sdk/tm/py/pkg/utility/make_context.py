# CloudflareEmailSecurity SDK utility: make_context

from projectname_sdk.core.context import CloudflareEmailSecurityContext


def make_context_util(ctxmap, basectx):
    return CloudflareEmailSecurityContext(ctxmap, basectx)
