-- CloudflareEmailSecurity SDK exists test

local sdk = require("cloudflare-email-security_sdk")

describe("CloudflareEmailSecuritySDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
