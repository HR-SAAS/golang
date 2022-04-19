package company

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/page-web/global"
	"hr-saas-go/page-web/proto"
	"hr-saas-go/page-web/request"
	"hr-saas-go/page-web/utils"
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
	res, err := global.CompanyServCon.CreateCompany(
		context.Background(), &proto.CreateCompanyRequest{
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
	companyId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("参数错误"))
		return
	}
	page, limit := utils.GetPage(ctx)
	realations, err := global.CompanyServCon.GetCompanyUserIdList(ctx, &proto.GetCompanyUserListRequest{
		Page:      page,
		Limit:     limit,
		CompanyId: int64(companyId),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	// 使用用户服务进行获取userList
	ids := make([]int64, 0, limit)

	for _, v := range realations.Data {
		ids = append(ids, v.UserId)
	}

	usersRes, err := global.UserServCon.GetUserListByIds(ctx, &proto.GetUserListByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	users := map[int64]interface{}{}

	for _, v := range usersRes.Data {
		users[v.Id] = v
	}

	result := make([]interface{}, 0, limit)
	for _, v := range realations.Data {
		result = append(result, map[string]interface{}{
			"user":       users[v.UserId],
			"user_id":    v.UserId,
			"nick_name":  v.NickName,
			"info":       v.Info,
			"status":     v.Status,
			"department": v.DepartmentId,
			"remark":     v.Remark,
			"created_at": v.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
			"updated_at": v.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"data":  result,
		"total": realations.Total,
	}))
}
