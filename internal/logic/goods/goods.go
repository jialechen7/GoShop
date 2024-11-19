package goods

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

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

// GetListBackend 查询商品列表
func (s *sGoods) GetListBackend(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsInfo
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

// UpdateBackend 更新商品
func (s *sGoods) UpdateBackend(ctx context.Context, in model.GoodsUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.GoodsInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.GoodsInfo.Columns().Id, in.Id).Update()
	return err
}

// AddBackend 添加商品
func (s *sGoods) AddBackend(ctx context.Context, in model.GoodsAddInput) (out *model.GoodsAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.GoodsAddOutput{GoodsId: int(lastInsertID)}, err
}

// DeleteBackend 删除商品
func (s *sGoods) DeleteBackend(ctx context.Context, id int) error {
	return dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除商品
		_, err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{
			dao.GoodsInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// GetListFrontend 查询商品列表
func (s *sGoods) GetListFrontend(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsInfo
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

// GetListByLevelFrontend 查询商品列表（2级分类）
func (s *sGoods) GetListByLevelFrontend(ctx context.Context, in model.GoodsGetListByLevelInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.Where(dao.GoodsInfo.Columns().Level2CategoryId, in.LevelId)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsInfo
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

// AddFrontend 添加商品
func (s *sGoods) AddFrontend(ctx context.Context, in model.GoodsAddInput) (out *model.GoodsAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.GoodsAddOutput{GoodsId: int(lastInsertID)}, err
}

// DetailFrontend 查询商品详情
func (s *sGoods) DetailFrontend(ctx context.Context, id int) (out *model.GoodsDetailOutput, err error) {
	var goodsInfo entity.GoodsInfo
	err = dao.GoodsInfo.Ctx(ctx).Where(dao.GoodsInfo.Columns().Id, id).Scan(&goodsInfo)
	if err != nil {
		return nil, err
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	goodsId := goodsInfo.Id

	// 判断当前用户是否对该商品点赞
	isPraise, _ := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		dao.PraiseInfo.Columns().Type:     consts.PraiseGoodsType,
		dao.PraiseInfo.Columns().UserId:   userId,
		dao.PraiseInfo.Columns().ObjectId: goodsId,
	}).Count()

	// 判断当前用户是否对该商品收藏
	isCollect, _ := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().Type:     consts.CollectionGoodsType,
		dao.CollectionInfo.Columns().UserId:   userId,
		dao.CollectionInfo.Columns().ObjectId: goodsId,
	}).Count()

	return &model.GoodsDetailOutput{
		Id:               goodsId,
		PicUrl:           goodsInfo.PicUrl,
		Name:             goodsInfo.Name,
		Price:            goodsInfo.Price,
		Level1CategoryId: goodsInfo.Level1CategoryId,
		Level2CategoryId: goodsInfo.Level2CategoryId,
		Level3CategoryId: goodsInfo.Level3CategoryId,
		Brand:            goodsInfo.Brand,
		Stock:            goodsInfo.Stock,
		Sale:             goodsInfo.Sale,
		Tags:             goodsInfo.Tags,
		DetailInfo:       goodsInfo.DetailInfo,
		IsPraise:         isPraise,
		IsCollect:        isCollect,
		TimeCommon: model.TimeCommon{
			CreatedAt: goodsInfo.CreatedAt,
			UpdatedAt: goodsInfo.UpdatedAt,
		},
	}, nil
}

// DetailBackend 查询商品详情（管理员）
func (s *sGoods) DetailBackend(ctx context.Context, id int) (out *model.GoodsDetailOutput, err error) {
	var goodsInfo entity.GoodsInfo
	err = dao.GoodsInfo.Ctx(ctx).Where(dao.GoodsInfo.Columns().Id, id).Scan(&goodsInfo)
	if err != nil {
		return nil, err
	}

	return &model.GoodsDetailOutput{
		Id:               goodsInfo.Id,
		PicUrl:           goodsInfo.PicUrl,
		Name:             goodsInfo.Name,
		Price:            goodsInfo.Price,
		Level1CategoryId: goodsInfo.Level1CategoryId,
		Level2CategoryId: goodsInfo.Level2CategoryId,
		Level3CategoryId: goodsInfo.Level3CategoryId,
		Brand:            goodsInfo.Brand,
		Stock:            goodsInfo.Stock,
		Sale:             goodsInfo.Sale,
		Tags:             goodsInfo.Tags,
		DetailInfo:       goodsInfo.DetailInfo,
		TimeCommon: model.TimeCommon{
			CreatedAt: goodsInfo.CreatedAt,
			UpdatedAt: goodsInfo.UpdatedAt,
		},
	}, nil
}
