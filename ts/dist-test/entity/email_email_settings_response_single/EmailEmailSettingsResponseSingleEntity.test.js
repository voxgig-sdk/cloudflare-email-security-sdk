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
(0, node_test_1.describe)('EmailEmailSettingsResponseSingleEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CloudflareEmailSecuritySDK.test();
        const ent = testsdk.EmailEmailSettingsResponseSingle();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
        for (const op of ['create', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'email_email_settings_response_single.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "enabled": { "a": true, "h": "Enabled", "n": "enabled", "r": false, "sh": "State of your zone Email Routing settings.", "t": "`$BOOLEAN`", "key$": "enabled", "index$": 0 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Domain of your zone.", "t": "`$STRING`", "key$": "name", "index$": 1 }, "result": { "a": true, "h": "Result", "n": "result", "r": false, "t": "`$OBJECT`", "key$": "result", "index$": 2 }, "skip_wizard": { "a": true, "h": "Skip Wizard", "n": "skip_wizard", "r": false, "sh": "Flag to check if the user skipped the configuration wizard.", "t": "`$BOOLEAN`", "key$": "skip_wizard", "index$": 3 }, "support_subaddress": { "a": true, "h": "Support Subaddress", "n": "support_subaddress", "r": false, "sh": "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.", "t": "`$BOOLEAN`", "key$": "support_subaddress", "index$": 4 } }, "name": "email_email_settings_response_single", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /zones/{zone_id}/email/routing/disable", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/zones/{zone_id}/email/routing/disable", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "disable" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "POST /zones/{zone_id}/email/routing/dns", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/zones/{zone_id}/email/routing/dns", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "dns" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "POST /zones/{zone_id}/email/routing/enable", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/zones/{zone_id}/email/routing/enable", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "enable" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }, { "a": true, "co": { "id": "POST /zones/{zone_id}/email/routing/unlock", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/zones/{zone_id}/email/routing/unlock", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "unlock" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 3 }], "key$": "create" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /zones/{zone_id}/email/routing", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/zones/{zone_id}/email/routing", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /zones/{zone_id}/email/routing", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/zones/{zone_id}/email/routing", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "PATCH /zones/{zone_id}/email/routing/dns", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/zones/{zone_id}/email/routing/dns", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "dns" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "patch" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /zones/{zone_id}/email/routing/dns", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "DELETE", "o": "/zones/{zone_id}/email/routing/dns", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }, { "lit": "dns" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /zones/{zone_id}/email/routing", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "023e105f4ecef8ad9ca31a8372d0c353", "k": "param", "n": "zone_id", "or": "zone_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PUT", "o": "/zones/{zone_id}/email/routing", "q": { "exist": ["zone_id"] }, "r": {}, "s": [{ "lit": "zones" }, { "var": "zone_id" }, { "lit": "email" }, { "lit": "routing" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [] }, "key$": "email_email_settings_response_single", "name__orig": "email_email_settings_response_single", "Name": "EmailEmailSettingsResponseSingle", "name_": "email_email_settings_response_single", "name-": "email-email-settings-response-single", "NAME": "EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE", "index$": 7 }, { "active": true, "entity": "email_email_settings_response_single", "key$": "BasicEmailEmailSettingsResponseSingleFlow", "kind": "basic", "name": "BasicEmailEmailSettingsResponseSingleFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "email_email_settings_response_single_ref01" }, "m": { "zone_id": "zone01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "email_email_settings_response_single_ref01", "srcdatavar": "email_email_settings_response_single_ref01_data", "suffix": "_up0", "textfield": "name" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-email_email_settings_response_single_ref01" } }], "v": [], "index$": 1 }, { "a": true, "d": {}, "i": { "ref": "email_email_settings_response_single_ref01", "srcdatavar": "email_email_settings_response_single_ref01_data", "suffix": "_dt0" }, "m": { "id": "email_email_settings_response_single01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-email_email_settings_response_single_ref01" } }], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "email_email_settings_response_single_ref01", "suffix": "_rm0" }, "m": { "id": "email_email_settings_response_single01" }, "o": "remove", "s": [], "v": [], "index$": 3 }] }, 'EmailEmailSettingsResponseSingle', { "POST /zones/{zone_id}/email/routing/disable": { "protocol": "http", "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "POST /zones/{zone_id}/email/routing/dns": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "name": { "description": "Domain of your zone.", "type": "string", "example": "example.net", "x-auditable": true, "x-ref": "#/components/schemas/email_email_setting_name", "key$": "name" } }, "nullable": true, "x-ref": "#/components/schemas/email_email_setting_dns_request_body", "index$": 1 } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "POST /zones/{zone_id}/email/routing/enable": { "protocol": "http", "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "POST /zones/{zone_id}/email/routing/unlock": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "name": { "description": "Domain of your zone.", "type": "string", "example": "example.net", "x-auditable": true, "x-ref": "#/components/schemas/email_email_setting_name", "key$": "name" } }, "nullable": true, "x-ref": "#/components/schemas/email_email_setting_dns_request_body", "index$": 1 } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "GET /zones/{zone_id}/email/routing": { "protocol": "http", "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "PATCH /zones/{zone_id}/email/routing": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "properties": { "enabled": { "description": "State of your zone Email Routing settings. No-op on this endpoint - use `POST`/`DELETE /zones/{zone_id}/email/routing/dns`.", "type": "boolean", "example": true, "enum": [true, false], "key$": "enabled" }, "skip_wizard": { "description": "Flag to check if the user skipped the configuration wizard.", "type": "boolean", "example": true, "enum": [true, false], "key$": "skip_wizard" }, "support_subaddress": { "description": "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.", "type": "boolean", "example": true, "enum": [true, false], "key$": "support_subaddress" } }, "x-ref": "#/components/schemas/email_update_email_routing_settings_properties", "index$": 1 } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "PATCH /zones/{zone_id}/email/routing/dns": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "name": { "description": "Domain of your zone.", "type": "string", "example": "example.net", "x-auditable": true, "x-ref": "#/components/schemas/email_email_setting_name", "key$": "name" } }, "nullable": true, "x-ref": "#/components/schemas/email_email_setting_dns_request_body", "index$": 1 } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "DELETE /zones/{zone_id}/email/routing/dns": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "name": { "description": "Domain of your zone.", "type": "string", "example": "example.net", "x-auditable": true, "x-ref": "#/components/schemas/email_email_setting_name", "key$": "name" } }, "nullable": true, "x-ref": "#/components/schemas/email_email_setting_dns_request_body" } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] }, "PUT /zones/{zone_id}/email/routing": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "properties": { "enabled": { "description": "State of your zone Email Routing settings. No-op on this endpoint - use `POST`/`DELETE /zones/{zone_id}/email/routing/dns`.", "type": "boolean", "example": true, "enum": [true, false], "key$": "enabled" }, "skip_wizard": { "description": "Flag to check if the user skipped the configuration wizard.", "type": "boolean", "example": true, "enum": [true, false], "key$": "skip_wizard" }, "support_subaddress": { "description": "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.", "type": "boolean", "example": true, "enum": [true, false], "key$": "support_subaddress" } }, "x-ref": "#/components/schemas/email_update_email_routing_settings_properties", "index$": 1 } } } }, "parameters": [{ "name": "zone_id", "in": "path", "required": true, "schema": { "description": "Identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "x-auditable": true, "x-ref": "#/components/schemas/email_identifier" }, "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const email_email_settings_response_single_ref01_ent = client.EmailEmailSettingsResponseSingle();
        let email_email_settings_response_single_ref01_data = setup.data.new.email_email_settings_response_single['email_email_settings_response_single_ref01'];
        email_email_settings_response_single_ref01_data['zone_id'] = setup.idmap['zone01'];
        email_email_settings_response_single_ref01_data = (await email_email_settings_response_single_ref01_ent.create(email_email_settings_response_single_ref01_data)).data();
        (0, node_assert_1.default)(null != email_email_settings_response_single_ref01_data);
        // UPDATE
        const email_email_settings_response_single_ref01_data_up0 = {};
        const email_email_settings_response_single_ref01_markdef_up0 = { name: 'name', value: 'Mark01-email_email_settings_response_single_ref01_' + setup.now };
        email_email_settings_response_single_ref01_data_up0[email_email_settings_response_single_ref01_markdef_up0.name] = email_email_settings_response_single_ref01_markdef_up0.value;
        const email_email_settings_response_single_ref01_resdata_up0 = (await email_email_settings_response_single_ref01_ent.update(email_email_settings_response_single_ref01_data_up0)).data();
        (0, node_assert_1.default)(null != email_email_settings_response_single_ref01_resdata_up0);
        (0, node_assert_1.default)(email_email_settings_response_single_ref01_resdata_up0[email_email_settings_response_single_ref01_markdef_up0.name] === email_email_settings_response_single_ref01_markdef_up0.value);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/email_email_settings_response_single/EmailEmailSettingsResponseSingleTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CloudflareEmailSecuritySDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['email_email_settings_response_single01', 'email_email_settings_response_single02', 'email_email_settings_response_single03', 'zone01'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID': idmap,
        'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
        'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
    });
    idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID'];
    const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_EMAIL_SETTINGS_RESPONSE_SINGLE_ENTID'];
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
//# sourceMappingURL=EmailEmailSettingsResponseSingleEntity.test.js.map