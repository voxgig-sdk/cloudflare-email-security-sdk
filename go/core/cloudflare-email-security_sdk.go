package core

import (
	"fmt"
	"strings"

	vs "github.com/voxgig-sdk/cloudflare-email-security-sdk/go/utility/struct"
)

type CloudflareEmailSecuritySDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewCloudflareEmailSecuritySDK(options map[string]any) *CloudflareEmailSecuritySDK {
	sdk := &CloudflareEmailSecuritySDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := SharedConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath(sdk.options, []any{"feature", "test", "active"}) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath(sdk.options, []any{"__derived__", "featureorder"}).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *CloudflareEmailSecuritySDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *CloudflareEmailSecuritySDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *CloudflareEmailSecuritySDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *CloudflareEmailSecuritySDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

// Raw endpoint access is operator-controllable, like every entity op.
// Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
// either one reaches the same endpoint.
func (sdk *CloudflareEmailSecuritySDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	if !sdk.opAllowed("direct") {
		return sdk.opDenied("direct"), nil
	}

	return sdk.rawRequest(fetchargs)
}

// Is this raw-access op permitted by the SDK's allow.op option?
func (sdk *CloudflareEmailSecuritySDK) opAllowed(op string) bool {
	allowOp, _ := vs.GetPath(sdk.options, []any{"allow", "op"}).(string)
	return strings.Contains(allowOp, op)
}

func (sdk *CloudflareEmailSecuritySDK) opDenied(op string) map[string]any {
	allowOp, _ := vs.GetPath(sdk.options, []any{"allow", "op"}).(string)
	return map[string]any{
		"ok": false,
		"err": fmt.Errorf("CloudflareEmailSecuritySDK: %s: operation not allowed by"+
			" SDK option allow.op value: \"%s\"", op, allowOp),
	}
}

// Ungated request path shared by Direct and Graphql, each of which checks
// its own allow.op token first. Unexported, rather than a flag on fetchargs:
// a caller-supplied marker would let anyone opt straight back out of the
// gate by passing it.
func (sdk *CloudflareEmailSecuritySDK) rawRequest(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}

func (sdk *CloudflareEmailSecuritySDK) Graphql(
	query string, variables map[string]any, ctrl map[string]any,
) (map[string]any, error) {
	if !sdk.opAllowed("graphql") {
		return sdk.opDenied("graphql"), nil
	}

	if variables == nil {
		variables = map[string]any{}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	res, err := sdk.rawRequest(map[string]any{
		"method":  "POST",
		"headers": map[string]any{"content-type": "application/json"},
		"body":    map[string]any{"query": query, "variables": variables},
		"ctrl":    ctrl,
	})

	if err != nil {
		return res, err
	}

	// Errors are read BEFORE any status check: a GraphQL parse or validation
	// failure comes back as HTTP 400 carrying the standard { errors: [...] }
	// body, and the raw path represents a non-2xx as ok:false with no err —
	// so returning early on status would discard the server's own
	// diagnostics, which are the only useful part of that response.
	errors, _ := vs.GetPath(res, []any{"data", "errors"}).([]any)

	if 0 < len(errors) {
		msg, _ := vs.GetProp(errors[0], "message").(string)
		if msg == "" {
			msg = "graphql error"
		}
		res["ok"] = false
		res["err"] = fmt.Errorf("CloudflareEmailSecuritySDK: graphql: %s", msg)
		res["graphql"] = errors
	}

	return res, nil
}


// EmailAccountRulesPlanResponseSingle returns a EmailAccountRulesPlanResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailAccountRulesPlanResponseSingle(nil).List(nil, nil) or
// client.EmailAccountRulesPlanResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailAccountRulesPlanResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailAccountRulesPlanResponseSingleEntityFunc(sdk, data)
}


// EmailAccountRulesResponseCollection returns a EmailAccountRulesResponseCollection entity bound to this client.
// Idiomatic usage: client.EmailAccountRulesResponseCollection(nil).List(nil, nil) or
// client.EmailAccountRulesResponseCollection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailAccountRulesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailAccountRulesResponseCollectionEntityFunc(sdk, data)
}


// EmailAuth returns a EmailAuth entity bound to this client.
// Idiomatic usage: client.EmailAuth(nil).List(nil, nil) or
// client.EmailAuth(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailAuth(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailAuthEntityFunc(sdk, data)
}


// EmailCatchAllRuleResponseSingle returns a EmailCatchAllRuleResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailCatchAllRuleResponseSingle(nil).List(nil, nil) or
// client.EmailCatchAllRuleResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailCatchAllRuleResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailCatchAllRuleResponseSingleEntityFunc(sdk, data)
}


// EmailDestinationAddressResponseSingle returns a EmailDestinationAddressResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailDestinationAddressResponseSingle(nil).List(nil, nil) or
// client.EmailDestinationAddressResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailDestinationAddressResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailDestinationAddressResponseSingleEntityFunc(sdk, data)
}


// EmailDestinationAddressesResponseCollection returns a EmailDestinationAddressesResponseCollection entity bound to this client.
// Idiomatic usage: client.EmailDestinationAddressesResponseCollection(nil).List(nil, nil) or
// client.EmailDestinationAddressesResponseCollection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailDestinationAddressesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailDestinationAddressesResponseCollectionEntityFunc(sdk, data)
}


// EmailDnsSettingsResponseCollection returns a EmailDnsSettingsResponseCollection entity bound to this client.
// Idiomatic usage: client.EmailDnsSettingsResponseCollection(nil).List(nil, nil) or
// client.EmailDnsSettingsResponseCollection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailDnsSettingsResponseCollection(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailDnsSettingsResponseCollectionEntityFunc(sdk, data)
}


