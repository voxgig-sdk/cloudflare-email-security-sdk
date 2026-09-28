# EmailSendingReputationResponseSingle entity test

import json
import os
import time

import pytest

from cloudflareemailsecurity_sdk.utility.voxgig_struct import voxgig_struct as vs
from cloudflareemailsecurity_sdk import CloudflareEmailSecuritySDK
from cloudflareemailsecurity_sdk.core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestEmailSendingReputationResponseSingleEntity:

    def test_should_create_instance(self):
        testsdk = CloudflareEmailSecuritySDK.test(None, None)
        ent = testsdk.EmailSendingReputationResponseSingle(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _email_sending_reputation_response_single_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "email_sending_reputation_response_single." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_REPUTATION_RESPONSE_SINGLE_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        email_sending_reputation_response_single_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.email_sending_reputation_response_single")))
        email_sending_reputation_response_single_ref01_data = None
        if len(email_sending_reputation_response_single_ref01_data_raw) > 0:
            email_sending_reputation_response_single_ref01_data = helpers.to_map(email_sending_reputation_response_single_ref01_data_raw[0][1])

        # LOAD
        email_sending_reputation_response_single_ref01_ent = client.EmailSendingReputationResponseSingle(None)
        email_sending_reputation_response_single_ref01_match_dt0 = {}
        email_sending_reputation_response_single_ref01_data_dt0_loaded = email_sending_reputation_response_single_ref01_ent.load(email_sending_reputation_response_single_ref01_match_dt0, None)
        assert email_sending_reputation_response_single_ref01_data_dt0_loaded is not None



def _email_sending_reputation_response_single_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/email_sending_reputation_response_single/EmailSendingReputationResponseSingleTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = CloudflareEmailSecuritySDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["email_sending_reputation_response_single01", "email_sending_reputation_response_single02", "email_sending_reputation_response_single03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_REPUTATION_RESPONSE_SINGLE_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_REPUTATION_RESPONSE_SINGLE_ENTID": idmap,
        "CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE": "FALSE",
        "CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN": "FALSE",
        "CLOUDFLARE_EMAIL_SECURITY_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_REPUTATION_RESPONSE_SINGLE_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("CLOUDFLARE_EMAIL_SECURITY_APIKEY"),
            },
            extra or {},
        ])
        client = CloudflareEmailSecuritySDK(helpers.to_map(merged_opts))

    _live = env.get("CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
