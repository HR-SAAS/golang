package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/company-web/global"
	"hr-saas-go/company-web/utils"
	"net/http"
)

type CompanySaveRequest struct {
	Name     string   `json:"name" form:"name" binding:"required,min=3"`
	Desc     string   `json:"desc" form:"desc" binding:""`
	Website  string   `json:"website" form:"website" binding:""`
	Config   string   `json:"config" form:"config" binding:""`
	Address  string   `json:"address" form:"address" binding:""`
	Tags     []string `json:"tags" form:"tags" binding:""`
	Info     string   `json:"info" form:"info" binding:""`
	ParentId int64    `json:"parent_id" form:"creator_id" binding:""`
	Logo     string   `json:"logo" form:"logo" binding:""`
}

func CompanySaveRequestGet(c *gin.Context) (CompanySaveRequest, error) {
	var request CompanySaveRequest
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
