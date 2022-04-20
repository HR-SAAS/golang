package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hr-saas-go/resume-web/global"
	"hr-saas-go/resume-web/proto"
	"hr-saas-go/resume-web/request"
	"hr-saas-go/resume-web/utils"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	search := make(map[string]string)

	userId := ctx.GetInt64("userId")
	search["user_id"] = strconv.FormatInt(userId, 10)

	page, limit := utils.GetPage(ctx)
	zap.S().Infof("%d,%d", page, limit)
	list, err := global.ResumeServCon.GetResumeList(ctx, &proto.GetResumeListRequest{
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
			"tag":        v.Tag,
			"type":       v.Type,
			"status":     v.Status,
			"post_count": v.PostCount,
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

	data, err := global.ResumeServCon.GetResumeDetail(ctx, &proto.GetResumeDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id":         data.Id,
		"name":       data.Name,
		"tag":        data.Tag,
		"type":       data.Type,
		"status":     data.Status,
		"content":    data.Content,
		"post_count": data.PostCount,
		"created_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

func Create(ctx *gin.Context) {
	req, err := request.ResumeSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := global.ResumeServCon.CreateResume(
		context.Background(), &proto.CreateResumeRequest{
			UserId:  userId,
			Name:    req.Name,
			Type:    req.Type,
			Tag:     req.Tag,
			Content: req.Content,
			Status:  req.Status,
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
	req, err := request.ResumeSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.ResumeServCon.UpdateResume(ctx, &proto.UpdateResumeRequest{
		Id:      int64(id),
		UserId:  userId,
		Name:    req.Name,
		Type:    req.Type,
		Tag:     req.Tag,
		Content: req.Content,
		Status:  req.Status,
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

	data, err := global.ResumeServCon.DeleteResume(ctx, &proto.DeleteResumeRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
