package initial

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
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

	initRecruit()
	initCompany()
	initUser()
}

func initUser() {
	con, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python",
			global.Config.ConsulConfig.Host,
			global.Config.ConsulConfig.Port,
			global.Config.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	global.UserServCon = proto.NewUserClient(con)
}

func initCompany() {
	con, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python",
			global.Config.ConsulConfig.Host,
			global.Config.ConsulConfig.Port,
			global.Config.CompanySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	global.DepartmentServCon = proto.NewDepartmentClient(con)
	global.CompanyServCon = proto.NewCompanyClient(con)
}

func initRecruit() {
	con, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python",
			global.Config.ConsulConfig.Host,
			global.Config.ConsulConfig.Port,
			global.Config.RecruitSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	global.RecruitCounterServiceServCon = proto.NewRecruitCounterServiceClient(con)
	global.PostServCon = proto.NewPostClient(con)
	global.UserPostServCon = proto.NewUserPostClient(con)
}
