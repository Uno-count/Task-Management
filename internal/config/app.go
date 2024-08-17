package config

import (
	"github.com/Uno-count/Task-Management/internal/infrastructure/pbclient"
	"github.com/Uno-count/Task-Management/internal/interfaces/http"
	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type App struct {
	PB     *pb.PocketBase
	Router *http.Router
}

func NewApp() *App {
	pocketBase := pb.New()
	pbClient := pbclient.NewClient(pocketBase)
	router := http.NewRouter(pbClient)

	return &App{
		PB:     pocketBase,
		Router: router,
	}
}

func (a *App) Start() error {
	a.PB.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		a.Router.SetupRoutes(e)
		return nil
	})

	return a.PB.Start()
}
