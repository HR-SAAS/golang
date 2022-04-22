package post

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

func List(ctx *gin.Context) {
	search := make(map[string]string)

	userId := ctx.GetInt64("userId")
	search["user_id"] = strconv.FormatInt(userId, 10)

	page, limit := utils.GetPage(ctx)
	list, err := global.PostServCon.GetPostList(ctx, &proto.GetPostListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	resumes := make([]interface{}, 0, limit)
	for _, v := range list.Data {
		resumes = append(resumes, map[string]interface{}{
			"id":         v.Id,
			"name":       v.Name,
			"type":       v.Type,
			"content":    v.Content,
			"created_at": v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
			"updated_at": v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		})
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"data":  resumes,
		"total": list.Total,
	}))
}

func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}
	// police

	data, err := global.PostServCon.GetPostDetail(ctx, &proto.GetPostDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	// 获取company

	// 获取department

	// 获取creator_id

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id":         data.Id,
		"name":       data.Name,
		"type":       data.Type,
		"content":    data.Content,
		"created_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

func Create(ctx *gin.Context) {
	req, err := request.PostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")

	startAt, err := utils.FormatTimeToStampPb(req.StartAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}

	endAt, err := utils.FormatTimeToStampPb(req.EndAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}
	res, err := global.PostServCon.CreatePost(
		context.Background(), &proto.CreatePostRequest{
			CompanyId:    req.CompanyId,
			DepartmentId: req.DepartmentId,
			CreatorId:    userId,
			Type:         req.Type,
			Name:         req.Name,
			Desc:         req.Desc,
			Content:      req.Content,
			Experience:   req.Experience,
			Education:    req.Education,
			Address:      req.Address,
			StartAt:      startAt,
			EndAt:        endAt,
		})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Id))
	return
}

func Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.PostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	startAt, err := utils.FormatTimeToStampPb(req.StartAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}

	endAt, err := utils.FormatTimeToStampPb(req.EndAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}
	_, err = global.PostServCon.UpdatePost(ctx, &proto.UpdatePostRequest{
		Id:           int64(id),
		CompanyId:    req.CompanyId,
		DepartmentId: req.DepartmentId,
		CreatorId:    userId,
		Type:         req.Type,
		Name:         req.Name,
		Desc:         req.Desc,
		Content:      req.Content,
		Experience:   req.Experience,
		Education:    req.Education,
		Address:      req.Address,
		StartAt:      startAt,
		EndAt:        endAt,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
		return
	}
	// 修改状态为取消投递 -1
	data, err := global.PostServCon.DeletePost(ctx, &proto.DeletePostRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
