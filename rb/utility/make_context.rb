# CloudflareEmailSecurity SDK utility: make_context
require_relative '../core/context'
module CloudflareEmailSecurityUtilities
  MakeContext = ->(ctxmap, basectx) {
    CloudflareEmailSecurityContext.new(ctxmap, basectx)
  }
end
