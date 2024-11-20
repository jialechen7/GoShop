package user_coupon

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"github.com/gogf/gf/util/gconv"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

// GetList 查询优惠券列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var (
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.Where(dao.UserCouponInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxUserId)))
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.UserCouponInfo
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

// Add 添加优惠券
func (s *sUserCoupon) Add(ctx context.Context, in model.UserCouponAddInput) (out *model.UserCouponAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.UserCouponInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.UserCouponAddOutput{UserCouponId: int(lastInsertID)}, err
}

// Delete 删除优惠券
func (s *sUserCoupon) Delete(ctx context.Context, id int) error {
	return dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除优惠券
		_, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
			dao.UserCouponInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 更新优惠券
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	_, err := dao.UserCouponInfo.Ctx(ctx).Where(dao.UserCouponInfo.Columns().Id, in.Id).OmitEmpty().Data(in).Update()
	return err
}
