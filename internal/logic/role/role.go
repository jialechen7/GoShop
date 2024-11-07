package role

import (
	"context"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 添加角色
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(lastInsertID)}, err
}

// AddPermissions 添加角色权限
func (s *sRole) AddPermissions(ctx context.Context, in model.RoleAddPermissionsInput) (err error) {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, permissionId := range in.PermissionIds {
			_, err = dao.RolePermissionInfo.Ctx(ctx).Data(g.Map{
				dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
				dao.RolePermissionInfo.Columns().PermissionId: permissionId,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// DeletePermission 删除角色权限
func (s *sRole) DeletePermissions(ctx context.Context, in model.RoleDeletePermissionsInput) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, permissionId := range in.PermissionIds {
			_, err := dao.RolePermissionInfo.Ctx(ctx).Where(g.Map{
				dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
				dao.RolePermissionInfo.Columns().PermissionId: permissionId,
			}).Unscoped().Delete()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, id int) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除角色
		_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
			dao.RoleInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改角色
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		// 执行更行
		_, err := dao.RoleInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RoleInfo.Columns().Id).
			Where(dao.RoleInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询角色列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.RoleInfo
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
