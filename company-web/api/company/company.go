package company

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		zap.S().Errorf("company service call err: %s", err.Error())
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(list))
}

func Show(ctx *gin.Context) {
	// 是否展示创建者
	id := ctx.Param("id")
	if idInt, err := strconv.Atoi(id); err == nil {
		data, err := global.CompanyServCon.GetCompanyDetail(ctx, &proto.GetCompanyDetailRequest{
			Id: int64(idInt),
		})
		if err != nil {
			zap.S().Errorf("err: %s", err)
			ctx.JSON(http.StatusOK, utils.ErrorJson("系统错误"))
			return
		}
		ctx.JSON(http.StatusOK, utils.SuccessJson(data))
		return
	}
	ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
	return
}

func Create(ctx *gin.Context) {
	req, err := request.CompanySaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("系统错误"))
		return
	}
	Tags := ""
	if req.Tags != nil {
		tags, _ := json.Marshal(req.Tags)
		Tags = string(tags)
	}
	res, err := global.CompanyServCon.CreateCompany(context.Background(), &proto.CreateCompanyRequest{
		Name:      req.Name,
		Desc:      req.Desc,
		Website:   req.Website,
		Config:    req.Config,
		Tags:      Tags,
		Address:   req.Address,
		Info:      req.Info,
		CreatorId: userId.(int64),
		ParentId:  req.ParentId,
		Status:    1,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	_, err = global.CompanyServCon.CreateUserCompany(context.Background(), &proto.SaveUserCompanyRequest{
		UserId:       ctx.GetInt64("userId"),
		CompanyId:    res.Id,
		Status:       1,
		DepartmentId: 0,
		Remark:       "Boss",
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Id))
	return
}

func Update(ctx *gin.Context) {
	req, err := request.CompanySaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.CompanyServCon.UpdateCompany(ctx, &proto.UpdateCompanyRequest{
		Name:    req.Name,
		Desc:    req.Desc,
		Website: req.Website,
		Config:  req.Config,
		//Tags:      req.Tags,
		Address:   req.Address,
		Info:      req.Info,
		CreatorId: userId.(int64),
		ParentId:  req.ParentId,
		Status:    1,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

func Delete(ctx *gin.Context) {
	// 是否展示创建者
	id := ctx.Param("id")
	if idInt, err := strconv.Atoi(id); err == nil {
		data, err := global.CompanyServCon.DeleteCompany(ctx, &proto.DeleteCompanyRequest{
			Id: int64(idInt),
		})
		if err != nil {
			zap.S().Errorf("err: %s", err)
			ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
			return
		}
		ctx.JSON(http.StatusOK, utils.SuccessJson(data))

		return
	}
	ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
	return
}

func MyCompany(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	res, err := global.CompanyServCon.GetMyCompanyList(ctx, &proto.GetMyCompanyListRequest{
		UserId: userId.(int64),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res))
	return
}
