package category

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

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// GetList 查询分类列表（管理员查询全部）
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CategoryInfo
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

// GetListFrontend 查询分类列表
func (s *sCategory) GetListFrontend(ctx context.Context, in model.CategoryGetListWithParentIdInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	queryMap := g.Map{
		dao.CategoryInfo.Columns().ParentId: in.ParentId,
	}
	listModel := m.Ctx(ctx).Where(queryMap).OrderDesc(dao.CategoryInfo.Columns().Sort).Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CategoryInfo
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

// Add 添加分类
func (s *sCategory) Add(ctx context.Context, in model.CategoryAddInput) (out *model.CategoryAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.CategoryAddOutput{CategoryId: int(lastInsertID)}, err
}

// Delete 删除分类
func (s *sCategory) Delete(ctx context.Context, id int) error {
	return dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除分类
		_, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
			dao.CategoryInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}
