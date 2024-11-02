// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure for table permission_info.
type PermissionInfo struct {
	Id        int         `json:"id"        orm:"id"         description:"权限ID"`
	Name      string      `json:"name"      orm:"name"       description:"权限名称"`
	Path      string      `json:"path"      orm:"path"       description:"权限路径，指向具体的API或页面"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间，用于软删除"`
}
