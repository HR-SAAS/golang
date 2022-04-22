package api

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
	"hr-saas-go/recruit-web/utils"
	"net/http"
)

func Count(ctx *gin.Context) {
	// 统计服务

	res, err := global.RecruitCounterServiceServCon.CountPost(ctx, &proto.CountPostRequest{
		CompanyId: 0,
		Search:    nil,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int64{
		"count": res.Count,
	}))
}
