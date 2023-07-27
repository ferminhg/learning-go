package bootstrap

import "github.com/ferminhg/learning-go/internal/infra/server"

func Run(host string, port uint) error {
	srv := server.New(host, port)
	return srv.Run()
}
