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

func TestEmailSendingSubdomainDnsStatusEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EmailSendingSubdomainDnsStatus(nil)
		if ent == nil {
			t.Fatal("expected non-nil EmailSendingSubdomainDnsStatusEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := email_sending_subdomain_dns_statusBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "email_sending_subdomain_dns_status." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_DNS_STATUS_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		emailSendingSubdomainDnsStatusRef01Ent := client.EmailSendingSubdomainDnsStatus(nil)
		emailSendingSubdomainDnsStatusRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "email_sending_subdomain_dns_status"}), "email_sending_subdomain_dns_status_ref01"))
		emailSendingSubdomainDnsStatusRef01Data["subdomain_id"] = setup.idmap["subdomain01"]
		emailSendingSubdomainDnsStatusRef01Data["zone_id"] = setup.idmap["zone01"]

		emailSendingSubdomainDnsStatusRef01DataResult, err := emailSendingSubdomainDnsStatusRef01Ent.Create(emailSendingSubdomainDnsStatusRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		emailSendingSubdomainDnsStatusRef01Data = core.ToMapAny(entityData(emailSendingSubdomainDnsStatusRef01DataResult))
		if emailSendingSubdomainDnsStatusRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// LOAD
		emailSendingSubdomainDnsStatusRef01MatchDt0 := map[string]any{}
		emailSendingSubdomainDnsStatusRef01DataDt0Loaded, err := emailSendingSubdomainDnsStatusRef01Ent.Load(emailSendingSubdomainDnsStatusRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if emailSendingSubdomainDnsStatusRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func email_sending_subdomain_dns_statusBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "email_sending_subdomain_dns_status", "EmailSendingSubdomainDnsStatusTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read email_sending_subdomain_dns_status test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse email_sending_subdomain_dns_status test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"email_sending_subdomain_dns_status01", "email_sending_subdomain_dns_status02", "email_sending_subdomain_dns_status03", "subdomain01", "subdomain02", "subdomain03", "zone01"},
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
	entidEnvRaw := os.Getenv("CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_DNS_STATUS_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_DNS_STATUS_ENTID": idmap,
		"CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE":      "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN":   "FALSE",
		"CLOUDFLARE_EMAIL_SECURITY_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_DNS_STATUS_ENTID"])
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
