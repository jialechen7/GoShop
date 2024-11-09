package consts

const (
	Version            = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey         = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GtokenAdminPrefix  = "Admin:"             // gtoken登录管理后台前缀，用于区分admin和user的用户名相同的情况
	GtokenUserPrefix   = "User:"              // gtoken登录客户端前缀，用于区分admin和user的用户名相同的情况
	CtxAdminId         = "CtxAdminId"
	CtxAdminName       = "CtxAdminName"
	CtxAdminIsAdmin    = "CtxAdminIsAdmin"
	CtxAdminRoleIds    = "CtxAdminRoleIds"
	CtxUserId          = "CtxUserId"
	CtxUserName        = "CtxUserName"
	CtxUserAvatar      = "CtxUserAvatar"
	CtxUserSign        = "CtxUserSign"
	CtxUserSex         = "CtxUserSex"
	CtxUserStatus      = "CtxUserStatus"
	UserStatusBlacked  = 2 // 用户状态：2-黑名单
	ErrSecretAnswer    = "密保答案错误"
	ErrUserStatus      = "您的账号被冻结拉黑，请联系管理员"
	ErrNoPermission    = "您没有权限访问"
	MinPasswordLength  = 6 // 密码最小长度
)
