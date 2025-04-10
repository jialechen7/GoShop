// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleInfoDao is the data access object for table article_info.
type ArticleInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ArticleInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleInfoColumns defines and stores column names for table article_info.
type ArticleInfoColumns struct {
	Id         string //
	UserId     string // 作者id
	Title      string // 标题
	Desc       string // 摘要
	PicUrl     string // 封面图
	IsAdmin    string // 1后台管理员发布 2前台用户发布
	Praise     string // 点赞数
	Collection string // 收藏数
	Detail     string // 文章详情
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

// articleInfoColumns holds the columns for table article_info.
var articleInfoColumns = ArticleInfoColumns{
	Id:         "id",
	UserId:     "user_id",
	Title:      "title",
	Desc:       "desc",
	PicUrl:     "pic_url",
	IsAdmin:    "is_admin",
	Praise:     "praise",
	Collection: "collection",
	Detail:     "detail",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewArticleInfoDao creates and returns a new DAO object for table data access.
func NewArticleInfoDao() *ArticleInfoDao {
	return &ArticleInfoDao{
		group:   "default",
		table:   "article_info",
		columns: articleInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleInfoDao) Columns() ArticleInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
