<?php
declare(strict_types=1);

// CloudflareEmailSecurity SDK base feature

class CloudflareEmailSecurityBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(CloudflareEmailSecurityContext $ctx, array $options): void {}
    public function PostConstruct(CloudflareEmailSecurityContext $ctx): void {}
    public function PostConstructEntity(CloudflareEmailSecurityContext $ctx): void {}
    public function SetData(CloudflareEmailSecurityContext $ctx): void {}
    public function GetData(CloudflareEmailSecurityContext $ctx): void {}
    public function GetMatch(CloudflareEmailSecurityContext $ctx): void {}
    public function SetMatch(CloudflareEmailSecurityContext $ctx): void {}
    public function PrePoint(CloudflareEmailSecurityContext $ctx): void {}
    public function PreSpec(CloudflareEmailSecurityContext $ctx): void {}
    public function PreRequest(CloudflareEmailSecurityContext $ctx): void {}
    public function PreResponse(CloudflareEmailSecurityContext $ctx): void {}
    public function PreResult(CloudflareEmailSecurityContext $ctx): void {}
    public function PreDone(CloudflareEmailSecurityContext $ctx): void {}
    public function PreUnexpected(CloudflareEmailSecurityContext $ctx): void {}
}
