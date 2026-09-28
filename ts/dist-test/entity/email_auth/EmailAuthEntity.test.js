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
(0, node_test_1.describe)('EmailAuthEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CloudflareEmailSecuritySDK.test();
        const ent = testsdk.EmailAuth();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
        for (const op of ['update', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'email_auth.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "enabled": { "a": true, "h": "Enabled", "n": "enabled", "r": false, "sh": "Enable or disable DMARC reports for this zone", "t": "`$BOOLEAN`", "key$": "enabled", "index$": 0 }, "result": { "a": true, "h": "Result", "n": "result", "op": { "load": { "req": false, "type": "`$OBJECT`" } }, "r": true, "sh": "Recursive SPF inspection tree", "t": "`$OBJECT`", "key$": "result", "index$": 1 }, "skip_wizard": { "a": true, "h": "Skip Wizard", "n": "skip_wizard", "r": false, "sh": "Skip the DMARC setup wizard", "t": "`$BOOLEAN`", "key$": "skip_wizard", "index$": 2 } }, "name": "email_auth", "op": { "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /zones/{zone_id}/email/auth/spf/inspect", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "ex": "e5bb46707a802688812d5d1c9f7977d4", "k": "query", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/zones/{zone_id}/email/auth/spf/inspect", "q": { "exist": ["id", "zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "auth" }, { "lit": "spf" }, { "lit": "inspect" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "GET /zones/{zone_id}/email/auth/dmarc-reports", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/zones/{zone_id}/email/auth/dmarc-reports", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "auth" }, { "lit": "dmarc-reports" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "load" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PATCH /zones/{zone_id}/email/auth/dmarc-reports", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/zones/{zone_id}/email/auth/dmarc-reports", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "auth" }, { "lit": "dmarc-reports" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [] }, "key$": "email_auth", "name__orig": "email_auth", "Name": "EmailAuth", "name_": "email_auth", "name-": "email-auth", "NAME": "EMAIL_AUTH", "index$": 2 }, { "active": true, "entity": "email_auth", "key$": "BasicEmailAuthFlow", "kind": "basic", "name": "BasicEmailAuthFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "email_auth_ref01", "srcdatavar": "email_auth_ref01_data", "suffix": "_up0" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-email_auth_ref01" } }], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "email_auth_ref01", "srcdatavar": "email_auth_ref01_data", "suffix": "_dt0" }, "m": { "id": "email_auth01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-email_auth_ref01" } }], "index$": 1 }] }, 'EmailAuth', { "GET /zones/{zone_id}/email/auth/spf/inspect": { "protocol": "http", "parameters": [{ "description": "Zone identifier.", "in": "path", "name": "zone_id", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email-auth_identifier" }, "x-ref": "#/components/parameters/email-auth_zone_id", "index$": 0 }, { "description": "DNS record ID (rec_tag) to inspect", "example": "e5bb46707a802688812d5d1c9f7977d4", "in": "query", "name": "id", "required": true, "schema": { "type": "string", "pattern": "^[a-zA-Z0-9-]{1,64}$" }, "x-ref": "#/components/parameters/email-auth_record_id", "index$": 1 }] }, "GET /zones/{zone_id}/email/auth/dmarc-reports": { "protocol": "http", "parameters": [{ "description": "Zone identifier.", "in": "path", "name": "zone_id", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email-auth_identifier" }, "x-ref": "#/components/parameters/email-auth_zone_id", "index$": 0 }] }, "PATCH /zones/{zone_id}/email/auth/dmarc-reports": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "description": "Request body for PATCH /dmarc-reports", "type": "object", "properties": { "enabled": { "description": "Enable or disable DMARC reports for this zone", "type": "boolean", "example": true, "nullable": true, "key$": "enabled" }, "skip_wizard": { "description": "Skip the DMARC setup wizard", "type": "boolean", "example": false, "nullable": true, "key$": "skip_wizard" } }, "x-ref": "#/components/schemas/email-auth_ConfigureDmarcReportsRequest", "index$": 1 } } } }, "parameters": [{ "description": "Zone identifier.", "in": "path", "name": "zone_id", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email-auth_identifier" }, "x-ref": "#/components/parameters/email-auth_zone_id", "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let email_auth_ref01_data = Object.values(setup.data.existing.email_auth)[0];
        // UPDATE
        const email_auth_ref01_ent = client.EmailAuth();
        const email_auth_ref01_data_up0 = {};
        const email_auth_ref01_resdata_up0 = (await email_auth_ref01_ent.update(email_auth_ref01_data_up0)).data();
        (0, node_assert_1.default)(null != email_auth_ref01_resdata_up0);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/email_auth/EmailAuthTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CloudflareEmailSecuritySDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['email_auth01', 'email_auth02', 'email_auth03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_AUTH_ENTID': idmap,
        'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
    });
    idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_AUTH_ENTID'];
    const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_AUTH_ENTID'];
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
//# sourceMappingURL=EmailAuthEntity.test.js.map