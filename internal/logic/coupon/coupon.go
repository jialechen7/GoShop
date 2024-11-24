package coupon

import (
	"context"
	"fmt"
	"goshop/api/frontend"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"slices"
	"strings"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

// GetListBackend 查询优惠券列表
func (s *sCoupon) GetListBackend(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
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

// GetListFrontend 查询可抢优惠券列表
func (s *sCoupon) GetListFrontend(ctx context.Context, in model.CouponGetListWithGoodsIdInput) (out *model.CouponGetListWithGoodsIdOutput, err error) {
	fields := []string{"c.*", "sc.stock", "sc.start_time", "sc.end_time"}
	m := g.Model(dao.CouponInfo.Table(), "c")
	m = m.LeftJoin(dao.SeckillCouponInfo.Table(), "sc", "c.id=sc.coupon_id").
		Where(m.Builder().Where(fmt.Sprintf("FIND_IN_SET(%s, c.goods_ids) != 0", gconv.String(in.GoodsId))).WhereOr("c.goods_ids=''")).
		Where(m.Builder().Where("c.type=1 and sc.end_time>now()").WhereOr("c.type=0")).
		OrderDesc("c.type").
		OrderAsc("sc.start_time").
		Fields(fields)
	out = &model.CouponGetListWithGoodsIdOutput{}
	if err = m.Scan(&out.List); err != nil {
		return nil, err
	}
	return
}

// GetListAvailable 获取可用优惠券列表
func (s *sCoupon) GetListAvailable(ctx context.Context, in model.CouponGetListAvailableInput) (out *model.CouponGetListAvailableOutput, err error) {
	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	// 获取用户所有的优惠券
	m := dao.UserCouponInfo.Ctx(ctx).With(model.CouponInfo{}).With(model.SeckillCouponInfo{})
	m = m.Where(dao.UserCouponInfo.Columns().UserId, userId)
	userCouponList := &model.UserCouponGetListAllOutput{}
	if err = m.Scan(&userCouponList.List); err != nil {
		return nil, err
	}
	out = &model.CouponGetListAvailableOutput{}
	for _, userCoupon := range userCouponList.List {
		coupon := userCoupon.CouponInfo
		if coupon == nil {
			continue
		}
		goodsIds := make([]int, 0)
		if coupon.GoodsIds != "" {
			goodsIds = gconv.SliceInt(strings.Split(coupon.GoodsIds, ","))
		}
		sum := 0
		for _, orderGoodsInfo := range in.OrderGoodsInfos.([]*frontend.OrderGoodsInfo) {
			if slices.Contains(goodsIds, orderGoodsInfo.GoodsId) {
				sum += orderGoodsInfo.Price * orderGoodsInfo.Count
			}
		}

		item := model.CouponGetListOutputItem{
			Id:         coupon.Id,
			Name:       coupon.Name,
			Condition:  coupon.Condition,
			Price:      coupon.Price,
			GoodsIds:   coupon.GoodsIds,
			CategoryId: coupon.CategoryId,
			Type:       coupon.Type,
			TimeCommon: model.TimeCommon{
				CreatedAt: coupon.CreatedAt,
				UpdatedAt: coupon.UpdatedAt,
			},
		}
		if coupon.Type == consts.CouponTypeSeckill && coupon.SeckillCouponInfo != nil {
			item.Stock = coupon.SeckillCouponInfo.Stock
			item.StartTime = coupon.SeckillCouponInfo.StartTime
			item.EndTime = coupon.SeckillCouponInfo.EndTime
		}
		if sum >= coupon.Condition && userCoupon.Status == consts.CouponStatusAvailable {
			out.AvailableList = append(out.AvailableList, item)
		} else {
			if userCoupon.Status == consts.CouponStatusUsed {
				item.Reason = consts.CouponStatusUsedText
			} else if userCoupon.Status == consts.CouponStatusExpired {
				item.Reason = consts.CouponStatusExpiredText
			} else if sum < coupon.Condition {
				item.Reason = fmt.Sprintf("未达到满减条件：%s元，差%s元", gconv.String(coupon.Condition/100.0), gconv.String((coupon.Condition-sum)/100.0))
			}
			out.UnavailableList = append(out.UnavailableList, item)
		}
	}
	return
}

// Add 添加优惠券
func (s *sCoupon) Add(ctx context.Context, in model.CouponAddInput) (out *model.CouponAddOutput, err error) {
	in.Type = consts.CouponTypeCommon
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
