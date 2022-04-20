package initial

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"hr-saas-go/message-web/global"
	"hr-saas-go/message-web/proto"
)

func InitialCon() {
	//host := ""
	//port := 0

	//services := utils.FilterService(fmt.Sprintf(`Service == "%s"`, global.Config.UserSrvInfo.Name))
	//
	//for _, item := range services {
	//	host = item.Address
	//	port = item.Port
	//	break
	//}
	//
	//con, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())

	initResume()
}

func initResume() {
	con, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python",
			global.Config.ConsulConfig.Host,
			global.Config.ConsulConfig.Port,
			global.Config.UserMessageSrv.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	global.UserMessageServCon = proto.NewUserMessageClient(con)
	// 统计服务
	global.MessageCounterService = proto.NewMessageCounterClient(con)
}
