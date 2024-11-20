package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminGetListCommonReq struct {
	g.Meta `path:"/admin/list" tags:"管理员后台" method:"get" summary:"管理员列表接口"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AdminReq struct {
	g.Meta   `path:"/admin/add" tags:"管理员后台" method:"post" summary:"创建管理员接口"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_id" dc:"角色ids"`
	IsAdmin  int    `json:"is_admin" dc:"是否是管理员"`
}
type AdminRes struct {
	AdminId int `json:"adminId"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/admin/delete" method:"delete" tags:"管理员后台" summary:"删除管理员接口"`
	Id     int `v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/admin/update" method:"post" tags:"管理员后台" summary:"修改管理员接口"`
	Id       int    `json:"id"      v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name" dc:"用户名"`
	Password string `json:"password" dc:"密码"`
	RoleIds  string `json:"role_id" dc:"角色ids"`
	IsAdmin  int    `json:"is_admin" dc:"是否是管理员"`
}
type AdminUpdateRes struct{}

type AdminGetInfoReq struct {
	g.Meta `path:"/admin/info" method:"get" tags:"管理员后台" summary:"获取管理员信息接口"`
}

//// AdminGetInfoRes 用于JWT
//type AdminGetInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// AdminGetInfoRes 用于gtoken
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}

type AdminUpdatePasswordReq struct {
	g.Meta   `path:"/admin/update/my/password" method:"post" tags:"管理员后台" summary:"修改管理员密码接口"`
	Password string `json:"password" dc:"密码"`
}

type AdminUpdatePasswordRes struct{}
