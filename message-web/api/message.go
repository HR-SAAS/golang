package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/message-web/global"
	"hr-saas-go/message-web/proto"
	"hr-saas-go/message-web/request"
	"hr-saas-go/message-web/utils"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	search := make(map[string]string)

	userId := ctx.GetInt64("userId")
	search["user_id"] = strconv.FormatInt(userId, 10)

	page, limit := utils.GetPage(ctx)
	list, err := global.UserMessageServCon.GetMessageList(ctx, &proto.GetMessageListRequest{
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
		resumes = append(
			resumes,
			map[string]interface{}{
				"id":            v.Id,
				"type":          v.Type,
				"source_type":   v.SourceType,
				"source_id":     v.SourceId,
				"is_read":       v.IsRead,
				"relation_id":   v.RelationId,
				"relation_type": v.RelationType,
				"user_id":       v.UserId,
				"status":        v.Status,
				"content":       v.Content,
				"created_at":    v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
				"updated_at":    v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
			},
		)
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

	data, err := global.UserMessageServCon.GetMessageDetail(ctx, &proto.GetMessageDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id":            data.Id,
		"type":          data.Type,
		"source_type":   data.SourceType,
		"source_id":     data.SourceId,
		"is_read":       data.IsRead,
		"relation_id":   data.RelationId,
		"relation_type": data.RelationType,
		"user_id":       data.UserId,
		"status":        data.Status,
		"content":       data.Content,
		"created_at":    data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at":    data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

func SendTo(ctx *gin.Context) {
	req, err := request.ResumeSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := global.UserMessageServCon.CreateMessage(
		context.Background(), &proto.CreateMessageRequest{
			Id:           0,
			UserId:       0,
			SourceType:   "user",
			SourceId:     userId,
			Content:      req.Content,
			IsRead:       false,
			Status:       0,
			RelationId:   0,
			RelationType: "",
		},
	)
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

	//todo 权限控制
	req, err := request.ResumeSaveRequestGet(ctx)
	if err != nil {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.UserMessageServCon.UpdateMessage(ctx, &proto.UpdateMessageRequest{
		Id:           int64(id),
		UserId:       0,
		SourceType:   "",
		SourceId:     0,
		Type:         "",
		Content:      req.Content,
		IsRead:       false,
		Status:       req.Status,
		RelationId:   0,
		RelationType: "",
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

	data, err := global.UserMessageServCon.DeleteMessage(ctx, &proto.DeleteMessageRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
