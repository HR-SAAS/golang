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
	// 登录
}

//用于返回信息
func GetInfo() {
	// 获取id
}

// 获取某角色信息
func Show() {

}

// LoginBySms 手机号验证码登录
func LoginBySms(ctx *gin.Context) {
	// 获取手机和验证码
	// 获取手机号
	mobileRequest, err := request.MobileLoginRequestGet(ctx)
	// 发送短信
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("手机号不正确"))
		return
	}

	// 获取缓存

	_, _ = ctx.GetPostForm("code")
	// 校验缓存
	user, err := global.UserServCon.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: mobileRequest.Mobile})
	if err != nil {
		// 走注册逻辑
		_, err := global.UserServCon.CreateUser(context.Background(), &proto.UserRequest{
			Mobile:   mobileRequest.Mobile,
			Password: "",
			Sex:      3,
		})
		if err != nil {
			ctx.JSON(http.StatusOK, utils.ErrorJson("服务器异常"))
			return
		}
	}

	j := middleware.NewJWT()
	token, err := j.CreateToken(models.CustomClaims{
		ID:       user.Id,
		NickName: user.NickName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "blank",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  nil,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("登录失败"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]string{
		"token": token,
	}))
	return
}

func LoginByMobile(ctx *gin.Context) {
	req, err := request.MobileLoginRequestGet(ctx)
	if err != nil {
		return
	}

	user, err := global.UserServCon.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: req.Mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("手机号或密码错误"))
		return
	}

	checkResult, err := global.UserServCon.CheckPassword(context.Background(), &proto.CheckPasswordRequest{Password: req.Password, Encrypt: user.Password})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	if !base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.Captcha, true) {
		ctx.JSON(http.StatusOK, utils.ErrorJson("验证码错误"))
		return
	}

	if checkResult.Result {
		// jwt
		j := middleware.NewJWT()
		token, err := j.CreateToken(models.CustomClaims{
			ID:       user.Id,
			NickName: user.NickName,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "blank",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  nil,
			},
		})
		if err != nil {
			ctx.JSON(http.StatusOK, utils.ErrorJson("登录失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]string{
			"token": token,
		}))
		return
	} else {
		ctx.JSON(http.StatusOK, utils.ErrorJson("手机号或密码错误"))
		return
	}
}
