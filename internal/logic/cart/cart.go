package cart

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
)

type sCart struct{}

func init() {
	service.RegisterCart(New())
}

func New() *sCart {
	return &sCart{}
}

// GetListFrontend 查询购物车列表
func (s *sCart) GetListFrontend(ctx context.Context, in model.CartGetListInput) (out *model.CartGetListOutput, err error) {
	var (
		m = dao.CartInfo.Ctx(ctx)
	)
	out = &model.CartGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.Where(dao.CartInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxUserId))).WithAll()
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CartInfo
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

// AddFrontend 添加购物车
func (s *sCart) AddFrontend(ctx context.Context, in model.CartAddInput) (out *model.CartAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	in.CartCreateUpdateBase.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.CartInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.CartAddOutput{CartId: int(lastInsertID)}, err
}

// DeleteFrontend 删除购物车
func (s *sCart) DeleteFrontend(ctx context.Context, ids []int) error {
	err := dao.CartInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, id := range ids {
			var cartInfo entity.CartInfo
			err := dao.CartInfo.Ctx(ctx).Where(dao.CartInfo.Columns().Id, id).Scan(&cartInfo)
			if err != nil {
				return err
			}

			userId := gconv.Int(ctx.Value(consts.CtxUserId))
			// 判断当前用户是否有权限对该购物车操作
			if userId != cartInfo.UserId {
				return gerror.New(consts.ErrNoPermission)
			}
			_, err = dao.CartInfo.Ctx(ctx).Where(g.Map{
				dao.CartInfo.Columns().Id: id,
			}).Unscoped().Delete()
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// UpdateFrontend 更新购物车
func (s *sCart) UpdateFrontend(ctx context.Context, in model.CartUpdateInput) (out model.CartUpdateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	in.CartCreateUpdateBase.UserId = userId
	// 判断当前用户是否有权限对该购物车操作
	var cartInfo entity.CartInfo
	err = dao.CartInfo.Ctx(ctx).Where(dao.CartInfo.Columns().Id, in.Id).Scan(&cartInfo)
	if err != nil {
		return out, err
	}
	if userId != cartInfo.UserId {
		return out, gerror.New(consts.ErrNoPermission)
	}

	_, err = dao.CartInfo.Ctx(ctx).Where(dao.CartInfo.Columns().Id, in.Id).Data(in).Update()
	return
}
