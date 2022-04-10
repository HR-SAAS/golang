package api

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码失败",
		})
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
		ctx.JSON(http.StatusOK, utils.ErrorJson("手机号不正确"))
		return
	}
	// 获取缓存

	// 发送手机验证码
	client, err := dysmsapi.NewClientWithAccessKey("cn-qingdao", "LTAI5tAAjLhEXmxpigb7BFYz", "XwpYel3yJ98sx8D0zcoRl9SUiOK2cM")
	/* use STS Token
	client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobileRequest.Mobile //接收短信的手机号码
	request.SignName = "yblog"                  //短信签名名称
	request.TemplateCode = "SMS_148865003"      //短信模板ID
	// 随机验证码
	request.TemplateParam = "{\"code\":\"" + generateCode(5) + "\"}"
	_, err = client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	// redis 存储次数和验证码内容
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
