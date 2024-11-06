package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserGetListCommonReq struct {
	g.Meta `path:"/user/list" tags:"User" method:"get" summary:"获取用户列表接口"`
	CommonPaginationReq
}

type UserGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type UserUpdateReq struct {
	g.Meta       `path:"/user/update" method:"post" tags:"用户" summary:"修改用户接口"`
	Id           int    `json:"id" v:"min:1#请选择需要修改的用户" dc:"用户Id"`
	Name         string `json:"name" dc:"用户名"`
	Password     string `json:"password" dc:"密码"`
	UserSalt     string `json:"user_salt" dc:"用户盐"`
	Avatar       string `json:"avatar" dc:"头像"`
	Sign         string `json:"sign" dc:"个性签名"`
	Sex          int    `json:"sex" dc:"性别"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题答案"`
	Status       int    `json:"status" dc:"用户状态"`
}

type UserUpdateRes struct{}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"User" summary:"删除用户接口"`
	Id     int `v:"min:1#请选择需要删除的用户" dc:"用户id"`
}
type UserDeleteRes struct{}
