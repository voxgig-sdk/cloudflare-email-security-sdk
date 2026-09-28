package core

type CloudflareEmailSecurityError struct {
	IsCloudflareEmailSecurityError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewCloudflareEmailSecurityError(code string, msg string, ctx *Context) *CloudflareEmailSecurityError {
	return &CloudflareEmailSecurityError{
		IsCloudflareEmailSecurityError: true,
		Sdk:              "CloudflareEmailSecurity",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *CloudflareEmailSecurityError) Error() string {
	return e.Msg
}
