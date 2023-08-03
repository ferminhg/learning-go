package bootstrap

import "github.com/ferminhg/learning-go/internal/infra/server"

func Run(host string, port uint, brokerList []string) error {
	srv := server.New(host, port, brokerList)
	return srv.Run()
}
