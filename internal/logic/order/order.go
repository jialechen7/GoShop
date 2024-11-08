package order

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"

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
