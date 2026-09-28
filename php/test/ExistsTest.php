<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK exists test

require_once __DIR__ . '/../cloudflareemailsecurity_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = CloudflareEmailSecuritySDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
