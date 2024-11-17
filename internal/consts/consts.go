package consts

const (
	ProjectName              = "goshop"             // 项目名称
	ProjectUsage             = "个人学习使用"             // 项目用途
	ProjectBrief             = "基于gf框架的商城系统"        // 项目简介
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GtokenAdminPrefix        = "Admin:"             // gtoken登录管理后台前缀，用于区分admin和user的用户名相同的情况
	GtokenUserPrefix         = "User:"              // gtoken登录客户端前缀，用于区分admin和user的用户名相同的情况
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CtxUserId                = "CtxUserId"
	CtxUserName              = "CtxUserName"
	CtxUserAvatar            = "CtxUserAvatar"
	CtxUserSign              = "CtxUserSign"
	CtxUserSex               = "CtxUserSex"
	CtxUserStatus            = "CtxUserStatus"
	UserStatusBlacked        = 2 // 用户状态：2-黑名单
	ErrSecretAnswer          = "密保答案错误"
	ErrUserStatus            = "您的账号被冻结拉黑，请联系管理员"
	ErrNoPermission          = "您没有权限访问"
	ErrUserNotExist          = "用户不存在"
	ErrPassword              = "密码错误"
	MinPasswordLength        = 6 // 密码最小长度
	ArticlePublisherAdmin    = 1 // 后台管理员发布文章
	ArticlePublisherFrontend = 2 // 前台用户发布文章
	PraiseGoodsType          = 1 // 点赞商品类型
	PraiseArticleType        = 2 // 点赞查询文章类型
	CollectionGoodsType      = 1 // 收藏商品类型
	CollectionArticleType    = 2 // 收藏文章类型
	ConsigneeNotDefault      = 0
	ConsigneeDefault         = 1                // 默认收货人
	CacheModeRedis           = 2                // 缓存模式：2-Redis
	MultiLogin               = true             // 是否支持多点登录
	GtokenExpireIn           = 7 * 24 * 60 * 60 // gtoken过期时间(单位:秒)
)
