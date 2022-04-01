package main

import (
	"context"
	"fmt"
	"github.com/ivahaev/go-logger"
	"github.com/joho/godotenv"
	"kadam_test/internal/application/loaders"
	"kadam_test/internal/migration"
	"kadam_test/internal/repository/clicks"
	"kadam_test/internal/service/clicker"
	"kadam_test/internal/transport/http/controller"
	"time"
)

const (
	shutDownDuration = 5 * time.Second
)

func main() {
	err := logger.SetLevel("debug")
	if err != nil {
		panic(fmt.Sprintf(`failed init logs, error: %s`, err.Error()))
	}

	err = godotenv.Load(".env")
	if err != nil {
		panic("fatal error loading .env file")
	}

	ctxApp, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := loaders.NewDataBase()
	if err != nil {
		panic(fmt.Sprintf(`fatal error init database, error: %s`, err.Error()))
	}

	err = migration.CreateSchema(db)
	if err != nil {
		panic(fmt.Sprintf(`fatal error migration database, error: %s`, err.Error()))
	}

	// NewController - init controller http server.
	go controller.NewController(
		loaders.InitHTTPServer(&ctxApp),
		clicker.NewClickerService(),
		clicks.NewClickRepository(db),
	)

	<-loaders.GracefulShutdown()
	_, forceCancel := context.WithTimeout(ctxApp, shutDownDuration)

	logger.Info("Graceful Shutdown")
	defer forceCancel()
}
