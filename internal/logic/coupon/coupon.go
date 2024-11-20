package coupon

import (
	"context"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

// GetList 查询优惠券列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var (
		m = dao.CouponInfo.Ctx(ctx)
	)
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.OrderDesc(dao.CouponInfo.Columns().Price)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CouponInfo
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
func (s *sCoupon) Add(ctx context.Context, in model.CouponAddInput) (out *model.CouponAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CouponInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.CouponAddOutput{CouponId: int(lastInsertID)}, err
}

// Delete 删除优惠券
func (s *sCoupon) Delete(ctx context.Context, id int) error {
	return dao.CouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除优惠券
		_, err := dao.CouponInfo.Ctx(ctx).Where(g.Map{
			dao.CouponInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 更新优惠券
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.CouponInfo.Ctx(ctx).Where(dao.CouponInfo.Columns().Id, in.Id).OmitEmpty().Data(in).Update()
	return err
}
