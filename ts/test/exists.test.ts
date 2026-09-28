
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { CloudflareEmailSecuritySDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = CloudflareEmailSecuritySDK.test()
    equal(testsdk instanceof CloudflareEmailSecuritySDK, true,
      'CloudflareEmailSecuritySDK.test() must return a client synchronously')
  })

})
