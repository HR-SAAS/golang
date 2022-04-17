package initial

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"hr-saas-go/company-web/global"
	"hr-saas-go/company-web/proto"
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

	initCompany()
	initDepartment()
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

func initDepartment() {
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
	global.CompanyServCon = proto.NewCompanyClient(con)
}
