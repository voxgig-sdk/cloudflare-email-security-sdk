# EmailSecurity entity test

require "minitest/autorun"
require "json"
require_relative "../CloudflareEmailSecurity_sdk"
require_relative "runner"

class EmailSecurityEntityTest < Minitest::Test
  def test_create_instance
    testsdk = CloudflareEmailSecuritySDK.test(nil, nil)
    ent = testsdk.EmailSecurity(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "email_security" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = CloudflareEmailSecuritySDK.test(seed, nil)
    seen = base.EmailSecurity(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = CloudflareEmailSecurityConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = CloudflareEmailSecuritySDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.EmailSecurity(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = email_security_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "email_security." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    email_security_ref01_ent = client.EmailSecurity(nil)
    email_security_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.email_security"), "email_security_ref01"))
    email_security_ref01_data["account_id"] = setup[:idmap]["account01"]
    email_security_ref01_data["bulk_id"] = setup[:idmap]["bulk01"]
    email_security_ref01_data["investigate_id"] = setup[:idmap]["investigate01"]

    email_security_ref01_data_result = email_security_ref01_ent.create(email_security_ref01_data, nil)
    email_security_ref01_data = Helpers.to_map(email_security_ref01_data_result.respond_to?(:data_get) ? email_security_ref01_data_result.data_get : email_security_ref01_data_result)
    assert !email_security_ref01_data.nil?
    assert !email_security_ref01_data["id"].nil?

    # LIST
    email_security_ref01_match = {
      "account_id" => setup[:idmap]["account01"],
      "investigate_id" => setup[:idmap]["investigate01"],
    }

    email_security_ref01_list_result = email_security_ref01_ent.list(email_security_ref01_match, nil)
    assert email_security_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(email_security_ref01_list_result),
      { "id" => email_security_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # LOAD
    email_security_ref01_match_dt0 = {
      "id" => email_security_ref01_data["id"],
    }
    email_security_ref01_data_dt0_loaded = email_security_ref01_ent.load(email_security_ref01_match_dt0, nil)
    email_security_ref01_data_dt0_load_result = Helpers.to_map(email_security_ref01_data_dt0_loaded.respond_to?(:data_get) ? email_security_ref01_data_dt0_loaded.data_get : email_security_ref01_data_dt0_loaded)
    assert !email_security_ref01_data_dt0_load_result.nil?
    assert_equal email_security_ref01_data_dt0_load_result["id"], email_security_ref01_data["id"]

    # REMOVE
    email_security_ref01_match_rm0 = {
      "id" => email_security_ref01_data["id"],
    }
    email_security_ref01_ent.remove(email_security_ref01_match_rm0, nil)

    # LIST
    email_security_ref01_match_rt0 = {
      "account_id" => setup[:idmap]["account01"],
      "investigate_id" => setup[:idmap]["investigate01"],
    }

    email_security_ref01_list_rt0_result = email_security_ref01_ent.list(email_security_ref01_match_rt0, nil)
    assert email_security_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(email_security_ref01_list_rt0_result),
      { "id" => email_security_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def email_security_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "email_security", "EmailSecurityTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CloudflareEmailSecuritySDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["email_security01", "email_security02", "email_security03", "account01", "bulk01", "investigate01"],
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
  entid_env_raw = ENV["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID" => idmap,
    "CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE" => "FALSE",
    "CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN" => "FALSE",
    "CLOUDFLARE_EMAIL_SECURITY_APIKEY" => "",
  })

  idmap_resolved = Helpers.to_map(
    env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
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
