package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/utils"
	"net/http"
)

type PostSaveRequest struct {
	Name      string   `json:"name" form:"name" binding:""`
	Type      string   `json:"type" form:"type" binding:"required"`
	Tag       []string `json:"tag" form:"tag" binding:""`
	Content   string   `json:"content" form:"content" binding:"required"`
	Status    int32    `json:"status" form:"status" binding:"oneof=0 1"`
	CompanyId int64    `json:"company_id" form:"company_id"`
}

func PostSaveRequestGet(c *gin.Context) (PostSaveRequest, error) {
	var request PostSaveRequest
	if err := c.ShouldBind(&request); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		} else {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("类型错误"))
		}
		//handle error
		return request, err
	}
	return request, nil
}
