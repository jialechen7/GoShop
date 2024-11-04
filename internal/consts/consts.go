package consts

const (
	Version            = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey         = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GtokenAdminPrefix  = "Admin:"             // gtoken登录管理后台前缀，用于区分admin和user的用户名相同的情况
	CtxAdminId         = "CtxAdminId"
	CtxAdminName       = "CtxAdminName"
	CtxAdminIsAdmin    = "CtxAdminIsAdmin"
	CtxAdminRoleIds    = "CtxAdminRoleIds"
)
