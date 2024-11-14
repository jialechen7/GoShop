package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// GoodsOptions 商品规格管理
var GoodsOptions = cGoodsOptions{}

type cGoodsOptions struct{}

// ListBackend 查询商品规格列表
func (c *cGoodsOptions) ListBackend(ctx context.Context, req *backend.GoodsOptionsGetListCommonReq) (res *backend.GoodsOptionsGetListCommonRes, err error) {
	getListRes, err := service.GoodsOptions().GetListBackend(ctx, model.GoodsOptionsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListFrontend 查询商品规格列表
func (c *cGoodsOptions) ListFrontend(ctx context.Context, req *frontend.GoodsOptionsGetListCommonReq) (res *frontend.GoodsOptionsGetListCommonRes, err error) {
	g.Dump(req)
	getListRes, err := service.GoodsOptions().GetListFrontend(ctx, model.GoodsOptionsGetListInput{
		Page:    req.Page,
		Size:    req.Size,
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.GoodsOptionsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cGoodsOptions) DetailFrontend(ctx context.Context, req *frontend.GoodsOptionsDetailReq) (res *frontend.GoodsOptionsDetailRes, err error) {
	out, err := service.GoodsOptions().DetailFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.GoodsOptionsDetailRes{
		Id:      out.Id,
		GoodsId: out.GoodsId,
		PicUrl:  out.PicUrl,
		Name:    out.Name,
		Price:   out.Price,
		Stock:   out.Stock,
	}, nil
}
