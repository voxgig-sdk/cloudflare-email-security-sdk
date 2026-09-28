

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { CloudflareEmailSecuritySDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('EmailRuleResponseSingleEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailRuleResponseSingle()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['create', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_rule_response_single.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"actions":{"a":true,"h":"Actions","n":"actions","r":true,"sh":"List actions patterns.","t":"`$ARRAY`","key$":"actions","index$":0},"enabled":{"a":true,"h":"Enabled","n":"enabled","r":false,"sh":"Routing rule status.","t":"`$BOOLEAN`","key$":"enabled","index$":1},"matchers":{"a":true,"h":"Matchers","n":"matchers","r":true,"sh":"Matching patterns to forward to your actions.","t":"`$ARRAY`","key$":"matchers","index$":2},"name":{"a":true,"h":"Name","n":"name","r":false,"sh":"Routing rule name.","t":"`$STRING`","key$":"name","index$":3},"owner_worker_tag":{"a":true,"h":"Owner Worker Tag","n":"owner_worker_tag","r":false,"sh":"Public tag (script_tag) of the Worker that owns this rule.","t":"`$STRING`","wo":true,"key$":"owner_worker_tag","index$":4},"priority":{"a":true,"h":"Priority","n":"priority","r":false,"sh":"Priority of the routing rule.","t":"`$NUMBER`","key$":"priority","index$":5},"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$OBJECT`","key$":"result","index$":6},"source":{"a":true,"h":"Source","n":"source","r":false,"sh":"Who manages the rule.","t":"`$STRING`","key$":"source","index$":7}},"name":"email_rule_response_single","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /zones/{zone_id}/email/routing/rules","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/zones/{zone_id}/email/routing/rules","q":{"exist":["zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"routing"},{"lit":"rules"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /zones/{zone_id}/email/routing/rules/{rule_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"a7e6fb77503c41d8a7f3113c6918f10c","k":"param","n":"rule_identifier","or":"rule_identifier","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/routing/rules/{rule_identifier}","q":{"exist":["rule_identifier","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"routing"},{"lit":"rules"},{"var":"rule_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /zones/{zone_id}/email/routing/rules/{rule_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"a7e6fb77503c41d8a7f3113c6918f10c","k":"param","n":"rule_identifier","or":"rule_identifier","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/zones/{zone_id}/email/routing/rules/{rule_identifier}","q":{"exist":["rule_identifier","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"routing"},{"lit":"rules"},{"var":"rule_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PUT /zones/{zone_id}/email/routing/rules/{rule_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"a7e6fb77503c41d8a7f3113c6918f10c","k":"param","n":"rule_identifier","or":"rule_identifier","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PUT","o":"/zones/{zone_id}/email/routing/rules/{rule_identifier}","q":{"exist":["rule_identifier","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"routing"},{"lit":"rules"},{"var":"rule_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"email_rule_response_single","name__orig":"email_rule_response_single","Name":"EmailRuleResponseSingle","name_":"email_rule_response_single","name-":"email-rule-response-single","NAME":"EMAIL_RULE_RESPONSE_SINGLE","index$":9}, {"active":true,"entity":"email_rule_response_single","key$":"BasicEmailRuleResponseSingleFlow","kind":"basic","name":"BasicEmailRuleResponseSingleFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"email_rule_response_single_ref01"},"m":{"zone_id":"zone01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{"zone_id":"zone01"},"i":{"ref":"email_rule_response_single_ref01","srcdatavar":"email_rule_response_single_ref01_data","suffix":"_up0","textfield":"name"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_rule_response_single_ref01"}}],"v":[],"index$":1},{"a":true,"d":{},"i":{"ref":"email_rule_response_single_ref01","srcdatavar":"email_rule_response_single_ref01_data","suffix":"_dt0"},"m":{"id":"email_rule_response_single01","zone_id":"zone01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_rule_response_single_ref01"}}],"index$":2},{"a":true,"d":{},"i":{"ref":"email_rule_response_single_ref01","suffix":"_rm0"},"m":{"id":"email_rule_response_single01","zone_id":"zone01"},"o":"remove","s":[],"v":[],"index$":3}]}, 'EmailRuleResponseSingle', {"POST /zones/{zone_id}/email/routing/rules":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"actions":{"description":"List actions patterns.","type":"array","items":{"description":"Actions pattern.","type":"object","properties":{"type":{"description":"Type of supported action.","enum":[],"example":"forward","type":"string","x-auditable":true},"value":{"description":"List of values for the action. Currently limited to a single value.","items":{},"maxItems":1,"type":"array"}},"required":["type"],"x-ref":"#/components/schemas/email_rule_action"},"maxItems":1,"minItems":1,"x-ref":"#/components/schemas/email_rule_actions","key$":"actions"},"enabled":{"description":"Routing rule status.","type":"boolean","example":true,"default":true,"enum":[true,false],"x-auditable":true,"x-ref":"#/components/schemas/email_rule_enabled","key$":"enabled"},"matchers":{"description":"Matching patterns to forward to your actions.","type":"array","items":{"description":"Matching pattern to forward your actions.","type":"object","properties":{"field":{"description":"Field for type matcher.","enum":[],"example":"to","type":"string","x-auditable":true},"type":{"description":"Type of matcher.","enum":[],"example":"literal","type":"string","x-auditable":true},"value":{"description":"Value for matcher.","example":"test@example.com","maxLength":90,"type":"string","x-auditable":true}},"required":["type"],"x-ref":"#/components/schemas/email_rule_matcher"},"x-ref":"#/components/schemas/email_rule_matchers","key$":"matchers"},"name":{"description":"Routing rule name.","type":"string","example":"Send to user@example.net rule.","maxLength":256,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_name","key$":"name"},"owner_worker_tag":{"description":"Public tag (script_tag) of the Worker that owns this rule. Required when\n`source` is `wrangler`.\n","type":"string","example":"a7e6fb77503c41d8a7f3113c6918f10c","maxLength":32,"writeOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_owner_worker_tag","key$":"owner_worker_tag"},"priority":{"description":"Priority of the routing rule.","type":"number","default":0,"minimum":0,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_priority","key$":"priority"},"source":{"description":"Who manages the rule. `api` covers dashboard, generic API, and Terraform;\n`wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults\nto `api` when omitted on write.\n","type":"string","example":"api","default":"api","enum":["api","wrangler"],"x-auditable":true,"x-ref":"#/components/schemas/email_rule_source","key$":"source"}},"required":["actions","matchers"],"x-ref":"#/components/schemas/email_create_rule_properties","index$":1}}}},"parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":0}]},"GET /zones/{zone_id}/email/routing/rules/{rule_identifier}":{"protocol":"http","parameters":[{"name":"rule_identifier","in":"path","required":true,"schema":{"description":"Routing rule identifier.","type":"string","example":"a7e6fb77503c41d8a7f3113c6918f10c","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_identifier"},"index$":0},{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]},"DELETE /zones/{zone_id}/email/routing/rules/{rule_identifier}":{"protocol":"http","parameters":[{"name":"rule_identifier","in":"path","required":true,"schema":{"description":"Routing rule identifier.","type":"string","example":"a7e6fb77503c41d8a7f3113c6918f10c","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_identifier"},"index$":0},{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]},"PUT /zones/{zone_id}/email/routing/rules/{rule_identifier}":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"actions":{"description":"List actions patterns.","type":"array","items":{"description":"Actions pattern.","type":"object","properties":{"type":{"description":"Type of supported action.","enum":[],"example":"forward","type":"string","x-auditable":true},"value":{"description":"List of values for the action. Currently limited to a single value.","items":{},"maxItems":1,"type":"array"}},"required":["type"],"x-ref":"#/components/schemas/email_rule_action"},"maxItems":1,"minItems":1,"x-ref":"#/components/schemas/email_rule_actions","key$":"actions"},"enabled":{"description":"Routing rule status.","type":"boolean","example":true,"default":true,"enum":[true,false],"x-auditable":true,"x-ref":"#/components/schemas/email_rule_enabled","key$":"enabled"},"matchers":{"description":"Matching patterns to forward to your actions.","type":"array","items":{"description":"Matching pattern to forward your actions.","type":"object","properties":{"field":{"description":"Field for type matcher.","enum":[],"example":"to","type":"string","x-auditable":true},"type":{"description":"Type of matcher.","enum":[],"example":"literal","type":"string","x-auditable":true},"value":{"description":"Value for matcher.","example":"test@example.com","maxLength":90,"type":"string","x-auditable":true}},"required":["type"],"x-ref":"#/components/schemas/email_rule_matcher"},"x-ref":"#/components/schemas/email_rule_matchers","key$":"matchers"},"name":{"description":"Routing rule name.","type":"string","example":"Send to user@example.net rule.","maxLength":256,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_name","key$":"name"},"owner_worker_tag":{"description":"Public tag (script_tag) of the Worker that owns this rule. Required when\n`source` is `wrangler`.\n","type":"string","example":"a7e6fb77503c41d8a7f3113c6918f10c","maxLength":32,"writeOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_owner_worker_tag","key$":"owner_worker_tag"},"priority":{"description":"Priority of the routing rule.","type":"number","default":0,"minimum":0,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_priority","key$":"priority"},"source":{"description":"Who manages the rule. `api` covers dashboard, generic API, and Terraform;\n`wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults\nto `api` when omitted on write.\n","type":"string","example":"api","default":"api","enum":["api","wrangler"],"x-auditable":true,"x-ref":"#/components/schemas/email_rule_source","key$":"source"}},"required":["actions","matchers"],"x-ref":"#/components/schemas/email_update_rule_properties","index$":1}}}},"parameters":[{"name":"rule_identifier","in":"path","required":true,"schema":{"description":"Routing rule identifier.","type":"string","example":"a7e6fb77503c41d8a7f3113c6918f10c","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_rule_identifier"},"index$":0},{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const email_rule_response_single_ref01_ent = client.EmailRuleResponseSingle()
    let email_rule_response_single_ref01_data = setup.data.new.email_rule_response_single['email_rule_response_single_ref01']
    email_rule_response_single_ref01_data['zone_id'] = setup.idmap['zone01']

    email_rule_response_single_ref01_data = (await email_rule_response_single_ref01_ent.create(email_rule_response_single_ref01_data)).data()
    assert(null != email_rule_response_single_ref01_data)


    // UPDATE
    const email_rule_response_single_ref01_data_up0: any = {}
    email_rule_response_single_ref01_data_up0 ['zone_id'] = setup.idmap['zone_id']

    const email_rule_response_single_ref01_markdef_up0 = { name: 'name', value: 'Mark01-email_rule_response_single_ref01_' + setup.now }
    ;(email_rule_response_single_ref01_data_up0 as any)[email_rule_response_single_ref01_markdef_up0.name] = email_rule_response_single_ref01_markdef_up0.value

    const email_rule_response_single_ref01_resdata_up0 = (await email_rule_response_single_ref01_ent.update(email_rule_response_single_ref01_data_up0)).data()
    assert(null != email_rule_response_single_ref01_resdata_up0)

    assert((email_rule_response_single_ref01_resdata_up0 as any)[email_rule_response_single_ref01_markdef_up0.name] === email_rule_response_single_ref01_markdef_up0.value)




  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_rule_response_single/EmailRuleResponseSingleTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = CloudflareEmailSecuritySDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['email_rule_response_single01','email_rule_response_single02','email_rule_response_single03','zone01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_RULE_RESPONSE_SINGLE_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_RULE_RESPONSE_SINGLE_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_RULE_RESPONSE_SINGLE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new CloudflareEmailSecuritySDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
