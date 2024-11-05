// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id           int         `json:"id"           orm:"id"            description:"用户ID"`
	Name         string      `json:"name"         orm:"name"          description:"用户名"`
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"头像URL"`
	Password     string      `json:"password"     orm:"password"      description:"加密后的密码"`
	UserSalt     string      `json:"userSalt"     orm:"user_salt"     description:"加密盐，用于生成密码"`
	Sex          int         `json:"sex"          orm:"sex"           description:"性别：1表示男，2表示女"`
	Status       int         `json:"status"       orm:"status"        description:"状态：1表示正常，2表示拉黑冻结"`
	Sign         string      `json:"sign"         orm:"sign"          description:"个性签名"`
	SecretAnswer string      `json:"secretAnswer" orm:"secret_answer" description:"密保问题答案"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`
}
