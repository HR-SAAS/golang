package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/utils"
	"net/http"
)

type PostSaveRequest struct {
	CompanyId    int64    ` json:"company_id" form:"company_id" bindging:""`
	DepartmentId int64    ` json:"department_id" form:"department_id" bindging:""`
	CreatorId    int64    ` json:"creator_id," form:"creator_id," bindging:""`
	Type         int32    ` json:"type" form:"type" bindging:""`
	Name         string   ` json:"name" form:"name" bindging:""`
	Desc         string   ` json:"desc" form:"desc" bindging:""`
	Content      string   ` json:"content" form:"content" bindging:""`
	Experience   int32    ` json:"experience" form:"experience" bindging:""`
	Education    int32    ` json:"education" form:"education" bindging:""`
	Address      []string ` json:"address" form:"address" bindging:""`
	StartAt      string   ` json:"start_at" form:"start_at" bindging:""`
	EndAt        string   ` json:"end_at" form:"end_at" bindging:""`
	Status       int32    ` json:"status" form:"status" bindging:""`
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
