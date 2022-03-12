package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/middleware"
	"hr-saas-go/user-web/models"
	"hr-saas-go/user-web/proto"
	"hr-saas-go/user-web/request"
	"hr-saas-go/user-web/utils"
	"net/http"
	"time"
)

func GetUserList(ctx *gin.Context) {
	//proto
	//host := global.Config.UserSrvInfo.Host
	//port := global.Config.UserSrvInfo.Port

	list, err := global.UserServCon.GetUserList(ctx, &proto.PageInfo{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		zap.S().Errorf("user service call err: %s", err.Error())
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func GetUser(ctx *gin.Context) {
	//proto
	get, err := request.MobileLoginRequestGet(ctx)
	if err != nil {
		return
	}
	println(get.Password, get.Mobile)
}

func LoginByMobile(ctx *gin.Context) {
	req, err := request.MobileLoginRequestGet(ctx)
	if err != nil {
		return
	}

	user, err := global.UserServCon.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: req.Mobile})
	if err != nil {
		utils.HandleGrpcError(err, ctx, "手机号或密码错误")
		return
	}

	checkResult, err := global.UserServCon.CheckPassword(context.Background(), &proto.CheckPasswordRequest{Password: req.Password, Encrypt: user.Password})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	if !base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.Captcha, true) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	if checkResult.Result {
		// jwt
		j := middleware.NewJWT()
		token, err := j.CreateToken(models.CustomClaims{
			ID:       uint(user.Id),
			NickName: user.NickName,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "blank",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  nil,
			},
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "登录失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不正确",
		})
		return
	}

}
