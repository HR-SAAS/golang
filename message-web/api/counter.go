package api

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/message-web/global"
	"hr-saas-go/message-web/proto"
	"hr-saas-go/message-web/utils"
	"net/http"
	"strconv"
)

func Count(ctx *gin.Context) {
	// 统计服务
	userId := ctx.GetInt64("userId")
	status, err := strconv.Atoi(ctx.Query("status"))
	if err != nil {
		status = -1
	}
	search := make(map[string]string)
	if content := ctx.GetString("content"); content != "" {
		search["content"] = content
	}
	res, err := global.MessageCounterService.CountUserMessage(ctx, &proto.CountUserMessageRequest{
		UserId: userId,
		Search: search,
		Status: int32(status),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int64{
		"count": res.Count,
	}))
}
