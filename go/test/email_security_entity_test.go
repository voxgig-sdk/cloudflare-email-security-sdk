package sdktest

import (
	"encoding/json"
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

func TestEmailSecurityEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailSecurity(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailSecurityEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"email_security": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.EmailSecurity(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.EmailSecurity(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_securityBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_security." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailSecurityRef01Ent := client.EmailSecurity(nil)
		emailSecurityRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_security"}), "email_security_ref01"))
		emailSecurityRef01Data["account_id"] = setup.idmap["account01"]
		emailSecurityRef01Data["bulk_id"] = setup.idmap["bulk01"]
		emailSecurityRef01Data["investigate_id"] = setup.idmap["investigate01"]

		emailSecurityRef01DataResult, err := emailSecurityRef01Ent.Create(emailSecurityRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailSecurityRef01Data = core.ToMapAny(entityData(emailSecurityRef01DataResult))
		if emailSecurityRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if emailSecurityRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		emailSecurityRef01Match := map[string]any{
			"account_id": setup.idmap["account01"],
			"investigate_id": setup.idmap["investigate01"],
		}

		emailSecurityRef01ListResult, err := emailSecurityRef01Ent.List(emailSecurityRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		emailSecurityRef01List, emailSecurityRef01ListOk := emailSecurityRef01ListResult.([]any)
		if !emailSecurityRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", emailSecurityRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(emailSecurityRef01List), map[string]any{"id": emailSecurityRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		emailSecurityRef01MatchDt0 := map[string]any{
			"id": emailSecurityRef01Data["id"],
		}
		emailSecurityRef01DataDt0Loaded, err := emailSecurityRef01Ent.Load(emailSecurityRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		emailSecurityRef01DataDt0LoadResult := core.ToMapAny(entityData(emailSecurityRef01DataDt0Loaded))
		if emailSecurityRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if emailSecurityRef01DataDt0LoadResult["id"] != emailSecurityRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		emailSecurityRef01MatchRm0 := map[string]any{
			"id": emailSecurityRef01Data["id"],
		}
		_, err = emailSecurityRef01Ent.Remove(emailSecurityRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		emailSecurityRef01MatchRt0 := map[string]any{
			"account_id": setup.idmap["account01"],
			"investigate_id": setup.idmap["investigate01"],
		}

		emailSecurityRef01ListRt0Result, err := emailSecurityRef01Ent.List(emailSecurityRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		emailSecurityRef01ListRt0, emailSecurityRef01ListRt0Ok := emailSecurityRef01ListRt0Result.([]any)
		if !emailSecurityRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", emailSecurityRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(emailSecurityRef01ListRt0), map[string]any{"id": emailSecurityRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func email_securityBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_security", "EmailSecurityTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_security test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_security test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_security01", "email_security02", "email_security03", "account01", "bulk01", "investigate01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SECURITY_ENTID"])
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
