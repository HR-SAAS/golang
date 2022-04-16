package company

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/company-web/global"
	"hr-saas-go/company-web/proto"
	"hr-saas-go/company-web/request"
	"hr-saas-go/company-web/utils"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	list, err := global.CompanyServCon.GetCompanyList(ctx, &proto.GetCompanyListRequest{
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

	data, err := global.CompanyServCon.GetCompanyDetail(ctx, &proto.GetCompanyDetailRequest{
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
	req, err := request.CompanySaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := global.CompanyServCon.CreateCompany(context.Background(), &proto.CreateCompanyRequest{
		Name:      req.Name,
		Desc:      req.Desc,
		Website:   req.Website,
		Config:    req.Config,
		Address:   req.Address,
		Info:      req.Info,
		CreatorId: userId,
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
	req, err := request.CompanySaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.CompanyServCon.UpdateCompany(ctx, &proto.UpdateCompanyRequest{
		Id:        int64(id),
		Name:      req.Name,
		Desc:      req.Desc,
		Website:   req.Website,
		Config:    req.Config,
		Address:   req.Address,
		Info:      req.Info,
		CreatorId: userId,
		ParentId:  req.ParentId,
		Status:    1,
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
	// TODO 鉴权
	data, err := global.CompanyServCon.DeleteCompany(ctx, &proto.DeleteCompanyRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}

func MyCompany(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	page, limit := utils.GetPage(ctx)
	if userId == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("请登陆系统"))
		return
	}
	res, err := global.CompanyServCon.GetMyCompanyList(ctx, &proto.GetMyCompanyListRequest{
		Page:   page,
		Limit:  limit,
		UserId: userId,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"total": res.Total,
		"data":  res.Data,
	}))
	return
}

func GetCompanyUsers(ctx *gin.Context) {
	companyId, err := strconv.Atoi(ctx.Param("companyId"))
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("参数错误"))
		return
	}
	page, limit := utils.GetPage(ctx)
	res, err := global.CompanyServCon.GetCompanyUserIdList(ctx, &proto.GetCompanyUserListRequest{
		Page:      page,
		Limit:     limit,
		CompanyId: int64(companyId),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	// 使用用户服务进行获取userList
	print(res.Data)
}
