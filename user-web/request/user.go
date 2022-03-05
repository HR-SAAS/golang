package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
	"net/http"
	"regexp"
)

type MobileLoginRequest struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password string `json:"password" form:"password" binding:"required,min=3"`
}

func ValidateMobile(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	ok, _ := regexp.MatchString(`^1[235789]\d{9}$`, str)
	return ok
}

func MobileLoginRequestGet(c *gin.Context) (MobileLoginRequest, error) {
	var mobileLoginRequest MobileLoginRequest
	if err := c.ShouldBind(&mobileLoginRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": utils.RemoveTopName(e.Translate(global.Trans)),
			})
		}
		//handle error
		return mobileLoginRequest, err
	}
	return mobileLoginRequest, nil
}
