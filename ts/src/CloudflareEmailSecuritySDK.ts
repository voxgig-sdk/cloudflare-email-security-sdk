// CloudflareEmailSecurity Ts SDK

import { EmailAccountRulesPlanResponseSingleEntity } from './entity/EmailAccountRulesPlanResponseSingleEntity'
import { EmailAccountRulesResponseCollectionEntity } from './entity/EmailAccountRulesResponseCollectionEntity'
import { EmailAuthEntity } from './entity/EmailAuthEntity'
import { EmailCatchAllRuleResponseSingleEntity } from './entity/EmailCatchAllRuleResponseSingleEntity'
import { EmailDestinationAddressResponseSingleEntity } from './entity/EmailDestinationAddressResponseSingleEntity'
import { EmailDestinationAddressesResponseCollectionEntity } from './entity/EmailDestinationAddressesResponseCollectionEntity'
import { EmailDnsSettingsResponseCollectionEntity } from './entity/EmailDnsSettingsResponseCollectionEntity'
import { EmailEmailSettingsResponseSingleEntity } from './entity/EmailEmailSettingsResponseSingleEntity'
import { EmailRoutingSuppressionEntity } from './entity/EmailRoutingSuppressionEntity'
import { EmailRuleResponseSingleEntity } from './entity/EmailRuleResponseSingleEntity'
import { EmailRulesResponseCollectionEntity } from './entity/EmailRulesResponseCollectionEntity'
import { EmailSecurityEntity } from './entity/EmailSecurityEntity'
import { EmailSecurityAnalyticsEntity } from './entity/EmailSecurityAnalyticsEntity'
import { EmailSecuritySettingEntity } from './entity/EmailSecuritySettingEntity'
import { EmailSendingEntity } from './entity/EmailSendingEntity'
import { EmailSendingLimitsResponseSingleEntity } from './entity/EmailSendingLimitsResponseSingleEntity'
import { EmailSendingReputationResponseSingleEntity } from './entity/EmailSendingReputationResponseSingleEntity'
import { EmailSendingSubdomainDnsStatusEntity } from './entity/EmailSendingSubdomainDnsStatusEntity'
import { EmailSendingSubdomainPreviewEntity } from './entity/EmailSendingSubdomainPreviewEntity'
import { EmailSendingSubdomainReputationComplaintEntity } from './entity/EmailSendingSubdomainReputationComplaintEntity'
import { EmailSendingSubdomainResponseSingleEntity } from './entity/EmailSendingSubdomainResponseSingleEntity'
import { EmailSendingSubdomainsResponseCollectionEntity } from './entity/EmailSendingSubdomainsResponseCollectionEntity'
import { EmailSendingSuppressionEntity } from './entity/EmailSendingSuppressionEntity'
import { SubdomainEntity } from './entity/SubdomainEntity'

export type * from './CloudflareEmailSecurityTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { CloudflareEmailSecurityEntityBase } from './CloudflareEmailSecurityEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'



const stdutil = new Utility()


class CloudflareEmailSecuritySDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context
  

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const extend = this._options.extend || []

    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        // An active name with no generated class is legal when an
        // extend-supplied instance carries that name (station's adopt
        // path): the instance is added below, positioned by its own
        // __after__ entry, so skip it here rather than fail construction.
        if (!this._rootctx.config.hasFeature(fname) &&
          extend.some((f: any) => fname === f.name)) {
          continue
        }
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    for (let f of extend) {
      featureAdd(this._rootctx, f)
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }

  


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    const spec: any = {
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
    }

    ctx.spec = spec

    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    

    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs?: any) {
    if (!this._options.allow.op.includes('direct')) {
      return {
        ok: false,
        err: new Error('CloudflareEmailSecuritySDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs?: any) {
    const utility = this._utility

    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  async graphql(query: string, variables?: any, ctrl?: any) {
    const options = this._options

    if (!options.allow.op.includes('graphql')) {
      return {
        ok: false,
        err: new Error('CloudflareEmailSecuritySDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res: any = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    if (res instanceof Error) {
      return res
    }

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err: any = new Error('CloudflareEmailSecuritySDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.EmailAccountRulesPlanResponseSingle().list()` / `client.EmailAccountRulesPlanResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailAccountRulesPlanResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailAccountRulesPlanResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailAccountRulesResponseCollection().list()` / `client.EmailAccountRulesResponseCollection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailAccountRulesResponseCollection(entopts?: Record<string, any>) {
    const self = this
    return new EmailAccountRulesResponseCollectionEntity(self, entopts)
  }


  // Entity access: `client.EmailAuth().list()` / `client.EmailAuth().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailAuth(entopts?: Record<string, any>) {
    const self = this
    return new EmailAuthEntity(self, entopts)
  }


  // Entity access: `client.EmailCatchAllRuleResponseSingle().list()` / `client.EmailCatchAllRuleResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailCatchAllRuleResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailCatchAllRuleResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailDestinationAddressResponseSingle().list()` / `client.EmailDestinationAddressResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailDestinationAddressResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailDestinationAddressResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailDestinationAddressesResponseCollection().list()` / `client.EmailDestinationAddressesResponseCollection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailDestinationAddressesResponseCollection(entopts?: Record<string, any>) {
    const self = this
    return new EmailDestinationAddressesResponseCollectionEntity(self, entopts)
  }


  // Entity access: `client.EmailDnsSettingsResponseCollection().list()` / `client.EmailDnsSettingsResponseCollection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailDnsSettingsResponseCollection(entopts?: Record<string, any>) {
    const self = this
    return new EmailDnsSettingsResponseCollectionEntity(self, entopts)
  }


  // Entity access: `client.EmailEmailSettingsResponseSingle().list()` / `client.EmailEmailSettingsResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailEmailSettingsResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailEmailSettingsResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailRoutingSuppression().list()` / `client.EmailRoutingSuppression().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailRoutingSuppression(entopts?: Record<string, any>) {
    const self = this
    return new EmailRoutingSuppressionEntity(self, entopts)
  }


  // Entity access: `client.EmailRuleResponseSingle().list()` / `client.EmailRuleResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailRuleResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailRuleResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailRulesResponseCollection().list()` / `client.EmailRulesResponseCollection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailRulesResponseCollection(entopts?: Record<string, any>) {
    const self = this
    return new EmailRulesResponseCollectionEntity(self, entopts)
  }


  // Entity access: `client.EmailSecurity().list()` / `client.EmailSecurity().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSecurity(entopts?: Record<string, any>) {
    const self = this
    return new EmailSecurityEntity(self, entopts)
  }


  // Entity access: `client.EmailSecurityAnalytics().list()` / `client.EmailSecurityAnalytics().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSecurityAnalytics(entopts?: Record<string, any>) {
    const self = this
    return new EmailSecurityAnalyticsEntity(self, entopts)
  }


  // Entity access: `client.EmailSecuritySetting().list()` / `client.EmailSecuritySetting().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSecuritySetting(entopts?: Record<string, any>) {
    const self = this
    return new EmailSecuritySettingEntity(self, entopts)
  }


  // Entity access: `client.EmailSending().list()` / `client.EmailSending().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSending(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingLimitsResponseSingle().list()` / `client.EmailSendingLimitsResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingLimitsResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingLimitsResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingReputationResponseSingle().list()` / `client.EmailSendingReputationResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingReputationResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingReputationResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSubdomainDnsStatus().list()` / `client.EmailSendingSubdomainDnsStatus().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSubdomainDnsStatus(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSubdomainDnsStatusEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSubdomainPreview().list()` / `client.EmailSendingSubdomainPreview().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSubdomainPreview(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSubdomainPreviewEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSubdomainReputationComplaint().list()` / `client.EmailSendingSubdomainReputationComplaint().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSubdomainReputationComplaint(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSubdomainReputationComplaintEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSubdomainResponseSingle().list()` / `client.EmailSendingSubdomainResponseSingle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSubdomainResponseSingle(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSubdomainResponseSingleEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSubdomainsResponseCollection().list()` / `client.EmailSendingSubdomainsResponseCollection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSubdomainsResponseCollection(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSubdomainsResponseCollectionEntity(self, entopts)
  }


  // Entity access: `client.EmailSendingSuppression().list()` / `client.EmailSendingSuppression().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailSendingSuppression(entopts?: Record<string, any>) {
    const self = this
    return new EmailSendingSuppressionEntity(self, entopts)
  }


  // Entity access: `client.Subdomain().list()` / `client.Subdomain().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Subdomain(entopts?: Record<string, any>) {
    const self = this
    return new SubdomainEntity(self, entopts)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new CloudflareEmailSecuritySDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return CloudflareEmailSecuritySDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'CloudflareEmailSecurity' }
  }

  toString() {
    return 'CloudflareEmailSecurity ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = CloudflareEmailSecuritySDK


export {
  stdutil,
  config,
  

  BaseFeature,
  CloudflareEmailSecurityEntityBase,

  CloudflareEmailSecuritySDK,
  SDK,
}


