package controller

import (
	"context"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Collection 收藏管理
var Collection = cCollection{}

type cCollection struct{}

// ListFrontend 查询收藏列表（仅用户自己）
func (c *cCollection) ListFrontend(ctx context.Context, req *frontend.CollectionGetListCommonReq) (res *frontend.CollectionGetListCommonRes, err error) {
	getListRes, err := service.Collection().GetListFrontend(ctx, model.CollectionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cCollection) AddFrontend(ctx context.Context, req *frontend.CollectionAddReq) (res *frontend.CollectionAddRes, err error) {
	out, err := service.Collection().AddFrontend(ctx, model.CollectionAddInput{
		CollectionCreateUpdateBase: model.CollectionCreateUpdateBase{
			Type:     req.Type,
			ObjectId: req.ObjectId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionAddRes{
		CollectionId: out.CollectionId,
	}, nil
}

func (c *cCollection) DeleteFrontend(ctx context.Context, req *frontend.CollectionDeleteReq) (res *frontend.CollectionDeleteRes, err error) {
	err = service.Collection().DeleteFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionDeleteRes{}, nil
}

func (c *cCollection) DeleteByTypeFrontend(ctx context.Context, req *frontend.CollectionDeleteByTypeReq) (res *frontend.CollectionDeleteByTypeRes, err error) {
	err = service.Collection().DeleteByTypeFrontend(ctx, model.CollectionDeleteByTypeInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionDeleteByTypeRes{}, nil
}
