package user_post

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
	"hr-saas-go/recruit-web/request"
	"hr-saas-go/recruit-web/utils"
	"net/http"
	"strconv"
)

// List 投递列表,投递人信息,企业信息,简历信息
func List(ctx *gin.Context) {
	search := make(map[string]string)
	// 根据用户搜索

	// 根据公司搜索

	// 根据岗位搜索

	page, limit := utils.GetPage(ctx)
	var list, err = global.UserPostServCon.GetUserPostList(ctx, &proto.GetUserPostListRequest{
		Page:  page,
		Limit: limit,
		Sort: map[string]string{
			"created_at": "asc",
		},
		Search: search,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	resumes := make([]interface{}, 0, limit)
	// 获取其他信息

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"data":  resumes,
		"total": list.Total,
	}))
}

// Show 查看投递详情
func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}
	// police

	data, err := global.UserPostServCon.GetUserPostDetail(ctx, &proto.GetUserPostDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	// TODO 获取其他信息

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id": data.Id,
		// 信息
		"created_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

// Create 投递简历
func Create(ctx *gin.Context) {
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := global.UserPostServCon.CreateUserPost(
		context.Background(), &proto.CreateUserPostRequest{
			PostId:     req.PostId,
			UserId:     userId,
			ResumeId:   req.ResumeId,
			ResumeType: req.ResumeType,
			ResumeName: req.ResumeName,
			Resume:     req.Resume,
			ReviewId:   req.ReviewId,
			Status:     req.Status,
			CompanyId:  req.CompanyId,
			Remark:     req.Remark,
		})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Id))
	return
}

// Update 更新状态
func Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.UserPostServCon.UpdateUserPost(ctx, &proto.UpdateUserPostRequest{
		ReviewId:  userId,
		Status:    req.Status,
		CompanyId: req.CompanyId,
		Remark:    req.Remark,
		// 添加审核ID
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

// BatchHandle 批量处理
func BatchHandle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.UserPostServCon.BatchUpdateUserPost(ctx, &proto.BatchUpdateUserPostRequest{
		Ids:      []int64{1},
		ReviewId: userId,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

// Delete 删除投递记录
func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
		return
	}

	data, err := global.UserPostServCon.DeleteUserPost(ctx, &proto.DeleteUserPostRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
