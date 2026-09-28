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

func TestEmailSendingSubdomainResponseSingleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailSendingSubdomainResponseSingle(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailSendingSubdomainResponseSingleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_sending_subdomain_response_singleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_sending_subdomain_response_single." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_RESPONSE_SINGLE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailSendingSubdomainResponseSingleRef01Ent := client.EmailSendingSubdomainResponseSingle(nil)
		emailSendingSubdomainResponseSingleRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_sending_subdomain_response_single"}), "email_sending_subdomain_response_single_ref01"))
		emailSendingSubdomainResponseSingleRef01Data["zone_id"] = setup.idmap["zone01"]

		emailSendingSubdomainResponseSingleRef01DataResult, err := emailSendingSubdomainResponseSingleRef01Ent.Create(emailSendingSubdomainResponseSingleRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailSendingSubdomainResponseSingleRef01Data = core.ToMapAny(entityData(emailSendingSubdomainResponseSingleRef01DataResult))
		if emailSendingSubdomainResponseSingleRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// UPDATE
		emailSendingSubdomainResponseSingleRef01DataUp0Up := map[string]any{
			"zone_id": setup.idmap["zone_id"],
		}

		emailSendingSubdomainResponseSingleRef01MarkdefUp0Name := "name"
		emailSendingSubdomainResponseSingleRef01MarkdefUp0Value := fmt.Sprintf("Mark01-email_sending_subdomain_response_single_ref01_%d", setup.now)
		emailSendingSubdomainResponseSingleRef01DataUp0Up[emailSendingSubdomainResponseSingleRef01MarkdefUp0Name] = emailSendingSubdomainResponseSingleRef01MarkdefUp0Value

		emailSendingSubdomainResponseSingleRef01ResdataUp0Result, err := emailSendingSubdomainResponseSingleRef01Ent.Update(emailSendingSubdomainResponseSingleRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		emailSendingSubdomainResponseSingleRef01ResdataUp0 := core.ToMapAny(entityData(emailSendingSubdomainResponseSingleRef01ResdataUp0Result))
		if emailSendingSubdomainResponseSingleRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if emailSendingSubdomainResponseSingleRef01ResdataUp0[emailSendingSubdomainResponseSingleRef01MarkdefUp0Name] != emailSendingSubdomainResponseSingleRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", emailSendingSubdomainResponseSingleRef01MarkdefUp0Name, emailSendingSubdomainResponseSingleRef01ResdataUp0[emailSendingSubdomainResponseSingleRef01MarkdefUp0Name])
		}

		// LOAD
		emailSendingSubdomainResponseSingleRef01MatchDt0 := map[string]any{}
		emailSendingSubdomainResponseSingleRef01DataDt0Loaded, err := emailSendingSubdomainResponseSingleRef01Ent.Load(emailSendingSubdomainResponseSingleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if emailSendingSubdomainResponseSingleRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func email_sending_subdomain_response_singleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_sending_subdomain_response_single", "EmailSendingSubdomainResponseSingleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_sending_subdomain_response_single test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_sending_subdomain_response_single test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_sending_subdomain_response_single01", "email_sending_subdomain_response_single02", "email_sending_subdomain_response_single03", "subdomain01", "subdomain02", "subdomain03", "zone01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_RESPONSE_SINGLE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_RESPONSE_SINGLE_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_RESPONSE_SINGLE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add zone_id alias for update test.
	if idmapResolved["zone_id"] == nil {
		idmapResolved["zone_id"] = idmapResolved["zone01"]
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
