package rpc

import (
	"github.com/mixarchitecture/microp/server/rpc"
	protos "github.com/turistikrota/service.business/protos/business"
	"github.com/turistikrota/service.business/src/app"
	"google.golang.org/grpc"
)

type Server struct {
	port int
	app  app.Application
	protos.BusinessListServiceServer
}

type Config struct {
	Port int
	App  app.Application
}

func New(cnf Config) Server {
	return Server{
		app:  cnf.App,
		port: cnf.Port,
	}
}

func (h Server) Load() {
	rpc.RunServer(h.port, func(server *grpc.Server) {
		protos.RegisterBusinessListServiceServer(server, h)
	})
}
