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

func TestEmailEmailSettingsResponseSingleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailEmailSettingsResponseSingle(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailEmailSettingsResponseSingleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_email_settings_response_singleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_email_settings_response_single." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailEmailSettingsResponseSingleRef01Ent := client.EmailEmailSettingsResponseSingle(nil)
		emailEmailSettingsResponseSingleRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_email_settings_response_single"}), "email_email_settings_response_single_ref01"))
		emailEmailSettingsResponseSingleRef01Data["zone_id"] = setup.idmap["zone01"]

		emailEmailSettingsResponseSingleRef01DataResult, err := emailEmailSettingsResponseSingleRef01Ent.Create(emailEmailSettingsResponseSingleRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailEmailSettingsResponseSingleRef01Data = core.ToMapAny(entityData(emailEmailSettingsResponseSingleRef01DataResult))
		if emailEmailSettingsResponseSingleRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// UPDATE
		emailEmailSettingsResponseSingleRef01DataUp0Up := map[string]any{
		}

		emailEmailSettingsResponseSingleRef01MarkdefUp0Name := "name"
		emailEmailSettingsResponseSingleRef01MarkdefUp0Value := fmt.Sprintf("Mark01-email_email_settings_response_single_ref01_%d", setup.now)
		emailEmailSettingsResponseSingleRef01DataUp0Up[emailEmailSettingsResponseSingleRef01MarkdefUp0Name] = emailEmailSettingsResponseSingleRef01MarkdefUp0Value

		emailEmailSettingsResponseSingleRef01ResdataUp0Result, err := emailEmailSettingsResponseSingleRef01Ent.Update(emailEmailSettingsResponseSingleRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		emailEmailSettingsResponseSingleRef01ResdataUp0 := core.ToMapAny(entityData(emailEmailSettingsResponseSingleRef01ResdataUp0Result))
		if emailEmailSettingsResponseSingleRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if emailEmailSettingsResponseSingleRef01ResdataUp0[emailEmailSettingsResponseSingleRef01MarkdefUp0Name] != emailEmailSettingsResponseSingleRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", emailEmailSettingsResponseSingleRef01MarkdefUp0Name, emailEmailSettingsResponseSingleRef01ResdataUp0[emailEmailSettingsResponseSingleRef01MarkdefUp0Name])
		}

		// LOAD
		emailEmailSettingsResponseSingleRef01MatchDt0 := map[string]any{}
		emailEmailSettingsResponseSingleRef01DataDt0Loaded, err := emailEmailSettingsResponseSingleRef01Ent.Load(emailEmailSettingsResponseSingleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if emailEmailSettingsResponseSingleRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}


	})
}

func email_email_settings_response_singleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_email_settings_response_single", "EmailEmailSettingsResponseSingleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_email_settings_response_single test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_email_settings_response_single test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_email_settings_response_single01", "email_email_settings_response_single02", "email_email_settings_response_single03", "zone01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID"])
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
