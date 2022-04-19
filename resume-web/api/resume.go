package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/resume-web/global"
	"hr-saas-go/resume-web/proto"
	"hr-saas-go/resume-web/utils"
	"net/http"
	"strconv"
)

// 增删改查

func List(ctx *gin.Context) {
	list, err := global.ResumeServCon.GetResumeList(ctx, &proto.GetResumeListRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(list))
}

func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}

	data, err := global.ResumeServCon.GetResumeDetail(ctx, &proto.GetResumeDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
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
			Name:      req.Name,
			Desc:      req.Desc,
			Website:   req.Website,
			Config:    req.Config,
			Address:   req.Address,
			Tags:      req.Tags,
			Info:      req.Info,
			CreatorId: userId,
			Logo:      req.Logo,
			ParentId:  req.ParentId,
			Status:    1,
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
		Id:     int64(id),
		Name:   req.Name,
		Status: 1,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

func Delete(ctx *gin.Context) {
	// 是否展示创建者
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
