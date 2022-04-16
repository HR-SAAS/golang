package department

import (
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
	companyId, _ := strconv.Atoi(ctx.Query("company_id"))
	if companyId == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("错误到company_id"))
	}
	page, limit := utils.GetPage(ctx)
	var list, err = global.DepartmentServCon.GetDepartmentList(ctx, &proto.GetDepartmentListRequest{
		Page:      page,
		Limit:     limit,
		CompanyId: int64(companyId),
		Search:    "",
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}
	data, err := global.DepartmentServCon.GetDepartmentDetail(ctx, &proto.GetDepartmentDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		zap.S().Errorf("err: %s", err)
		ctx.JSON(http.StatusOK, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return
}

func Create(ctx *gin.Context) {
	req, err := request.DepartmentSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := global.DepartmentServCon.CreateDepartment(ctx, &proto.CreateDepartmentRequest{
		CompanyId: req.CompanyId,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
		Name:      req.Name,
		Remark:    req.Remark,
		Desc:      req.Desc,
		Info:      req.Info,
		CreatorId: userId,
		Status:    1,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res))
	return
}

func Update(ctx *gin.Context) {
	req, err := request.DepartmentSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("参数错误"))
		return
	}
	_, err = global.DepartmentServCon.UpdateDepartment(ctx, &proto.UpdateDepartmentRequest{
		Id:        int64(id),
		CompanyId: req.CompanyId,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
		Name:      req.Name,
		Remark:    req.Remark,
		Desc:      req.Desc,
		Info:      req.Info,
		CreatorId: userId,
		Status:    req.Status,
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
		return
	}

	data, err := global.DepartmentServCon.DeleteDepartment(ctx, &proto.DeleteDepartmentRequest{
		Id: int64(id),
	})
	if err != nil {
		zap.S().Errorf("err: %s", err)
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return
}

// MyDepartment 我加入的
func MyDepartment(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	res, err := global.DepartmentServCon.GetMyDepartmentList(ctx, &proto.GetMyDepartmentListRequest{
		UserId: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res))
	return
}
