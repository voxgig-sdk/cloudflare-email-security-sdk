"use strict";
// CloudflareEmailSecurity Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.CloudflareEmailSecuritySDK = exports.CloudflareEmailSecurityEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const EmailAccountRulesPlanResponseSingleEntity_1 = require("./entity/EmailAccountRulesPlanResponseSingleEntity");
const EmailAccountRulesResponseCollectionEntity_1 = require("./entity/EmailAccountRulesResponseCollectionEntity");
const EmailAuthEntity_1 = require("./entity/EmailAuthEntity");
const EmailCatchAllRuleResponseSingleEntity_1 = require("./entity/EmailCatchAllRuleResponseSingleEntity");
const EmailDestinationAddressResponseSingleEntity_1 = require("./entity/EmailDestinationAddressResponseSingleEntity");
const EmailDestinationAddressesResponseCollectionEntity_1 = require("./entity/EmailDestinationAddressesResponseCollectionEntity");
const EmailDnsSettingsResponseCollectionEntity_1 = require("./entity/EmailDnsSettingsResponseCollectionEntity");
const EmailEmailSettingsResponseSingleEntity_1 = require("./entity/EmailEmailSettingsResponseSingleEntity");
const EmailRoutingSuppressionEntity_1 = require("./entity/EmailRoutingSuppressionEntity");
const EmailRuleResponseSingleEntity_1 = require("./entity/EmailRuleResponseSingleEntity");
const EmailRulesResponseCollectionEntity_1 = require("./entity/EmailRulesResponseCollectionEntity");
const EmailSecurityEntity_1 = require("./entity/EmailSecurityEntity");
const EmailSecurityAnalyticsEntity_1 = require("./entity/EmailSecurityAnalyticsEntity");
const EmailSecuritySettingEntity_1 = require("./entity/EmailSecuritySettingEntity");
const EmailSendingEntity_1 = require("./entity/EmailSendingEntity");
const EmailSendingLimitsResponseSingleEntity_1 = require("./entity/EmailSendingLimitsResponseSingleEntity");
const EmailSendingReputationResponseSingleEntity_1 = require("./entity/EmailSendingReputationResponseSingleEntity");
const EmailSendingSubdomainDnsStatusEntity_1 = require("./entity/EmailSendingSubdomainDnsStatusEntity");
const EmailSendingSubdomainPreviewEntity_1 = require("./entity/EmailSendingSubdomainPreviewEntity");
const EmailSendingSubdomainReputationComplaintEntity_1 = require("./entity/EmailSendingSubdomainReputationComplaintEntity");
const EmailSendingSubdomainResponseSingleEntity_1 = require("./entity/EmailSendingSubdomainResponseSingleEntity");
const EmailSendingSubdomainsResponseCollectionEntity_1 = require("./entity/EmailSendingSubdomainsResponseCollectionEntity");
const EmailSendingSuppressionEntity_1 = require("./entity/EmailSendingSuppressionEntity");
const SubdomainEntity_1 = require("./entity/SubdomainEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const CloudflareEmailSecurityEntityBase_1 = require("./CloudflareEmailSecurityEntityBase");
Object.defineProperty(exports, "CloudflareEmailSecurityEntityBase", { enumerable: true, get: function () { return CloudflareEmailSecurityEntityBase_1.CloudflareEmailSecurityEntityBase; } });
const Utility_1 = require("./utility/Utility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class CloudflareEmailSecuritySDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const extend = this._options.extend || [];
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                // An active name with no generated class is legal when an
                // extend-supplied instance carries that name (station's adopt
                // path): the instance is added below, positioned by its own
                // __after__ entry, so skip it here rather than fail construction.
                if (!this._rootctx.config.hasFeature(fname) &&
                    extend.some((f) => fname === f.name)) {
                    continue;
                }
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        for (let f of extend) {
            featureAdd(this._rootctx, f);
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
        const spec = {
            base: options.base,
            prefix: options.prefix,
            suffix: options.suffix,
            path: fetchargs.path || '',
            method: fetchargs.method || 'GET',
            params: fetchargs.params || {},
            query: fetchargs.query || {},
            headers: prepareHeaders(ctx),
            body: fetchargs.body,
            step: 'start',
        };
        ctx.spec = spec;
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
    }
    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    // either one reaches the same endpoint.
    async direct(fetchargs) {
        if (!this._options.allow.op.includes('direct')) {
            return {
                ok: false,
                err: new Error('CloudflareEmailSecuritySDK: direct: operation not allowed by' +
                    ' SDK option allow.op value: "' + this._options.allow.op + '"'),
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return fetchdef;
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: fetched };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            if (!noBody) {
                try {
                    json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                }
                catch (parseErr) {
                    // Body wasn't valid JSON — surface the raw response rather than
                    // throwing. data stays undefined; callers can inspect status/headers.
                    json = undefined;
                }
            }
            return {
                ok: status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
            };
        }
        catch (err) {
            return { ok: false, err };
        }
    }
    async graphql(query, variables, ctrl) {
        const options = this._options;
        if (!options.allow.op.includes('graphql')) {
            return {
                ok: false,
                err: new Error('CloudflareEmailSecuritySDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        if (res instanceof Error) {
            return res;
        }
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('CloudflareEmailSecuritySDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.EmailAccountRulesPlanResponseSingle().list()` / `client.EmailAccountRulesPlanResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailAccountRulesPlanResponseSingle(entopts) {
        const self = this;
        return new EmailAccountRulesPlanResponseSingleEntity_1.EmailAccountRulesPlanResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailAccountRulesResponseCollection().list()` / `client.EmailAccountRulesResponseCollection().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailAccountRulesResponseCollection(entopts) {
        const self = this;
        return new EmailAccountRulesResponseCollectionEntity_1.EmailAccountRulesResponseCollectionEntity(self, entopts);
    }
    // Entity access: `client.EmailAuth().list()` / `client.EmailAuth().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailAuth(entopts) {
        const self = this;
        return new EmailAuthEntity_1.EmailAuthEntity(self, entopts);
    }
    // Entity access: `client.EmailCatchAllRuleResponseSingle().list()` / `client.EmailCatchAllRuleResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailCatchAllRuleResponseSingle(entopts) {
        const self = this;
        return new EmailCatchAllRuleResponseSingleEntity_1.EmailCatchAllRuleResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailDestinationAddressResponseSingle().list()` / `client.EmailDestinationAddressResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailDestinationAddressResponseSingle(entopts) {
        const self = this;
        return new EmailDestinationAddressResponseSingleEntity_1.EmailDestinationAddressResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailDestinationAddressesResponseCollection().list()` / `client.EmailDestinationAddressesResponseCollection().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailDestinationAddressesResponseCollection(entopts) {
        const self = this;
        return new EmailDestinationAddressesResponseCollectionEntity_1.EmailDestinationAddressesResponseCollectionEntity(self, entopts);
    }
    // Entity access: `client.EmailDnsSettingsResponseCollection().list()` / `client.EmailDnsSettingsResponseCollection().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailDnsSettingsResponseCollection(entopts) {
        const self = this;
        return new EmailDnsSettingsResponseCollectionEntity_1.EmailDnsSettingsResponseCollectionEntity(self, entopts);
    }
    // Entity access: `client.EmailEmailSettingsResponseSingle().list()` / `client.EmailEmailSettingsResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailEmailSettingsResponseSingle(entopts) {
        const self = this;
        return new EmailEmailSettingsResponseSingleEntity_1.EmailEmailSettingsResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailRoutingSuppression().list()` / `client.EmailRoutingSuppression().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailRoutingSuppression(entopts) {
        const self = this;
        return new EmailRoutingSuppressionEntity_1.EmailRoutingSuppressionEntity(self, entopts);
    }
    // Entity access: `client.EmailRuleResponseSingle().list()` / `client.EmailRuleResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailRuleResponseSingle(entopts) {
        const self = this;
        return new EmailRuleResponseSingleEntity_1.EmailRuleResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailRulesResponseCollection().list()` / `client.EmailRulesResponseCollection().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailRulesResponseCollection(entopts) {
        const self = this;
        return new EmailRulesResponseCollectionEntity_1.EmailRulesResponseCollectionEntity(self, entopts);
    }
    // Entity access: `client.EmailSecurity().list()` / `client.EmailSecurity().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSecurity(entopts) {
        const self = this;
        return new EmailSecurityEntity_1.EmailSecurityEntity(self, entopts);
    }
    // Entity access: `client.EmailSecurityAnalytics().list()` / `client.EmailSecurityAnalytics().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSecurityAnalytics(entopts) {
        const self = this;
        return new EmailSecurityAnalyticsEntity_1.EmailSecurityAnalyticsEntity(self, entopts);
    }
    // Entity access: `client.EmailSecuritySetting().list()` / `client.EmailSecuritySetting().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSecuritySetting(entopts) {
        const self = this;
        return new EmailSecuritySettingEntity_1.EmailSecuritySettingEntity(self, entopts);
    }
    // Entity access: `client.EmailSending().list()` / `client.EmailSending().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSending(entopts) {
        const self = this;
        return new EmailSendingEntity_1.EmailSendingEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingLimitsResponseSingle().list()` / `client.EmailSendingLimitsResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingLimitsResponseSingle(entopts) {
        const self = this;
        return new EmailSendingLimitsResponseSingleEntity_1.EmailSendingLimitsResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingReputationResponseSingle().list()` / `client.EmailSendingReputationResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingReputationResponseSingle(entopts) {
        const self = this;
        return new EmailSendingReputationResponseSingleEntity_1.EmailSendingReputationResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSubdomainDnsStatus().list()` / `client.EmailSendingSubdomainDnsStatus().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSubdomainDnsStatus(entopts) {
        const self = this;
        return new EmailSendingSubdomainDnsStatusEntity_1.EmailSendingSubdomainDnsStatusEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSubdomainPreview().list()` / `client.EmailSendingSubdomainPreview().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSubdomainPreview(entopts) {
        const self = this;
        return new EmailSendingSubdomainPreviewEntity_1.EmailSendingSubdomainPreviewEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSubdomainReputationComplaint().list()` / `client.EmailSendingSubdomainReputationComplaint().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSubdomainReputationComplaint(entopts) {
        const self = this;
        return new EmailSendingSubdomainReputationComplaintEntity_1.EmailSendingSubdomainReputationComplaintEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSubdomainResponseSingle().list()` / `client.EmailSendingSubdomainResponseSingle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSubdomainResponseSingle(entopts) {
        const self = this;
        return new EmailSendingSubdomainResponseSingleEntity_1.EmailSendingSubdomainResponseSingleEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSubdomainsResponseCollection().list()` / `client.EmailSendingSubdomainsResponseCollection().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSubdomainsResponseCollection(entopts) {
        const self = this;
        return new EmailSendingSubdomainsResponseCollectionEntity_1.EmailSendingSubdomainsResponseCollectionEntity(self, entopts);
    }
    // Entity access: `client.EmailSendingSuppression().list()` / `client.EmailSendingSuppression().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailSendingSuppression(entopts) {
        const self = this;
        return new EmailSendingSuppressionEntity_1.EmailSendingSuppressionEntity(self, entopts);
    }
    // Entity access: `client.Subdomain().list()` / `client.Subdomain().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Subdomain(entopts) {
        const self = this;
        return new SubdomainEntity_1.SubdomainEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new CloudflareEmailSecuritySDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return CloudflareEmailSecuritySDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'CloudflareEmailSecurity' };
    }
    toString() {
        return 'CloudflareEmailSecurity ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.CloudflareEmailSecuritySDK = CloudflareEmailSecuritySDK;
const SDK = CloudflareEmailSecuritySDK;
exports.SDK = SDK;
//# sourceMappingURL=CloudflareEmailSecuritySDK.js.map