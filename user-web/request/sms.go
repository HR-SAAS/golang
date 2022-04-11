package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
	"net/http"
)

type MobileRequest struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Type   string `json:"type" form:"type" binding:"required,oneof=register login'"`
}

func MobileSmsLoginRequestGet(c *gin.Context) (MobileRequest, error) {
	var mobileRequest MobileRequest
	if err := c.ShouldBind(&mobileRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		}
		//handle error
		return mobileRequest, err
	}
	return mobileRequest, nil
}
