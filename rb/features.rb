# CloudflareEmailSecurity SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module CloudflareEmailSecurityFeatures
  def self.make_feature(name)
    case name
    when "base"
      CloudflareEmailSecurityBaseFeature.new
    when "test"
      CloudflareEmailSecurityTestFeature.new
    else
      CloudflareEmailSecurityBaseFeature.new
    end
  end
end
