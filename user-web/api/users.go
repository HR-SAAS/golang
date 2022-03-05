package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/proto"
	"hr-saas-go/user-web/request"
	"net/http"
)

func MakeTrans(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func SuccessRet(msg string) {

}

func HandleGrpcError(err error, ctx *gin.Context) {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, "找不到", nil))
		case codes.InvalidArgument:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, "错误参数", nil))
		case codes.Internal:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, "内部错误", nil))
		default:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, "其他错误", nil))
		}
	}
}

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
		HandleGrpcError(err, ctx)
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
	get, err := request.MobileLoginRequestGet(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": get,
	})
}
