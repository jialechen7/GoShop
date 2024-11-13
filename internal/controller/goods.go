package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// Goods 商品管理
var Goods = cGoods{}

type cGoods struct{}

// ListBackend 查询商品列表
func (c *cGoods) ListBackend(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetListBackend(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cGoods) UpdateBackend(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	g.Dump(req)
	err = service.Goods().UpdateBackend(ctx, model.GoodsUpdateInput{
		Id: req.Id,
		GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
			PicUrl:           req.PicUrl,
			Name:             req.Name,
			Price:            req.Price,
			Level1CategoryId: req.Level1CategoryId,
			Level2CategoryId: req.Level2CategoryId,
			Level3CategoryId: req.Level3CategoryId,
			Brand:            req.Brand,
			Stock:            req.Stock,
			Sale:             req.Sale,
			Tags:             req.Tags,
			DetailInfo:       req.DetailInfo,
		},
	})
	return
}

// AddBackend 添加商品
func (c *cGoods) AddBackend(ctx context.Context, req *backend.GoodsAddReq) (res *backend.GoodsAddRes, err error) {
	out, err := service.Goods().AddBackend(ctx, model.GoodsAddInput{
		GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
			PicUrl:           req.PicUrl,
			Name:             req.Name,
			Price:            req.Price,
			Level1CategoryId: req.Level1CategoryId,
			Level2CategoryId: req.Level2CategoryId,
			Level3CategoryId: req.Level3CategoryId,
			Brand:            req.Brand,
			Stock:            req.Stock,
			Sale:             req.Sale,
			Tags:             req.Tags,
			DetailInfo:       req.DetailInfo,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsAddRes{
		Id: out.GoodsId,
	}, nil
}

// ListFrontend 查询商品列表
func (c *cGoods) ListFrontend(ctx context.Context, req *frontend.GoodsGetListCommonReq) (res *frontend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetListFrontend(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.GoodsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListByLevelFrontend 查询商品列表（根据2级分类）
func (c *cGoods) ListByLevelFrontend(ctx context.Context, req *frontend.GoodsGetListByLevelReq) (res *frontend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetListByLevelFrontend(ctx, model.GoodsGetListByLevelInput{
		Page:    req.Page,
		Size:    req.Size,
		LevelId: req.LevelId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.GoodsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cGoods) DeleteBackend(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().DeleteBackend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsDeleteRes{}, nil
}

func (c *cGoods) DetailFrontend(ctx context.Context, req *frontend.GoodsDetailReq) (res *frontend.GoodsDetailRes, err error) {
	out, err := service.Goods().DetailFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.GoodsDetailRes{
		Id:               out.Id,
		PicUrl:           out.PicUrl,
		Name:             out.Name,
		Price:            out.Price,
		Level1CategoryId: out.Level1CategoryId,
		Level2CategoryId: out.Level2CategoryId,
		Level3CategoryId: out.Level3CategoryId,
		Brand:            out.Brand,
		Stock:            out.Stock,
		Sale:             out.Sale,
		Tags:             out.Tags,
		DetailInfo:       out.DetailInfo,
		IsPraise:         out.IsPraise,
	}, nil
}

func (c *cGoods) DetailBackend(ctx context.Context, req *backend.GoodsDetailReq) (res *backend.GoodsDetailRes, err error) {
	out, err := service.Goods().DetailBackend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsDetailRes{
		Id:               out.Id,
		PicUrl:           out.PicUrl,
		Name:             out.Name,
		Price:            out.Price,
		Level1CategoryId: out.Level1CategoryId,
		Level2CategoryId: out.Level2CategoryId,
		Level3CategoryId: out.Level3CategoryId,
		Brand:            out.Brand,
		Stock:            out.Stock,
		Sale:             out.Sale,
		Tags:             out.Tags,
		DetailInfo:       out.DetailInfo,
	}, nil
}
