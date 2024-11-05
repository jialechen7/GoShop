// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfoDao is the data access object for table user_info.
type UserInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserInfoColumns defines and stores column names for table user_info.
type UserInfoColumns struct {
	Id           string // 用户ID
	Name         string // 用户名
	Avatar       string // 头像URL
	Password     string // 加密后的密码
	UserSalt     string // 加密盐，用于生成密码
	Sex          string // 性别：1表示男，2表示女
	Status       string // 状态：1表示正常，2表示拉黑冻结
	Sign         string // 个性签名
	SecretAnswer string // 密保问题答案
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// userInfoColumns holds the columns for table user_info.
var userInfoColumns = UserInfoColumns{
	Id:           "id",
	Name:         "name",
	Avatar:       "avatar",
	Password:     "password",
	UserSalt:     "user_salt",
	Sex:          "sex",
	Status:       "status",
	Sign:         "sign",
	SecretAnswer: "secret_answer",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewUserInfoDao creates and returns a new DAO object for table data access.
func NewUserInfoDao() *UserInfoDao {
	return &UserInfoDao{
		group:   "default",
		table:   "user_info",
		columns: userInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserInfoDao) Columns() UserInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
