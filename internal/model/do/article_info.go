// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleInfo is the golang structure of table article_info for DAO operations like Where/Data.
type ArticleInfo struct {
	g.Meta     `orm:"table:article_info, do:true"`
	Id         interface{} //
	UserId     interface{} // 作者id
	Title      interface{} // 标题
	Desc       interface{} // 摘要
	PicUrl     interface{} // 封面图
	IsAdmin    interface{} // 1后台管理员发布 2前台用户发布
	Praise     interface{} // 点赞数
	Collection interface{} // 收藏数
	Detail     interface{} // 文章详情
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
