// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleInfoDao is the data access object for table role_info.
type RoleInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RoleInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RoleInfoColumns defines and stores column names for table role_info.
type RoleInfoColumns struct {
	Id        string // 角色ID
	Name      string // 角色名称
	Desc      string // 角色描述
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// roleInfoColumns holds the columns for table role_info.
var roleInfoColumns = RoleInfoColumns{
	Id:        "id",
	Name:      "name",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRoleInfoDao creates and returns a new DAO object for table data access.
func NewRoleInfoDao() *RoleInfoDao {
	return &RoleInfoDao{
		group:   "default",
		table:   "role_info",
		columns: roleInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoleInfoDao) Columns() RoleInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
