package loaders

import (
	"context"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type HTTP struct {
	Iris *iris.Application
	CTX  *context.Context
}

// InitHTTPServer - init instance http server
func InitHTTPServer(ctx *context.Context) *HTTP {

	server := HTTP{
		Iris: iris.Default(),
		CTX:  ctx,
	}

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	})

	server.Iris.UseRouter(crs)

	return &server
}
