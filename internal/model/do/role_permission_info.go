// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure of table role_permission_info for DAO operations like Where/Data.
type RolePermissionInfo struct {
	g.Meta       `orm:"table:role_permission_info, do:true"`
	Id           interface{} // ID
	RoleId       interface{} // 角色ID
	PermissionId interface{} // 权限ID
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
