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
	Mobile    string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password  string `json:"password" form:"password" binding:"required,min=3"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" form:"captcha_id" binding:"required"`
}

func ValidateMobile(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	ok, _ := regexp.MatchString(`^1[235789]\d{9}$`, str)
	return ok
}

type MobileRequest struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Type   string `json:"type" form:"type" binding:"required,mobile"`
}

func MobileSmsLoginRequestGet(c *gin.Context) (MobileRequest, error) {
	var mobileRequest MobileRequest
	if err := c.ShouldBind(&mobileRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": utils.RemoveTopName(e.Translate(global.Trans)),
			})
		}
		//handle error
		return mobileRequest, err
	}
	return mobileRequest, nil
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
