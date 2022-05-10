package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"github.com/wujunyi792/where-it-go/config"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"io"
	"path"
)

func init() {
	//if !config.GetConfig().OSS.Use {
	//	panic("OSS not open, please check config")
	//}
	InitOSS()
	logger.Info.Println("OSS init SUCCESS ")
}

var client *oss.Client
var bucket *oss.Bucket

func InitOSS() {
	// 创建OSSClient实例。
	var err error
	conf := &config.GetConfig().OSS.Config
	client, err = oss.New(conf.EndPoint, conf.AccessKeyId, conf.AccessKeySecret)
	// 获取存储空间。
	if err != nil {
		logger.Error.Fatalln(err)
	}
	bucket, err = client.Bucket(conf.BucketName)
	if err != nil {
		logger.Error.Fatalln("阿里云图库连接失败: ", err)
	}
}

func UploadFileToOss(filename string, fd io.Reader) string {
	conf := &config.GetConfig().OSS.Config
	fname := uuid.NewV4().String() + path.Ext(filename)
	err := bucket.PutObject(conf.Path+fname, fd)
	pictureUrl := conf.BaseURL + conf.Path + fname
	if err != nil {
		logger.Error.Println("File upload to OSS fail，fileName：", pictureUrl, ", err: :", err)
		return ""
	}
	return pictureUrl
}
