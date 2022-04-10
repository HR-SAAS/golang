package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type Register interface {
	Register(name string, id string, address string, port int, tags []string) error
	FilterService(filter string) map[string]*api.AgentService
	Deregister(id string) error
}

func RemoveTopName(fieldErr map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fieldErr {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func SuccessJson(data interface{}) gin.H {
	return MakeTrans(0, "success", data)
}

func ErrorJson(msg string) gin.H {
	return MakeTrans(1, msg, nil)

}

func MakeTrans(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
func HandleGrpcError(err error, ctx *gin.Context, errMsgs ...string) {
	errMsgArr := []string{
		"找不到", "错误参数", "内部错误", "其他错误",
	}
	for index, errMsg := range errMsgs {
		zap.S().Infof("change err %s", errMsg)
		errMsgArr[index] = errMsg
	}
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, errMsgArr[0], nil))
		case codes.InvalidArgument:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, errMsgArr[1], nil))
		case codes.Internal:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, errMsgArr[2], nil))
		default:
			ctx.JSON(http.StatusNotFound, MakeTrans(1, errMsgArr[3], nil))
		}
	}
}
