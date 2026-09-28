<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK utility: result_headers

class CloudflareEmailSecurityResultHeaders
{
    public static function call(CloudflareEmailSecurityContext $ctx): ?CloudflareEmailSecurityResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
