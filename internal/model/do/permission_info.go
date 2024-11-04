// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure of table permission_info for DAO operations like Where/Data.
type PermissionInfo struct {
	g.Meta    `orm:"table:permission_info, do:true"`
	Id        interface{} // 权限ID
	Name      interface{} // 权限名称
	Path      interface{} // 权限路径，指向具体的API或页面
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
