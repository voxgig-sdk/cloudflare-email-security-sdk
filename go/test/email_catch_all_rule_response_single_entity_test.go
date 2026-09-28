package sdktest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/cloudflare-email-security-sdk/go"
	"github.com/voxgig-sdk/cloudflare-email-security-sdk/go/core"

	vs "github.com/voxgig-sdk/cloudflare-email-security-sdk/go/utility/struct"
)

func TestEmailCatchAllRuleResponseSingleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailCatchAllRuleResponseSingle(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailCatchAllRuleResponseSingleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_catch_all_rule_response_singleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_catch_all_rule_response_single." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_CATCH_ALL_RULE_RESPONSE_SINGLE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		emailCatchAllRuleResponseSingleRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath(setup.data, "existing.email_catch_all_rule_response_single")))
		var emailCatchAllRuleResponseSingleRef01Data map[string]any
		if len(emailCatchAllRuleResponseSingleRef01DataRaw) > 0 {
			emailCatchAllRuleResponseSingleRef01Data = core.ToMapAny(emailCatchAllRuleResponseSingleRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = emailCatchAllRuleResponseSingleRef01Data

		// UPDATE
		emailCatchAllRuleResponseSingleRef01Ent := client.EmailCatchAllRuleResponseSingle(nil)
		emailCatchAllRuleResponseSingleRef01DataUp0Up := map[string]any{
		}

		emailCatchAllRuleResponseSingleRef01MarkdefUp0Name := "name"
		emailCatchAllRuleResponseSingleRef01MarkdefUp0Value := fmt.Sprintf("Mark01-email_catch_all_rule_response_single_ref01_%d", setup.now)
		emailCatchAllRuleResponseSingleRef01DataUp0Up[emailCatchAllRuleResponseSingleRef01MarkdefUp0Name] = emailCatchAllRuleResponseSingleRef01MarkdefUp0Value

		emailCatchAllRuleResponseSingleRef01ResdataUp0Result, err := emailCatchAllRuleResponseSingleRef01Ent.Update(emailCatchAllRuleResponseSingleRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		emailCatchAllRuleResponseSingleRef01ResdataUp0 := core.ToMapAny(entityData(emailCatchAllRuleResponseSingleRef01ResdataUp0Result))
		if emailCatchAllRuleResponseSingleRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if emailCatchAllRuleResponseSingleRef01ResdataUp0[emailCatchAllRuleResponseSingleRef01MarkdefUp0Name] != emailCatchAllRuleResponseSingleRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", emailCatchAllRuleResponseSingleRef01MarkdefUp0Name, emailCatchAllRuleResponseSingleRef01ResdataUp0[emailCatchAllRuleResponseSingleRef01MarkdefUp0Name])
		}

		// LOAD
		emailCatchAllRuleResponseSingleRef01MatchDt0 := map[string]any{}
		emailCatchAllRuleResponseSingleRef01DataDt0Loaded, err := emailCatchAllRuleResponseSingleRef01Ent.Load(emailCatchAllRuleResponseSingleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if emailCatchAllRuleResponseSingleRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func email_catch_all_rule_response_singleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_catch_all_rule_response_single", "EmailCatchAllRuleResponseSingleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_catch_all_rule_response_single test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_catch_all_rule_response_single test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_catch_all_rule_response_single01", "email_catch_all_rule_response_single02", "email_catch_all_rule_response_single03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_CATCH_ALL_RULE_RESPONSE_SINGLE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_CATCH_ALL_RULE_RESPONSE_SINGLE_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_CATCH_ALL_RULE_RESPONSE_SINGLE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["CLOUDFLARE_EMAIL_SECURITY_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewCloudflareEmailSecuritySDK(core.ToMapAny(mergedOpts))
	}

	live := env["CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
