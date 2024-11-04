package file

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"time"

	"github.com/gogf/gf/os/gfile"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// Upload 上传文件
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败，上传路径不存在")
	}

	if in.Name != "" {
		in.File.Filename = in.Name
		in.RandomName = false
	}
	// 安全性校验：每人1分钟内只能上传consts.FileMaxUploadCountMinute次
	count, err := dao.FileInfo.Ctx(ctx).Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).Count()
	if err != nil {
		return nil, gerror.New("上传失败")
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传次数过多，请稍后再试")
	}

	// 定义临时保存路径
	dirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dirName), in.RandomName)
	if err != nil {
		return nil, err
	}

	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dirName, fileName),
		Url:    "/upload/" + dirName + "/" + fileName,
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	// 入库
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   int(id),
		Name: data.Name,
		Url:  data.Url,
		Src:  data.Src,
	}, nil
}
