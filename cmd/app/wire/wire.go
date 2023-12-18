package wire

import (
	"fiber-wire-template/internal/handler"
	"fiber-wire-template/internal/repository"
	"fiber-wire-template/internal/service"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/gredis"
	"fiber-wire-template/pkg/jwt"
	"fiber-wire-template/pkg/log"
	"fiber-wire-template/pkg/ozzodb"
	"fiber-wire-template/pkg/server"
	"fiber-wire-template/pkg/util/sid"
	"fiber-wire-template/route"
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	ozzodb.NewDb,
	gredis.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func NewApp(*config.Config, *log.Logger) (*server.FiberServer, error) {
	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server.NewFiberServer,
		route.SetupRoutes,
		sid.NewSid,
		jwt.NewJwt,
	))
	return &server.FiberServer{}, nil
}
