

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


describe('EmailSendingSubdomainReputationComplaintEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailSendingSubdomainReputationComplaint()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_sending_subdomain_reputation_complaint.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"result":{"a":true,"h":"Result","n":"result","r":true,"t":"`$OBJECT`","key$":"result","index$":0}},"name":"email_sending_subdomain_reputation_complaint","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"aabbccdd11223344aabbccdd11223344","k":"param","n":"subdomain_id","or":"subdomain_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"023e105f4ecef8ad9ca31a8372d0c353","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}],"query":[{"a":true,"k":"query","n":"end_at","or":"end_at","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"start_at","or":"start_at","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints","q":{"exist":["end_at","start_at","subdomain_id","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"subdomains"},{"var":"subdomain_id"},{"lit":"reputation"},{"lit":"complaints"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["$.main.kit.entity.subdomain"]]},"key$":"email_sending_subdomain_reputation_complaint","name__orig":"email_sending_subdomain_reputation_complaint","Name":"EmailSendingSubdomainReputationComplaint","name_":"email_sending_subdomain_reputation_complaint","name-":"email-sending-subdomain-reputation-complaint","NAME":"EMAIL_SENDING_SUBDOMAIN_REPUTATION_COMPLAINT","index$":19}, {"active":true,"entity":"email_sending_subdomain_reputation_complaint","key$":"BasicEmailSendingSubdomainReputationComplaintFlow","kind":"basic","name":"BasicEmailSendingSubdomainReputationComplaintFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"email_sending_subdomain_reputation_complaint_ref01","srcdatavar":"email_sending_subdomain_reputation_complaint_ref01_data","suffix":"_dt0"},"m":{"id":"email_sending_subdomain_reputation_complaint01","zone_id":"zone01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_sending_subdomain_reputation_complaint_ref01"}}],"index$":0}]}, 'EmailSendingSubdomainReputationComplaint', {"GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/reputation/complaints":{"protocol":"http","parameters":[{"name":"subdomain_id","in":"path","required":true,"schema":{"description":"Sending subdomain identifier.","type":"string","example":"aabbccdd11223344aabbccdd11223344","maxLength":32,"readOnly":true,"x-auditable":true,"x-ref":"#/components/schemas/email_sending_subdomain_identifier"},"index$":0},{"name":"zone_id","in":"path","required":true,"schema":{"description":"Identifier.","type":"string","example":"023e105f4ecef8ad9ca31a8372d0c353","maxLength":32,"x-auditable":true,"x-ref":"#/components/schemas/email_identifier"},"index$":1},{"name":"start_at","in":"query","required":true,"schema":{"type":"string","format":"date-time"},"index$":2},{"name":"end_at","in":"query","required":true,"schema":{"type":"string","format":"date-time"},"index$":3}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let email_sending_subdomain_reputation_complaint_ref01_data = Object.values(setup.data.existing.email_sending_subdomain_reputation_complaint)[0] as any

    // LOAD: skipped — no entity id field and load requires path params.
    // Entity-var is declared here so later flow steps still compile.
    const email_sending_subdomain_reputation_complaint_ref01_ent = client.EmailSendingSubdomainReputationComplaint()


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_sending_subdomain_reputation_complaint/EmailSendingSubdomainReputationComplaintTestData.json')

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
    ['email_sending_subdomain_reputation_complaint01','email_sending_subdomain_reputation_complaint02','email_sending_subdomain_reputation_complaint03','subdomain01','subdomain02','subdomain03','zone01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_REPUTATION_COMPLAINT_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_REPUTATION_COMPLAINT_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUBDOMAIN_REPUTATION_COMPLAINT_ENTID']
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
  
