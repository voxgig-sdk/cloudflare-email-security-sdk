

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


describe('EmailSendingSuppressionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE=TRUE.
  afterEach(liveDelay('CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    const ent = testsdk.EmailSendingSuppression()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'email_sending_suppression.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"sh":"When the suppression was created.","t":"`$STRING`","key$":"created_at","index$":0},"email":{"a":true,"fo":"email","h":"Email","n":"email","r":true,"sh":"The email address to suppress.","t":"`$STRING`","key$":"email","index$":1},"errors":{"a":true,"h":"Errors","n":"errors","r":true,"t":"`$ARRAY`","key$":"errors","index$":2},"expires_at":{"a":true,"fo":"date-time","h":"Expires At","n":"expires_at","op":{"create":{"req":false,"type":"`$STRING`"},"update":{"req":false,"type":"`$STRING`"}},"r":true,"sh":"Expiration timestamp for the suppression.","t":"`$STRING`","key$":"expires_at","index$":3},"id":{"a":true,"fo":"uuid","h":"Id","n":"id","r":true,"sh":"Unique identifier for this suppression.","t":"`$STRING`","key$":"id","index$":4},"items":{"a":true,"h":"Items","n":"items","r":true,"sh":"Suppressions to import.","t":"`$ARRAY`","union":{"branches":2,"count":1,"depth":3},"key$":"items","index$":5},"messages":{"a":true,"h":"Messages","n":"messages","r":true,"t":"`$ARRAY`","key$":"messages","index$":6},"note":{"a":true,"h":"Note","n":"note","r":false,"sh":"Advisory note for this suppression.","t":"`$STRING`","key$":"note","index$":7},"read_only":{"a":true,"h":"Read Only","n":"read_only","r":true,"sh":"Whether clients may mutate this suppression.","t":"`$BOOLEAN`","key$":"read_only","index$":8},"reason":{"a":true,"h":"Reason","n":"reason","r":true,"sh":"Why the address is suppressed: `manual`, `complaint`, `hard_bounce`, `soft_bounce`, or `policy`.","t":"`$STRING`","key$":"reason","index$":9},"result":{"a":true,"h":"Result","n":"result","r":true,"t":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"key$":"result","index$":10},"scope":{"a":true,"h":"Scope","n":"scope","r":false,"sh":"Where the suppression applies.","t":"`$ANY`","union":{"branches":2,"count":1,"depth":0},"key$":"scope","index$":11},"success":{"a":true,"h":"Success","n":"success","r":true,"t":"`$BOOLEAN`","key$":"success","index$":12},"zones":{"a":true,"h":"Zones","n":"zones","r":false,"t":"`$ARRAY`","key$":"zones","index$":13}},"id":{"field":"id","name":"id"},"name":"email_sending_suppression","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /accounts/{account_id}/email/sending/suppression","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"54442216","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/sending/suppression","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":0},{"a":true,"co":{"id":"POST /accounts/{account_id}/email/sending/suppressions","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/sending/suppressions","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1},{"a":true,"co":{"id":"POST /accounts/{account_id}/email/sending/suppressions/bulk","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/accounts/{account_id}/email/sending/suppressions/bulk","q":{"exist":["account_id"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"},{"lit":"bulk"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":2},{"a":true,"co":{"id":"POST /zones/{zone_id}/email/sending/suppression","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"88228822","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/zones/{zone_id}/email/sending/suppression","q":{"exist":["zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":3}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/email/sending/suppressions","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"eyJ0Ijo0LCJwaCI6ImRvbWFpbiIsInAiOjEyMywiYyI6ImMzZGM1ZjBiMzRhMTRmZjhlMWIzZWMwNDg5NWUxYjIyIn0","k":"query","n":"cursor","or":"cursor","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":"user@example.com","k":"query","n":"email","or":"email","r":false,"t":"`$STRING`","index$":1},{"a":true,"ex":100,"k":"query","n":"per_page","or":"per_page","r":false,"t":"`$INTEGER`","index$":2},{"a":true,"ex":"hard_bounce","k":"query","n":"reason","or":"reason","r":false,"t":"`$STRING`","index$":3},{"a":true,"ex":"sending_domain","k":"query","n":"scope_type","or":"scope_type","r":false,"t":"`$STRING`","index$":4},{"a":true,"ex":"mail.example.com","k":"query","n":"scope_value","or":"scope_value","r":false,"t":"`$STRING`","index$":5},{"a":true,"ex":"billing@","k":"query","n":"search","or":"search","r":false,"t":"`$STRING`","index$":6}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/sending/suppressions","q":{"exist":["account_id","cursor","email","per_page","reason","scope_type","scope_value","search"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /accounts/{account_id}/email/sending/suppression","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"54442216","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"desc","k":"query","n":"direction","or":"direction","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":"created_at","k":"query","n":"order","or":"order","r":false,"t":"`$STRING`","index$":1},{"a":true,"ex":1,"k":"query","n":"page","or":"page","r":false,"t":"`$INTEGER`","index$":2},{"a":true,"ex":100,"k":"query","n":"per_page","or":"per_page","r":false,"t":"`$INTEGER`","index$":3}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/sending/suppression","q":{"exist":["account_id","direction","order","page","per_page"]},"r":{},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":1},{"a":true,"co":{"id":"GET /zones/{zone_id}/email/sending/suppression","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"88228822","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"desc","k":"query","n":"direction","or":"direction","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":"created_at","k":"query","n":"order","or":"order","r":false,"t":"`$STRING`","index$":1},{"a":true,"ex":1,"k":"query","n":"page","or":"page","r":false,"t":"`$INTEGER`","index$":2},{"a":true,"ex":100,"k":"query","n":"per_page","or":"per_page","r":false,"t":"`$INTEGER`","index$":3}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/sending/suppression","q":{"exist":["direction","order","page","per_page","zone_id"]},"r":{},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":2}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /accounts/{account_id}/email/sending/suppression/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"54442216","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/sending/suppression/{suppression_id}","q":{"exist":["account_id","id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":0},{"a":true,"co":{"id":"GET /accounts/{account_id}/email/sending/suppressions/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/accounts/{account_id}/email/sending/suppressions/{suppression_id}","q":{"exist":["account_id","id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1},{"a":true,"co":{"id":"GET /zones/{zone_id}/email/sending/suppression/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"88228822","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/zones/{zone_id}/email/sending/suppression/{suppression_id}","q":{"exist":["id","zone_id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body.result`"},"index$":2}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /accounts/{account_id}/email/sending/suppression/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"54442216","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/accounts/{account_id}/email/sending/suppression/{suppression_id}","q":{"exist":["account_id","id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"DELETE /accounts/{account_id}/email/sending/suppressions/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/accounts/{account_id}/email/sending/suppressions/{suppression_id}","q":{"exist":["account_id","id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1},{"a":true,"co":{"id":"DELETE /zones/{zone_id}/email/sending/suppression/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"88228822","k":"param","n":"zone_id","or":"zone_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/zones/{zone_id}/email/sending/suppression/{suppression_id}","q":{"exist":["id","zone_id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"zones"},{"var":"zone_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppression"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":2}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PATCH /accounts/{account_id}/email/sending/suppressions/{suppression_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"12345678","k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"396a5436-d4b0-42a6-b3fc-48e8fa522321","k":"param","n":"id","or":"suppression_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PATCH","o":"/accounts/{account_id}/email/sending/suppressions/{suppression_id}","q":{"exist":["account_id","id"]},"r":{"param":{"suppression_id":"id"}},"s":[{"lit":"accounts"},{"var":"account_id"},{"lit":"email"},{"lit":"sending"},{"lit":"suppressions"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"email_sending_suppression","name__orig":"email_sending_suppression","Name":"EmailSendingSuppression","name_":"email_sending_suppression","name-":"email-sending-suppression","NAME":"EMAIL_SENDING_SUPPRESSION","index$":22}, {"active":true,"entity":"email_sending_suppression","key$":"BasicEmailSendingSuppressionFlow","kind":"basic","name":"BasicEmailSendingSuppressionFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"email_sending_suppression_ref01"},"m":{"account_id":"account01","zone_id":"zone01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{"zone_id":"zone01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"email_sending_suppression_ref01"}}],"index$":1},{"a":true,"d":{"account_id":"account01"},"i":{"ref":"email_sending_suppression_ref01","srcdatavar":"email_sending_suppression_ref01_data","suffix":"_up0","textfield":"created_at"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_sending_suppression_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"email_sending_suppression_ref01","srcdatavar":"email_sending_suppression_ref01_data","suffix":"_dt0"},"m":{"id":"email_sending_suppression01","zone_id":"zone01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-email_sending_suppression_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"email_sending_suppression_ref01","suffix":"_rm0"},"m":{"id":"email_sending_suppression01","zone_id":"zone01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{"zone_id":"zone01"},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"email_sending_suppression_ref01"}}],"index$":5}]}, 'EmailSendingSuppression', {"POST /accounts/{account_id}/email/sending/suppression":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"email":{"type":"string","format":"email","example":"user@example.com","key$":"email"},"expires_at":{"type":"string","format":"date-time","nullable":true,"key$":"expires_at"}},"required":["email"],"index$":1}}}},"parameters":[{"name":"account_id","in":"path","required":true,"schema":{"type":"string","example":"54442216"},"index$":0}]},"POST /accounts/{account_id}/email/sending/suppressions":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"email":{"description":"The email address to suppress.","type":"string","format":"email","example":"user@example.com","key$":"email"},"expires_at":{"description":"Expiration timestamp for the suppression. Omit or set to null for a permanent suppression that never expires.","type":"string","format":"date-time","example":"2027-01-01T00:00:00Z","nullable":true,"key$":"expires_at"},"note":{"description":"Advisory note for this suppression. Not enforced or validated beyond length.","type":"string","example":"Imported from CRM","maxLength":1000,"key$":"note"},"scope":{"description":"Where the suppression applies. Omit for `{ \"type\": \"account\" }`, which blocks the recipient for every sending domain of the account.","example":{"type":"sending_domain","value":"mail.example.com"},"oneOf":[{"properties":{"type":{}},"required":["type"],"type":"object"},{"properties":{"type":{},"value":{}},"required":["type","value"],"type":"object"}],"key$":"scope"}},"required":["email"],"index$":1}}}},"parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0}]},"POST /accounts/{account_id}/email/sending/suppressions/bulk":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"items":{"description":"Suppressions to import. Items with the same email address and scope are deduplicated before processing.","type":"array","items":{"properties":{"email":{"description":"The email address to suppress.","type":"string","example":"user@example.com"},"expires_at":{"description":"Expiration timestamp for the suppression. Omit or set to null for a permanent suppression that never expires.","type":"string","format":"date-time","example":"2027-01-01T00:00:00Z","nullable":true},"note":{"description":"Advisory note for this suppression. Not enforced or validated beyond length.","type":"string","example":"Imported from CRM","maxLength":1000},"scope":{"description":"Where the suppression applies. Omit for `{ \"type\": \"account\" }`, which blocks the recipient for every sending domain of the account.","example":{},"oneOf":[]}},"required":["email"],"type":"object"},"example":[{"email":"user@example.com"},{"email":"other@example.com","expires_at":"2027-01-01T00:00:00Z","note":"Imported from CRM","scope":{"type":"sending_domain","value":"mail.example.com"}}],"maxItems":1000,"key$":"items"}},"required":["items"],"index$":1}}}},"parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0}]},"POST /zones/{zone_id}/email/sending/suppression":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"email":{"type":"string","format":"email","example":"user@example.com","key$":"email"},"expires_at":{"type":"string","format":"date-time","nullable":true,"key$":"expires_at"}},"required":["email"],"index$":1}}}},"parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"type":"string","example":"88228822"},"index$":0}]},"GET /accounts/{account_id}/email/sending/suppressions":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0},{"name":"per_page","in":"query","schema":{"description":"Maximum number of suppressions to return per page.","type":"integer","example":100,"default":100,"maximum":1000,"minimum":1},"index$":1},{"name":"cursor","in":"query","schema":{"description":"Opaque pagination cursor returned as `result_info.next_cursor`. It carries the filters that produced it.","type":"string","example":"eyJ0Ijo0LCJwaCI6ImRvbWFpbiIsInAiOjEyMywiYyI6ImMzZGM1ZjBiMzRhMTRmZjhlMWIzZWMwNDg5NWUxYjIyIn0","maxLength":2048},"index$":2},{"name":"email","in":"query","schema":{"description":"Exact email-address filter.","type":"string","format":"email","example":"user@example.com"},"index$":3},{"name":"search","in":"query","schema":{"description":"A complete address is an exact match; a value ending in `@` matches that username across every domain. Prefix searches may return short intermediate pages while the bounded account scan advances.","type":"string","example":"billing@","maxLength":320,"minLength":1},"index$":4},{"name":"reason","in":"query","schema":{"description":"Filter to suppressions with this reason.","type":"string","example":"hard_bounce","enum":["manual","complaint","hard_bounce","soft_bounce","policy"]},"index$":5},{"name":"scope_type","in":"query","schema":{"description":"Filter by scope: `account` returns only account-wide suppressions, `sending_domain` only sending-domain suppressions. Omit to list both, sending-domain suppressions first.","type":"string","example":"sending_domain","enum":["account","sending_domain"]},"index$":6},{"name":"scope_value","in":"query","schema":{"description":"Exact sending-domain filter. Requires `scope_type=sending_domain`.","type":"string","example":"mail.example.com","maxLength":1024},"index$":7}]},"GET /accounts/{account_id}/email/sending/suppression":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"type":"string","example":"54442216"},"index$":0},{"name":"page","in":"query","schema":{"type":"integer","default":1,"minimum":0},"index$":1},{"name":"per_page","in":"query","schema":{"type":"integer","default":100,"maximum":1000,"minimum":0},"index$":2},{"name":"order","in":"query","schema":{"type":"string","default":"created_at","enum":["email","expires_at","created_at"]},"index$":3},{"name":"direction","in":"query","schema":{"type":"string","default":"desc","enum":["asc","desc"]},"index$":4}]},"GET /zones/{zone_id}/email/sending/suppression":{"protocol":"http","parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"type":"string","example":"88228822"},"index$":0},{"name":"page","in":"query","schema":{"type":"integer","default":1,"minimum":0},"index$":1},{"name":"per_page","in":"query","schema":{"type":"integer","default":100,"maximum":1000,"minimum":0},"index$":2},{"name":"order","in":"query","schema":{"type":"string","default":"created_at","enum":["email","expires_at","created_at"]},"index$":3},{"name":"direction","in":"query","schema":{"type":"string","default":"desc","enum":["asc","desc"]},"index$":4}]},"GET /accounts/{account_id}/email/sending/suppression/{suppression_id}":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"type":"string","example":"54442216"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"type":"string","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"GET /accounts/{account_id}/email/sending/suppressions/{suppression_id}":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"description":"The suppression's identifier.","type":"string","format":"uuid","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"GET /zones/{zone_id}/email/sending/suppression/{suppression_id}":{"protocol":"http","parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"type":"string","example":"88228822"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"type":"string","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"DELETE /accounts/{account_id}/email/sending/suppression/{suppression_id}":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"type":"string","example":"54442216"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"type":"string","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"DELETE /accounts/{account_id}/email/sending/suppressions/{suppression_id}":{"protocol":"http","parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"description":"The suppression's identifier.","type":"string","format":"uuid","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"DELETE /zones/{zone_id}/email/sending/suppression/{suppression_id}":{"protocol":"http","parameters":[{"name":"zone_id","in":"path","required":true,"schema":{"type":"string","example":"88228822"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"type":"string","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]},"PATCH /accounts/{account_id}/email/sending/suppressions/{suppression_id}":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"expires_at":{"description":"New expiry. Send `null` to make the suppression permanent; omit to leave it unchanged.","type":"string","format":"date-time","example":"2027-01-01T00:00:00Z","nullable":true,"key$":"expires_at"},"note":{"description":"Replacement advisory note. Send an empty string to clear it; omit to leave it unchanged.","type":"string","example":"Customer re-confirmed opt-in","maxLength":1000,"key$":"note"},"scope":{"description":"Not editable. Scope is fixed when the suppression is created; any value returns 400 with code `scope_immutable`. Delete and recreate the suppression to change it.","example":{"type":"account"},"nullable":true,"key$":"scope"}},"index$":1}}}},"parameters":[{"name":"account_id","in":"path","required":true,"schema":{"description":"Cloudflare account ID.","type":"string","example":"12345678"},"index$":0},{"name":"suppression_id","in":"path","required":true,"schema":{"description":"The suppression's identifier.","type":"string","format":"uuid","example":"396a5436-d4b0-42a6-b3fc-48e8fa522321"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const email_sending_suppression_ref01_ent = client.EmailSendingSuppression()
    let email_sending_suppression_ref01_data = setup.data.new.email_sending_suppression['email_sending_suppression_ref01']
    email_sending_suppression_ref01_data['account_id'] = setup.idmap['account01']
    email_sending_suppression_ref01_data['zone_id'] = setup.idmap['zone01']

    email_sending_suppression_ref01_data = (await email_sending_suppression_ref01_ent.create(email_sending_suppression_ref01_data)).data()
    assert(null != email_sending_suppression_ref01_data.id)


    // LIST
    const email_sending_suppression_ref01_match: any = {}
    email_sending_suppression_ref01_match['zone_id'] = setup.idmap['zone01']

    const email_sending_suppression_ref01_list = (await email_sending_suppression_ref01_ent.list(email_sending_suppression_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(email_sending_suppression_ref01_list, { id: email_sending_suppression_ref01_data.id })))


    // UPDATE
    const email_sending_suppression_ref01_data_up0: any = {}
    email_sending_suppression_ref01_data_up0.id = email_sending_suppression_ref01_data.id
    email_sending_suppression_ref01_data_up0 ['account_id'] = setup.idmap['account_id']

    const email_sending_suppression_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-email_sending_suppression_ref01_' + setup.now }
    ;(email_sending_suppression_ref01_data_up0 as any)[email_sending_suppression_ref01_markdef_up0.name] = email_sending_suppression_ref01_markdef_up0.value

    const email_sending_suppression_ref01_resdata_up0 = (await email_sending_suppression_ref01_ent.update(email_sending_suppression_ref01_data_up0)).data()
    assert(email_sending_suppression_ref01_resdata_up0.id === email_sending_suppression_ref01_data_up0.id)

    assert((email_sending_suppression_ref01_resdata_up0 as any)[email_sending_suppression_ref01_markdef_up0.name] === email_sending_suppression_ref01_markdef_up0.value)


    // LOAD
    const email_sending_suppression_ref01_match_dt0: any = {}
    email_sending_suppression_ref01_match_dt0.id = email_sending_suppression_ref01_data.id
    const email_sending_suppression_ref01_data_dt0 = (await email_sending_suppression_ref01_ent.load(email_sending_suppression_ref01_match_dt0)).data()
    assert(email_sending_suppression_ref01_data_dt0.id === email_sending_suppression_ref01_data.id)


    // REMOVE
    const email_sending_suppression_ref01_match_rm0: any = { id: email_sending_suppression_ref01_data.id }
    await email_sending_suppression_ref01_ent.remove(email_sending_suppression_ref01_match_rm0)
  

    // LIST
    const email_sending_suppression_ref01_match_rt0: any = {}
    email_sending_suppression_ref01_match_rt0['zone_id'] = setup.idmap['zone01']

    const email_sending_suppression_ref01_list_rt0 = (await email_sending_suppression_ref01_ent.list(email_sending_suppression_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(email_sending_suppression_ref01_list_rt0, { id: email_sending_suppression_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/email_sending_suppression/EmailSendingSuppressionTestData.json')

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
    ['email_sending_suppression01','email_sending_suppression02','email_sending_suppression03','account01','zone01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUPPRESSION_ENTID': idmap,
    'CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_TEST_EXPLAIN': 'FALSE',
    'CLOUDFLARE_EMAIL_SECURITY_APIKEY': '',
  })

  idmap = env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUPPRESSION_ENTID']

  const live = 'TRUE' === env.CLOUDFLARE_EMAIL_SECURITY_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CLOUDFLARE_EMAIL_SECURITY_TEST_EMAIL_SENDING_SUPPRESSION_ENTID']
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
  
