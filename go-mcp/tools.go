package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/voxgig-sdk/cloudflare-email-security-sdk/go"
)

// Args is the common argument shape for both tools. `entity` selects
// the SDK entity to operate on; `query` is the optional reqmatch /
// reqdata map passed through to the SDK. For load, `query` should be
// `{"id": <value>}`. For list, omit `query` or pass an empty map.
type Args struct {
	Entity string         `json:"entity" jsonschema:"email_account_rules_plan_response_single | email_account_rules_response_collection | email_auth | email_catch_all_rule_response_single | email_destination_address_response_single | email_destination_addresses_response_collection | email_dns_settings_response_collection | email_email_settings_response_single | email_routing_suppression | email_rule_response_single | email_rules_response_collection | email_security | email_security_analytics | email_security_setting | email_sending | email_sending_limits_response_single | email_sending_reputation_response_single | email_sending_subdomain_dns_status | email_sending_subdomain_preview | email_sending_subdomain_reputation_complaint | email_sending_subdomain_response_single | email_sending_subdomains_response_collection | email_sending_suppression | subdomain"`
	Query  map[string]any `json:"query,omitempty" jsonschema:"optional match map e.g. {\"id\":1} for load, omit for list"`
}

func registerTools(server *mcp.Server, client *sdk.CloudflareEmailSecuritySDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "cloudflare-email-security_list",
		Description: "List records from CloudflareEmailSecurity. " +
			"Args: entity (one of the supported SDK entities), query (optional filter map). " +
			"Returns the first page of records as JSON.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args Args) (*mcp.CallToolResult, any, error) {
		return runOp(client, "list", args)
	})

	mcp.AddTool(server, &mcp.Tool{
		Name: "cloudflare-email-security_load",
		Description: "Load a single record from CloudflareEmailSecurity. " +
			"Args: entity, query ({\"id\":N} required). Returns the record as JSON.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args Args) (*mcp.CallToolResult, any, error) {
		return runOp(client, "load", args)
	})
}

func runOp(client *sdk.CloudflareEmailSecuritySDK, op string, args Args) (*mcp.CallToolResult, any, error) {
	ent, err := entityFor(client, args.Entity)
	if err != nil {
		return toolError(err.Error())
	}

	var result any
	switch op {
	case "list":
		result, err = ent.List(args.Query, nil)
	case "load":
		result, err = ent.Load(args.Query, nil)
	default:
		return toolError(fmt.Sprintf("unknown op %q", op))
	}
	if err != nil {
		return toolError(err.Error())
	}

	// SDK returns *Entity wrappers; unwrap each via .Data() to get a
	// plain map[string]any (or []any of maps for list) suitable for
	// JSON marshalling.
	data := extractData(result)
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Sprintf("marshal: %v", err))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, data, nil
}

// entityFor dispatches on the lowercase entity name. The generator
// emits one `case "<name>":` per entity defined in the SDK model.
func entityFor(client *sdk.CloudflareEmailSecuritySDK, name string) (sdk.CloudflareEmailSecurityEntity, error) {
	switch strings.ToLower(name) {
	case "email_account_rules_plan_response_single":
		return client.EmailAccountRulesPlanResponseSingle(nil), nil
	case "email_account_rules_response_collection":
		return client.EmailAccountRulesResponseCollection(nil), nil
	case "email_auth":
		return client.EmailAuth(nil), nil
	case "email_catch_all_rule_response_single":
		return client.EmailCatchAllRuleResponseSingle(nil), nil
	case "email_destination_address_response_single":
		return client.EmailDestinationAddressResponseSingle(nil), nil
	case "email_destination_addresses_response_collection":
		return client.EmailDestinationAddressesResponseCollection(nil), nil
	case "email_dns_settings_response_collection":
		return client.EmailDnsSettingsResponseCollection(nil), nil
	case "email_email_settings_response_single":
		return client.EmailEmailSettingsResponseSingle(nil), nil
	case "email_routing_suppression":
		return client.EmailRoutingSuppression(nil), nil
	case "email_rule_response_single":
		return client.EmailRuleResponseSingle(nil), nil
	case "email_rules_response_collection":
		return client.EmailRulesResponseCollection(nil), nil
	case "email_security":
		return client.EmailSecurity(nil), nil
	case "email_security_analytics":
		return client.EmailSecurityAnalytics(nil), nil
	case "email_security_setting":
		return client.EmailSecuritySetting(nil), nil
	case "email_sending":
		return client.EmailSending(nil), nil
	case "email_sending_limits_response_single":
		return client.EmailSendingLimitsResponseSingle(nil), nil
	case "email_sending_reputation_response_single":
		return client.EmailSendingReputationResponseSingle(nil), nil
	case "email_sending_subdomain_dns_status":
		return client.EmailSendingSubdomainDnsStatus(nil), nil
	case "email_sending_subdomain_preview":
		return client.EmailSendingSubdomainPreview(nil), nil
	case "email_sending_subdomain_reputation_complaint":
		return client.EmailSendingSubdomainReputationComplaint(nil), nil
	case "email_sending_subdomain_response_single":
		return client.EmailSendingSubdomainResponseSingle(nil), nil
	case "email_sending_subdomains_response_collection":
		return client.EmailSendingSubdomainsResponseCollection(nil), nil
	case "email_sending_suppression":
		return client.EmailSendingSuppression(nil), nil
	case "subdomain":
		return client.Subdomain(nil), nil

	}
	return nil, fmt.Errorf("unknown entity %q", name)
}

func extractData(x any) any {
	switch v := x.(type) {
	case sdk.Entity:
		return extractData(v.Data())
	case []any:
		out := make([]any, len(v))
		for i, e := range v {
			out[i] = extractData(e)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, vv := range v {
			out[k] = extractData(vv)
		}
		return out
	}
	return x
}

func toolError(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}
