package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(ctx *gin.Context) (page, limit int32) {
	page = 1
	limit = 15

	if p, _ := strconv.Atoi(ctx.Query("page")); p > 0 {
		page = int32(p)
	}

	if l, _ := strconv.Atoi(ctx.Query("limit")); l > 0 {
		page = int32(l)
	}
	return
}
