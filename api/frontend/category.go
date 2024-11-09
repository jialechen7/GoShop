package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CategoryGetListCommonReq struct {
	g.Meta   `path:"/category/list" tags:"Category" method:"get" summary:"分类列表接口"`
	ParentId int `json:"parent_id" in:"query" v:"required#请选择父级ID"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CategoryAddReq struct {
	g.Meta `path:"/category/add" tags:"Category" method:"post" summary:"创建分类接口"`
}

type CategoryAddRes struct {
	CategoryId int `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"Category" summary:"删除分类接口"`
	Id     int `v:"min:1#请选择需要删除的分类" dc:"分类id"`
}
type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" method:"post" tags:"分类" summary:"修改分类接口"`
}
type CategoryUpdateRes struct{}
