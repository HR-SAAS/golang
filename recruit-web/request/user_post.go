package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/utils"
	"net/http"
)

type UserPostSaveRequest struct {
	PostId     int64  ` json:"post_id" form:"post_id" binding:""`
	ResumeId   int64  ` json:"resume_id" form:"resume_id" binding:""`
	ResumeType string ` json:"resume_type" form:"resume_type" binding:""`
	ResumeName string ` json:"resume_name" form:"resume_name" binding:""`
	Resume     string ` json:"resume" form:"resume" binding:""`
	ReviewId   int64  ` json:"review_id" form:"review_id" binding:""`
	Status     int32  ` json:"status" form:"status" binding:""`
	CompanyId  int64  ` json:"company_id" form:"company_id" binding:""`
	Remark     string ` json:"remark" form:"remark" binding:""`
}

func UserPostSaveRequestGet(c *gin.Context) (UserPostSaveRequest, error) {
	var request UserPostSaveRequest
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
