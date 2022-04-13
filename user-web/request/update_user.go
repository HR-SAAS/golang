package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
	"net/http"
)

type UpdateUserRequest struct {
	Id            int64  `json:"id" form:"id" binding:""`
	NickName      string `json:"nick_name" form:"nick_name" binding:"required"`
	Avatar        string `json:"avatar" form:"avatar" binding:""`
	Sex           int32  `json:"sex" form:"sex" binding:"min=0,max=3"`
	OldPassword   string `json:"old_password" form:"old_password" binding:""`
	NewPassword   string `json:"new_password" form:"new_password" binding:""`
	CheckPassword string `json:"check_password" form:"check_password"`
}

func UpdateUserRequestGet(c *gin.Context) (UpdateUserRequest, error) {
	var mobileLoginRequest UpdateUserRequest
	if err := c.ShouldBind(&mobileLoginRequest); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ErrorJson("验证错误", utils.RemoveTopName(e.Translate(global.Trans))))
		} else {
			c.JSON(http.StatusBadRequest, utils.ErrorJson(err.Error()))
		}
		//handle error
		return mobileLoginRequest, err
	}
	return mobileLoginRequest, nil
}
