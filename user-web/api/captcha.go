package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
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
