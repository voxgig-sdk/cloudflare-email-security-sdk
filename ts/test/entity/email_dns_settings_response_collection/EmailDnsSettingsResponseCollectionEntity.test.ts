

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


describe('EmailDnsSettingsResponseCollectionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailDnsSettingsResponseCollection()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_dns_settings_response_collection.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$ARRAY`","union":{"branches":2,"count":1,"depth":3},"key$":"result","index$":0}},"name":"email_dns_settings_response_collection","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /zones/{zone_id}/email/routing/dns","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"example.net","k":"query","n":"subdomain","or":"subdomain","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/routing/dns","q":{"exist":["subdomain","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"routing"},{"lit":"dns"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"aabbccdd11223344aabbccdd11223344","k":"param","n":"subdomain_id","or":"subdomain_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns","q":{"exist":["subdomain_id","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"subdomains"},{"var":"subdomain_id"},{"lit":"dns"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"}},"relations":{"ancestors":[["$.main.kit.entity.subdomain"]]},"key$":"email_dns_settings_response_collection","name__orig":"email_dns_settings_response_collection","Name":"EmailDnsSettingsResponseCollection","name_":"email_dns_settings_response_collection","name-":"email-dns-settings-response-collection","NAME":"EMAIL_DNS_SETTINGS_RESPONSE_COLLECTION","index$":6}, {"active":true,"entity":"email_dns_settings_response_collection","key$":"BasicEmailDnsSettingsResponseCollectionFlow","kind":"basic","name":"BasicEmailDnsSettingsResponseCollectionFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{"subdomain_id":"subdomain01","zone_id":"zone01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"email_dns_settings_response_collection_ref01"}}],"index$":0}]}, 'EmailDnsSettingsResponseCollection', {"GET /zones/{zone_id}/email/routing/dns":{"protocol":"http","parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":0},{"name":"subdomain","in":"query","description":"Deprecated. When supplied, the response shape differs from the documented default and is not modeled in generated SDKs. Do not rely on this parameter.","schema":{"description":"Domain of your zone.","type":"string","example":"example.net","x-auditable":true,"x-ref":"#/components/schemas/email_email_setting_name"},"deprecated":true,"index$":1}]},"GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns":{"protocol":"http","parameters":[{"name":"subdomain_id","in":"path","required":true,"schema":{"description":"Sending subdomain identifier.","type":"string","example":"aabbccdd11223344aabbccdd11223344","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_sending_subdomain_identifier"},"index$":0},{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let email_dns_settings_response_collection_ref01_data = Object.values(setup.data.existing.email_dns_settings_response_collection)[0] as any

    // LIST
    const email_dns_settings_response_collection_ref01_ent = client.EmailDnsSettingsResponseCollection()
    const email_dns_settings_response_collection_ref01_match: any = {}
    email_dns_settings_response_collection_ref01_match['subdomain_id'] = setup.idmap['subdomain01']
    email_dns_settings_response_collection_ref01_match['zone_id'] = setup.idmap['zone01']

    const email_dns_settings_response_collection_ref01_list = (await email_dns_settings_response_collection_ref01_ent.list(email_dns_settings_response_collection_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_dns_settings_response_collection/EmailDnsSettingsResponseCollectionTestData.json')

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
    ['email_dns_settings_response_collection01','email_dns_settings_response_collection02','email_dns_settings_response_collection03','subdomain01','subdomain02','subdomain03','zone01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DNS_SETTINGS_RESPONSE_COLLECTION_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DNS_SETTINGS_RESPONSE_COLLECTION_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DNS_SETTINGS_RESPONSE_COLLECTION_ENTID']
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
  
