package controller

import (
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris/v12"
	"kadam_test/internal/application/loaders"
	"kadam_test/internal/repository/clicks"
	"kadam_test/internal/service/clicker"
	"kadam_test/internal/transport/http/router"
)

type Controller struct {
	Transport        *loaders.HTTP
	ClickerService   *clicker.Service
	ClicksRepository *clicks.Repository
}

func NewController(http *loaders.HTTP, clickerService *clicker.Service, clicksRepository *clicks.Repository) {
	controller := Controller{
		Transport:        http,
		ClickerService:   clickerService,
		ClicksRepository: clicksRepository,
	}

	iris.RegisterOnInterrupt(func() {
		err := controller.Transport.Iris.Shutdown(*controller.Transport.CTX)
		if err != nil {
			logger.Error(err)
		}
	})

	irisRouter := router.NewRouter(controller.Transport.Iris, controller.ClickerService, controller.ClicksRepository)

	err := irisRouter.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		logger.Error(err)
	}

}
