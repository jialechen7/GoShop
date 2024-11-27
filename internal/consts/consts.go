package consts

import (
	"os"
	"sync"
)

const (
	ProjectName               = "goshop"      // 项目名称
	ProjectUsage              = "个人学习使用"      // 项目用途
	ProjectBrief              = "基于gf框架的商城系统" // 项目简介
	ContextKey                = "ContextKey"  // 上下文变量存储键名，前后端系统共享
	GtokenAdminPrefix         = "Admin:"      // gtoken登录管理后台前缀，用于区分admin和user的用户名相同的情况
	GtokenUserPrefix          = "User:"       // gtoken登录客户端前缀，用于区分admin和user的用户名相同的情况
	CtxAdminId                = "CtxAdminId"
	CtxAdminName              = "CtxAdminName"
	CtxAdminIsAdmin           = "CtxAdminIsAdmin"
	CtxAdminRoleIds           = "CtxAdminRoleIds"
	CtxUserId                 = "CtxUserId"
	CtxUserName               = "CtxUserName"
	CtxUserAvatar             = "CtxUserAvatar"
	CtxUserSign               = "CtxUserSign"
	CtxUserSex                = "CtxUserSex"
	CtxUserStatus             = "CtxUserStatus"
	ErrSecretAnswer           = "密保答案错误"
	ErrUserStatus             = "您的账号被冻结拉黑，请联系管理员"
	ErrNoPermission           = "您没有权限访问"
	ErrUserNotExist           = "用户不存在"
	ErrPassword               = "密码错误"
	ErrStockNotEnough         = "库存不足"
	ErrCaptcha                = "验证码错误"
	ErrParams                 = "参数错误"
	ErrStock                  = "库存不足"
	ErrSeckillNotStart        = "秒杀活动未开始"
	ErrSeckillEnd             = "秒杀活动已结束"
	ErrDecreaseStock          = "扣减库存失败"
	ErrHasSeckill             = "请勿重复领取"
	ErrSeckill                = "秒杀失败"
	MinPasswordLength         = 6 // 密码最小长度
	ArticlePublisherAdmin     = 1 // 后台管理员发布文章
	ArticlePublisherFrontend  = 2 // 前台用户发布文章
	PraiseGoodsType           = 1 // 点赞商品类型
	PraiseArticleType         = 2 // 点赞查询文章类型
	CollectionGoodsType       = 1 // 收藏商品类型
	CollectionArticleType     = 2 // 收藏文章类型
	CategoryLevel1            = 1 // 一级分类
	CategoryLevel2            = 2 // 二级分类
	CategoryLevel3            = 3 // 二级分类
	ConsigneeNotDefault       = 0
	ConsigneeDefault          = 1                // 默认收货人
	CacheModeCache            = 1                // 缓存模式：1-内存
	CacheModeRedis            = 2                // 缓存模式：2-Redis
	CacheModeFile             = 3                // 缓存模式：3-文件
	MultiLogin                = true             // 是否支持多点登录
	GtokenExpireIn            = 7 * 24 * 60 * 60 // gtoken过期时间(单位:秒)
	CaptchaLength             = 6                // 验证码长度
	CaptchaPrefix             = "captcha:"       // 验证码前缀
	CaptchaExpire             = 60 * 5           // 验证码过期时间(单位:秒)
	OrderIdKey                = "order_id"
	RedisLockKey              = "lock:"
	UserCouponIdKey           = "user_coupon:"   // 用户优惠券key
	CouponTypeCommon          = 0                // 优惠券类型：0-普通优惠券
	CouponTypeSeckill         = 1                // 优惠券类型：1-秒杀优惠券
	CouponStatusAvailable     = 1                // 优惠券状态：1-可用
	CouponStatusUsed          = 2                // 优惠券状态：2-已使用
	CouponStatusExpired       = 3                // 优惠券状态：3-已过期
	CouponStatusUsedText      = "已使用"            // 优惠券状态文本：已使用
	CouponStatusExpiredText   = "已过期"            // 优惠券状态文本：已过期
	UserStatusNormal          = 1                // 用户状态：1-正常
	UserStatusBlacked         = 2                // 用户状态：2-黑名单
	SeckillStockRedisPrefix   = "seckill:stock:" // 秒杀库存前缀
	UserHasSeckillRedisPrefix = "seckill:user:"
	SessionKeyAdmin           = "SessionKeyAdmin" // 用户信息存放在Session中的Key
	FileMaxUploadCountMinute  = 10                // 同一用户1分钟之内最大上传数量
	// Response Code
	UserNameOrPasswordError  = 3                    // 用户名或密码错误
	StreamUserCoupon         = "stream.user_coupon" // 用户优惠券消息队列
	StreamUserCouponGroup    = "g1"                 // 用户优惠券消息队列分组
	StreamUserCouponConsumer = "c1"                 // 用户优惠券消息队列消费者
	StreamUserCouponOnceRead = 1                    // 用户优惠券消息队列一次读取数量
	StreamUserCouponBlock    = 2000                 // 用户优惠券消息队列阻塞时间(单位:毫秒)
	StreamReadLatest         = ">"                  // 读取最新消息
	StreamReadPendingList    = "0"                  // 读取待处理消息
)

var (
	LuaSeckillScript string
	seckillLoadOnce  sync.Once
)

func LoadLuaSeckillScript() error {
	var err error
	seckillLoadOnce.Do(func() {
		var script []byte
		script, err = os.ReadFile("hack/seckill.lua")
		if err == nil {
			LuaSeckillScript = string(script)
		}
	})
	return err
}
