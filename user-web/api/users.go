package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/proto"
	"hr-saas-go/user-web/request"
	"hr-saas-go/user-web/utils"
	"net/http"
)

func GetUserList(ctx *gin.Context) {
	//proto
	host := global.Config.UserSrvInfo.Host
	port := global.Config.UserSrvInfo.Port
	con, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("connect user service err: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "调用失败",
			"err": err.Error(),
		})
	}
	usr := proto.NewUserClient(con)
	list, err := usr.GetUserList(ctx, &proto.PageInfo{
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
	request, err := request.MobileLoginRequestGet(ctx)
	if err != nil {
		return
	}
	//登录逻辑
	host := global.Config.UserSrvInfo.Host
	port := global.Config.UserSrvInfo.Port
	con, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("connect user service err: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "调用失败",
			"err": err.Error(),
		})
	}
	usr := proto.NewUserClient(con)
	user, err := usr.FindUserByMobile(context.Background(), &proto.MobileRequest{Mobile: request.Mobile})
	if err != nil {
		utils.HandleGrpcError(err, ctx, "手机号或密码错误")
		return
	}

	checkResult, err := usr.CheckPassword(context.Background(), &proto.CheckPasswordRequest{Password: request.Password, Encrypt: user.Password})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	if checkResult.Result {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不正确",
		})
		return
	}

}
