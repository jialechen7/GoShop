package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CategoryGetListCommonReq struct {
	g.Meta `path:"/category/level/list" tags:"分类后台" method:"get" summary:"分类列表接口"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CategoryAddReq struct {
	g.Meta   `path:"/category/add" tags:"分类后台" method:"post" summary:"创建分类接口"`
	Name     string `v:"required#请输入分类名称" json:"name" form:"name" description:"分类名称"`
	PicUrl   string `json:"pic_url" form:"pic_url" description:"分类图片"`
	ParentId int    `json:"parent_id" v:"required#请选择父级分类" form:"parent_id" description:"父级分类"`
	Level    int    `json:"level" v:"required#请选择分类级别" form:"level" description:"分类级别"`
	Sort     int    `json:"sort" form:"sort" description:"排序"`
}

type CategoryAddRes struct {
	CategoryId int `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"分类后台" summary:"删除分类接口"`
	Id     int `v:"min:1#请选择需要删除的分类" dc:"分类id"`
}
type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" method:"post" tags:"分类后台" summary:"修改分类接口"`
}
type CategoryUpdateRes struct{}