// EmailEmailSettingsResponseSingle returns a EmailEmailSettingsResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailEmailSettingsResponseSingle(nil).List(nil, nil) or
// client.EmailEmailSettingsResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailEmailSettingsResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailEmailSettingsResponseSingleEntityFunc(sdk, data)
}


// EmailRoutingSuppression returns a EmailRoutingSuppression entity bound to this client.
// Idiomatic usage: client.EmailRoutingSuppression(nil).List(nil, nil) or
// client.EmailRoutingSuppression(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailRoutingSuppression(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailRoutingSuppressionEntityFunc(sdk, data)
}


// EmailRuleResponseSingle returns a EmailRuleResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailRuleResponseSingle(nil).List(nil, nil) or
// client.EmailRuleResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailRuleResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailRuleResponseSingleEntityFunc(sdk, data)
}


// EmailRulesResponseCollection returns a EmailRulesResponseCollection entity bound to this client.
// Idiomatic usage: client.EmailRulesResponseCollection(nil).List(nil, nil) or
// client.EmailRulesResponseCollection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailRulesResponseCollection(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailRulesResponseCollectionEntityFunc(sdk, data)
}


// EmailSecurity returns a EmailSecurity entity bound to this client.
// Idiomatic usage: client.EmailSecurity(nil).List(nil, nil) or
// client.EmailSecurity(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSecurity(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSecurityEntityFunc(sdk, data)
}


// EmailSecurityAnalytics returns a EmailSecurityAnalytics entity bound to this client.
// Idiomatic usage: client.EmailSecurityAnalytics(nil).List(nil, nil) or
// client.EmailSecurityAnalytics(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSecurityAnalytics(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSecurityAnalyticsEntityFunc(sdk, data)
}


// EmailSecuritySetting returns a EmailSecuritySetting entity bound to this client.
// Idiomatic usage: client.EmailSecuritySetting(nil).List(nil, nil) or
// client.EmailSecuritySetting(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSecuritySetting(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSecuritySettingEntityFunc(sdk, data)
}


// EmailSending returns a EmailSending entity bound to this client.
// Idiomatic usage: client.EmailSending(nil).List(nil, nil) or
// client.EmailSending(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSending(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingEntityFunc(sdk, data)
}


// EmailSendingLimitsResponseSingle returns a EmailSendingLimitsResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailSendingLimitsResponseSingle(nil).List(nil, nil) or
// client.EmailSendingLimitsResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingLimitsResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingLimitsResponseSingleEntityFunc(sdk, data)
}


// EmailSendingReputationResponseSingle returns a EmailSendingReputationResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailSendingReputationResponseSingle(nil).List(nil, nil) or
// client.EmailSendingReputationResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingReputationResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingReputationResponseSingleEntityFunc(sdk, data)
}


// EmailSendingSubdomainDnsStatus returns a EmailSendingSubdomainDnsStatus entity bound to this client.
// Idiomatic usage: client.EmailSendingSubdomainDnsStatus(nil).List(nil, nil) or
// client.EmailSendingSubdomainDnsStatus(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSubdomainDnsStatus(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSubdomainDnsStatusEntityFunc(sdk, data)
}


// EmailSendingSubdomainPreview returns a EmailSendingSubdomainPreview entity bound to this client.
// Idiomatic usage: client.EmailSendingSubdomainPreview(nil).List(nil, nil) or
// client.EmailSendingSubdomainPreview(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSubdomainPreview(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSubdomainPreviewEntityFunc(sdk, data)
}


// EmailSendingSubdomainReputationComplaint returns a EmailSendingSubdomainReputationComplaint entity bound to this client.
// Idiomatic usage: client.EmailSendingSubdomainReputationComplaint(nil).List(nil, nil) or
// client.EmailSendingSubdomainReputationComplaint(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSubdomainReputationComplaint(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSubdomainReputationComplaintEntityFunc(sdk, data)
}


// EmailSendingSubdomainResponseSingle returns a EmailSendingSubdomainResponseSingle entity bound to this client.
// Idiomatic usage: client.EmailSendingSubdomainResponseSingle(nil).List(nil, nil) or
// client.EmailSendingSubdomainResponseSingle(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSubdomainResponseSingle(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSubdomainResponseSingleEntityFunc(sdk, data)
}


// EmailSendingSubdomainsResponseCollection returns a EmailSendingSubdomainsResponseCollection entity bound to this client.
// Idiomatic usage: client.EmailSendingSubdomainsResponseCollection(nil).List(nil, nil) or
// client.EmailSendingSubdomainsResponseCollection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSubdomainsResponseCollection(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSubdomainsResponseCollectionEntityFunc(sdk, data)
}


// EmailSendingSuppression returns a EmailSendingSuppression entity bound to this client.
// Idiomatic usage: client.EmailSendingSuppression(nil).List(nil, nil) or
// client.EmailSendingSuppression(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) EmailSendingSuppression(data map[string]any) CloudflareEmailSecurityEntity {
	return NewEmailSendingSuppressionEntityFunc(sdk, data)
}


// Subdomain returns a Subdomain entity bound to this client.
// Idiomatic usage: client.Subdomain(nil).List(nil, nil) or
// client.Subdomain(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *CloudflareEmailSecuritySDK) Subdomain(data map[string]any) CloudflareEmailSecurityEntity {
	return NewSubdomainEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *CloudflareEmailSecuritySDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewCloudflareEmailSecuritySDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
