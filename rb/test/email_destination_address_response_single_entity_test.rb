# EmailDestinationAddressResponseSingle entity test

require "minitest/autorun"
require "json"
require_relative "../CloudflareEmailSecurity_sdk"
require_relative "runner"

class EmailDestinationAddressResponseSingleEntityTest < Minitest::Test
  def test_create_instance
    testsdk = CloudflareEmailSecuritySDK.test(nil, nil)
    ent = testsdk.EmailDestinationAddressResponseSingle(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = email_destination_address_response_single_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "email_destination_address_response_single." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    email_destination_address_response_single_ref01_ent = client.EmailDestinationAddressResponseSingle(nil)
    email_destination_address_response_single_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.email_destination_address_response_single"), "email_destination_address_response_single_ref01"))
    email_destination_address_response_single_ref01_data["account_id"] = setup[:idmap]["account01"]

    email_destination_address_response_single_ref01_data_result = email_destination_address_response_single_ref01_ent.create(email_destination_address_response_single_ref01_data, nil)
    email_destination_address_response_single_ref01_data = Helpers.to_map(email_destination_address_response_single_ref01_data_result.respond_to?(:data_get) ? email_destination_address_response_single_ref01_data_result.data_get : email_destination_address_response_single_ref01_data_result)
    assert !email_destination_address_response_single_ref01_data.nil?

    # UPDATE
    email_destination_address_response_single_ref01_data_up0_up = {
      "account_id" => setup[:idmap]["account_id"],
    }

    email_destination_address_response_single_ref01_markdef_up0_name = "email"
    email_destination_address_response_single_ref01_markdef_up0_value = "Mark01-email_destination_address_response_single_ref01_#{setup[:now]}"
    email_destination_address_response_single_ref01_data_up0_up[email_destination_address_response_single_ref01_markdef_up0_name] = email_destination_address_response_single_ref01_markdef_up0_value

    email_destination_address_response_single_ref01_resdata_up0_result = email_destination_address_response_single_ref01_ent.update(email_destination_address_response_single_ref01_data_up0_up, nil)
    email_destination_address_response_single_ref01_resdata_up0 = Helpers.to_map(email_destination_address_response_single_ref01_resdata_up0_result.respond_to?(:data_get) ? email_destination_address_response_single_ref01_resdata_up0_result.data_get : email_destination_address_response_single_ref01_resdata_up0_result)
    assert !email_destination_address_response_single_ref01_resdata_up0.nil?
    assert_equal email_destination_address_response_single_ref01_resdata_up0[email_destination_address_response_single_ref01_markdef_up0_name], email_destination_address_response_single_ref01_markdef_up0_value

    # LOAD
    email_destination_address_response_single_ref01_match_dt0 = {}
    email_destination_address_response_single_ref01_data_dt0_loaded = email_destination_address_response_single_ref01_ent.load(email_destination_address_response_single_ref01_match_dt0, nil)
    assert !email_destination_address_response_single_ref01_data_dt0_loaded.nil?


  end
end

def email_destination_address_response_single_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "email_destination_address_response_single", "EmailDestinationAddressResponseSingleTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CloudflareEmailSecuritySDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["email_destination_address_response_single01", "email_destination_address_response_single02", "email_destination_address_response_single03", "account01"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID" => idmap,
    "CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE" => "FALSE",
    "CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN" => "FALSE",
    "CLOUDFLARE_EMAIL_SECURITY_APIKEY" => "",
  })

  idmap_resolved = Helpers.to_map(
    env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end
  if idmap_resolved["account_id"].nil?
    idmap_resolved["account_id"] = idmap_resolved["account01"]
  end

  if env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      # FIRST, so the generated fields below win: sdk-test-control.json's
      # test.client.options adds to the live client, it does not redirect it.
      Runner.live_client_options,
      {
        "apikey" => env["CLOUDFLARE_EMAIL_SECURITY_APIKEY"],
      },
      extra || {},
    ])
    client = CloudflareEmailSecuritySDK.new(Helpers.to_map(merged_opts))
  end

  live = env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
