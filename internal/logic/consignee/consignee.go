package consignee

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

type sConsignee struct{}

func init() {
	service.RegisterConsignee(New())
}

func New() *sConsignee {
	return &sConsignee{}
}

// GetListFrontend 查询收货人列表
func (s *sConsignee) GetListFrontend(ctx context.Context, in model.ConsigneeGetListInput) (out *model.ConsigneeGetListOutput, err error) {
	var (
		m = dao.ConsigneeInfo.Ctx(ctx)
	)
	out = &model.ConsigneeGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	m = m.Where(dao.ConsigneeInfo.Columns().UserId, userId)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ConsigneeInfo
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

// GetListBackend 查询文章列表
func (s *sConsignee) GetListBackend(ctx context.Context, in model.ConsigneeGetListBackendInput) (out *model.ConsigneeGetListOutput, err error) {
	var (
		m = dao.ConsigneeInfo.Ctx(ctx)
	)
	out = &model.ConsigneeGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	queryMap := g.Map{}
	if in.Name != "" {
		queryMap[dao.ConsigneeInfo.Columns().Name] = in.Name
	}
	if in.Phone != "" {
		queryMap[dao.ConsigneeInfo.Columns().Phone] = in.Phone
	}

	m = m.Where(queryMap)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ConsigneeInfo
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

// UpdateBackend 更新收货人
func (s *sConsignee) UpdateBackend(ctx context.Context, in model.ConsigneeUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.ConsigneeInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.ConsigneeInfo.Columns().Id, in.Id).Update()
	return err
}

// AddBackend 添加收货人
func (s *sConsignee) AddBackend(ctx context.Context, in model.ConsigneeAddInput) (out *model.ConsigneeAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.ConsigneeInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.ConsigneeAddOutput{ConsigneeId: int(lastInsertID)}, err
}

// DeleteBackend 删除收货人
func (s *sConsignee) DeleteBackend(ctx context.Context, id int) error {
	return dao.ConsigneeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除收货人
		_, err := dao.ConsigneeInfo.Ctx(ctx).Where(g.Map{
			dao.ConsigneeInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// AddFrontend 添加收货人
func (s *sConsignee) AddFrontend(ctx context.Context, in model.ConsigneeAddInput) (out *model.ConsigneeAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	var lastInsertID int64
	err = dao.ConsigneeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if in.IsDefault == consts.ConsigneeDefault {
			err := s.UnsetDefault(ctx)
			if err != nil {
				return err
			}
		}
		lastInsertID, err = dao.ConsigneeInfo.Ctx(ctx).Data(in).InsertAndGetId()
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &model.ConsigneeAddOutput{ConsigneeId: int(lastInsertID)}, nil
}

// DeleteFrontend 删除收货人
func (s *sConsignee) DeleteFrontend(ctx context.Context, id int) error {
	var consigneeInfo entity.ConsigneeInfo
	err := dao.ConsigneeInfo.Ctx(ctx).Where(dao.ConsigneeInfo.Columns().Id, id).Scan(&consigneeInfo)
	if err != nil {
		return err
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	if userId != consigneeInfo.UserId {
		return gerror.New(consts.ErrNoPermission)
	}

	return dao.ConsigneeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除收货人
		_, err := dao.ConsigneeInfo.Ctx(ctx).Where(g.Map{
			dao.ConsigneeInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// UpdateFrontend 更新收货人
func (s *sConsignee) UpdateFrontend(ctx context.Context, in model.ConsigneeUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	return dao.ConsigneeInfo.Transaction(ctx, func(ctx context.Context, gdb gdb.TX) error {
		if in.IsDefault == consts.ConsigneeDefault {
			err := s.UnsetDefault(ctx)
			if err != nil {
				return err
			}
		}
		_, err := dao.ConsigneeInfo.Ctx(ctx).Data(in).Where(dao.ConsigneeInfo.Columns().Id, in.Id).Update()
		if err != nil {
			return err
		}
		return err
	})
}

func (s *sConsignee) UnsetDefault(ctx context.Context) error {
	// 取消默认
	_, err := dao.ConsigneeInfo.Ctx(ctx).Data(g.Map{
		dao.ConsigneeInfo.Columns().IsDefault: consts.ConsigneeNotDefault,
	}).Where(g.Map{
		dao.ConsigneeInfo.Columns().UserId:    gconv.Int(ctx.Value(consts.CtxUserId)),
		dao.ConsigneeInfo.Columns().IsDefault: consts.ConsigneeDefault,
	}).Update()
	if err != nil {
		return err
	}
	return err
}
