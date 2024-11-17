package collection

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

// GetListFrontend 查询收藏列表（仅用户发表的收藏）
func (s *sCollection) GetListFrontend(ctx context.Context, in model.CollectionGetListInput) (out *model.CollectionGetListOutput, err error) {
	var (
		m = dao.CollectionInfo.Ctx(ctx)
	)
	out = &model.CollectionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	queryMap := g.Map{
		dao.CollectionInfo.Columns().UserId: gconv.Int(ctx.Value(consts.CtxUserId)),
		dao.CollectionInfo.Columns().Type:   in.Type,
	}
	m = m.Where(queryMap)
	if in.Type == consts.CollectionArticleType {
		m = m.With(model.ArticleInfo{})
	} else if in.Type == consts.CollectionGoodsType {
		m = m.With(model.GoodsInfo{})
	} else {
		m = m.WithAll()
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CollectionInfo
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
	g.Dump(out.List)
	return
}

// AddFrontend 添加收藏
func (s *sCollection) AddFrontend(ctx context.Context, in model.CollectionAddInput) (out *model.CollectionAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	var lastInsertID int64
	err = dao.CollectionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入收藏
		lastInsertID, err = dao.CollectionInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
		if err != nil {
			return err
		}

		// 更新文章收藏数
		if in.Type == consts.CollectionArticleType {
			_, err = dao.ArticleInfo.Ctx(ctx).Where(dao.ArticleInfo.Columns().Id, in.ObjectId).Increment(dao.ArticleInfo.Columns().Collection, 1)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &model.CollectionAddOutput{CollectionId: int(lastInsertID)}, err
}

// DeleteFrontend 删除收藏
func (s *sCollection) DeleteFrontend(ctx context.Context, id int) error {
	return dao.CollectionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除收藏
		_, err := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
			dao.CollectionInfo.Columns().Id: id,
		}).Unscoped().Delete()

		return err
	})
}

// DeleteByTypeFrontend 删除收藏（根据类型）
func (s *sCollection) DeleteByTypeFrontend(ctx context.Context, in model.CollectionDeleteByTypeInput) error {
	return dao.CollectionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除收藏
		_, err := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
			dao.CollectionInfo.Columns().UserId:   in.UserId,
			dao.CollectionInfo.Columns().Type:     in.Type,
			dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		}).Unscoped().Delete()

		// 更新文章收藏数
		_, err = dao.ArticleInfo.Ctx(ctx).Where(g.Map{
			dao.ArticleInfo.Columns().Id: in.ObjectId,
		}).Decrement(dao.ArticleInfo.Columns().Collection, 1)
		if err != nil {
			return err
		}
		return err
	})
}
