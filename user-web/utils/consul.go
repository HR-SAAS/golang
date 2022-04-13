package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulRegister struct {
	host string
	port int
}

func getClient(host string, port int) *api.Client {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

func NewRegister(host string, port int) Register {
	return &ConsulRegister{
		host: host,
		port: port,
	}
}

func (c *ConsulRegister) Register(name string, id string, address string, port int, tags []string) error {
	client := getClient(c.host, c.port)
	register := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{
			Name:                           name,
			Interval:                       "10s",
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

func (c *ConsulRegister) FilterService(filter string) map[string]*api.AgentService {
	client := getClient(c.host, c.port)
	withFilter, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil
	}
	return withFilter
}

func (c *ConsulRegister) Deregister(id string) error {
	client := getClient(c.host, c.port)
	err := client.Agent().ServiceDeregister(id)

	return err
}
