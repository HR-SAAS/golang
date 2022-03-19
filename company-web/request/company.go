package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/company-web/global"
	"hr-saas-go/company-web/utils"
	"net/http"
)

type CompanySaveRequest struct {
	Name     string   `json:"name" form:"name" binding:"required,mobile"`
	Desc     string   `json:"desc" form:"desc" binding:"required,min=3"`
	Website  string   `json:"website" form:"website" binding:"required"`
	Config   string   `json:"config" form:"config" binding:"required"`
	Address  string   `json:"address" form:"address" binding:"required"`
	Tags     []string `json:"tags" form:"tags" binding:"required"`
	Info     string   `json:"info" form:"info" binding:"required"`
	ParentId int64    `json:"parent_id" form:"creator_id" binding:"required"`
}

func CompanySaveRequestGet(c *gin.Context) (CompanySaveRequest, error) {
	var request CompanySaveRequest
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
