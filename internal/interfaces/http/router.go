package http

import (
	"github.com/Uno-count/Task-Management/internal/application"
	"github.com/Uno-count/Task-Management/internal/infrastructure/pbclient"
	"github.com/Uno-count/Task-Management/internal/interfaces/http/handlers"
	"github.com/pocketbase/pocketbase/core"
)

type Router struct {
	pbClient *pbclient.Client
}

func NewRouter(pbClient *pbclient.Client) *Router {
	return &Router{pbClient: pbClient}
}

func (r *Router) SetupRoutes(e *core.ServeEvent) {
	userRepo := pbclient.NewUserRepository(r.pbClient)
	userService := application.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := e.Router.Group("/api")
	api.POST("/register", userHandler.Register)
}
