# CloudflareEmailSecurity SDK exists test

import pytest
from cloudflareemailsecurity_sdk import CloudflareEmailSecuritySDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = CloudflareEmailSecuritySDK.test(None, None)
        assert testsdk is not None
