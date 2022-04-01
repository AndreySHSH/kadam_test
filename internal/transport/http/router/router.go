package router

import (
	"github.com/kataras/iris/v12"
	"kadam_test/internal/repository/clicks"
	"kadam_test/internal/service/clicker"
	"kadam_test/internal/transport/http/handlers"
)

type Router struct {
	iris *iris.Application
}

func NewRouter(iris *iris.Application, clickerService *clicker.Service, clicksRepository *clicks.Repository) *iris.Application {

	handler := handlers.Handlers{
		ClickerService:   clickerService,
		ClicksRepository: clicksRepository,
	}

	iris.Handle("GET", "/{bodyUrl}", handler.ClickBanner)

	return iris
}
