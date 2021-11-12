package utils

import (
	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// GetUploadToken 获取上传令牌
func GetUploadToken() string {
	c := g.Config()
	accessKey := c.GetString("qiniu.AccessKey")
	secretKey := c.GetString("qiniu.SecretKey")
	bucket := c.GetString("qiniu.Bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	return putPolicy.UploadToken(mac)
}
