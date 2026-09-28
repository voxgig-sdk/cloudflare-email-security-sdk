# CloudflareEmailSecurity SDK configuration

module CloudflareEmailSecurityConfig
  # Return the process-wide config, built once on first use. The SDK reads
  # the config on every request and never writes to it, so one instance is
  # shared by every client rather than rebuilt per client.
  #
  # The returned hash is shared: treat it as read-only. Callers that need to
  # mutate should use make_config, which always returns a fresh copy.
  def self.shared_config
    @shared_config ||= make_config
  end


  # Build a fresh, fully materialised config hash. Every call rebuilds the
  # whole structure, so prefer shared_config unless you need a private copy
  # you intend to mutate.
  def self.make_config
    {
      "main" => {
        "name" => "CloudflareEmailSecurity",
        "slug" => "cloudflare-email-security",
        "version" => "0.0.1",
        "target" => "rb",
      },
      "feature" => {
        "test" => {
          "options" => {
            "active" => false,
          },
          "optspec" => {
            "entity" => "`$MAP`",
            "net" => "`$MAP`",
          },
          "strict" => false,
          "transport" => "base",
        },
      },
      "options" => {
        "base" => "https://api.cloudflare.com/client/v4",
        "auth" => {
          "prefix" => "",
          "name" => "X-Auth-Email",
        },
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "email_account_rules_plan_response_single" => {},
          "email_account_rules_response_collection" => {},
          "email_auth" => {},
          "email_catch_all_rule_response_single" => {},
          "email_destination_address_response_single" => {},
          "email_destination_addresses_response_collection" => {},
          "email_dns_settings_response_collection" => {},
          "email_email_settings_response_single" => {},
          "email_routing_suppression" => {},
          "email_rule_response_single" => {},
          "email_rules_response_collection" => {},
          "email_security" => {},
          "email_security_analytics" => {},
          "email_security_setting" => {},
          "email_sending" => {},
          "email_sending_limits_response_single" => {},
          "email_sending_reputation_response_single" => {},
          "email_sending_subdomain_dns_status" => {},
          "email_sending_subdomain_preview" => {},
          "email_sending_subdomain_reputation_complaint" => {},
          "email_sending_subdomain_response_single" => {},
          "email_sending_subdomains_response_collection" => {},
          "email_sending_suppression" => {},
          "subdomain" => {},
        },
      },
      "entity" => {
        "email_account_rules_plan_response_single" => {
          "fields" => [
            {
              "name" => "catch_all_rules",
              "title" => "Catch All Rules",
              "type" => "`$ARRAY`",
              "short" => "Desired catch-all Email Routing rules managed by the deploying Worker.",
            },
            {
              "name" => "owner_worker_tag",
              "title" => "Owner Worker Tag",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Public tag (script_tag) of the Worker that owns this rule.",
              "writeOnly" => true,
            },
            {
              "name" => "rules",
              "title" => "Rules",
              "type" => "`$ARRAY`",
              "short" => "Desired normal Email Routing rules managed by the deploying Worker.",
            },
          ],
          "name" => "email_account_rules_plan_response_single",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/routing/rules/plan",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "lit" => "plan",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "rules",
                    "plan",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_account_rules_response_collection" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "result_info",
              "title" => "Result Info",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_account_rules_response_collection",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/routing/rules",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "rules",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "enabled",
                        "orig" => "enabled",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                        "example" => true,
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "enabled",
                      "page",
                      "per_page",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_auth" => {
          "fields" => [
            {
              "name" => "enabled",
              "title" => "Enabled",
              "type" => "`$BOOLEAN`",
              "short" => "Enable or disable DMARC reports for this zone",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
              "op" => {
                "load" => {
                  "type" => "`$OBJECT`",
                },
              },
              "short" => "Recursive SPF inspection tree",
            },
            {
              "name" => "skip_wizard",
              "title" => "Skip Wizard",
              "type" => "`$BOOLEAN`",
              "short" => "Skip the DMARC setup wizard",
            },
          ],
          "name" => "email_auth",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/auth/spf/inspect",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "auth",
                    },
                    {
                      "lit" => "spf",
                    },
                    {
                      "lit" => "inspect",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "auth",
                    "spf",
                    "inspect",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "id",
                        "orig" => "id",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "reqd" => true,
                        "example" => "e5bb46707a802688812d5d1c9f7977d4",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/auth/dmarc-reports",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "auth",
                    },
                    {
                      "lit" => "dmarc-reports",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "auth",
                    "dmarc-reports",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/zones/{zone_id}/email/auth/dmarc-reports",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "auth",
                    },
                    {
                      "lit" => "dmarc-reports",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "auth",
                    "dmarc-reports",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_catch_all_rule_response_single" => {
          "fields" => [
            {
              "name" => "actions",
              "title" => "Actions",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List actions for the catch-all routing rule.",
            },
            {
              "name" => "enabled",
              "title" => "Enabled",
              "type" => "`$BOOLEAN`",
              "short" => "Routing rule status.",
            },
            {
              "name" => "matchers",
              "title" => "Matchers",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List of matchers for the catch-all routing rule.",
            },
            {
              "name" => "name",
              "title" => "Name",
              "type" => "`$STRING`",
              "short" => "Routing rule name.",
            },
            {
              "name" => "owner_worker_tag",
              "title" => "Owner Worker Tag",
              "type" => "`$STRING`",
              "short" => "Public tag (script_tag) of the Worker that owns this rule.",
              "writeOnly" => true,
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
            {
              "name" => "source",
              "title" => "Source",
              "type" => "`$STRING`",
              "short" => "Who manages the rule.",
            },
          ],
          "name" => "email_catch_all_rule_response_single",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/rules/catch_all",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "lit" => "catch_all",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                    "catch_all",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/zones/{zone_id}/email/routing/rules/catch_all",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "lit" => "catch_all",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                    "catch_all",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_destination_address_response_single" => {
          "fields" => [
            {
              "name" => "email",
              "title" => "Email",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The contact email address of the user.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
            {
              "name" => "status",
              "title" => "Status",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Destination address status.",
            },
          ],
          "name" => "email_destination_address_response_single",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/routing/addresses",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "addresses",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "addresses",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "addresses",
                    },
                    {
                      "var" => "destination_address_identifier",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "addresses",
                    "{destination_address_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "destination_address_identifier",
                        "orig" => "destination_address_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "ea95132c15732412d22c1476fa83f27a",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "destination_address_identifier",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "addresses",
                    },
                    {
                      "var" => "destination_address_identifier",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "addresses",
                    "{destination_address_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "destination_address_identifier",
                        "orig" => "destination_address_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "ea95132c15732412d22c1476fa83f27a",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "destination_address_identifier",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "addresses",
                    },
                    {
                      "var" => "destination_address_identifier",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "addresses",
                    "{destination_address_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "destination_address_identifier",
                        "orig" => "destination_address_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "ea95132c15732412d22c1476fa83f27a",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "destination_address_identifier",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_destination_addresses_response_collection" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "result_info",
              "title" => "Result Info",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_destination_addresses_response_collection",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/routing/addresses",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "addresses",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "addresses",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "asc",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "verified",
                        "orig" => "verified",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                        "example" => true,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "page",
                      "per_page",
                      "verified",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_dns_settings_response_collection" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ARRAY`",
            },
          ],
          "name" => "email_dns_settings_response_collection",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "subdomain",
                        "orig" => "subdomain",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "example.net",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain",
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "$.main.kit.entity.subdomain",
              ],
            ],
          },
        },
        "email_email_settings_response_single" => {
          "fields" => [
            {
              "name" => "enabled",
              "title" => "Enabled",
              "type" => "`$BOOLEAN`",
              "short" => "State of your zone Email Routing settings.",
            },
            {
              "name" => "name",
              "title" => "Name",
              "type" => "`$STRING`",
              "short" => "Domain of your zone.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
            {
              "name" => "skip_wizard",
              "title" => "Skip Wizard",
              "type" => "`$BOOLEAN`",
              "short" => "Flag to check if the user skipped the configuration wizard.",
            },
            {
              "name" => "support_subaddress",
              "title" => "Support Subaddress",
              "type" => "`$BOOLEAN`",
              "short" => "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.",
            },
          ],
          "name" => "email_email_settings_response_single",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/disable",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "disable",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "disable",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/enable",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "enable",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "enable",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/unlock",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "unlock",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "unlock",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/zones/{zone_id}/email/routing",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/zones/{zone_id}/email/routing/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/zones/{zone_id}/email/routing/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/zones/{zone_id}/email/routing",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_routing_suppression" => {
          "fields" => [
            {
              "name" => "created_at",
              "title" => "Created At",
              "type" => "`$STRING`",
              "req" => true,
              "format" => "date-time",
            },
            {
              "name" => "email",
              "title" => "Email",
              "type" => "`$STRING`",
              "req" => true,
              "format" => "email",
            },
            {
              "name" => "expires_at",
              "title" => "Expires At",
              "type" => "`$STRING`",
              "req" => true,
              "op" => {
                "create" => {
                  "type" => "`$STRING`",
                },
              },
              "format" => "date-time",
            },
            {
              "name" => "id",
              "title" => "Id",
              "type" => "`$STRING`",
              "req" => true,
              "format" => "uuid",
            },
            {
              "name" => "reason",
              "title" => "Reason",
              "type" => "`$STRING`",
              "req" => true,
            },
            {
              "name" => "zones",
              "title" => "Zones",
              "type" => "`$ARRAY`",
            },
          ],
          "id" => {
            "field" => "id",
            "name" => "id",
          },
          "name" => "email_routing_suppression",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/routing/suppression",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/suppression",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/routing/suppression",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "desc",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "created_at",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 100,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "order",
                      "page",
                      "per_page",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/suppression",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "desc",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "created_at",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 100,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "direction",
                      "order",
                      "page",
                      "per_page",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/routing/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email/routing/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "routing",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/zones/{zone_id}/email/routing/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_rule_response_single" => {
          "fields" => [
            {
              "name" => "actions",
              "title" => "Actions",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List actions patterns.",
            },
            {
              "name" => "enabled",
              "title" => "Enabled",
              "type" => "`$BOOLEAN`",
              "short" => "Routing rule status.",
            },
            {
              "name" => "matchers",
              "title" => "Matchers",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "Matching patterns to forward to your actions.",
            },
            {
              "name" => "name",
              "title" => "Name",
              "type" => "`$STRING`",
              "short" => "Routing rule name.",
            },
            {
              "name" => "owner_worker_tag",
              "title" => "Owner Worker Tag",
              "type" => "`$STRING`",
              "short" => "Public tag (script_tag) of the Worker that owns this rule.",
              "writeOnly" => true,
            },
            {
              "name" => "priority",
              "title" => "Priority",
              "type" => "`$NUMBER`",
              "short" => "Priority of the routing rule.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
            {
              "name" => "source",
              "title" => "Source",
              "type" => "`$STRING`",
              "short" => "Who manages the rule.",
            },
          ],
          "name" => "email_rule_response_single",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/routing/rules",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "var" => "rule_identifier",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                    "{rule_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "rule_identifier",
                        "orig" => "rule_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "a7e6fb77503c41d8a7f3113c6918f10c",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "rule_identifier",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "var" => "rule_identifier",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                    "{rule_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "rule_identifier",
                        "orig" => "rule_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "a7e6fb77503c41d8a7f3113c6918f10c",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "rule_identifier",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/zones/{zone_id}/email/routing/rules/{rule_identifier}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                    {
                      "var" => "rule_identifier",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                    "{rule_identifier}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "rule_identifier",
                        "orig" => "rule_identifier",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "a7e6fb77503c41d8a7f3113c6918f10c",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "rule_identifier",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_rules_response_collection" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "result_info",
              "title" => "Result Info",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_rules_response_collection",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/routing/rules",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "routing",
                    },
                    {
                      "lit" => "rules",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "routing",
                    "rules",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "enabled",
                        "orig" => "enabled",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                        "example" => true,
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$NUMBER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "enabled",
                      "page",
                      "per_page",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_security" => {
          "fields" => [
            {
              "name" => "action",
              "title" => "Action",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The action the job performs on every message matching the search parameters.",
            },
            {
              "name" => "comment",
              "title" => "Comment",
              "type" => "`$STRING`",
              "short" => "Optional note describing the job.",
            },
            {
              "name" => "destination",
              "title" => "Destination",
              "type" => "`$ANY`",
              "op" => {
                "create" => {
                  "req" => true,
                  "type" => "`$STRING`",
                },
              },
              "short" => "Required when action is 'MOVE'.",
            },
            {
              "name" => "errors",
              "title" => "Errors",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "expected_disposition",
              "title" => "Expected Disposition",
              "type" => "`$STRING`",
              "short" => "Nonfunctional field.",
              "deprecated" => true,
            },
            {
              "name" => "id",
              "title" => "Id",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Unique identifier for a message retrieved from investigation.",
            },
            {
              "name" => "ids",
              "title" => "Ids",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List of message IDs to move.",
            },
            {
              "name" => "messages",
              "title" => "Messages",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "postfix_ids",
              "title" => "Postfix Ids",
              "type" => "`$ARRAY`",
              "short" => "Deprecated, use `ids` instead.",
              "deprecated" => true,
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "result_info",
              "title" => "Result Info",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "search_params",
              "title" => "Search Params",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "success",
              "title" => "Success",
              "type" => "`$BOOLEAN`",
              "req" => true,
              "short" => "Whether the API call was successful.",
            },
          ],
          "id" => {
            "field" => "id",
            "name" => "id",
          },
          "name" => "email_security",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                    {
                      "var" => "bulk_id",
                    },
                    {
                      "lit" => "cancel",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                    "{bulk_id}",
                    "cancel",
                  ],
                  "rename" => {
                    "param" => {
                      "job_id" => "bulk_id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "bulk_id",
                        "orig" => "job_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "cancel",
                    "exist" => [
                      "account_id",
                      "bulk_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/move",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "move",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "move",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "move",
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "reclassify",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "reclassify",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "reclassify",
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/move",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "move",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "move",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/preview",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "preview",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "preview",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/investigate/release",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "release",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "release",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/submissions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "submissions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "submissions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "submission",
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "alert_id",
                        "orig" => "alert_id",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "cursor",
                        "orig" => "cursor",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "delivery_status",
                        "orig" => "delivery_status",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "detections_only",
                        "orig" => "detections_only",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                        "example" => true,
                      },
                      {
                        "name" => "domain",
                        "orig" => "domain",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "example.com",
                      },
                      {
                        "name" => "end",
                        "orig" => "end",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2022-07-25T14:30:00Z",
                      },
                      {
                        "name" => "final_disposition",
                        "orig" => "final_disposition",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "message_action",
                        "orig" => "message_action",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "message_id",
                        "orig" => "message_id",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "metric",
                        "orig" => "metric",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "query",
                        "orig" => "query",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "bob jones",
                      },
                      {
                        "name" => "recipient",
                        "orig" => "recipient",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "me@example.com",
                      },
                      {
                        "name" => "sender",
                        "orig" => "sender",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "noreply@example.com",
                      },
                      {
                        "name" => "smtp_helo_ip",
                        "orig" => "smtp_helo_ip",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "192.0.2.1",
                      },
                      {
                        "name" => "start",
                        "orig" => "start",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2022-06-25T14:30:00Z",
                      },
                      {
                        "name" => "subject",
                        "orig" => "subject",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "investigate",
                    "exist" => [
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
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/submissions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "submissions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "submissions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "end",
                        "orig" => "end",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2022-07-25T14:30:00Z",
                      },
                      {
                        "name" => "escalated_from_user",
                        "orig" => "escalated_from_user",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "requested_at",
                      },
                      {
                        "name" => "original_disposition",
                        "orig" => "original_disposition",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "outcome_disposition",
                        "orig" => "outcome_disposition",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "query",
                        "orig" => "query",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "requested_disposition",
                        "orig" => "requested_disposition",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "start",
                        "orig" => "start",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2022-06-25T14:30:00Z",
                      },
                      {
                        "name" => "status",
                        "orig" => "status",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "submission_id",
                        "orig" => "submission_id",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "type",
                        "orig" => "type",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "$action" => "submission",
                    "exist" => [
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
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/phishguard/reports",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "phishguard",
                    },
                    {
                      "lit" => "reports",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "phishguard",
                    "reports",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "end",
                        "orig" => "end",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2020-09-01T00:00:00Z",
                      },
                      {
                        "name" => "from_date",
                        "orig" => "from_date",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2020-08-01T00:00:00Z",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "start",
                        "orig" => "start",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2020-08-01T00:00:00Z",
                      },
                      {
                        "name" => "to_date",
                        "orig" => "to_date",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "2020-09-01T00:00:00Z",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "end",
                      "from_date",
                      "page",
                      "per_page",
                      "start",
                      "to_date",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "action_type",
                        "orig" => "action_type",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "status",
                        "orig" => "status",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "action_type",
                      "page",
                      "per_page",
                      "status",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk/{job_id}/messages",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                    {
                      "var" => "bulk_id",
                    },
                    {
                      "lit" => "messages",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                    "{bulk_id}",
                    "messages",
                  ],
                  "rename" => {
                    "param" => {
                      "job_id" => "bulk_id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "bulk_id",
                        "orig" => "job_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "status",
                        "orig" => "status",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "bulk_id",
                      "page",
                      "per_page",
                      "status",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/action_log",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "action_log",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "action_log",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                      "page",
                      "per_page",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/detections",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "detections",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "detections",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/preview",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "preview",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "preview",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/raw",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "raw",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "raw",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}/trace",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                    {
                      "lit" => "trace",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                    "trace",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/{investigate_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "var" => "investigate_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "{investigate_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "investigate_id",
                        "orig" => "investigate_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "4Njp3P0STMz2c02Q-2024-01-05T10:00:00-12345678",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "submission",
                        "orig" => "submission",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "investigate_id",
                      "submission",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk/{job_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                    {
                      "var" => "job_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                    "{job_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "job_id",
                        "orig" => "job_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "job_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/investigate/bulk/{job_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "investigate",
                    },
                    {
                      "lit" => "bulk",
                    },
                    {
                      "var" => "job_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "investigate",
                    "bulk",
                    "{job_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "job_id",
                        "orig" => "job_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "job_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_security_analytics" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_security_analytics",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/analytics/monthly_report",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "analytics",
                    },
                    {
                      "lit" => "monthly_report",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "analytics",
                    "monthly_report",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_security_setting" => {
          "fields" => [
            {
              "name" => "allowed_delivery_modes",
              "title" => "Allowed Delivery Modes",
              "type" => "`$ARRAY`",
              "req" => true,
              "op" => {
                "patch" => {
                  "type" => "`$ARRAY`",
                },
              },
              "short" => "Delivery modes to onboard the domain through.",
            },
            {
              "name" => "comments",
              "title" => "Comments",
              "type" => "`$STRING`",
              "short" => "Optional note describing the reason for the ignore pattern.",
            },
            {
              "name" => "deletes",
              "title" => "Deletes",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "IDs of the allow policies to delete.",
            },
            {
              "name" => "domain",
              "title" => "Domain",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The email domain to protect.",
            },
            {
              "name" => "drop_dispositions",
              "title" => "Drop Dispositions",
              "type" => "`$ARRAY`",
              "req" => true,
              "op" => {
                "patch" => {
                  "type" => "`$ARRAY`",
                },
              },
              "short" => "Dispositions to drop instead of delivering, e.g.",
            },
            {
              "name" => "folder",
              "title" => "Folder",
              "type" => "`$ANY`",
              "short" => "The mailbox folder to scan, for API-scanning domains.",
            },
            {
              "name" => "integration_id",
              "title" => "Integration Id",
              "type" => "`$STRING`",
              "short" => "Identifier of the CASB integration that authorizes this domain.",
              "format" => "uuid",
            },
            {
              "name" => "ip_restrictions",
              "title" => "Ip Restrictions",
              "type" => "`$ARRAY`",
              "req" => true,
              "op" => {
                "patch" => {
                  "type" => "`$ARRAY`",
                },
              },
              "short" => "Source IP ranges mail is accepted from.",
            },
            {
              "name" => "lookback_hops",
              "title" => "Lookback Hops",
              "type" => "`$INTEGER`",
              "short" => "Number of hops to trace back through received headers when reconstructing the original message (1-20).",
            },
            {
              "name" => "patches",
              "title" => "Patches",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "Partial updates to apply — each entry carries the policy's ID and only the fields to change.",
            },
            {
              "name" => "pattern",
              "title" => "Pattern",
              "type" => "`$STRING`",
              "req" => true,
              "op" => {
                "patch" => {
                  "type" => "`$STRING`",
                },
              },
              "short" => "Regular expression identifying URLs to exempt from rewriting.",
            },
            {
              "name" => "posts",
              "title" => "Posts",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "Allow policies to create.",
            },
            {
              "name" => "puts",
              "title" => "Puts",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "Full replacements to apply — each entry carries the policy's ID and every field of its new value.",
            },
            {
              "name" => "regions",
              "title" => "Regions",
              "type" => "`$ARRAY`",
              "req" => true,
              "op" => {
                "patch" => {
                  "type" => "`$ARRAY`",
                },
              },
              "short" => "Regions that process messages for this domain, e.g.",
            },
            {
              "name" => "require_tls_inbound",
              "title" => "Require Tls Inbound",
              "type" => "`$BOOLEAN`",
              "short" => "Require TLS on inbound connections.",
            },
            {
              "name" => "require_tls_outbound",
              "title" => "Require Tls Outbound",
              "type" => "`$BOOLEAN`",
              "short" => "Require TLS on outbound connections.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "op" => {
                "load" => {
                  "req" => true,
                  "type" => "`$OBJECT`",
                },
              },
              "short" => "An impersonation registry entry.",
            },
            {
              "name" => "transport",
              "title" => "Transport",
              "type" => "`$STRING`",
              "short" => "The mail transport hostname for MX/Inline delivery — the MX record Cloudflare delivers email to (e.g.",
            },
          ],
          "name" => "email_security_setting",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => {
                      "allowed_delivery_modes" => "`reqdata.allowed_delivery_mode`",
                      "domain" => "`reqdata.domain`",
                      "drop_dispositions" => "`reqdata.drop_disposition`",
                      "folder" => "`reqdata.folder`",
                      "integration_id" => "`reqdata.integration_id`",
                      "ip_restrictions" => "`reqdata.ip_restriction`",
                      "lookback_hops" => "`reqdata.lookback_hop`",
                      "regions" => "`reqdata.region`",
                      "require_tls_inbound" => "`reqdata.require_tls_inbound`",
                      "require_tls_outbound" => "`reqdata.require_tls_outbound`",
                      "transport" => "`reqdata.transport`",
                    },
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/impersonation_registry",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "impersonation_registry",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "impersonation_registry",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => {
                      "comments" => "`reqdata.comment`",
                      "pattern" => "`reqdata.pattern`",
                    },
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns/batch",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                    {
                      "lit" => "batch",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                    "batch",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "is_acceptable_sender",
                        "orig" => "is_acceptable_sender",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "is_exempt_recipient",
                        "orig" => "is_exempt_recipient",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "is_trusted_sender",
                        "orig" => "is_trusted_sender",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "pattern",
                        "orig" => "pattern",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "pattern_type",
                        "orig" => "pattern_type",
                        "type" => "`$ANY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "verify_sender",
                        "orig" => "verify_sender",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
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
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "active_delivery_mode",
                        "orig" => "active_delivery_mode",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "allowed_delivery_mode",
                        "orig" => "allowed_delivery_mode",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "domain",
                        "orig" => "domain",
                        "type" => "`$ARRAY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "integration_id",
                        "orig" => "integration_id",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "status",
                        "orig" => "status",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
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
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "is_recent",
                        "orig" => "is_recent",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "is_similarity",
                        "orig" => "is_similarity",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "pattern",
                        "orig" => "pattern",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "is_recent",
                      "is_similarity",
                      "order",
                      "page",
                      "pattern",
                      "per_page",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "enabled",
                        "orig" => "enabled",
                        "type" => "`$BOOLEAN`",
                        "kind" => "query",
                      },
                      {
                        "name" => "name",
                        "orig" => "name",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "enabled",
                      "name",
                      "order",
                      "page",
                      "per_page",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "pattern",
                        "orig" => "pattern",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "pattern_type",
                        "orig" => "pattern_type",
                        "type" => "`$ANY`",
                        "kind" => "query",
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "order",
                      "page",
                      "pattern",
                      "pattern_type",
                      "per_page",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/impersonation_registry",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "impersonation_registry",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "impersonation_registry",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "provenance",
                        "orig" => "provenance",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "order",
                      "page",
                      "per_page",
                      "provenance",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "order",
                      "page",
                      "per_page",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 20,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "page",
                      "per_page",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "var" => "domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "{domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "domain_id",
                        "orig" => "domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/{domain_id}/verification",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "var" => "domain_id",
                    },
                    {
                      "lit" => "verification",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "{domain_id}",
                    "verification",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "domain_id",
                        "orig" => "domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "impersonation_registry",
                    },
                    {
                      "var" => "impersonation_registry_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "impersonation_registry",
                    "{impersonation_registry_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "impersonation_registry_id",
                        "orig" => "impersonation_registry_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "impersonation_registry_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                    {
                      "var" => "sending_domain_restriction_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                    "{sending_domain_restriction_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "sending_domain_restriction_id",
                        "orig" => "sending_domain_restriction_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "sending_domain_restriction_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                    {
                      "var" => "trusted_domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                    "{trusted_domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "trusted_domain_id",
                        "orig" => "trusted_domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "trusted_domain_id",
                    ],
                  },
                },
              ],
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "var" => "domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "{domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "domain_id",
                        "orig" => "domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "impersonation_registry",
                    },
                    {
                      "var" => "impersonation_registry_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "impersonation_registry",
                    "{impersonation_registry_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "impersonation_registry_id",
                        "orig" => "impersonation_registry_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "impersonation_registry_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => {
                      "comments" => "`reqdata.comment`",
                      "pattern" => "`reqdata.pattern`",
                    },
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                    {
                      "var" => "sending_domain_restriction_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                    "{sending_domain_restriction_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "sending_domain_restriction_id",
                        "orig" => "sending_domain_restriction_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "sending_domain_restriction_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                    {
                      "var" => "trusted_domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                    "{trusted_domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "trusted_domain_id",
                        "orig" => "trusted_domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "trusted_domain_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "var" => "domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "{domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "domain_id",
                        "orig" => "domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/impersonation_registry/{impersonation_registry_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "impersonation_registry",
                    },
                    {
                      "var" => "impersonation_registry_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "impersonation_registry",
                    "{impersonation_registry_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "impersonation_registry_id",
                        "orig" => "impersonation_registry_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "impersonation_registry_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "url_ignore_patterns",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "url_ignore_patterns",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/content_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "content_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "content_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "sending_domain_restrictions",
                    },
                    {
                      "var" => "sending_domain_restriction_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "sending_domain_restrictions",
                    "{sending_domain_restriction_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "sending_domain_restriction_id",
                        "orig" => "sending_domain_restriction_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "sending_domain_restriction_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                    {
                      "var" => "trusted_domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                    "{trusted_domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "trusted_domain_id",
                        "orig" => "trusted_domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "trusted_domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/accounts/{account_id}/email-security/settings/domains/{domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "domains",
                    },
                    {
                      "var" => "domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "domains",
                    "{domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "domain_id",
                        "orig" => "domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "domain_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/accounts/{account_id}/email-security/settings/block_senders/{pattern_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "block_senders",
                    },
                    {
                      "var" => "pattern_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "block_senders",
                    "{pattern_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "pattern_id",
                        "orig" => "pattern_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "pattern_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/accounts/{account_id}/email-security/settings/allow_policies/{policy_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "allow_policies",
                    },
                    {
                      "var" => "policy_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "allow_policies",
                    "{policy_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "policy_id",
                        "orig" => "policy_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "policy_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "PUT",
                  "orig" => "/accounts/{account_id}/email-security/settings/trusted_domains/{trusted_domain_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email-security",
                    },
                    {
                      "lit" => "settings",
                    },
                    {
                      "lit" => "trusted_domains",
                    },
                    {
                      "var" => "trusted_domain_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email-security",
                    "settings",
                    "trusted_domains",
                    "{trusted_domain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                      {
                        "name" => "trusted_domain_id",
                        "orig" => "trusted_domain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "trusted_domain_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending" => {
          "fields" => [
            {
              "name" => "attachments",
              "title" => "Attachments",
              "type" => "`$ARRAY`",
              "short" => "File attachments and inline images.",
            },
            {
              "name" => "bcc",
              "title" => "Bcc",
              "type" => "`$ANY`",
              "short" => "Blind carbon copy recipient(s).",
            },
            {
              "name" => "cc",
              "title" => "Cc",
              "type" => "`$ANY`",
              "short" => "Carbon copy recipient(s).",
            },
            {
              "name" => "errors",
              "title" => "Errors",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "from",
              "title" => "From",
              "type" => "`$ANY`",
              "req" => true,
              "short" => "Sender email address.",
            },
            {
              "name" => "headers",
              "title" => "Headers",
              "type" => "`$OBJECT`",
              "short" => "Custom email headers as key-value pairs.",
            },
            {
              "name" => "html",
              "title" => "Html",
              "type" => "`$STRING`",
              "short" => "HTML body of the email.",
            },
            {
              "name" => "messages",
              "title" => "Messages",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "mime_message",
              "title" => "Mime Message",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The full MIME-encoded email message.",
            },
            {
              "name" => "recipients",
              "title" => "Recipients",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "List of recipient email addresses.",
            },
            {
              "name" => "reply_to",
              "title" => "Reply To",
              "type" => "`$ANY`",
              "short" => "Reply-to address.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "result_info",
              "title" => "Result Info",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "subject",
              "title" => "Subject",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Email subject line.",
            },
            {
              "name" => "success",
              "title" => "Success",
              "type" => "`$BOOLEAN`",
              "req" => true,
            },
            {
              "name" => "text",
              "title" => "Text",
              "type" => "`$STRING`",
              "short" => "Plain text body of the email.",
            },
            {
              "name" => "to",
              "title" => "To",
              "type" => "`$ANY`",
              "short" => "Recipient(s).",
            },
          ],
          "name" => "email_sending",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/sending/send",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "send",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "send",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => {
                      "attachments" => "`reqdata.attachment`",
                      "bcc" => "`reqdata.bcc`",
                      "cc" => "`reqdata.cc`",
                      "from" => "`reqdata.from`",
                      "headers" => "`reqdata.header`",
                      "html" => "`reqdata.html`",
                      "reply_to" => "`reqdata.reply_to`",
                      "subject" => "`reqdata.subject`",
                      "text" => "`reqdata.text`",
                      "to" => "`reqdata.to`",
                    },
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/sending/send_raw",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "send_raw",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "send_raw",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => {
                      "from" => "`reqdata.from`",
                      "mime_message" => "`reqdata.mime_message`",
                      "recipients" => "`reqdata.recipient`",
                    },
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/messages/{message_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "messages",
                    },
                    {
                      "var" => "message_id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "messages",
                    "{message_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345",
                      },
                      {
                        "name" => "message_id",
                        "orig" => "message_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "<aB3xK9mP@example.com>",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "message_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending_limits_response_single" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_sending_limits_response_single",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/limits",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "limits",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "limits",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending_reputation_response_single" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
            },
          ],
          "name" => "email_sending_reputation_response_single",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/reputation",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "reputation",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "reputation",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending_subdomain_dns_status" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
            },
          ],
          "name" => "email_sending_subdomain_dns_status",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                    {
                      "lit" => "dns",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                    "dns",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns/status",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                    {
                      "lit" => "dns",
                    },
                    {
                      "lit" => "status",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                    "dns",
                    "status",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "$.main.kit.entity.subdomain",
              ],
            ],
          },
        },
        "email_sending_subdomain_preview" => {
          "fields" => [
            {
              "name" => "name",
              "title" => "Name",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The domain name within the zone.",
            },
          ],
          "name" => "email_sending_subdomain_preview",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/preview",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "lit" => "preview",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "preview",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending_subdomain_reputation_complaint" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
            },
          ],
          "name" => "email_sending_subdomain_reputation_complaint",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                    {
                      "lit" => "reputation",
                    },
                    {
                      "lit" => "complaints",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                    "reputation",
                    "complaints",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "end_at",
                        "orig" => "end_at",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "reqd" => true,
                      },
                      {
                        "name" => "start_at",
                        "orig" => "start_at",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "reqd" => true,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "end_at",
                      "start_at",
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "$.main.kit.entity.subdomain",
              ],
            ],
          },
        },
        "email_sending_subdomain_response_single" => {
          "fields" => [
            {
              "name" => "drop_suppressed_recipients",
              "title" => "Drop Suppressed Recipients",
              "type" => "`$BOOLEAN`",
              "short" => "Whether a send request that includes a recipient suppressed on this subdomain drops that recipient and still delivers to the rest, instead of failing the entire request.",
            },
            {
              "name" => "name",
              "title" => "Name",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The domain name within the zone.",
            },
            {
              "name" => "preview_enabled",
              "title" => "Preview Enabled",
              "type" => "`$BOOLEAN`",
              "short" => "Whether sent messages from this subdomain can be previewed in the activity log.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ANY`",
            },
          ],
          "name" => "email_sending_subdomain_response_single",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "subdomain_id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{subdomain_id}",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "subdomain_id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "subdomain_id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "$.main.kit.entity.subdomain",
              ],
            ],
          },
        },
        "email_sending_subdomains_response_collection" => {
          "fields" => [
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$ARRAY`",
            },
          ],
          "name" => "email_sending_subdomains_response_collection",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "email_sending_suppression" => {
          "fields" => [
            {
              "name" => "created_at",
              "title" => "Created At",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "When the suppression was created.",
              "format" => "date-time",
            },
            {
              "name" => "email",
              "title" => "Email",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "The email address to suppress.",
              "format" => "email",
            },
            {
              "name" => "errors",
              "title" => "Errors",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "expires_at",
              "title" => "Expires At",
              "type" => "`$STRING`",
              "req" => true,
              "op" => {
                "create" => {
                  "type" => "`$STRING`",
                },
                "update" => {
                  "type" => "`$STRING`",
                },
              },
              "short" => "Expiration timestamp for the suppression.",
              "format" => "date-time",
            },
            {
              "name" => "id",
              "title" => "Id",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Unique identifier for this suppression.",
              "format" => "uuid",
            },
            {
              "name" => "items",
              "title" => "Items",
              "type" => "`$ARRAY`",
              "req" => true,
              "short" => "Suppressions to import.",
            },
            {
              "name" => "messages",
              "title" => "Messages",
              "type" => "`$ARRAY`",
              "req" => true,
            },
            {
              "name" => "note",
              "title" => "Note",
              "type" => "`$STRING`",
              "short" => "Advisory note for this suppression.",
            },
            {
              "name" => "read_only",
              "title" => "Read Only",
              "type" => "`$BOOLEAN`",
              "req" => true,
              "short" => "Whether clients may mutate this suppression.",
            },
            {
              "name" => "reason",
              "title" => "Reason",
              "type" => "`$STRING`",
              "req" => true,
              "short" => "Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`.",
            },
            {
              "name" => "result",
              "title" => "Result",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "scope",
              "title" => "Scope",
              "type" => "`$ANY`",
              "short" => "Where the suppression applies.",
            },
            {
              "name" => "success",
              "title" => "Success",
              "type" => "`$BOOLEAN`",
              "req" => true,
            },
            {
              "name" => "zones",
              "title" => "Zones",
              "type" => "`$ARRAY`",
            },
          ],
          "id" => {
            "field" => "id",
            "name" => "id",
          },
          "name" => "email_sending_suppression",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/sending/suppression",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions/bulk",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                    {
                      "lit" => "bulk",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                    "bulk",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/zones/{zone_id}/email/sending/suppression",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "cursor",
                        "orig" => "cursor",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "eyJ0Ijo0LCJwaCI6ImRvbWFpbiIsInAiOjEyMywiYyI6ImMzZGM1ZjBiMzRhMTRmZjhlMWIzZWMwNDg5NWUxYjIyIn0",
                      },
                      {
                        "name" => "email",
                        "orig" => "email",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "user@example.com",
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 100,
                      },
                      {
                        "name" => "reason",
                        "orig" => "reason",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "hard_bounce",
                      },
                      {
                        "name" => "scope_type",
                        "orig" => "scope_type",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "sending_domain",
                      },
                      {
                        "name" => "scope_value",
                        "orig" => "scope_value",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "mail.example.com",
                      },
                      {
                        "name" => "search",
                        "orig" => "search",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "billing@",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "cursor",
                      "email",
                      "per_page",
                      "reason",
                      "scope_type",
                      "scope_value",
                      "search",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/suppression",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "desc",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "created_at",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 100,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "direction",
                      "order",
                      "page",
                      "per_page",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/suppression",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "suppression",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "direction",
                        "orig" => "direction",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "desc",
                      },
                      {
                        "name" => "order",
                        "orig" => "order",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "created_at",
                      },
                      {
                        "name" => "page",
                        "orig" => "page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 1,
                      },
                      {
                        "name" => "per_page",
                        "orig" => "per_page",
                        "type" => "`$INTEGER`",
                        "kind" => "query",
                        "example" => 100,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "direction",
                      "order",
                      "page",
                      "per_page",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/zones/{zone_id}/email/sending/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.result`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email/sending/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "54442216",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/zones/{zone_id}/email/sending/suppression/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppression",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "suppression",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "88228822",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "PATCH",
                  "orig" => "/accounts/{account_id}/email/sending/suppressions/{suppression_id}",
                  "segments" => [
                    {
                      "lit" => "accounts",
                    },
                    {
                      "var" => "account_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "suppressions",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "accounts",
                    "{account_id}",
                    "email",
                    "sending",
                    "suppressions",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "suppression_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "account_id",
                        "orig" => "account_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "12345678",
                      },
                      {
                        "name" => "id",
                        "orig" => "suppression_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "396a5436-d4b0-42a6-b3fc-48e8fa522321",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "account_id",
                      "id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "subdomain" => {
          "fields" => [
            {
              "name" => "id",
              "title" => "Id",
              "type" => "`$STRING`",
            },
          ],
          "id" => {
            "field" => "id",
            "name" => "id",
          },
          "name" => "subdomain",
          "op" => {
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "DELETE",
                  "orig" => "/zones/{zone_id}/email/sending/subdomains/{subdomain_id}",
                  "segments" => [
                    {
                      "lit" => "zones",
                    },
                    {
                      "var" => "zone_id",
                    },
                    {
                      "lit" => "email",
                    },
                    {
                      "lit" => "sending",
                    },
                    {
                      "lit" => "subdomains",
                    },
                    {
                      "var" => "id",
                    },
                  ],
                  "parts" => [
                    "zones",
                    "{zone_id}",
                    "email",
                    "sending",
                    "subdomains",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subdomain_id" => "id",
                    },
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "params" => [
                      {
                        "name" => "id",
                        "orig" => "subdomain_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "aabbccdd11223344aabbccdd11223344",
                      },
                      {
                        "name" => "zone_id",
                        "orig" => "zone_id",
                        "type" => "`$STRING`",
                        "kind" => "param",
                        "reqd" => true,
                        "example" => "023e105f4ecef8ad9ca31a8372d0c353",
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "zone_id",
                    ],
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    CloudflareEmailSecurityFeatures.make_feature(name)
  end
end
