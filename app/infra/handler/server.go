package handler

import (
	"cloud.google.com/go/spanner"
	"practice.dev/api"
	"practice.dev/infra/repository/database"
	"practice.dev/usecase/service"
)

type Server struct {
	actionService *service.ActionService

	api.UnimplementedActionServiceServer
}

func NewServer(cli *spanner.Client) *Server {
	return &Server{
		actionService: service.NewActionService(
			database.NewMissionRepository(cli),
			database.NewUserRepository(cli),
		),
	}
}
