package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"hr-saas-go/user-web/global"
)

func getClient() *api.Client {
	config := global.Config.ConsulConfig
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

func Register(name string, id string, address string, port int, tags []string) error {

	client := getClient()

	register := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{
			Name:                           name,
			Interval:                       "5s",
			Timeout:                        "5s",
			HTTP:                           fmt.Sprintf("http://%s:%d/api/health", address, port),
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err := client.Agent().ServiceRegister(&register)
	if err != nil {
		panic(err)
	}
	return nil
}

func FilterService(filter string) map[string]*api.AgentService {
	client := getClient()
	withFilter, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil
	}
	return withFilter
}
