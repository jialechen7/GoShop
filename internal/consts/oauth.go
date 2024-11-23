package consts

const (
	GithubStatePrefix = "GithubState:" // Github登录状态前缀
	StateExpireIn     = 60 * 10        // 状态过期时间(10分钟)
	ErrState          = "状态错误"
	ErrEncryptToken   = "token加密失败"
	AdminDefaultPwd   = "123456"
)
