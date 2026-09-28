

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


describe('EmailDestinationAddressResponseSingleEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailDestinationAddressResponseSingle()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['create', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_destination_address_response_single.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"email":{"a":true,"h":"Email","n":"email","r":true,"sh":"The contact email address of the user.","t":"`$STRING`","key$":"email","index$":0},"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$OBJECT`","key$":"result","index$":1},"status":{"a":true,"h":"Status","n":"status","r":true,"sh":"Destination address status.","t":"`$STRING`","key$":"status","index$":2}},"name":"email_destination_address_response_single","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /accounts/{account_id}/email/routing/addresses","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/routing/addresses","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"routing"},{"lit":"addresses"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"ea95132c15732412d22c1476fa83f27a","k":"param","n":"destination_address_identifier","or":"destination_address_identifier","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","q":{"exist":["account_id","destination_address_identifier"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"routing"},{"lit":"addresses"},{"var":"destination_address_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"ea95132c15732412d22c1476fa83f27a","k":"param","n":"destination_address_identifier","or":"destination_address_identifier","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","q":{"exist":["account_id","destination_address_identifier"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"routing"},{"lit":"addresses"},{"var":"destination_address_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PATCH /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"ea95132c15732412d22c1476fa83f27a","k":"param","n":"destination_address_identifier","or":"destination_address_identifier","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PATCH","o":"/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}","q":{"exist":["account_id","destination_address_identifier"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"routing"},{"lit":"addresses"},{"var":"destination_address_identifier"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"email_destination_address_response_single","name__orig":"email_destination_address_response_single","Name":"EmailDestinationAddressResponseSingle","name_":"email_destination_address_response_single","name-":"email-destination-address-response-single","NAME":"EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE","index$":4}, {"active":true,"entity":"email_destination_address_response_single","key$":"BasicEmailDestinationAddressResponseSingleFlow","kind":"basic","name":"BasicEmailDestinationAddressResponseSingleFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"email_destination_address_response_single_ref01"},"m":{"account_id":"account01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{"account_id":"account01"},"i":{"ref":"email_destination_address_response_single_ref01","srcdatavar":"email_destination_address_response_single_ref01_data","suffix":"_up0","textfield":"email"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_destination_address_response_single_ref01"}}],"v":[],"index$":1},{"a":true,"d":{},"i":{"ref":"email_destination_address_response_single_ref01","srcdatavar":"email_destination_address_response_single_ref01_data","suffix":"_dt0"},"m":{"account_id":"account01","id":"email_destination_address_response_single01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_destination_address_response_single_ref01"}}],"index$":2},{"a":true,"d":{},"i":{"ref":"email_destination_address_response_single_ref01","suffix":"_rm0"},"m":{"account_id":"account01","id":"email_destination_address_response_single01"},"o":"remove","s":[],"v":[],"index$":3}]}, 'EmailDestinationAddressResponseSingle', {"POST /accounts/{account_id}/email/routing/addresses":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"email":{"description":"The contact email address of the user.","type":"string","example":"user@example.com","maxLength":90,"x-auditable":true,"x-ref":"#/components/schemas/email_email","key$":"email"}},"required":["email"],"x-ref":"#/components/schemas/email_create_destination_address_properties","index$":1}}}},"parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":0}]},"GET /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}":{"protocol":"http","parameters":[{"name":"destination_address_identifier","in":"path","required":true,"schema":{"description":"Destination address identifier.","type":"string","example":"ea95132c15732412d22c1476fa83f27a","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_destination_address_identifier"},"index$":0},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]},"DELETE /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}":{"protocol":"http","parameters":[{"name":"destination_address_identifier","in":"path","required":true,"schema":{"description":"Destination address identifier.","type":"string","example":"ea95132c15732412d22c1476fa83f27a","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_destination_address_identifier"},"index$":0},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]},"PATCH /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"status":{"description":"Destination address status. Non-admin callers may only set verified addresses back to unverified; setting to verified requires admin privileges.","type":"string","example":"verified","enum":["unverified","verified"],"x-auditable":true,"key$":"status"}},"required":["status"],"x-ref":"#/components/schemas/email_update_destination_address_properties","index$":1}}}},"parameters":[{"name":"destination_address_identifier","in":"path","required":true,"schema":{"description":"Destination address identifier.","type":"string","example":"ea95132c15732412d22c1476fa83f27a","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_destination_address_identifier"},"index$":0},{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const email_destination_address_response_single_ref01_ent = client.EmailDestinationAddressResponseSingle()
    let email_destination_address_response_single_ref01_data = setup.data.new.email_destination_address_response_single['email_destination_address_response_single_ref01']
    email_destination_address_response_single_ref01_data['account_id'] = setup.idmap['account01']

    email_destination_address_response_single_ref01_data = (await email_destination_address_response_single_ref01_ent.create(email_destination_address_response_single_ref01_data)).data()
    assert(null != email_destination_address_response_single_ref01_data)


    // UPDATE
    const email_destination_address_response_single_ref01_data_up0: any = {}
    email_destination_address_response_single_ref01_data_up0 ['account_id'] = setup.idmap['account_id']

    const email_destination_address_response_single_ref01_markdef_up0 = { name: 'email', value: 'Mark01-email_destination_address_response_single_ref01_' + setup.now }
    ;(email_destination_address_response_single_ref01_data_up0 as any)[email_destination_address_response_single_ref01_markdef_up0.name] = email_destination_address_response_single_ref01_markdef_up0.value

    const email_destination_address_response_single_ref01_resdata_up0 = (await email_destination_address_response_single_ref01_ent.update(email_destination_address_response_single_ref01_data_up0)).data()
    assert(null != email_destination_address_response_single_ref01_resdata_up0)

    assert((email_destination_address_response_single_ref01_resdata_up0 as any)[email_destination_address_response_single_ref01_markdef_up0.name] === email_destination_address_response_single_ref01_markdef_up0.value)




  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_destination_address_response_single/EmailDestinationAddressResponseSingleTestData.json')

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
    ['email_destination_address_response_single01','email_destination_address_response_single02','email_destination_address_response_single03','account01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESS_RESPONSE_SINGLE_ENTID']
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
  
