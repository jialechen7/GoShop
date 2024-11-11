package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PraiseGetListCommonReq struct {
	g.Meta `path:"/praise/list" tags:"Praise" method:"get" summary:"点赞列表接口"`
	CommonPaginationReq
	Type int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
}

type PraiseGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type PraiseAddReq struct {
	g.Meta   `path:"/praise/add" tags:"Praise" method:"post" summary:"创建点赞接口"`
	Type     int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
	ObjectId int `json:"object_id" form:"object_id" v:"required#请选择对象" dc:"对象ID"`
}

type PraiseAddRes struct {
	PraiseId int `json:"praise_id"`
}

type PraiseDeleteReq struct {
	g.Meta `path:"/praise/delete" method:"delete" tags:"Praise" summary:"删除点赞接口"`
	Id     int `v:"min:1#请选择需要删除的点赞" dc:"点赞id"`
}
type PraiseDeleteRes struct{}

type PraiseDeleteByTypeReq struct {
	g.Meta   `path:"/praise/deleteByType" method:"delete" tags:"Praise" summary:"删除点赞接口"`
	Type     int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
	ObjectId int `json:"object_id" form:"object_id" v:"required#请选择对象" dc:"对象"`
}

type PraiseDeleteByTypeRes struct{}

type PraiseUpdateReq struct {
	g.Meta `path:"/praise/update" method:"post" tags:"点赞" summary:"修改点赞接口"`
}
type PraiseUpdateRes struct{}

type PraiseDetailReq struct {
	g.Meta `path:"/praise/detail" method:"get" tags:"点赞" summary:"点赞详情接口"`
	Id     int `v:"min:1#请选择需要查询的点赞" form:"query" dc:"点赞id"`
}

type PraiseDetailRes struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Detail    string `json:"detail"`
	PicUrl    string `json:"pic_url"`
	IsAdmin   int    `json:"is_admin"`
	Praise    int    `json:"praise"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
