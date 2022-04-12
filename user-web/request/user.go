package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
	"net/http"
	"regexp"
)

type MobileLoginByPasswordRequest struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password string `json:"password" form:"password" binding:"required,min=3"`
	//Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	//CaptchaId string `json:"captcha_id" form:"captcha_id" binding:"required"`
}

func MobileLoginRequestGet(c *gin.Context) (MobileLoginByPasswordRequest, error) {
	var mobileLoginRequest MobileLoginByPasswordRequest
	if err := c.ShouldBind(&mobileLoginRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		}
		//handle error
		return mobileLoginRequest, err
	}
	return mobileLoginRequest, nil
}

type MobileLoginByCodeRequest struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Code   string `json:"code" form:"code" binding:"required"`
}

func MobileLoginByCodeRequestGet(c *gin.Context) (MobileLoginByCodeRequest, error) {
	var mobileLoginRequest MobileLoginByCodeRequest
	if err := c.ShouldBind(&mobileLoginRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		}
		//handle error
		return mobileLoginRequest, err
	}
	return mobileLoginRequest, nil
}

func ValidateMobile(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	ok, _ := regexp.MatchString(`^1[235789]\d{9}$`, str)
	return ok
}

type MobileRegisterRequest struct {
	Mobile      string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password    string `json:"password" form:"password" binding:"required,min=3"`
	Code        string `json:"code" form:"code" binding:"required"`
	CurrentRole int32  `json:"current_role" form:"current_role" binding:"required,oneof=1 2"`
}

func MobileRegisterRequestGet(c *gin.Context) (MobileRegisterRequest, error) {
	var mobileRegister MobileRegisterRequest
	if err := c.ShouldBind(&mobileRegister); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		} else {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("错误的请求"))
		}
		return mobileRegister, err
	}
	return mobileRegister, nil
}
