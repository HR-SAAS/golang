package utils

import "net"

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
