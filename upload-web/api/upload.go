package api

import (
	"crypto/md5"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hr-saas-go/upload-web/global"
	"hr-saas-go/upload-web/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func UploadFile(t string) func(ctx *gin.Context) {
	formName := "file"
	return func(ctx *gin.Context) {
		file, _ := ctx.FormFile(formName)
		if file == nil {
			ctx.JSON(http.StatusOK, utils.ErrorJson(fmt.Sprintf(" %s 不能为空", formName)))
			return
		}
		path := makePath(t, ctx.GetInt64("userId"))
		// 将文件发送到oss
		res, err := uploadToAliOss(file, path)
		if err != nil {
			zap.S().Errorf("upload err %v", err)
			ctx.JSON(http.StatusOK, utils.ErrorJson("上传失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
			"url": res,
		}))
	}
}

func uploadToAliOss(file *multipart.FileHeader, path string) (string, error) {
	client, err := oss.New(
		global.Config.OssConfig.Endpoint, global.Config.OssConfig.AccessId, global.Config.OssConfig.AccessKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如exampleBucket。
	bucket, err := client.Bucket(global.Config.OssConfig.Bucket)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)
	filename, err := makeFilename(file)
	if err != nil {
		return "", err
	}
	targetFileName := path + filename
	err = bucket.PutObject(targetFileName, src)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("//%s.%s/%s", global.Config.OssConfig.Bucket, global.Config.OssConfig.Endpoint, targetFileName), nil
}

func makeFilename(file *multipart.FileHeader) (string, error) {
	fieExtend := filepath.Ext(file.Filename)
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	sum, err := md5sum(f)
	if err != nil {
		return "", err
	}
	return sum + fieExtend, nil
}

func md5sum(file multipart.File) (string, error) {
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("%x", md5.Sum(body))
	runtime.GC()
	return res, nil
}

func makePath(fileType string, uid int64) string {
	return time.Now().Format(fmt.Sprintf("%s/06/01/02/%d/", fileType, uid))
}
