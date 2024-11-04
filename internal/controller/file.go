package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/errors/gerror"
)

type cFile struct{}

var File = cFile{}

// FileUpload 文件上传
func (c *cFile) FileUpload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	out, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &backend.FileUploadRes{
		Name: out.Name,
		Url:  out.Url,
	}, nil
}
