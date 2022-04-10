package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/company-web/global"
	"hr-saas-go/company-web/utils"
	"net/http"
)

type DepartmentSaveRequest struct {
	Id        int64  ` form:"id" binding:"required" json:"id"`
	CompanyId int64  ` form:"company_id" binding:"required" json:"company_id"`
	ParentId  int64  ` form:"parent_id" binding:"required" json:"parent_id"`
	Icon      int64  ` form:"icon" binding:"" json:"icon"`
	Name      string ` form:"name" binding:"required" json:"name"`
	Remark    string ` form:"remark" binding:"" json:"remark"`
	Desc      string ` form:"desc" binding:"" json:"desc"`
	Info      string ` form:"info" binding:"" json:"info"`
	Status    int32  `form:"status" json:"status"`
}

func DepartmentSaveRequestGet(c *gin.Context) (DepartmentSaveRequest, error) {
	var request DepartmentSaveRequest
	if err := c.ShouldBind(&request); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": utils.RemoveTopName(e.Translate(global.Trans)),
			})
		}
		//handle error
		return request, err
	}
	return request, nil
}
