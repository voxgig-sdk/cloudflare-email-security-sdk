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

func TestEmailRoutingSuppressionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailRoutingSuppression(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailRoutingSuppressionEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"email_routing_suppression": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.EmailRoutingSuppression(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.EmailRoutingSuppression(nil).Stream("list", nil, nil) {
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
		setup := email_routing_suppressionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_routing_suppression." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailRoutingSuppressionRef01Ent := client.EmailRoutingSuppression(nil)
		emailRoutingSuppressionRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_routing_suppression"}), "email_routing_suppression_ref01"))
		emailRoutingSuppressionRef01Data["account_id"] = setup.idmap["account01"]
		emailRoutingSuppressionRef01Data["zone_id"] = setup.idmap["zone01"]

		emailRoutingSuppressionRef01DataResult, err := emailRoutingSuppressionRef01Ent.Create(emailRoutingSuppressionRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailRoutingSuppressionRef01Data = core.ToMapAny(entityData(emailRoutingSuppressionRef01DataResult))
		if emailRoutingSuppressionRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if emailRoutingSuppressionRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		emailRoutingSuppressionRef01Match := map[string]any{
			"zone_id": setup.idmap["zone01"],
		}

		emailRoutingSuppressionRef01ListResult, err := emailRoutingSuppressionRef01Ent.List(emailRoutingSuppressionRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		emailRoutingSuppressionRef01List, emailRoutingSuppressionRef01ListOk := emailRoutingSuppressionRef01ListResult.([]any)
		if !emailRoutingSuppressionRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", emailRoutingSuppressionRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(emailRoutingSuppressionRef01List), map[string]any{"id": emailRoutingSuppressionRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		emailRoutingSuppressionRef01MatchDt0 := map[string]any{
			"id": emailRoutingSuppressionRef01Data["id"],
		}
		emailRoutingSuppressionRef01DataDt0Loaded, err := emailRoutingSuppressionRef01Ent.Load(emailRoutingSuppressionRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		emailRoutingSuppressionRef01DataDt0LoadResult := core.ToMapAny(entityData(emailRoutingSuppressionRef01DataDt0Loaded))
		if emailRoutingSuppressionRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if emailRoutingSuppressionRef01DataDt0LoadResult["id"] != emailRoutingSuppressionRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		emailRoutingSuppressionRef01MatchRm0 := map[string]any{
			"id": emailRoutingSuppressionRef01Data["id"],
		}
		_, err = emailRoutingSuppressionRef01Ent.Remove(emailRoutingSuppressionRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		emailRoutingSuppressionRef01MatchRt0 := map[string]any{
			"zone_id": setup.idmap["zone01"],
		}

		emailRoutingSuppressionRef01ListRt0Result, err := emailRoutingSuppressionRef01Ent.List(emailRoutingSuppressionRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		emailRoutingSuppressionRef01ListRt0, emailRoutingSuppressionRef01ListRt0Ok := emailRoutingSuppressionRef01ListRt0Result.([]any)
		if !emailRoutingSuppressionRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", emailRoutingSuppressionRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(emailRoutingSuppressionRef01ListRt0), map[string]any{"id": emailRoutingSuppressionRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func email_routing_suppressionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_routing_suppression", "EmailRoutingSuppressionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_routing_suppression test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_routing_suppression test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_routing_suppression01", "email_routing_suppression02", "email_routing_suppression03", "account01", "zone01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ROUTING_SUPPRESSION_ENTID"])
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
