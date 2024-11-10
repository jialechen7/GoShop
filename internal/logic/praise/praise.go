package praise

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

// GetListFrontend 查询点赞列表（仅用户发表的点赞）
func (s *sPraise) GetListFrontend(ctx context.Context, in model.PraiseGetListInput) (out *model.PraiseGetListOutput, err error) {
	var (
		m = dao.PraiseInfo.Ctx(ctx)
	)
	out = &model.PraiseGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	queryMap := g.Map{
		dao.PraiseInfo.Columns().UserId: gconv.Int(ctx.Value(consts.CtxUserId)),
		dao.PraiseInfo.Columns().Type:   in.Type,
	}
	listModel := m.Where(queryMap).Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PraiseInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// AddFrontend 添加点赞
func (s *sPraise) AddFrontend(ctx context.Context, in model.PraiseAddInput) (out *model.PraiseAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.PraiseAddOutput{PraiseId: int(lastInsertID)}, err
}

// DeleteFrontend 删除点赞
func (s *sPraise) DeleteFrontend(ctx context.Context, id int) error {
	return dao.PraiseInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除点赞
		_, err := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
			dao.PraiseInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}
