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

func TestEmailDestinationAddressResponseSingleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailDestinationAddressResponseSingle(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailDestinationAddressResponseSingleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_destination_address_response_singleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_destination_address_response_single." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailDestinationAddressResponseSingleRef01Ent := client.EmailDestinationAddressResponseSingle(nil)
		emailDestinationAddressResponseSingleRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_destination_address_response_single"}), "email_destination_address_response_single_ref01"))
		emailDestinationAddressResponseSingleRef01Data["account_id"] = setup.idmap["account01"]

		emailDestinationAddressResponseSingleRef01DataResult, err := emailDestinationAddressResponseSingleRef01Ent.Create(emailDestinationAddressResponseSingleRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailDestinationAddressResponseSingleRef01Data = core.ToMapAny(entityData(emailDestinationAddressResponseSingleRef01DataResult))
		if emailDestinationAddressResponseSingleRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// UPDATE
		emailDestinationAddressResponseSingleRef01DataUp0Up := map[string]any{
			"account_id": setup.idmap["account_id"],
		}

		emailDestinationAddressResponseSingleRef01MarkdefUp0Name := "email"
		emailDestinationAddressResponseSingleRef01MarkdefUp0Value := fmt.Sprintf("Mark01-email_destination_address_response_single_ref01_%d", setup.now)
		emailDestinationAddressResponseSingleRef01DataUp0Up[emailDestinationAddressResponseSingleRef01MarkdefUp0Name] = emailDestinationAddressResponseSingleRef01MarkdefUp0Value

		emailDestinationAddressResponseSingleRef01ResdataUp0Result, err := emailDestinationAddressResponseSingleRef01Ent.Update(emailDestinationAddressResponseSingleRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		emailDestinationAddressResponseSingleRef01ResdataUp0 := core.ToMapAny(entityData(emailDestinationAddressResponseSingleRef01ResdataUp0Result))
		if emailDestinationAddressResponseSingleRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if emailDestinationAddressResponseSingleRef01ResdataUp0[emailDestinationAddressResponseSingleRef01MarkdefUp0Name] != emailDestinationAddressResponseSingleRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", emailDestinationAddressResponseSingleRef01MarkdefUp0Name, emailDestinationAddressResponseSingleRef01ResdataUp0[emailDestinationAddressResponseSingleRef01MarkdefUp0Name])
		}

		// LOAD
		emailDestinationAddressResponseSingleRef01MatchDt0 := map[string]any{}
		emailDestinationAddressResponseSingleRef01DataDt0Loaded, err := emailDestinationAddressResponseSingleRef01Ent.Load(emailDestinationAddressResponseSingleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if emailDestinationAddressResponseSingleRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}


	})
}

func email_destination_address_response_singleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_destination_address_response_single", "EmailDestinationAddressResponseSingleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_destination_address_response_single test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_destination_address_response_single test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_destination_address_response_single01", "email_destination_address_response_single02", "email_destination_address_response_single03", "account01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add account_id alias for update test.
	if idmapResolved["account_id"] == nil {
		idmapResolved["account_id"] = idmapResolved["account01"]
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
