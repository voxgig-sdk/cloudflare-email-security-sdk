-- CloudflareEmailSecurity SDK error

local CloudflareEmailSecurityError = {}
CloudflareEmailSecurityError.__index = CloudflareEmailSecurityError


function CloudflareEmailSecurityError.new(code, msg, ctx)
  local self = setmetatable({}, CloudflareEmailSecurityError)
  self.is_sdk_error = true
  self.sdk = "CloudflareEmailSecurity"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function CloudflareEmailSecurityError:error()
  return self.msg
end


function CloudflareEmailSecurityError:__tostring()
  return self.msg
end


return CloudflareEmailSecurityError
