package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "CloudflareEmailSecurity",
			"slug": "cloudflare-email-security",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://api.cloudflare.com/client/v4",
			"auth": map[string]any{
				"prefix": "",
				"name": "X-Auth-Email",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"email_account_rules_plan_response_single": map[string]any{},
				"email_account_rules_response_collection": map[string]any{},
				"email_auth": map[string]any{},
				"email_catch_all_rule_response_single": map[string]any{},
				"email_destination_address_response_single": map[string]any{},
				"email_destination_addresses_response_collection": map[string]any{},
				"email_dns_settings_response_collection": map[string]any{},
				"email_email_settings_response_single": map[string]any{},
				"email_routing_suppression": map[string]any{},
				"email_rule_response_single": map[string]any{},
				"email_rules_response_collection": map[string]any{},
				"email_security": map[string]any{},
				"email_security_analytics": map[string]any{},
				"email_security_setting": map[string]any{},
				"email_sending": map[string]any{},
				"email_sending_limits_response_single": map[string]any{},
				"email_sending_reputation_response_single": map[string]any{},
				"email_sending_subdomain_dns_status": map[string]any{},
				"email_sending_subdomain_preview": map[string]any{},
				"email_sending_subdomain_reputation_complaint": map[string]any{},
				"email_sending_subdomain_response_single": map[string]any{},
				"email_sending_subdomains_response_collection": map[string]any{},
				"email_sending_suppression": map[string]any{},
				"subdomain": map[string]any{},
			},
		},
		"entity": map[string]any{
			"email_account_rules_plan_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "catch_all_rules",
						"title": "Catch All Rules",
						"type": "`$ARRAY`",
						"short": "Desired catch-all Email Routing rules managed by the deploying Worker.",
					},
					map[string]any{
						"name": "owner_worker_tag",
						"title": "Owner Worker Tag",
						"type": "`$STRING`",
						"req": true,
						"short": "Public tag (script_tag) of the Worker that owns this rule.",
						"writeOnly": true,
					},
					map[string]any{
						"name": "rules",
						"title": "Rules",
						"type": "`$ARRAY`",
						"short": "Desired normal Email Routing rules managed by the deploying Worker.",
					},
				},
				"name": "email_account_rules_plan_response_single",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/routing/rules/plan",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"lit": "plan",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"rules",
									"plan",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_account_rules_response_collection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result_info",
						"title": "Result Info",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_account_rules_response_collection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/routing/rules",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"rules",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "enabled",
											"orig": "enabled",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 20,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"enabled",
										"page",
										"per_page",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_auth": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "enabled",
						"title": "Enabled",
						"type": "`$BOOLEAN`",
						"short": "Enable or disable DMARC reports for this zone",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
						"op": map[string]any{
							"load": map[string]any{
								"type": "`$OBJECT`",
							},
						},
						"short": "Recursive SPF inspection tree",
					},
					map[string]any{
						"name": "skip_wizard",
						"title": "Skip Wizard",
						"type": "`$BOOLEAN`",
						"short": "Skip the DMARC setup wizard",
					},
				},
				"name": "email_auth",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/auth/spf/inspect",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "auth",
									},
									map[string]any{
										"lit": "spf",
									},
									map[string]any{
										"lit": "inspect",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"auth",
									"spf",
									"inspect",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
											"example": "e5bb46707a802688812d5d1c9f7977d4",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/auth/dmarc-reports",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "auth",
									},
									map[string]any{
										"lit": "dmarc-reports",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"auth",
									"dmarc-reports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/zones/{zone_id}/email/auth/dmarc-reports",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "auth",
									},
									map[string]any{
										"lit": "dmarc-reports",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"auth",
									"dmarc-reports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_catch_all_rule_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "actions",
						"title": "Actions",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List actions for the catch-all routing rule.",
					},
					map[string]any{
						"name": "enabled",
						"title": "Enabled",
						"type": "`$BOOLEAN`",
						"short": "Routing rule status.",
					},
					map[string]any{
						"name": "matchers",
						"title": "Matchers",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of matchers for the catch-all routing rule.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Routing rule name.",
					},
					map[string]any{
						"name": "owner_worker_tag",
						"title": "Owner Worker Tag",
						"type": "`$STRING`",
						"short": "Public tag (script_tag) of the Worker that owns this rule.",
						"writeOnly": true,
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"short": "Who manages the rule.",
					},
				},
				"name": "email_catch_all_rule_response_single",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/rules/catch_all",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"lit": "catch_all",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
									"catch_all",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/zones/{zone_id}/email/routing/rules/catch_all",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"lit": "catch_all",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
									"catch_all",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_destination_address_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"req": true,
						"short": "The contact email address of the user.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"req": true,
						"short": "Destination address status.",
					},
				},
				"name": "email_destination_address_response_single",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/routing/addresses",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "addresses",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"addresses",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "addresses",
									},
									map[string]any{
										"var": "destination_address_identifier",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"addresses",
									"{destination_address_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "destination_address_identifier",
											"orig": "destination_address_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "ea95132c15732412d22c1476fa83f27a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"destination_address_identifier",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "addresses",
									},
									map[string]any{
										"var": "destination_address_identifier",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"addresses",
									"{destination_address_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "destination_address_identifier",
											"orig": "destination_address_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "ea95132c15732412d22c1476fa83f27a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"destination_address_identifier",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "addresses",
									},
									map[string]any{
										"var": "destination_address_identifier",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"addresses",
									"{destination_address_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "destination_address_identifier",
											"orig": "destination_address_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "ea95132c15732412d22c1476fa83f27a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"destination_address_identifier",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_destination_addresses_response_collection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result_info",
						"title": "Result Info",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_destination_addresses_response_collection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/routing/addresses",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "addresses",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"addresses",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "asc",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "verified",
											"orig": "verified",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"page",
										"per_page",
										"verified",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_dns_settings_response_collection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ARRAY`",
					},
				},
				"name": "email_dns_settings_response_collection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "subdomain",
											"orig": "subdomain",
											"type": "`$STRING`",
											"kind": "query",
											"example": "example.net",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain",
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.subdomain",
						},
					},
				},
			},
			"email_email_settings_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "enabled",
						"title": "Enabled",
						"type": "`$BOOLEAN`",
						"short": "State of your zone Email Routing settings.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Domain of your zone.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "skip_wizard",
						"title": "Skip Wizard",
						"type": "`$BOOLEAN`",
						"short": "Flag to check if the user skipped the configuration wizard.",
					},
					map[string]any{
						"name": "support_subaddress",
						"title": "Support Subaddress",
						"type": "`$BOOLEAN`",
						"short": "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.",
					},
				},
				"name": "email_email_settings_response_single",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/disable",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "disable",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"disable",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/enable",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "enable",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"enable",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/unlock",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "unlock",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"unlock",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/zones/{zone_id}/email/routing",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/zones/{zone_id}/email/routing/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/zones/{zone_id}/email/routing/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/zones/{zone_id}/email/routing",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_routing_suppression": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"req": true,
						"format": "email",
					},
					map[string]any{
						"name": "expires_at",
						"title": "Expires At",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$STRING`",
							},
						},
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"format": "uuid",
					},
					map[string]any{
						"name": "reason",
						"title": "Reason",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "zones",
						"title": "Zones",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "email_routing_suppression",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/routing/suppression",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/suppression",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/routing/suppression",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "desc",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "created_at",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"order",
										"page",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/suppression",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "desc",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "created_at",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"direction",
										"order",
										"page",
										"per_page",
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/routing/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email/routing/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"routing",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/zones/{zone_id}/email/routing/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_rule_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "actions",
						"title": "Actions",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List actions patterns.",
					},
					map[string]any{
						"name": "enabled",
						"title": "Enabled",
						"type": "`$BOOLEAN`",
						"short": "Routing rule status.",
					},
					map[string]any{
						"name": "matchers",
						"title": "Matchers",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Matching patterns to forward to your actions.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Routing rule name.",
					},
					map[string]any{
						"name": "owner_worker_tag",
						"title": "Owner Worker Tag",
						"type": "`$STRING`",
						"short": "Public tag (script_tag) of the Worker that owns this rule.",
						"writeOnly": true,
					},
					map[string]any{
						"name": "priority",
						"title": "Priority",
						"type": "`$NUMBER`",
						"short": "Priority of the routing rule.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"short": "Who manages the rule.",
					},
				},
				"name": "email_rule_response_single",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/routing/rules",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"var": "rule_identifier",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
									"{rule_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "rule_identifier",
											"orig": "rule_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "a7e6fb77503c41d8a7f3113c6918f10c",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"rule_identifier",
										"zone_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"var": "rule_identifier",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
									"{rule_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "rule_identifier",
											"orig": "rule_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "a7e6fb77503c41d8a7f3113c6918f10c",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"rule_identifier",
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"var": "rule_identifier",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
									"{rule_identifier}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "rule_identifier",
											"orig": "rule_identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "a7e6fb77503c41d8a7f3113c6918f10c",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"rule_identifier",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_rules_response_collection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result_info",
						"title": "Result Info",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_rules_response_collection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/routing/rules",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "routing",
									},
									map[string]any{
										"lit": "rules",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"routing",
									"rules",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "enabled",
											"orig": "enabled",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 20,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"enabled",
										"page",
										"per_page",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_security": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "action",
						"title": "Action",
						"type": "`$STRING`",
						"req": true,
						"short": "The action the job performs on every message matching the search parameters.",
					},
					map[string]any{
						"name": "comment",
						"title": "Comment",
						"type": "`$STRING`",
						"short": "Optional note describing the job.",
					},
					map[string]any{
						"name": "destination",
						"title": "Destination",
						"type": "`$ANY`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Required when action is 'MOVE'.",
					},
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "expected_disposition",
						"title": "Expected Disposition",
						"type": "`$STRING`",
						"short": "Nonfunctional field.",
						"deprecated": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier for a message retrieved from investigation.",
					},
					map[string]any{
						"name": "ids",
						"title": "Ids",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of message IDs to move.",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "postfix_ids",
						"title": "Postfix Ids",
						"type": "`$ARRAY`",
						"short": "Deprecated, use `ids` instead.",
						"deprecated": true,
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "result_info",
						"title": "Result Info",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "search_params",
						"title": "Search Params",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the API call was successful.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "email_security",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
									map[string]any{
										"var": "bulk_id",
									},
									map[string]any{
										"lit": "cancel",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
									"{bulk_id}",
									"cancel",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"job_id": "bulk_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "bulk_id",
											"orig": "job_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"$action": "cancel",
									"exist": []any{
										"account_id",
										"bulk_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/move",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "move",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"move",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"$action": "move",
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "reclassify",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"reclassify",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"$action": "reclassify",
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/move",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "move",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"move",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/preview",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "preview",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"preview",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/investigate/release",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "release",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"release",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/submissions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "submissions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"submissions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"$action": "submission",
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "alert_id",
											"orig": "alert_id",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "delivery_status",
											"orig": "delivery_status",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "detections_only",
											"orig": "detections_only",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "domain",
											"orig": "domain",
											"type": "`$STRING`",
											"kind": "query",
											"example": "example.com",
										},
										map[string]any{
											"name": "end",
											"orig": "end",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-07-25T14:30:00Z",
										},
										map[string]any{
											"name": "final_disposition",
											"orig": "final_disposition",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "message_action",
											"orig": "message_action",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "message_id",
											"orig": "message_id",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "metric",
											"orig": "metric",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
											"example": "bob jones",
										},
										map[string]any{
											"name": "recipient",
											"orig": "recipient",
											"type": "`$STRING`",
											"kind": "query",
											"example": "me@example.com",
										},
										map[string]any{
											"name": "sender",
											"orig": "sender",
											"type": "`$STRING`",
											"kind": "query",
											"example": "noreply@example.com",
										},
										map[string]any{
											"name": "smtp_helo_ip",
											"orig": "smtp_helo_ip",
											"type": "`$STRING`",
											"kind": "query",
											"example": "192.0.2.1",
										},
										map[string]any{
											"name": "start",
											"orig": "start",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-06-25T14:30:00Z",
										},
										map[string]any{
											"name": "subject",
											"orig": "subject",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "investigate",
									"exist": []any{
										"account_id",
										"alert_id",
										"cursor",
										"delivery_status",
										"detections_only",
										"domain",
										"end",
										"final_disposition",
										"message_action",
										"message_id",
										"metric",
										"page",
										"per_page",
										"query",
										"recipient",
										"sender",
										"smtp_helo_ip",
										"start",
										"subject",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/submissions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "submissions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"submissions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "end",
											"orig": "end",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-07-25T14:30:00Z",
										},
										map[string]any{
											"name": "escalated_from_user",
											"orig": "escalated_from_user",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "requested_at",
										},
										map[string]any{
											"name": "original_disposition",
											"orig": "original_disposition",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "outcome_disposition",
											"orig": "outcome_disposition",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "requested_disposition",
											"orig": "requested_disposition",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "start",
											"orig": "start",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-06-25T14:30:00Z",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "submission_id",
											"orig": "submission_id",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "submission",
									"exist": []any{
										"account_id",
										"direction",
										"end",
										"escalated_from_user",
										"order",
										"original_disposition",
										"outcome_disposition",
										"page",
										"per_page",
										"query",
										"requested_disposition",
										"start",
										"status",
										"submission_id",
										"type",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/phishguard/reports",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "phishguard",
									},
									map[string]any{
										"lit": "reports",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"phishguard",
									"reports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "end",
											"orig": "end",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2020-09-01T00:00:00Z",
										},
										map[string]any{
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2020-08-01T00:00:00Z",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "start",
											"orig": "start",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2020-08-01T00:00:00Z",
										},
										map[string]any{
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2020-09-01T00:00:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"end",
										"from_date",
										"page",
										"per_page",
										"start",
										"to_date",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "action_type",
											"orig": "action_type",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"action_type",
										"page",
										"per_page",
										"status",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk/{job_id}/messages",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
									map[string]any{
										"var": "bulk_id",
									},
									map[string]any{
										"lit": "messages",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
									"{bulk_id}",
									"messages",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"job_id": "bulk_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "bulk_id",
											"orig": "job_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"bulk_id",
										"page",
										"per_page",
										"status",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/action_log",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "action_log",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"action_log",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
										"page",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/detections",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "detections",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"detections",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/preview",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "preview",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"preview",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/raw",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "raw",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"raw",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}/trace",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
									map[string]any{
										"lit": "trace",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
									"trace",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/{investigate_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"var": "investigate_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"{investigate_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "investigate_id",
											"orig": "investigate_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
										},
									},
									"query": []any{
										map[string]any{
											"name": "submission",
											"orig": "submission",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"investigate_id",
										"submission",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk/{job_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
									map[string]any{
										"var": "job_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
									"{job_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "job_id",
											"orig": "job_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"job_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/investigate/bulk/{job_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "investigate",
									},
									map[string]any{
										"lit": "bulk",
									},
									map[string]any{
										"var": "job_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"investigate",
									"bulk",
									"{job_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "job_id",
											"orig": "job_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"job_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_security_analytics": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_security_analytics",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/analytics/monthly_report",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "analytics",
									},
									map[string]any{
										"lit": "monthly_report",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"analytics",
									"monthly_report",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_security_setting": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "allowed_delivery_modes",
						"title": "Allowed Delivery Modes",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"short": "Delivery modes to onboard the domain through.",
					},
					map[string]any{
						"name": "comments",
						"title": "Comments",
						"type": "`$STRING`",
						"short": "Optional note describing the reason for the ignore pattern.",
					},
					map[string]any{
						"name": "deletes",
						"title": "Deletes",
						"type": "`$ARRAY`",
						"req": true,
						"short": "IDs of the allow policies to delete.",
					},
					map[string]any{
						"name": "domain",
						"title": "Domain",
						"type": "`$STRING`",
						"req": true,
						"short": "The email domain to protect.",
					},
					map[string]any{
						"name": "drop_dispositions",
						"title": "Drop Dispositions",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"short": "Dispositions to drop instead of delivering, e.g.",
					},
					map[string]any{
						"name": "folder",
						"title": "Folder",
						"type": "`$ANY`",
						"short": "The mailbox folder to scan, for API-scanning domains.",
					},
					map[string]any{
						"name": "integration_id",
						"title": "Integration Id",
						"type": "`$STRING`",
						"short": "Identifier of the CASB integration that authorizes this domain.",
						"format": "uuid",
					},
					map[string]any{
						"name": "ip_restrictions",
						"title": "Ip Restrictions",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"short": "Source IP ranges mail is accepted from.",
					},
					map[string]any{
						"name": "lookback_hops",
						"title": "Lookback Hops",
						"type": "`$INTEGER`",
						"short": "Number of hops to trace back through received headers when reconstructing the original message (1-20).",
					},
					map[string]any{
						"name": "patches",
						"title": "Patches",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Partial updates to apply — each entry carries the policy's ID and only the fields to change.",
					},
					map[string]any{
						"name": "pattern",
						"title": "Pattern",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Regular expression identifying URLs to exempt from rewriting.",
					},
					map[string]any{
						"name": "posts",
						"title": "Posts",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Allow policies to create.",
					},
					map[string]any{
						"name": "puts",
						"title": "Puts",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Full replacements to apply — each entry carries the policy's ID and every field of its new value.",
					},
					map[string]any{
						"name": "regions",
						"title": "Regions",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"short": "Regions that process messages for this domain, e.g.",
					},
					map[string]any{
						"name": "require_tls_inbound",
						"title": "Require Tls Inbound",
						"type": "`$BOOLEAN`",
						"short": "Require TLS on inbound connections.",
					},
					map[string]any{
						"name": "require_tls_outbound",
						"title": "Require Tls Outbound",
						"type": "`$BOOLEAN`",
						"short": "Require TLS on outbound connections.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"op": map[string]any{
							"load": map[string]any{
								"req": true,
								"type": "`$OBJECT`",
							},
						},
						"short": "An impersonation registry entry.",
					},
					map[string]any{
						"name": "transport",
						"title": "Transport",
						"type": "`$STRING`",
						"short": "The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g.",
					},
				},
				"name": "email_security_setting",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"allowed_delivery_modes": "`reqdata.allowed_delivery_mode`",
										"domain": "`reqdata.domain`",
										"drop_dispositions": "`reqdata.drop_disposition`",
										"folder": "`reqdata.folder`",
										"integration_id": "`reqdata.integration_id`",
										"ip_restrictions": "`reqdata.ip_restriction`",
										"lookback_hops": "`reqdata.lookback_hop`",
										"regions": "`reqdata.region`",
										"require_tls_inbound": "`reqdata.require_tls_inbound`",
										"require_tls_outbound": "`reqdata.require_tls_outbound`",
										"transport": "`reqdata.transport`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/domains/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/impersonation_registry",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "impersonation_registry",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"impersonation_registry",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"comments": "`reqdata.comment`",
										"pattern": "`reqdata.pattern`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns/batch",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_acceptable_sender",
											"orig": "is_acceptable_sender",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_exempt_recipient",
											"orig": "is_exempt_recipient",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_trusted_sender",
											"orig": "is_trusted_sender",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "pattern",
											"orig": "pattern",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "pattern_type",
											"orig": "pattern_type",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "verify_sender",
											"orig": "verify_sender",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"is_acceptable_sender",
										"is_exempt_recipient",
										"is_trusted_sender",
										"order",
										"page",
										"pattern",
										"pattern_type",
										"per_page",
										"search",
										"verify_sender",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "active_delivery_mode",
											"orig": "active_delivery_mode",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "allowed_delivery_mode",
											"orig": "allowed_delivery_mode",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "domain",
											"orig": "domain",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "integration_id",
											"orig": "integration_id",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"active_delivery_mode",
										"allowed_delivery_mode",
										"direction",
										"domain",
										"integration_id",
										"order",
										"page",
										"per_page",
										"search",
										"status",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_recent",
											"orig": "is_recent",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_similarity",
											"orig": "is_similarity",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "pattern",
											"orig": "pattern",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"is_recent",
										"is_similarity",
										"order",
										"page",
										"pattern",
										"per_page",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "enabled",
											"orig": "enabled",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "name",
											"orig": "name",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"enabled",
										"name",
										"order",
										"page",
										"per_page",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "pattern",
											"orig": "pattern",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "pattern_type",
											"orig": "pattern_type",
											"type": "`$ANY`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"order",
										"page",
										"pattern",
										"pattern_type",
										"per_page",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/impersonation_registry",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "impersonation_registry",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"impersonation_registry",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "provenance",
											"orig": "provenance",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"order",
										"page",
										"per_page",
										"provenance",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"order",
										"page",
										"per_page",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"page",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"{domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/domains/{domain_id}/verification",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
									map[string]any{
										"lit": "verification",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"{domain_id}",
									"verification",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "impersonation_registry",
									},
									map[string]any{
										"var": "impersonation_registry_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"impersonation_registry",
									"{impersonation_registry_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "impersonation_registry_id",
											"orig": "impersonation_registry_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"impersonation_registry_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
									map[string]any{
										"var": "sending_domain_restriction_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
									"{sending_domain_restriction_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "sending_domain_restriction_id",
											"orig": "sending_domain_restriction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"sending_domain_restriction_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
									map[string]any{
										"var": "trusted_domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
									"{trusted_domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "trusted_domain_id",
											"orig": "trusted_domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"trusted_domain_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"{domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "impersonation_registry",
									},
									map[string]any{
										"var": "impersonation_registry_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"impersonation_registry",
									"{impersonation_registry_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "impersonation_registry_id",
											"orig": "impersonation_registry_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"impersonation_registry_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"comments": "`reqdata.comment`",
										"pattern": "`reqdata.pattern`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
									map[string]any{
										"var": "sending_domain_restriction_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
									"{sending_domain_restriction_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "sending_domain_restriction_id",
											"orig": "sending_domain_restriction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"sending_domain_restriction_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
									map[string]any{
										"var": "trusted_domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
									"{trusted_domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "trusted_domain_id",
											"orig": "trusted_domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"trusted_domain_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"{domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "impersonation_registry",
									},
									map[string]any{
										"var": "impersonation_registry_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"impersonation_registry",
									"{impersonation_registry_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "impersonation_registry_id",
											"orig": "impersonation_registry_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"impersonation_registry_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "url_ignore_patterns",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"url_ignore_patterns",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "content_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"content_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "sending_domain_restrictions",
									},
									map[string]any{
										"var": "sending_domain_restriction_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"sending_domain_restrictions",
									"{sending_domain_restriction_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "sending_domain_restriction_id",
											"orig": "sending_domain_restriction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"sending_domain_restriction_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
									map[string]any{
										"var": "trusted_domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
									"{trusted_domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "trusted_domain_id",
											"orig": "trusted_domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"trusted_domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email-security/settings/domains",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"domains",
									"{domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"domain_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "block_senders",
									},
									map[string]any{
										"var": "pattern_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"block_senders",
									"{pattern_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "pattern_id",
											"orig": "pattern_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"pattern_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "allow_policies",
									},
									map[string]any{
										"var": "policy_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"allow_policies",
									"{policy_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "policy_id",
											"orig": "policy_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"policy_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email-security",
									},
									map[string]any{
										"lit": "settings",
									},
									map[string]any{
										"lit": "trusted_domains",
									},
									map[string]any{
										"var": "trusted_domain_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email-security",
									"settings",
									"trusted_domains",
									"{trusted_domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
										map[string]any{
											"name": "trusted_domain_id",
											"orig": "trusted_domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"trusted_domain_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"short": "File attachments and inline images.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ANY`",
						"short": "Blind carbon copy recipient(s).",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ANY`",
						"short": "Carbon copy recipient(s).",
					},
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$ANY`",
						"req": true,
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "headers",
						"title": "Headers",
						"type": "`$OBJECT`",
						"short": "Custom email headers as key-value pairs.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "HTML body of the email.",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "mime_message",
						"title": "Mime Message",
						"type": "`$STRING`",
						"req": true,
						"short": "The full MIME-encoded email message.",
					},
					map[string]any{
						"name": "recipients",
						"title": "Recipients",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of recipient email addresses.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ANY`",
						"short": "Reply-to address.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "result_info",
						"title": "Result Info",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"req": true,
						"short": "Email subject line.",
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "Plain text body of the email.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ANY`",
						"short": "Recipient(s).",
					},
				},
				"name": "email_sending",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/sending/send",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "send",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"send",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"attachments": "`reqdata.attachment`",
										"bcc": "`reqdata.bcc`",
										"cc": "`reqdata.cc`",
										"from": "`reqdata.from`",
										"headers": "`reqdata.header`",
										"html": "`reqdata.html`",
										"reply_to": "`reqdata.reply_to`",
										"subject": "`reqdata.subject`",
										"text": "`reqdata.text`",
										"to": "`reqdata.to`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/sending/send_raw",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "send_raw",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"send_raw",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"from": "`reqdata.from`",
										"mime_message": "`reqdata.mime_message`",
										"recipients": "`reqdata.recipient`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/messages/{message_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "messages",
									},
									map[string]any{
										"var": "message_id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"messages",
									"{message_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345",
										},
										map[string]any{
											"name": "message_id",
											"orig": "message_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "<aB3xK9mP@example.com>",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"message_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending_limits_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_sending_limits_response_single",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/limits",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "limits",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"limits",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending_reputation_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"name": "email_sending_reputation_response_single",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/reputation",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "reputation",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"reputation",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending_subdomain_dns_status": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
				},
				"name": "email_sending_subdomain_dns_status",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
									map[string]any{
										"lit": "dns",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
									"dns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns/status",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
									map[string]any{
										"lit": "dns",
									},
									map[string]any{
										"lit": "status",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
									"dns",
									"status",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.subdomain",
						},
					},
				},
			},
			"email_sending_subdomain_preview": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The domain name within the zone.",
					},
				},
				"name": "email_sending_subdomain_preview",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/sending/subdomains/preview",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"lit": "preview",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"preview",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending_subdomain_reputation_complaint": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"name": "email_sending_subdomain_reputation_complaint",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
									map[string]any{
										"lit": "reputation",
									},
									map[string]any{
										"lit": "complaints",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
									"reputation",
									"complaints",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
									"query": []any{
										map[string]any{
											"name": "end_at",
											"orig": "end_at",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
										map[string]any{
											"name": "start_at",
											"orig": "start_at",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"end_at",
										"start_at",
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.subdomain",
						},
					},
				},
			},
			"email_sending_subdomain_response_single": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "drop_suppressed_recipients",
						"title": "Drop Suppressed Recipients",
						"type": "`$BOOLEAN`",
						"short": "Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The domain name within the zone.",
					},
					map[string]any{
						"name": "preview_enabled",
						"title": "Preview Enabled",
						"type": "`$BOOLEAN`",
						"short": "Whether sent messages from this subdomain can be previewed in the activity log.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ANY`",
					},
				},
				"name": "email_sending_subdomain_response_single",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/sending/subdomains",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "subdomain_id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{subdomain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "subdomain_id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"subdomain_id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.subdomain",
						},
					},
				},
			},
			"email_sending_subdomains_response_collection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$ARRAY`",
					},
				},
				"name": "email_sending_subdomains_response_collection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/subdomains",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"email_sending_suppression": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "When the suppression was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"req": true,
						"short": "The email address to suppress.",
						"format": "email",
					},
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "expires_at",
						"title": "Expires At",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Expiration timestamp for the suppression.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier for this suppression.",
						"format": "uuid",
					},
					map[string]any{
						"name": "items",
						"title": "Items",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Suppressions to import.",
					},
					map[string]any{
						"name": "messages",
						"title": "Messages",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "note",
						"title": "Note",
						"type": "`$STRING`",
						"short": "Advisory note for this suppression.",
					},
					map[string]any{
						"name": "read_only",
						"title": "Read Only",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether clients may mutate this suppression.",
					},
					map[string]any{
						"name": "reason",
						"title": "Reason",
						"type": "`$STRING`",
						"req": true,
						"short": "Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`.",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "scope",
						"title": "Scope",
						"type": "`$ANY`",
						"short": "Where the suppression applies.",
					},
					map[string]any{
						"name": "success",
						"title": "Success",
						"type": "`$BOOLEAN`",
						"req": true,
					},
					map[string]any{
						"name": "zones",
						"title": "Zones",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "email_sending_suppression",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/sending/suppression",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/sending/suppressions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/accounts/{account_id}/email/sending/suppressions/bulk",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"lit": "bulk",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
									"bulk",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/zones/{zone_id}/email/sending/suppression",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"zone_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/suppressions",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
									},
									"query": []any{
										map[string]any{
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
											"kind": "query",
											"example": "eyJ0Ijo0LCJwaCI6ImRvbWFpbiIsInAiOjEyMywiYyI6ImMzZGM1ZjBiMzRhMTRmZjhlMWIzZWMwNDg5NWUxYjIyIn0",
										},
										map[string]any{
											"name": "email",
											"orig": "email",
											"type": "`$STRING`",
											"kind": "query",
											"example": "user@example.com",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
										map[string]any{
											"name": "reason",
											"orig": "reason",
											"type": "`$STRING`",
											"kind": "query",
											"example": "hard_bounce",
										},
										map[string]any{
											"name": "scope_type",
											"orig": "scope_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "sending_domain",
										},
										map[string]any{
											"name": "scope_value",
											"orig": "scope_value",
											"type": "`$STRING`",
											"kind": "query",
											"example": "mail.example.com",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
											"example": "billing@",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"cursor",
										"email",
										"per_page",
										"reason",
										"scope_type",
										"scope_value",
										"search",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/suppression",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "desc",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "created_at",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"direction",
										"order",
										"page",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/suppression",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"suppression",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
									"query": []any{
										map[string]any{
											"name": "direction",
											"orig": "direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "desc",
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "created_at",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"direction",
										"order",
										"page",
										"per_page",
										"zone_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/zones/{zone_id}/email/sending/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email/sending/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "54442216",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/zones/{zone_id}/email/sending/suppression/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppression",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"suppression",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "88228822",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"accounts",
									"{account_id}",
									"email",
									"sending",
									"suppressions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "12345678",
										},
										map[string]any{
											"name": "id",
											"orig": "suppression_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "396a5436-d4b0-42a6-b3fc-48e8fa522321",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"subdomain": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "subdomain",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
								"segments": []any{
									map[string]any{
										"lit": "zones",
									},
									map[string]any{
										"var": "zone_id",
									},
									map[string]any{
										"lit": "email",
									},
									map[string]any{
										"lit": "sending",
									},
									map[string]any{
										"lit": "subdomains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"zones",
									"{zone_id}",
									"email",
									"sending",
									"subdomains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subdomain_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subdomain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "aabbccdd11223344aabbccdd11223344",
										},
										map[string]any{
											"name": "zone_id",
											"orig": "zone_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "023e105f4ecef8ad9ca31a8372d0c353",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"zone_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
