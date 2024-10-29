package admin

import (
	"context"
	"goshop/internal/model/entity"
	"goshop/utility"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"

	"goshop/internal/dao"
	"goshop/internal/model"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// GetList 查询管理员列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.AdminInfo
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

// Create 创建管理员
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	// 生成用户盐并加密密码
	in.UserSalt = grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, in.UserSalt)

	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertID)}, err
}

// Delete 删除管理员
func (s *sAdmin) Delete(ctx context.Context, id int) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除管理员
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改管理员
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		// 判断是否修改了密码
		if in.Password != "" {
			in.UserSalt = grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, in.UserSalt)
		}

		// 执行更行
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
