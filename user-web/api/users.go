package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"hr-saas-go/user-web/proto"
	"net/http"
)

func GetUserList(ctx *gin.Context) {
	//proto
	host := "127.0.0.1"
	port := 5001
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "调用失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func GetUser(ctx *gin.Context) {
	//proto
}
