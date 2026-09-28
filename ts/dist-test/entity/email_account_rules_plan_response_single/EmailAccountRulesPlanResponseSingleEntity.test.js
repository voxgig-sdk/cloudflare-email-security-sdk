"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('EmailAccountRulesPlanResponseSingleEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CloudflareEmailSecuritySDK.test();
        const ent = testsdk.EmailAccountRulesPlanResponseSingle();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'email_account_rules_plan_response_single.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "catch_all_rules": { "a": true, "h": "Catch All Rules", "n": "catch_all_rules", "r": false, "sh": "Desired catch-all Email Routing rules managed by the deploying Worker.", "t": "`$ARRAY`", "key$": "catch_all_rules", "index$": 0 }, "owner_worker_tag": { "a": true, "h": "Owner Worker Tag", "n": "owner_worker_tag", "r": true, "sh": "Public tag (script_tag) of the Worker that owns this rule.", "t": "`$STRING`", "wo": true, "key$": "owner_worker_tag", "index$": 1 }, "rules": { "a": true, "h": "Rules", "n": "rules", "r": false, "sh": "Desired normal Email Routing rules managed by the deploying Worker.", "t": "`$ARRAY`", "key$": "rules", "index$": 2 } }, "name": "email_account_rules_plan_response_single", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /accounts/{account_id}/email/routing/rules/plan", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "account_id", "or": "account_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/accounts/{account_id}/email/routing/rules/plan", "q": { "exist": ["account_id"] }, "r": {}, "s": [{ "lit": "accounts" }, { "var": "account_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "rules" }, { "lit": "plan" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "email_account_rules_plan_response_single", "name__orig": "email_account_rules_plan_response_single", "Name": "EmailAccountRulesPlanResponseSingle", "name_": "email_account_rules_plan_response_single", "name-": "email-account-rules-plan-response-single", "NAME": "EMAIL_ACCOUNT_RULES_PLAN_RESPONSE_SINGLE", "index$": 0 }, { "active": true, "entity": "email_account_rules_plan_response_single", "key$": "BasicEmailAccountRulesPlanResponseSingleFlow", "kind": "basic", "name": "BasicEmailAccountRulesPlanResponseSingleFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "email_account_rules_plan_response_single_ref01" }, "m": { "account_id": "account01" }, "o": "create", "s": [], "v": [], "index$": 0 }] }, 'EmailAccountRulesPlanResponseSingle', { "POST /accounts/{account_id}/email/routing/rules/plan": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "properties": { "catch_all_rules": { "description": "Desired catch-all Email Routing rules managed by the deploying Worker.", "type": "array", "items": { "type": "object", "properties": { "rule": { "type": "object", "properties": {}, "required": [], "x-ref": "#/components/schemas/email_account_rules_plan_catch_all_rule" }, "target": { "description": "Catch-all target to plan for, using the `*@domain` shape.", "type": "string", "example": "*@example.com" } }, "required": ["target", "rule"], "x-ref": "#/components/schemas/email_account_rules_plan_catch_all_entry" }, "key$": "catch_all_rules" }, "owner_worker_tag": { "description": "Public tag (script_tag) of the Worker that owns this rule. Required when\n`source` is `wrangler`.\n", "type": "string", "example": "a7e6fb77503c41d8a7f3113c6918f10c", "maxLength": 32, "writeOnly": true, "x-auditable": true, "x-ref": "#/components/schemas/email_rule_owner_worker_tag", "key$": "owner_worker_tag" }, "rules": { "description": "Desired normal Email Routing rules managed by the deploying Worker.", "type": "array", "items": { "type": "object", "properties": { "actions": { "description": "List actions patterns.", "type": "array", "items": {}, "maxItems": 1, "minItems": 1, "x-ref": "#/components/schemas/email_rule_actions" }, "enabled": { "description": "Routing rule status.", "type": "boolean", "example": true, "default": true, "enum": [], "x-auditable": true, "x-ref": "#/components/schemas/email_rule_enabled" }, "matchers": { "description": "Matching patterns to forward to your actions.", "type": "array", "items": {}, "x-ref": "#/components/schemas/email_rule_matchers" } }, "required": ["actions", "matchers"], "x-ref": "#/components/schemas/email_account_rules_plan_rule" }, "key$": "rules" } }, "example": { "catch_all_rules": [{ "rule": { "actions": [{}], "matchers": [{}] }, "target": "*@example.com" }], "owner_worker_tag": "a7e6fb77503c41d8a7f3113c6918f10c", "rules": [{ "actions": [{ "type": "worker", "value": [] }], "matchers": [{ "field": "to", "type": "literal", "value": "support@example.com" }] }, { "actions": [{ "type": "worker", "value": [] }], "matchers": [{ "field": "to", "type": "literal", "value": "info@example.com" }] }, { "actions": [{ "type": "worker", "value": [] }], "matchers": [{ "field": "to", "type": "literal", "value": "billing@example.com" }] }] }, "required": ["owner_worker_tag"], "x-ref": "#/components/schemas/email_account_rules_plan_request", "index$": 1 } } } }, "parameters": [{ "name": "account_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const email_account_rules_plan_response_single_ref01_ent = client.EmailAccountRulesPlanResponseSingle();
        let email_account_rules_plan_response_single_ref01_data = setup.data.new.email_account_rules_plan_response_single['email_account_rules_plan_response_single_ref01'];
        email_account_rules_plan_response_single_ref01_data['account_id'] = setup.idmap['account01'];
        email_account_rules_plan_response_single_ref01_data = (await email_account_rules_plan_response_single_ref01_ent.create(email_account_rules_plan_response_single_ref01_data)).data();
        (0, node_assert_1.default)(null != email_account_rules_plan_response_single_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/email_account_rules_plan_response_single/EmailAccountRulesPlanResponseSingleTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CloudflareEmailSecuritySDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['email_account_rules_plan_response_single01', 'email_account_rules_plan_response_single02', 'email_account_rules_plan_response_single03', 'account01'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ACCOUNT_RULES_PLAN_RESPONSE_SINGLE_ENTID': idmap,
        'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
    });
    idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ACCOUNT_RULES_PLAN_RESPONSE_SINGLE_ENTID'];
    const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_ACCOUNT_RULES_PLAN_RESPONSE_SINGLE_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.CloudflareEmailSecuritySDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.CLOUDFLARE_EMAIL_SECURITY_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=EmailAccountRulesPlanResponseSingleEntity.test.js.map