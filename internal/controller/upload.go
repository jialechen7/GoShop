package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/utility/upload"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type cUpload struct{}

var Upload = cUpload{}

func (c *cUpload) UploadImgToCloud(ctx context.Context, req *backend.UploadImgToCloudReq) (res *backend.UploadImgToCloudRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	ret, err := upload.UploadImgToCloud(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &backend.UploadImgToCloudRes{
		Url: ret,
	}, nil
}
