package utils

import (
	"errors"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer tcp.Close()
	return tcp.Addr().(*net.TCPAddr).Port, nil
}

func GetCurrentHost() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("获取ip失败")
}
