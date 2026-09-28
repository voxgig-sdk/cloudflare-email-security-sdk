# CloudflareEmailSecurity SDK utility: make_fetch_def

from __future__ import annotations
from cloudflareemailsecurity_sdk.utility.voxgig_struct import voxgig_struct as vs


def make_fetch_def_util(ctx):
    spec = ctx.spec
    if spec is None:
        return None, ctx.make_error("fetchdef_no_spec",
            "Expected context spec property to be defined.")

    from cloudflareemailsecurity_sdk.core.result import CloudflareEmailSecurityResult
    if ctx.result is None:
        ctx.result = CloudflareEmailSecurityResult({})

    spec.step = "prepare"

    url, err = ctx.utility.make_url(ctx)
    if err is not None:
        return None, err

    spec.url = url

    fetchdef = {
        "url": url,
        "method": spec.method,
        "headers": spec.headers,
    }

    if spec.body is not None:
        if isinstance(spec.body, dict):
            fetchdef["body"] = vs.jsonify(spec.body)
        else:
            fetchdef["body"] = spec.body

    return fetchdef, None
