package api

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/resume-web/global"
	"hr-saas-go/resume-web/proto"
	"hr-saas-go/resume-web/utils"
	"net/http"
	"strconv"
)

func Count(ctx *gin.Context) {
	// 统计服务
	userId := ctx.GetInt64("userId")

	t, err := strconv.Atoi(ctx.Query("type"))
	if err != nil {
		t = -1
	}
	status, err := strconv.Atoi(ctx.Query("status"))
	if err != nil {
		status = -1
	}
	res, err := global.ResumeCountServCon.CountResume(ctx, &proto.CountResumeRequest{
		UserId: userId,
		Name:   ctx.Query("name"),
		Type:   int32(t),
		Status: int32(status),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, map[string]int64{
		"count": res.Count,
	})
}
