package frontend

import "github.com/gogf/gf/v2/frame/g"

type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" tags:"Rotation" method:"get" summary:"轮播图列表接口"`
	Sort   int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
