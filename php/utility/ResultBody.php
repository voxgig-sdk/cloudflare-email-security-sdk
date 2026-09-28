<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK utility: result_body

class CloudflareEmailSecurityResultBody
{
    public static function call(CloudflareEmailSecurityContext $ctx): ?CloudflareEmailSecurityResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
