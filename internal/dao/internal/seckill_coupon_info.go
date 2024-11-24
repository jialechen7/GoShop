// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SeckillCouponInfoDao is the data access object for table seckill_coupon_info.
type SeckillCouponInfoDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SeckillCouponInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SeckillCouponInfoColumns defines and stores column names for table seckill_coupon_info.
type SeckillCouponInfoColumns struct {
	Id        string // 秒杀优惠券id
	CouponId  string // 优惠券id
	Stock     string // 库存
	StartTime string // 开始时间
	EndTime   string // 结束时间
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// seckillCouponInfoColumns holds the columns for table seckill_coupon_info.
var seckillCouponInfoColumns = SeckillCouponInfoColumns{
	Id:        "id",
	CouponId:  "coupon_id",
	Stock:     "stock",
	StartTime: "start_time",
	EndTime:   "end_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSeckillCouponInfoDao creates and returns a new DAO object for table data access.
func NewSeckillCouponInfoDao() *SeckillCouponInfoDao {
	return &SeckillCouponInfoDao{
		group:   "default",
		table:   "seckill_coupon_info",
		columns: seckillCouponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SeckillCouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SeckillCouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SeckillCouponInfoDao) Columns() SeckillCouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SeckillCouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SeckillCouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SeckillCouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
