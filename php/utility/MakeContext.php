<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class CloudflareEmailSecurityMakeContext
{
    public static function call(array $ctxmap, ?CloudflareEmailSecurityContext $basectx): CloudflareEmailSecurityContext
    {
        return new CloudflareEmailSecurityContext($ctxmap, $basectx);
    }
}
