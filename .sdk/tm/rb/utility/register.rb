# CloudflareEmailSecurity SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

CloudflareEmailSecurityUtility.registrar = ->(u) {
  u.clean = CloudflareEmailSecurityUtilities::Clean
  u.done = CloudflareEmailSecurityUtilities::Done
  u.make_error = CloudflareEmailSecurityUtilities::MakeError
  u.feature_add = CloudflareEmailSecurityUtilities::FeatureAdd
  u.feature_hook = CloudflareEmailSecurityUtilities::FeatureHook
  u.feature_init = CloudflareEmailSecurityUtilities::FeatureInit
  u.fetcher = CloudflareEmailSecurityUtilities::Fetcher
  u.make_fetch_def = CloudflareEmailSecurityUtilities::MakeFetchDef
  u.make_context = CloudflareEmailSecurityUtilities::MakeContext
  u.make_options = CloudflareEmailSecurityUtilities::MakeOptions
  u.make_request = CloudflareEmailSecurityUtilities::MakeRequest
  u.make_response = CloudflareEmailSecurityUtilities::MakeResponse
  u.make_result = CloudflareEmailSecurityUtilities::MakeResult
  u.make_point = CloudflareEmailSecurityUtilities::MakePoint
  u.make_spec = CloudflareEmailSecurityUtilities::MakeSpec
  u.make_url = CloudflareEmailSecurityUtilities::MakeUrl
  u.param = CloudflareEmailSecurityUtilities::Param
  u.prepare_auth = CloudflareEmailSecurityUtilities::PrepareAuth
  u.prepare_body = CloudflareEmailSecurityUtilities::PrepareBody
  u.prepare_headers = CloudflareEmailSecurityUtilities::PrepareHeaders
  u.prepare_method = CloudflareEmailSecurityUtilities::PrepareMethod
  u.prepare_params = CloudflareEmailSecurityUtilities::PrepareParams
  u.prepare_path = CloudflareEmailSecurityUtilities::PreparePath
  u.prepare_query = CloudflareEmailSecurityUtilities::PrepareQuery
  u.graphql_body = CloudflareEmailSecurityUtilities::GraphqlBody
  u.graphql_errors = CloudflareEmailSecurityUtilities::GraphqlErrors
  u.result_basic = CloudflareEmailSecurityUtilities::ResultBasic
  u.result_body = CloudflareEmailSecurityUtilities::ResultBody
  u.result_headers = CloudflareEmailSecurityUtilities::ResultHeaders
  u.transform_request = CloudflareEmailSecurityUtilities::TransformRequest
  u.transform_response = CloudflareEmailSecurityUtilities::TransformResponse
}
