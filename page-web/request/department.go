package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/page-web/global"
	"hr-saas-go/page-web/utils"
	"net/http"
)

type DepartmentSaveRequest struct {
	Id        int64  ` form:"id" binding:"" json:"id"`
	CompanyId int64  ` form:"company_id" binding:"" json:"company_id"`
	ParentId  int64  ` form:"parent_id" binding:"" json:"parent_id"`
	Icon      string ` form:"icon" binding:"" json:"icon"`
	Name      string ` form:"name" binding:"required" json:"name"`
	Remark    string ` form:"remark" binding:"" json:"remark"`
	Desc      string ` form:"desc" binding:"" json:"desc"`
	Info      string ` form:"info" binding:"" json:"info"`
	Status    int32  ` form:"status" json:"status"`
}

func DepartmentSaveRequestGet(c *gin.Context) (DepartmentSaveRequest, error) {
	var request DepartmentSaveRequest
	if err := c.ShouldBind(&request); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		}
		//handle error
		return request, err
	}
	return request, nil
}
