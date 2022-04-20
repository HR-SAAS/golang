package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/message-web/global"
	"hr-saas-go/message-web/utils"
	"net/http"
)

type ResumeSaveRequest struct {
	Name    string   `json:"name" form:"name" binding:""`
	Type    int32    `json:"type" form:"type" binding:"required"`
	Tag     []string `json:"tag" form:"tag" binding:""`
	Content string   `json:"content" form:"content" binding:"required"`
	Status  int32    `json:"status" form:"status" binding:"oneof=0 1"`
}

func ResumeSaveRequestGet(c *gin.Context) (ResumeSaveRequest, error) {
	var request ResumeSaveRequest
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
