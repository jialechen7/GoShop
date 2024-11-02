// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleInfo is the golang structure of table role_info for DAO operations like Where/Data.
type RoleInfo struct {
	g.Meta    `orm:"table:role_info, do:true"`
	Id        interface{} // 角色ID
	Name      interface{} // 角色名称
	Desc      interface{} // 角色描述
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
