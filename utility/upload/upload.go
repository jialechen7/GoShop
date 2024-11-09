package upload

import (
	"context"
	"os"

	"github.com/qiniu/go-sdk/v7/storagev2/region"

	"github.com/qiniu/go-sdk/v7/storagev2/credentials"

	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadImgToCloud 上传图片到七牛云工具方法
func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	dirPath := g.Cfg().MustGet(ctx, "qiniu.tempDir").String()
	fileName, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}
	// 定义本地文件路径
	localFile := dirPath + "/" + fileName
	// 获取七牛云配置文件
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	// 创建七牛云凭证，配置选项参数
	mac := credentials.NewCredentials(accessKey, secretKey)
	options := uploader.UploadManagerOptions{
		Options: http_client.Options{
			// z2: 华南
			Regions:     region.GetRegionByID("z2", false),
			Credentials: mac,
		},
	}
	uploadManager := uploader.NewUploadManager(&options)
	key := fileName

	// 上传文件到七牛云
	err = uploadManager.UploadFile(context.Background(), localFile, &uploader.ObjectOptions{
		FileName:   fileName,
		BucketName: bucket,
		ObjectName: &key,
	}, nil)
	if err != nil {
		return "", err
	}
	// 删除本地文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	// 返回七牛云文件地址，若使用的是七牛云的免费域名，则仅有30天有效期，且不支持https
	url = g.Cfg().MustGet(ctx, "qiniu.domain").String() + "/" + fileName
	return url, nil
}
