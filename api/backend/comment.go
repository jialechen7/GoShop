package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommentGetListCommonReq struct {
	g.Meta `path:"/comment/list" tags:"评论后台" method:"get" summary:"评论列表接口"`
	CommonPaginationReq
}

type CommentGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CommentDeleteReq struct {
	g.Meta `path:"/comment/delete" method:"delete" tags:"评论后台" summary:"删除评论接口"`
	Id     int `v:"min:1#请选择需要删除的评论" dc:"评论id"`
}
type CommentDeleteRes struct{}

type CommentUpdateReq struct {
	g.Meta `path:"/comment/update" method:"post" tags:"评论后台" summary:"修改评论接口"`
}
type CommentUpdateRes struct{}
