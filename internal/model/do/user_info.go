// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta       `orm:"table:user_info, do:true"`
	Id           interface{} // 用户ID
	Name         interface{} // 用户名
	Avatar       interface{} // 头像URL
	Password     interface{} // 加密后的密码
	UserSalt     interface{} // 加密盐，用于生成密码
	Sex          interface{} // 性别：1表示男，2表示女
	Status       interface{} // 状态：1表示正常，2表示拉黑冻结
	Sign         interface{} // 个性签名
	SecretAnswer interface{} // 密保问题答案
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
