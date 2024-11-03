package permission

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

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// GetList 查询权限列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PermissionInfo
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

// Create 创建权限
func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: int(lastInsertID)}, err
}

// Delete 删除权限
func (s *sPermission) Delete(ctx context.Context, id int) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除权限
		_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
			dao.PermissionInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改权限
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		// 执行更行
		_, err := dao.PermissionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PermissionInfo.Columns().Id).
			Where(dao.PermissionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
