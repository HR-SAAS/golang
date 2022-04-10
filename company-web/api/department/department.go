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
	companyIdStr := ctx.Query("company_id")
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	companyId, _ := strconv.Atoi(companyIdStr)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	var list, err = global.DepartmentServCon.GetDepartmentList(ctx, &proto.GetDepartmentListRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		CompanyId: int64(companyId),
		Search:    "",
		Order:     "",
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
		data, err := global.DepartmentServCon.GetDepartmentDetail(ctx, &proto.GetDepartmentDetailRequest{
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
	req, err := request.DepartmentSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("系统错误"))
		return
	}
	res, err := global.DepartmentServCon.CreateDepartment(ctx, &proto.CreateDepartmentRequest{
		CompanyId: req.CompanyId,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
		Name:      req.Name,
		Remark:    req.Remark,
		Desc:      req.Desc,
		Info:      req.Info,
		CreatorId: userId.(int64),
		Status:    1,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusInternalServerError, utils.SuccessJson(res))
	return
}

func Update(ctx *gin.Context) {
	req, err := request.DepartmentSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	id, _ := ctx.Get("id")
	_, err = global.DepartmentServCon.UpdateDepartment(ctx, &proto.UpdateDepartmentRequest{
		Id:        id.(int64),
		CompanyId: req.CompanyId,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
		Name:      req.Name,
		Remark:    req.Remark,
		Desc:      req.Desc,
		Info:      req.Info,
		CreatorId: userId.(int64),
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
	id := ctx.Param("id")
	if idInt, err := strconv.Atoi(id); err == nil {
		data, err := global.DepartmentServCon.DeleteDepartment(ctx, &proto.DeleteDepartmentRequest{
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

// MyDepartment 我加入的
func MyDepartment(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	res, err := global.DepartmentServCon.GetMyDepartmentList(ctx, &proto.GetMyDepartmentListRequest{
		UserId: userId.(int64),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res))
	return
}
