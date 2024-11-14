package goods_options

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

type sGoodsOptions struct{}

func init() {
	service.RegisterGoodsOptions(New())
}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

// GetListBackend 查询商品规格列表
func (s *sGoodsOptions) GetListBackend(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var (
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsOptionsInfo
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

// UpdateBackend 更新商品规格
func (s *sGoodsOptions) UpdateBackend(ctx context.Context, in model.GoodsOptionsUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).Update()
	return err
}

// AddBackend 添加商品规格
func (s *sGoodsOptions) AddBackend(ctx context.Context, in model.GoodsOptionsAddInput) (out *model.GoodsOptionsAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.GoodsOptionsAddOutput{GoodsOptionsId: int(lastInsertID)}, err
}

// DeleteBackend 删除商品规格
func (s *sGoodsOptions) DeleteBackend(ctx context.Context, id int) error {
	return dao.GoodsOptionsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除商品规格
		_, err := dao.GoodsOptionsInfo.Ctx(ctx).Where(g.Map{
			dao.GoodsOptionsInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// GetListFrontend 查询商品规格列表
func (s *sGoodsOptions) GetListFrontend(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var (
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.Where(dao.GoodsOptionsInfo.Columns().GoodsId, in.GoodsId)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsOptionsInfo
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

// AddFrontend 添加商品规格
func (s *sGoodsOptions) AddFrontend(ctx context.Context, in model.GoodsOptionsAddInput) (out *model.GoodsOptionsAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.GoodsOptionsAddOutput{GoodsOptionsId: int(lastInsertID)}, err
}

// DetailFrontend 查询商品规格详情
func (s *sGoodsOptions) DetailFrontend(ctx context.Context, id int) (out *model.GoodsOptionsDetailOutput, err error) {
	var goods_optionsInfo entity.GoodsOptionsInfo
	err = dao.GoodsOptionsInfo.Ctx(ctx).Where(dao.GoodsOptionsInfo.Columns().Id, id).Scan(&goods_optionsInfo)
	if err != nil {
		return nil, err
	}

	return &model.GoodsOptionsDetailOutput{
		Id:      goods_optionsInfo.Id,
		GoodsId: goods_optionsInfo.GoodsId,
		PicUrl:  goods_optionsInfo.PicUrl,
		Name:    goods_optionsInfo.Name,
		Price:   goods_optionsInfo.Price,
		Stock:   goods_optionsInfo.Stock,
		TimeCommon: model.TimeCommon{
			CreatedAt: goods_optionsInfo.CreatedAt,
			UpdatedAt: goods_optionsInfo.UpdatedAt,
		},
	}, nil
}
