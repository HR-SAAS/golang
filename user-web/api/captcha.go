package api

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/request"
	"hr-saas-go/user-web/utils"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetCaptcha(ctx *gin.Context) {
	c := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	// 返回图片验证码
	id, b64s, err := c.Generate()
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("生成验证码失败"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"base64": b64s,
	})
	return
}

func GetSmsCaptcha(ctx *gin.Context) {
	// 获取手机号
	mobileRequest, err := request.MobileSmsLoginRequestGet(ctx)
	// 发送短信
	if err != nil {
		return
	}
	rdb := global.Rdb

	key := fmt.Sprintf("%s_%s", mobileRequest.Mobile, mobileRequest.Type)
	// 获取缓存
	// lock
	if data := rdb.Get(context.Background(), fmt.Sprintf("%s_lock", mobileRequest.Mobile)); data.Val() != "" {
		ctx.JSON(http.StatusOK, utils.ErrorJson("请勿重复请求"))
		return
	}
	// times次数
	if data, _ := rdb.Get(context.Background(), mobileRequest.Mobile).Int(); data > 3 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("请勿重复请求"))
		return
	}
	code := 0
	if global.Debug {
		code = 12345
	} else {
		client, err := dysmsapi.NewClientWithAccessKey("cn-qingdao", "LTAI5tAAjLhEXmxpigb7BFYz", "XwpYel3yJ98sx8D0zcoRl9SUiOK2cM")
		/* use STS Token
		client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
		*/
		sendSmsRequest := dysmsapi.CreateSendSmsRequest()
		sendSmsRequest.Scheme = "https"
		sendSmsRequest.PhoneNumbers = mobileRequest.Mobile //接收短信的手机号码
		sendSmsRequest.SignName = "yblog"                  //短信签名名称
		sendSmsRequest.TemplateCode = "SMS_148865003"      //短信模板ID
		code := generateCode(5)
		sendSmsRequest.TemplateParam = "{\"code\":\"" + code + "\"}"
		_, err = client.SendSms(sendSmsRequest)
		if err != nil {
			ctx.JSON(http.StatusOK, utils.ErrorJson("发送失败"))
			return
		}
	}

	// 发送手机验证码

	// redis 存储次数和验证码内容
	rdb.Set(context.Background(), key, code, time.Minute*5)
	rdb.Set(context.Background(), fmt.Sprintf("%s_lock", mobileRequest.Mobile), code, time.Minute)

	rdb.Incr(context.Background(), mobileRequest.Mobile)
	ctx.JSON(http.StatusOK, utils.SuccessJson("发送成功"))
	return
}

func generateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
