package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionGetListCommonReq struct {
	g.Meta `path:"/permission/list" tags:"权限后台" method:"get" summary:"权限列表接口"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type PermissionReq struct {
	g.Meta `path:"/permission/add" tags:"权限后台" method:"post" summary:"创建权限接口"`
	Name   string `json:"name" v:"required#权限名不能为空" dc:"权限名"`
	Path   string `json:"path" v:"required#路径不能为空" dc:"路径"`
}
type PermissionRes struct {
	PermissionId int `json:"permissionId"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete" method:"delete" tags:"权限后台" summary:"删除权限接口"`
	Id     int `v:"min:1#请选择需要删除的权限" dc:"权限id"`
}
type PermissionDeleteRes struct{}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update" method:"post" tags:"权限后台" summary:"修改权限接口"`
	Id     int    `json:"id"      v:"min:1#请选择需要修改的权限" dc:"权限Id"`
	Name   string `json:"name" v:"required#权限名不能为空" dc:"权限"`
	Path   string `json:"path" v:"required#路径不能为空" dc:"路径"`
}
type PermissionUpdateRes struct{}
