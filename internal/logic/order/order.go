package order

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"goshop/utility"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"goshop/internal/dao"
	"goshop/internal/model"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

// GetList 查询订单列表
func (s *sOrder) GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error) {
	var (
		m = dao.OrderInfo.Ctx(ctx)
	)
	out = &model.OrderGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.OrderInfo
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

// GetListFrontend 查询订单列表（仅用户自己的订单）
func (s *sOrder) GetListFrontend(ctx context.Context, in model.OrderGetListWithStatusInput) (out *model.OrderGetListOutput, err error) {
	var (
		m = dao.OrderInfo.Ctx(ctx)
	)
	out = &model.OrderGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	queryMap := g.Map{
		dao.OrderInfo.Columns().UserId: userId,
	}
	if in.Status != 0 {
		queryMap[dao.OrderInfo.Columns().Status] = in.Status
	}
	listModel := m.Ctx(ctx).Where(queryMap).Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.OrderInfo
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

// AddFrontend 添加订单
func (s *sOrder) AddFrontend(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	var orderId int
	in.Number = utility.GetOrderNum()
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	in.PayAt = gtime.Now()
	err = dao.OrderInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 插入订单信息
		orderId, err := dao.OrderInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
		if err != nil {
			return err
		}

		// 2. 插入订单商品信息
		for _, goodsInfo := range in.OrderAddGoodsInfos {
			goodsInfo.OrderId = int(orderId)
			_, err = dao.OrderGoodsInfo.Ctx(ctx).OmitEmpty().Data(goodsInfo).Insert()
			if err != nil {
				return err
			}
		}

		// 3. 更新商品销量和库存
		for _, goodsInfo := range in.OrderAddGoodsInfos {
			var goodsEntity *entity.GoodsInfo
			err = dao.GoodsInfo.Ctx(ctx).WherePri(goodsInfo.GoodsId).Scan(&goodsEntity)
			if err != nil {
				return err
			}
			if goodsEntity.Stock < goodsInfo.Count {
				return gerror.New(consts.ErrStockNotEnough)
			}
			_, err = dao.GoodsInfo.Ctx(ctx).WherePri(goodsInfo.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, goodsInfo.Count)
			if err != nil {
				return err
			}
			_, err = dao.GoodsInfo.Ctx(ctx).WherePri(goodsInfo.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, goodsInfo.Count)
			if err != nil {
				return err
			}
			_, err = dao.GoodsOptionsInfo.Ctx(ctx).WherePri(goodsInfo.GoodsOptionsId).Decrement(dao.GoodsOptionsInfo.Columns().Stock, goodsInfo.Count)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return out, err
	}
	return &model.OrderAddOutput{OrderId: int(orderId)}, err
}
