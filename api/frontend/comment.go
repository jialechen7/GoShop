package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommentGetListCommonReq struct {
	g.Meta `path:"/comment/list" tags:"Comment" method:"get" summary:"评论列表接口"`
	CommonPaginationReq
	Type     int `json:"type" query:"type" v:"required#请选择评论类型" dc:"评论类型"`
	ObjectId int `json:"object_id" query:"object_id" v:"required#请选择评论ID" dc:"评论对象id"`
}

type CommentGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CommentAddReq struct {
	g.Meta   `path:"/comment/add" tags:"Comment" method:"post" summary:"创建评论接口"`
	ParentId int    `json:"parent_id" form:"parent_id" v:"required#请选择父级评论" dc:"父级评论id"`
	ObjectId int    `json:"object_id" form:"object_id" v:"required#请选择评论对象" dc:"评论对象id"`
	Content  string `json:"content" form:"content" v:"required#请输入评论内容" dc:"评论内容"`
	Type     int    `json:"type" form:"type" v:"required#请选择评论类型" dc:"评论类型"`
}

type CommentAddRes struct {
	CommentId int `json:"comment_id"`
}

type CommentDeleteReq struct {
	g.Meta `path:"/comment/delete" method:"delete" tags:"Comment" summary:"删除评论接口"`
	Id     int `v:"min:1#请选择需要删除的评论" dc:"评论id"`
}
type CommentDeleteRes struct{}

type CommentUpdateReq struct {
	g.Meta `path:"/comment/update" method:"post" tags:"评论" summary:"修改评论接口"`
}
type CommentUpdateRes struct{}
