package company

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/company-web/utils"
	"net/http"
)

// 统计服务

func UserCountByCompany(ctx *gin.Context) {
	// 用户统计
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int{
		"count": 1,
	}))
}

func DepartmentCount(ctx *gin.Context) {
	// 部门统计
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int{
		"count": 10,
	}))
}

func ResumeCount(ctx *gin.Context) {
	// 简历统计
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]int{
		"count": 100,
	}))
}

func PostCount(ctx *gin.Context) {
	//岗位统计
}
