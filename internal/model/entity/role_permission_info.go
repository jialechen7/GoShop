// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure for table role_permission_info.
type RolePermissionInfo struct {
	Id           int         `json:"id"           orm:"id"            description:"ID"`
	RoleId       int         `json:"roleId"       orm:"role_id"       description:"角色ID"`
	PermissionId int         `json:"permissionId" orm:"permission_id" description:"权限ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`
}
