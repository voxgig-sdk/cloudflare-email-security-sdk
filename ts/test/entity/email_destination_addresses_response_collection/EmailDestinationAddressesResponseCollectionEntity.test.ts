

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


describe('EmailDestinationAddressesResponseCollectionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailDestinationAddressesResponseCollection()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_destination_addresses_response_collection.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$ARRAY`","key$":"result","index$":0},"result_info":{"a":true,"h":"Result Info","n":"result_info","r":false,"t":"`$OBJECT`","key$":"result_info","index$":1}},"name":"email_destination_addresses_response_collection","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/email/routing/addresses","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"asc","k":"query","n":"direction","or":"direction","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":1,"k":"query","n":"page","or":"page","r":false,"t":"`$NUMBER`","index$":1},{"a":true,"ex":20,"k":"query","n":"per_page","or":"per_page","r":false,"t":"`$NUMBER`","index$":2},{"a":true,"ex":true,"k":"query","n":"verified","or":"verified","r":false,"t":"`$BOOLEAN`","index$":3}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/routing/addresses","q":{"exist":["account_id","direction","page","per_page","verified"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"routing"},{"lit":"addresses"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"email_destination_addresses_response_collection","name__orig":"email_destination_addresses_response_collection","Name":"EmailDestinationAddressesResponseCollection","name_":"email_destination_addresses_response_collection","name-":"email-destination-addresses-response-collection","NAME":"EMAIL_DESTINATION_ADDRESSES_RESPONSE_COLLECTION","index$":5}, {"active":true,"entity":"email_destination_addresses_response_collection","key$":"BasicEmailDestinationAddressesResponseCollectionFlow","kind":"basic","name":"BasicEmailDestinationAddressesResponseCollectionFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{"account_id":"account01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"email_destination_addresses_response_collection_ref01"}}],"index$":0}]}, 'EmailDestinationAddressesResponseCollection', {"GET /accounts/{account_id}/email/routing/addresses":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":0},{"name":"page","in":"query","schema":{"description":"Page number of paginated results.","type":"number","default":1,"minimum":1},"index$":1},{"name":"per_page","in":"query","schema":{"description":"Maximum number of results per page.","type":"number","default":20,"maximum":50,"minimum":5},"index$":2},{"name":"direction","in":"query","schema":{"description":"Sorts results in an ascending or descending order.","type":"string","example":"asc","default":"asc","enum":["asc","desc"]},"index$":3},{"name":"verified","in":"query","schema":{"description":"Filter by verified destination addresses.","type":"boolean","example":true,"default":true,"enum":[true,false]},"index$":4}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let email_destination_addresses_response_collection_ref01_data = Object.values(setup.data.existing.email_destination_addresses_response_collection)[0] as any

    // LIST
    const email_destination_addresses_response_collection_ref01_ent = client.EmailDestinationAddressesResponseCollection()
    const email_destination_addresses_response_collection_ref01_match: any = {}
    email_destination_addresses_response_collection_ref01_match['account_id'] = setup.idmap['account01']

    const email_destination_addresses_response_collection_ref01_list = (await email_destination_addresses_response_collection_ref01_ent.list(email_destination_addresses_response_collection_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_destination_addresses_response_collection/EmailDestinationAddressesResponseCollectionTestData.json')

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
    ['email_destination_addresses_response_collection01','email_destination_addresses_response_collection02','email_destination_addresses_response_collection03','account01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESSES_RESPONSE_COLLECTION_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESSES_RESPONSE_COLLECTION_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_DESTINATION_ADDRESSES_RESPONSE_COLLECTION_ENTID']
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
  
