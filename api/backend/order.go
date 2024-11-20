package backend

import "github.com/gogf/gf/v2/frame/g"

type OrderGetListCommonReq struct {
	g.Meta `path:"/order/list" tags:"订单后台" method:"get" summary:"订单列表接口"`
	CommonPaginationReq
}

type OrderGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
