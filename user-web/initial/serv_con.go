package initial

import (
	"fmt"
	"google.golang.org/grpc"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/proto"
	"hr-saas-go/user-web/utils"
)

func InitialCon() {
	host := ""
	port := 0

	services := utils.FilterService(fmt.Sprintf(`Service == "%s"`, global.Config.UserSrvInfo.Name))

	for _, item := range services {
		host = item.Address
		port = item.Port
		break
	}

	con, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	global.UserServCon = proto.NewUserClient(con)
}
