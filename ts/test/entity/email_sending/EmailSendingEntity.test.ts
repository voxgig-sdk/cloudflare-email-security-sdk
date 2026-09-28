

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


describe('EmailSendingEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailSending()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['create', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_sending.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"attachments":{"a":true,"h":"Attachments","n":"attachments","r":false,"sh":"File attachments and inline images.","t":"`$ARRAY`","key$":"attachments","index$":0},"bcc":{"a":true,"h":"Bcc","n":"bcc","r":false,"sh":"Blind carbon copy recipient(s).","t":"`$ANY`","union":{"branches":3,"count":2,"depth":3},"key$":"bcc","index$":1},"cc":{"a":true,"h":"Cc","n":"cc","r":false,"sh":"Carbon copy recipient(s).","t":"`$ANY`","union":{"branches":3,"count":2,"depth":3},"key$":"cc","index$":2},"errors":{"a":true,"h":"Errors","n":"errors","r":true,"t":"`$ARRAY`","key$":"errors","index$":3},"from":{"a":true,"h":"From","n":"from","r":true,"sh":"Sender email address.","t":"`$ANY`","union":{"branches":2,"count":1,"depth":0},"key$":"from","index$":4},"headers":{"a":true,"h":"Headers","n":"headers","r":false,"sh":"Custom email headers as key-value pairs.","t":"`$OBJECT`","key$":"headers","index$":5},"html":{"a":true,"h":"Html","n":"html","r":false,"sh":"HTML body of the email.","t":"`$STRING`","key$":"html","index$":6},"messages":{"a":true,"h":"Messages","n":"messages","r":true,"t":"`$ARRAY`","key$":"messages","index$":7},"mime_message":{"a":true,"h":"Mime Message","n":"mime_message","r":true,"sh":"The full MIME-encoded email message.","t":"`$STRING`","key$":"mime_message","index$":8},"recipients":{"a":true,"h":"Recipients","n":"recipients","r":true,"sh":"List of recipient email addresses.","t":"`$ARRAY`","key$":"recipients","index$":9},"reply_to":{"a":true,"h":"Reply To","n":"reply_to","r":false,"sh":"Reply-to address.","t":"`$ANY`","union":{"branches":2,"count":1,"depth":0},"key$":"reply_to","index$":10},"result":{"a":true,"h":"Result","n":"result","r":true,"t":"`$OBJECT`","key$":"result","index$":11},"result_info":{"a":true,"h":"Result Info","n":"result_info","r":true,"t":"`$OBJECT`","key$":"result_info","index$":12},"subject":{"a":true,"h":"Subject","n":"subject","r":true,"sh":"Email subject line.","t":"`$STRING`","key$":"subject","index$":13},"success":{"a":true,"h":"Success","n":"success","r":true,"t":"`$BOOLEAN`","key$":"success","index$":14},"text":{"a":true,"h":"Text","n":"text","r":false,"sh":"Plain text body of the email.","t":"`$STRING`","key$":"text","index$":15},"to":{"a":true,"h":"To","n":"to","r":false,"sh":"Recipient(s).","t":"`$ANY`","union":{"branches":3,"count":2,"depth":3},"key$":"to","index$":16}},"name":"email_sending","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /accounts/{account_id}/email/sending/send","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/sending/send","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"send"}],"t":{"req":{"attachments":"`reqdata.attachment`","bcc":"`reqdata.bcc`","cc":"`reqdata.cc`","from":"`reqdata.from`","headers":"`reqdata.header`","html":"`reqdata.html`","reply_to":"`reqdata.reply_to`","subject":"`reqdata.subject`","text":"`reqdata.text`","to":"`reqdata.to`"},"res":"`body`"},"index$":0},{"a":true,"co":{"id":"POST /accounts/{account_id}/email/sending/send_raw","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/sending/send_raw","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"send_raw"}],"t":{"req":{"from":"`reqdata.from`","mime_message":"`reqdata.mime_message`","recipients":"`reqdata.recipient`"},"res":"`body`"},"index$":1}],"key$":"create"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/email/sending/messages/{message_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"<aB3xK9mP@example.com>","k":"param","n":"message_id","or":"message_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/sending/messages/{message_id}","q":{"exist":["account_id","message_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"messages"},{"var":"message_id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"email_sending","name__orig":"email_sending","Name":"EmailSending","name_":"email_sending","name-":"email-sending","NAME":"EMAIL_SENDING","index$":14}, {"active":true,"entity":"email_sending","key$":"BasicEmailSendingFlow","kind":"basic","name":"BasicEmailSendingFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"email_sending_ref01"},"m":{"account_id":"account01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{"ref":"email_sending_ref01","srcdatavar":"email_sending_ref01_data","suffix":"_dt0"},"m":{"account_id":"account01","id":"email_sending01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_sending_ref01"}}],"index$":1}]}, 'EmailSending', {"POST /accounts/{account_id}/email/sending/send":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"attachments":{"description":"File attachments and inline images.","type":"array","items":{"discriminator":{"mapping":{"attachment":"#/components/schemas/email-sending_EmailAttachment","inline":"#/components/schemas/email-sending_EmailInlineAttachment"},"propertyName":"disposition"},"oneOf":[{"type":"object","properties":{},"example":{},"required":[],"x-ref":"#/components/schemas/email-sending_EmailInlineAttachment"},{"type":"object","properties":{},"example":{},"required":[],"x-ref":"#/components/schemas/email-sending_EmailAttachment"}]},"example":[{"content":"JVBERi0xLjQK...","disposition":"attachment","filename":"report.pdf","type":"application/pdf"}],"key$":"attachments"},"bcc":{"description":"Blind carbon copy recipient(s). Optional. A single email string, a named address object, or an array of either.","example":["bcc-recipient@example.com"],"oneOf":[{"description":"An email address as a plain string.","type":"string","example":"user@example.com","pattern":"^[\\x20-\\x7E]+$","x-ref":"#/components/schemas/email-sending_EmailAddressString"},{"type":"object","properties":{"address":{},"name":{}},"example":{"address":"user@example.com","name":"John Doe"},"required":["address"],"x-ref":"#/components/schemas/email-sending_EmailAddressObject"},{"items":{"oneOf":[]},"type":"array"}],"key$":"bcc"},"cc":{"description":"Carbon copy recipient(s). Optional. A single email string, a named address object, or an array of either.","example":["cc-recipient@example.com"],"oneOf":[{"description":"An email address as a plain string.","type":"string","example":"user@example.com","pattern":"^[\\x20-\\x7E]+$","x-ref":"#/components/schemas/email-sending_EmailAddressString"},{"type":"object","properties":{"address":{},"name":{}},"example":{"address":"user@example.com","name":"John Doe"},"required":["address"],"x-ref":"#/components/schemas/email-sending_EmailAddressObject"},{"items":{"oneOf":[]},"type":"array"}],"key$":"cc"},"from":{"description":"Sender email address. Either a plain string or an object with address and name.","example":"sender@example.com","oneOf":[{"description":"An email address as a plain string.","type":"string","example":"user@example.com","pattern":"^[\\x20-\\x7E]+$","x-ref":"#/components/schemas/email-sending_EmailAddressString"},{"type":"object","properties":{"address":{},"name":{}},"example":{"address":"user@example.com","name":"John Doe"},"required":["address"],"x-ref":"#/components/schemas/email-sending_EmailAddressObject"}],"key$":"from"},"headers":{"description":"Custom email headers as key-value pairs.","type":"object","example":{"X-Custom-Header":"value"},"additionalProperties":{"type":"string"},"key$":"headers"},"html":{"description":"HTML body of the email. Provide at least one of text or html (non-empty).","type":"string","example":"<h1>Hello</h1><p>Please find your report attached.</p>","key$":"html"},"reply_to":{"description":"Reply-to address. Either a plain string or an object with address and name.","example":"replies@example.com","oneOf":[{"description":"An email address as a plain string.","type":"string","example":"user@example.com","pattern":"^[\\x20-\\x7E]+$","x-ref":"#/components/schemas/email-sending_EmailAddressString"},{"type":"object","properties":{"address":{},"name":{}},"example":{"address":"user@example.com","name":"John Doe"},"required":["address"],"x-ref":"#/components/schemas/email-sending_EmailAddressObject"}],"key$":"reply_to"},"subject":{"description":"Email subject line.","type":"string","example":"Monthly Report","key$":"subject"},"text":{"description":"Plain text body of the email. Provide at least one of text or html (non-empty).","type":"string","example":"Hello\n\nPlease find your report attached.","key$":"text"},"to":{"description":"Recipient(s). Optional if cc or bcc is provided. A single email string, a named address object, or an array of either.","example":["recipient@example.com"],"oneOf":[{"description":"An email address as a plain string.","type":"string","example":"user@example.com","pattern":"^[\\x20-\\x7E]+$","x-ref":"#/components/schemas/email-sending_EmailAddressString"},{"type":"object","properties":{"address":{},"name":{}},"example":{"address":"user@example.com","name":"John Doe"},"required":["address"],"x-ref":"#/components/schemas/email-sending_EmailAddressObject"},{"items":{"oneOf":[]},"type":"array"}],"key$":"to"}},"example":{"attachments":[{"content":"JVBERi0xLjQK...","disposition":"attachment","filename":"report.pdf","type":"application/pdf"}],"from":"sender@example.com","html":"<h1>Hello</h1><p>Please find your report attached.</p>","subject":"Monthly Report","text":"Hello\n\nPlease find your report attached.","to":["recipient@example.com"]},"additionalProperties":false,"allOf":[{"anyOf":[{"required":["to"],"type":"object"},{"required":["cc"],"type":"object"},{"required":["bcc"],"type":"object"}]},{"anyOf":[{"required":["text"],"type":"object"},{"required":["html"],"type":"object"}]}],"required":["from","subject"],"x-ref":"#/components/schemas/email-sending_EmailBuilder","index$":1}}}},"parameters":[{"name":"account_id","in":"path","description":"Identifier of the account.","required":true,"schema":{"description":"Identifier of the account.","type":"string"},"index$":0}]},"POST /accounts/{account_id}/email/sending/send_raw":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","properties":{"from":{"description":"Sender email address.","type":"string","example":"sender@example.com","pattern":"^[\\x20-\\x7E]+$","key$":"from"},"mime_message":{"description":"The full MIME-encoded email message. Should include standard RFC 5322 headers such as From, To, Subject, and Content-Type. The from and recipients fields in the request body control SMTP envelope routing; the From and To headers in the MIME message control what the recipient's email client displays.","type":"string","example":"From: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Hello\r\nContent-Type: text/plain\r\n\r\nHello, World!","key$":"mime_message"},"recipients":{"description":"List of recipient email addresses.","type":"array","items":{"pattern":"^[\\x20-\\x7E]+$","type":"string"},"example":["recipient@example.com"],"key$":"recipients"}},"example":{"from":"sender@example.com","mime_message":"From: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Hello\r\nContent-Type: text/plain\r\n\r\nHello, World!","recipients":["recipient@example.com"]},"additionalProperties":false,"required":["from","recipients","mime_message"],"x-ref":"#/components/schemas/email-sending_SendRawRequest","index$":1}}}},"parameters":[{"name":"account_id","in":"path","description":"Identifier of the account.","required":true,"schema":{"description":"Identifier of the account.","type":"string"},"index$":0}]},"GET /accounts/{account_id}/email/sending/messages/{message_id}":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"type":"string","example":"12345","pattern":"^\\d+$"},"index$":0},{"name":"message_id","in":"path","required":true,"schema":{"type":"string","example":"<aB3xK9mP@example.com>"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const email_sending_ref01_ent = client.EmailSending()
    let email_sending_ref01_data = setup.data.new.email_sending['email_sending_ref01']
    email_sending_ref01_data['account_id'] = setup.idmap['account01']

    email_sending_ref01_data = (await email_sending_ref01_ent.create(email_sending_ref01_data)).data()
    assert(null != email_sending_ref01_data)



  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_sending/EmailSendingTestData.json')

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
    ['email_sending01','email_sending02','email_sending03','account01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_ENTID']
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
  
