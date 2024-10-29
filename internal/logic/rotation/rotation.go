package rotation

import (
	"context"
	"goshop/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"

	"goshop/internal/dao"
	"goshop/internal/model"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// GetList 查询轮播图列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	var (
		m = dao.RotationInfo.Ctx(ctx)
	)
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.RotationInfo.Columns().Sort)

	// 执行查询
	var list []*entity.RotationInfo
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

	// 区分Scan和ScanList用法
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// Create 创建轮播图
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

// Delete 删除轮播图
func (s *sRotation) Delete(ctx context.Context, id int) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除轮播图
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改轮播图
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
