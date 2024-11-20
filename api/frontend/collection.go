package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CollectionGetListCommonReq struct {
	g.Meta `path:"/collection/list" tags:"收藏前台" method:"get" summary:"收藏列表接口"`
	CommonPaginationReq
	Type int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
}

type CollectionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" tags:"收藏前台" method:"post" summary:"创建收藏接口"`
	Type     int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
	ObjectId int `json:"object_id" form:"object_id" v:"required#请选择对象" dc:"对象ID"`
}

type CollectionAddRes struct {
	CollectionId int `json:"collection_id"`
}

type CollectionDeleteReq struct {
	g.Meta `path:"/collection/delete" method:"delete" tags:"收藏前台" summary:"删除收藏接口"`
	Id     int `v:"min:1#请选择需要删除的收藏" dc:"收藏id"`
}
type CollectionDeleteRes struct{}

type CollectionDeleteByTypeReq struct {
	g.Meta   `path:"/collection/deleteByType" method:"delete" tags:"收藏前台" summary:"删除收藏接口"`
	Type     int `json:"type" form:"type" v:"required#请选择类型" dc:"类型"`
	ObjectId int `json:"object_id" form:"object_id" v:"required#请选择对象" dc:"对象"`
}

type CollectionDeleteByTypeRes struct{}

type CollectionUpdateReq struct {
	g.Meta `path:"/collection/update" method:"post" tags:"收藏前台" summary:"修改收藏接口"`
}
type CollectionUpdateRes struct{}

type CollectionDetailReq struct {
	g.Meta `path:"/collection/detail" method:"get" tags:"收藏前台" summary:"收藏详情接口"`
	Id     int `v:"min:1#请选择需要查询的收藏" form:"query" dc:"收藏id"`
}

type CollectionDetailRes struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Detail     string `json:"detail"`
	PicUrl     string `json:"pic_url"`
	IsAdmin    int    `json:"is_admin"`
	Collection int    `json:"collection"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
