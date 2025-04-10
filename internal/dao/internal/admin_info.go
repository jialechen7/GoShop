// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminInfoDao is the data access object for table admin_info.
type AdminInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AdminInfoColumns defines and stores column names for table admin_info.
type AdminInfoColumns struct {
	Id           string //
	Name         string // 用户名
	Password     string // 密码
	RoleIds      string // 角色ids
	UserSalt     string // 加密盐
	IsAdmin      string // 是否超级管理员
	GithubOpenid string // github openid
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// adminInfoColumns holds the columns for table admin_info.
var adminInfoColumns = AdminInfoColumns{
	Id:           "id",
	Name:         "name",
	Password:     "password",
	RoleIds:      "role_ids",
	UserSalt:     "user_salt",
	IsAdmin:      "is_admin",
	GithubOpenid: "github_openid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewAdminInfoDao creates and returns a new DAO object for table data access.
func NewAdminInfoDao() *AdminInfoDao {
	return &AdminInfoDao{
		group:   "default",
		table:   "admin_info",
		columns: adminInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminInfoDao) Columns() AdminInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
