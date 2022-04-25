package api

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
	"hr-saas-go/recruit-web/utils"
	"net/http"
	"strconv"
)

func CountPost(ctx *gin.Context) {
	// 统计统计有多少个岗位
	var companyId int64 = 0
	search := make(map[string]string)
	// 统计,
	if departmentId := ctx.Query("department_id"); departmentId != "" {
		search["department_id"] = departmentId
	}
	// 公司
	if cid, _ := strconv.Atoi(ctx.Query("company_id")); cid > 0 {
		companyId = int64(cid)
	}
	res, err := global.RecruitCounterServiceServCon.CountPost(ctx, &proto.CountPostRequest{
		CompanyId: companyId,
		Search:    search,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int64{
		"count": res.Count,
	}))
}

func CountUserPost(ctx *gin.Context) {
	// 统计统计有多少个岗位
	search := make(map[string]string)
	// 统计,
	req := &proto.CountUserPostRequest{
		Search: nil,
	}
	if departmentId := ctx.Query("department_id"); departmentId != "" {
		search["department_id"] = departmentId
	}
	// 公司
	if cid, _ := strconv.Atoi(ctx.Query("company_id")); cid > 0 {
		req.CompanyId = int64(cid)
	}
	if postId, _ := strconv.Atoi(ctx.Query("post_id")); postId > 0 {
		req.PostId = int64(postId)
	}
	res, err := global.RecruitCounterServiceServCon.CountUserPost(ctx, req)
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int64{
		"count": res.Count,
	}))
}
