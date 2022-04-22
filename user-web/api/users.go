package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/middleware"
	"hr-saas-go/user-web/models"
	"hr-saas-go/user-web/proto"
	"hr-saas-go/user-web/request"
	"hr-saas-go/user-web/utils"
	"net/http"
	"strconv"
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

func Update(ctx *gin.Context) {
	req, err := request.UpdateUserRequestGet(ctx)
	if err != nil {
		return
	}
	id := req.Id
	if id == 0 {
		id = ctx.GetInt64("userId")
	}
	//req
	_, err = global.UserServCon.UpdateUser(context.Background(), &proto.UpdateUserRequest{
		Id:          id,
		NickName:    req.NickName,
		Sex:         req.Sex,
		Avatar:      req.Avatar,
		CurrentRole: -1,
	})

	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
}

// GetInfo 用于返回信息
func GetInfo(ctx *gin.Context) {
	// 获取id
	id, _ := ctx.Get("userId")
	user, err := global.UserServCon.FindUserById(context.Background(), &proto.IdRequest{Id: id.(int64)})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id":           user.Id,
		"name":         user.Name,
		"nick_name":    user.NickName,
		"mobile":       user.Mobile,
		"sex":          user.Sex,
		"current_role": user.CurrentRole,
		"avatar":       user.Avatar,
	}))
}

// Show 获取某角色信息
func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := global.UserServCon.FindUserById(context.Background(), &proto.IdRequest{Id: int64(id)})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"name":         user.Name,
		"nick_name":    user.NickName,
		"mobile":       user.Mobile,
		"sex":          user.Sex,
		"current_role": user.CurrentRole,
	}))
}

func Register(ctx *gin.Context) {
	// 注册账户
	registerRequest, err := request.MobileRegisterRequestGet(ctx)
	if err != nil {
		return
	}
	// 获取验证码
	key := fmt.Sprintf("%s_register", registerRequest.Mobile)
	rdb := global.Rdb
	code := rdb.Get(context.Background(), key)
	// 验证码为空=>失效
	if code.Val() == "" {
		ctx.JSON(http.StatusOK, utils.ErrorJson("还未发送验证码"))
		return
	}

	if registerRequest.Code != code.Val() {
		ctx.JSON(http.StatusOK, utils.ErrorJson("验证码错误"))
		return
	}
	if user, _ := global.UserServCon.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: registerRequest.Mobile}); user != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("您已注册过，请勿重复注册"))
		return
	}
	// 注册账户
	_, err = global.UserServCon.CreateUser(context.Background(), &proto.UserRequest{
		// 随机名称
		Name:        "",
		Mobile:      registerRequest.Mobile,
		NickName:    "",
		Password:    registerRequest.Password,
		CurrentRole: registerRequest.CurrentRole,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	rdb.Del(context.Background(), key)
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
}

func LoginByCode(ctx *gin.Context) {
	mobileRequest, err := request.MobileLoginByCodeRequestGet(ctx)
	// 发送短信
	if err != nil {
		return
	}
	key := fmt.Sprintf("%s_login", mobileRequest.Mobile)
	rdb := global.Rdb
	code := rdb.Get(context.Background(), key)
	// 获取缓存
	if code.Val() == "" {
		ctx.JSON(http.StatusOK, utils.ErrorJson("还未发送验证码"))
		return
	}

	if mobileRequest.Code != code.Val() {
		ctx.JSON(http.StatusOK, utils.ErrorJson("验证码错误"))
		return
	}
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
		user, err = global.UserServCon.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: mobileRequest.Mobile})
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
	rdb.Del(context.Background(), key)
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]string{
		"token": token,
	}))
	return
}

// LoginByPassword 密码登录
func LoginByPassword(ctx *gin.Context) {
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

	//if !base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.Captcha, true) {
	//	ctx.JSON(http.StatusOK, utils.ErrorJson("验证码错误"))
	//	return
	//}

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
	}
	ctx.JSON(http.StatusOK, utils.ErrorJson("手机号或密码错误"))
	return
}
