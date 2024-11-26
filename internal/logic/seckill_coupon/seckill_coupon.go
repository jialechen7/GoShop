package coupon

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"goshop/utility/redis_lock"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sSeckillCoupon struct{}

func init() {
	service.RegisterSeckillCoupon(New())
}

func New() *sSeckillCoupon {
	return &sSeckillCoupon{}
}

// GetList 查询秒杀优惠券列表
func (s *sSeckillCoupon) GetList(ctx context.Context, in model.SeckillCouponGetListInput) (out *model.SeckillCouponGetListOutput, err error) {
	var (
		m = dao.SeckillCouponInfo.Ctx(ctx)
	)
	out = &model.SeckillCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.OrderAsc(dao.SeckillCouponInfo.Columns().StartTime)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.SeckillCouponInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// Add 添加秒杀优惠券
func (s *sSeckillCoupon) Add(ctx context.Context, in model.SeckillCouponAddInput) (out *model.SeckillCouponAddOutput, err error) {
	couponAddInput := model.CouponAddInput{}
	couponAddInput.Name = in.Name
	couponAddInput.Condition = in.Condition
	couponAddInput.Price = in.Price
	couponAddInput.GoodsIds = in.GoodsIds
	couponAddInput.CategoryId = in.CategoryId
	couponAddInput.Type = consts.CouponTypeSeckill
	couponId, err := dao.CouponInfo.Ctx(ctx).OmitEmpty().Data(couponAddInput).InsertAndGetId()
	if err != nil {
		return out, err
	}
	in.SeckillCouponCreateUpdateBase.CouponId = int(couponId)
	lastInsertID, err := dao.SeckillCouponInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.SeckillCouponAddOutput{SeckillCouponId: int(lastInsertID)}, err
}

// Delete 删除秒杀优惠券
func (s *sSeckillCoupon) Delete(ctx context.Context, id int) error {
	return dao.SeckillCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除秒杀优惠券
		_, err := dao.SeckillCouponInfo.Ctx(ctx).Where(g.Map{
			dao.SeckillCouponInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 更新秒杀优惠券
func (s *sSeckillCoupon) Update(ctx context.Context, in model.SeckillCouponUpdateInput) error {
	_, err := dao.SeckillCouponInfo.Ctx(ctx).Where(dao.SeckillCouponInfo.Columns().Id, in.Id).OmitEmpty().Data(in).Update()
	return err
}

// Kill 用户秒杀优惠券
func (s *sSeckillCoupon) Kill(ctx context.Context, coupon_id int) error {
	seckillCouponEntity := &entity.SeckillCouponInfo{}
	// 1. 根据id查到秒杀优惠券的信息
	if err := dao.SeckillCouponInfo.Ctx(ctx).Where(dao.SeckillCouponInfo.Columns().CouponId, coupon_id).Scan(seckillCouponEntity); err != nil {
		return err
	}
	// 2. 判断秒杀是否开始
	if gtime.Now().Before(seckillCouponEntity.StartTime) {
		return gerror.New(consts.ErrSeckillNotStart)
	}
	// 3. 判断秒杀是否结束
	if gtime.Now().After(seckillCouponEntity.EndTime) {
		return gerror.New(consts.ErrSeckillEnd)
	}
	// 4. 判断库存是否足够
	if seckillCouponEntity.Stock < 1 {
		return gerror.New(consts.ErrStock)
	}
	// 5. 一人一单
	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	// 5.1 分布式锁实现一人一单
	mutex := redis_lock.Rs.NewMutex(consts.RedisLockKey+consts.UserCouponIdKey+gconv.String(userId), redsync.WithExpiry(1200*time.Second))
	if err := mutex.Lock(); err != nil {
		return gerror.New(consts.ErrHasSeckill)
	}
	defer mutex.Unlock()
	// 5.2 判断是否已经秒杀过
	count, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
		dao.UserCouponInfo.Columns().UserId:   userId,
		dao.UserCouponInfo.Columns().CouponId: coupon_id,
	}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(consts.ErrHasSeckill)
	}
	err = dao.SeckillCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 6. 扣减库存
		result, err := dao.SeckillCouponInfo.Ctx(ctx).Where(dao.SeckillCouponInfo.Columns().CouponId, coupon_id).
			WhereGT(dao.SeckillCouponInfo.Columns().Stock, 0).Decrement(dao.SeckillCouponInfo.Columns().Stock, 1)
		if err != nil {
			return err
		}
		// 6.1 判断是否扣减成功
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return gerror.New(consts.ErrDecreaseStock)
		}
		// 7. 添加到user_coupon_info表中，表示秒杀成功
		_, err = dao.UserCouponInfo.Ctx(ctx).Insert(g.Map{
			dao.UserCouponInfo.Columns().UserId:   gconv.Int(ctx.Value(consts.CtxUserId)),
			dao.UserCouponInfo.Columns().CouponId: coupon_id,
			dao.UserCouponInfo.Columns().Status:   consts.CouponStatusAvailable,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
