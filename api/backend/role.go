package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleCreateReq struct {
	g.Meta `path:"/role/add" method:"post" summary:"添加角色" tags:"角色"`
	Name   string `json:"name" v:"required#请输入角色名称"`
	Desc   string `json:"desc" v:"required#请输入角色描述"`
}

type RoleCreateRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" summary:"更新角色" tags:"角色"`
	Id     int    `json:"id" v:"required#请输入角色ID"`
	Name   string `json:"name" v:"required#请输入角色名称"`
	Desc   string `json:"desc" v:"required#请输入角色描述"`
}

type RoleUpdateRes struct {
	RoleId int `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" summary:"删除角色" tags:"角色"`
	Id     int `json:"id" v:"required#请输入角色ID"`
}

type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta `path:"/role/list" tags:"角色" method:"get" summary:"角色列表接口"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type RoleAddPermissionsReq struct {
	g.Meta        `path:"/role/add/permission" method:"post" summary:"添加角色权限" tags:"角色"`
	RoleId        int   `json:"role_id" v:"required#请输入角色ID"`
	PermissionIds []int `json:"permission_ids" v:"required#请输入权限ID"`
}

type RoleAddPermissionsRes struct{}

type RoleDeletePermissionsReq struct {
	g.Meta        `path:"/role/delete/permission" method:"delete" summary:"删除角色权限" tags:"角色"`
	RoleId        int   `json:"role_id" v:"required#请输入角色ID"`
	PermissionIds []int `json:"permission_ids" v:"required#请输入权限ID"`
}

type RoleDeletePermissionsRes struct{}
