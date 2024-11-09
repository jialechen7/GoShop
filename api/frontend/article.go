package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" tags:"Article" method:"get" summary:"文章列表接口"`
	CommonPaginationReq
}

type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type ArticleAddReq struct {
	g.Meta  `path:"/article/add" tags:"Article" method:"post" summary:"创建文章接口"`
	UserId  int    `json:"user_id" form:"user_id" v:"required#请选择用户" dc:"用户ID"`
	Title   string `json:"title" form:"title" v:"required#请输入标题" dc:"标题"`
	Desc    string `json:"desc" form:"desc" v:"required#请输入描述" dc:"描述/摘要"`
	PicUrl  string `json:"pic_url" form:"pic_url" v:"required#请上传图片" dc:"图片地址"`
	Detail  string `json:"detail" form:"detail" v:"required#请输入内容" dc:"内容"`
	IsAdmin int    `json:"is_admin" form:"is_admin" v:"required#请选择是否管理员" dc:"是否管理员"`
}

type ArticleAddRes struct {
	ArticleId int `json:"article_id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"Article" summary:"删除文章接口"`
	Id     int `v:"min:1#请选择需要删除的文章" dc:"文章id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update" method:"post" tags:"文章" summary:"修改文章接口"`
}
type ArticleUpdateRes struct{}
