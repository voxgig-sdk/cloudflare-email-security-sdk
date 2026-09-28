-- EmailRoutingSuppression entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("cloudflare-email-security_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("EmailRoutingSuppressionEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:EmailRoutingSuppression(nil)
    assert.is_not_nil(ent)
  end)

  -- Feature #4: the entity stream(action, ...) method runs the op pipeline and
  -- returns an iterator over result items. With the streaming feature active it
  -- yields the feature's incremental output; otherwise it falls back to the
  -- materialised list so stream always yields.
  it("should stream", function()
    local seed = {
      entity = {
        ["email_routing_suppression"] = {
          s1 = { id = "s1" },
          s2 = { id = "s2" },
          s3 = { id = "s3" },
        },
      },
    }

    -- Fallback: streaming inactive -> yields the materialised list items.
    local base = sdk.test(seed, nil)
    local seen = {}
    for item in base:EmailRoutingSuppression(nil):stream("list", nil, nil) do
      table.insert(seen, item)
    end
    assert.are.equal(3, #seen)

    -- Inbound: streaming active -> yields each item from the feature.
    local config = require("config_shared")()
    if type(config.feature) == "table" and config.feature.streaming ~= nil then
      local streamsdk = sdk.test(seed, { feature = { streaming = { active = true } } })
      local got = {}
      for item in streamsdk:EmailRoutingSuppression(nil):stream("list", nil, nil) do
        if vs.islist(item) then
          for _, sub in ipairs(item) do
            table.insert(got, sub)
          end
        else
          table.insert(got, item)
        end
      end
      assert.are.equal(3, #got)
    end
  end)

  it("should run basic flow", function()
    local setup = email_routing_suppression_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create", "list", "load", "remove"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "email_routing_suppression." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local email_routing_suppression_ref01_ent = client:EmailRoutingSuppression(nil)
    local email_routing_suppression_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.email_routing_suppression"), "email_routing_suppression_ref01"))
    email_routing_suppression_ref01_data["account_id"] = setup.idmap["account01"]
    email_routing_suppression_ref01_data["zone_id"] = setup.idmap["zone01"]

    local email_routing_suppression_ref01_data_result, err = email_routing_suppression_ref01_ent:create(email_routing_suppression_ref01_data, nil)
    assert.is_nil(err)
    email_routing_suppression_ref01_data = helpers.to_map(type(email_routing_suppression_ref01_data_result) == 'table' and email_routing_suppression_ref01_data_result.data_get and email_routing_suppression_ref01_data_result:data_get() or email_routing_suppression_ref01_data_result)
    assert.is_not_nil(email_routing_suppression_ref01_data)
    assert.is_not_nil(email_routing_suppression_ref01_data["id"])

    -- LIST
    local email_routing_suppression_ref01_match = {
      ["zone_id"] = setup.idmap["zone01"],
    }

    local email_routing_suppression_ref01_list_result, err = email_routing_suppression_ref01_ent:list(email_routing_suppression_ref01_match, nil)
    assert.is_nil(err)
    assert.is_table(email_routing_suppression_ref01_list_result)

    local found_item = vs.select(
      runner.entity_list_to_data(email_routing_suppression_ref01_list_result),
      { id = email_routing_suppression_ref01_data["id"] })
    assert.is_false(vs.isempty(found_item))

    -- LOAD
    local email_routing_suppression_ref01_match_dt0 = {
      id = email_routing_suppression_ref01_data["id"],
    }
    local email_routing_suppression_ref01_data_dt0_loaded, err = email_routing_suppression_ref01_ent:load(email_routing_suppression_ref01_match_dt0, nil)
    assert.is_nil(err)
    local email_routing_suppression_ref01_data_dt0_load_result = helpers.to_map(type(email_routing_suppression_ref01_data_dt0_loaded) == 'table' and email_routing_suppression_ref01_data_dt0_loaded.data_get and email_routing_suppression_ref01_data_dt0_loaded:data_get() or email_routing_suppression_ref01_data_dt0_loaded)
    assert.is_not_nil(email_routing_suppression_ref01_data_dt0_load_result)
    assert.are.equal(email_routing_suppression_ref01_data_dt0_load_result["id"], email_routing_suppression_ref01_data["id"])

    -- REMOVE
    local email_routing_suppression_ref01_match_rm0 = {
      id = email_routing_suppression_ref01_data["id"],
    }
    local _, err = email_routing_suppression_ref01_ent:remove(email_routing_suppression_ref01_match_rm0, nil)
    assert.is_nil(err)

    -- LIST
    local email_routing_suppression_ref01_match_rt0 = {
      ["zone_id"] = setup.idmap["zone01"],
    }

    local email_routing_suppression_ref01_list_rt0_result, err = email_routing_suppression_ref01_ent:list(email_routing_suppression_ref01_match_rt0, nil)
    assert.is_nil(err)
    assert.is_table(email_routing_suppression_ref01_list_rt0_result)

    local not_found_item = vs.select(
      runner.entity_list_to_data(email_routing_suppression_ref01_list_rt0_result),
      { id = email_routing_suppression_ref01_data["id"] })
    assert.is_true(vs.isempty(not_found_item))

  end)
end)

function email_routing_suppression_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/email_routing_suppression/EmailRoutingSuppressionTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read email_routing_suppression test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "email_routing_suppression01", "email_routing_suppression02", "email_routing_suppression03", "account01", "zone01" },
    {
      ["`$PACK`"] = { "", {
        ["`$KEY`"] = "`$COPY`",
        ["`$VAL`"] = { "`$FORMAT`", "upper", "`$COPY`" },
      }},
    }
  )

  -- Detect ENTID env override before envOverride consumes it. When live
  -- mode is on without a real override, the basic test runs against synthetic
  -- IDs from the fixture and 4xx's. Surface this so the test can skip.
  local entid_env_raw = os.getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID"] = idmap,
    ["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] = "FALSE",
    ["CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN"] = "FALSE",
    ["CLOUDFLARE_EMAIL_SECURITY_APIKEY"] = "",
  })

  local idmap_resolved = helpers.to_map(
    env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
  end

  if env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE" then
    local merged_opts = vs.merge({
      -- FIRST, so the generated fields below win: sdk-test-control.json's
      -- test.client.options adds to the live client, it does not redirect it.
      runner.live_client_options(),
      {
        apikey = env["CLOUDFLARE_EMAIL_SECURITY_APIKEY"],
      },
      extra or {},
    })
    client = sdk.new(helpers.to_map(merged_opts))
  end

  local live = env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE"
  return {
    client = client,
    data = entity_data,
    idmap = idmap_resolved,
    env = env,
    explain = env["CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN"] == "TRUE",
    live = live,
    synthetic_only = live and not idmap_overridden,
    now = os.time() * 1000,
  }
end
