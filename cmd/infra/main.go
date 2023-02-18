package main

import (
	"fmt"
	"log"
	"os"

	authapi "github.com/darchlabs/infra/pkg/api/auth"
	authsapi "github.com/darchlabs/infra/pkg/api/auth"
	usersapi "github.com/darchlabs/infra/pkg/api/users"
	authservice "github.com/darchlabs/infra/pkg/service/auth"
	userservice "github.com/darchlabs/infra/pkg/service/users"
	"github.com/darchlabs/infra/pkg/storage"
	authstorage "github.com/darchlabs/infra/pkg/storage/auth"
	userstorage "github.com/darchlabs/infra/pkg/storage/users"
	"github.com/darchlabs/infra/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// get PORT environment value
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("invalid PORT environment value")
	}

	// get DSN environment value
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("invalid DSN environment value")
	}

	// initialize fiber
	api := fiber.New()
	api.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// initialize storage
	storage, err := storage.NewPostgres(dsn)
	if err != nil {
		log.Fatal("invalid DSN environment value")
	}

	// initialize service storages
	userStorage := userstorage.New(storage)
	authStorage := authstorage.New(storage)

	// initialize services
	userSvc := userservice.NewService(userStorage)
	authSvc := authservice.NewService(authStorage, userSvc)

	// initialize routers
	usersapi.Route(api, usersapi.HandlerContext{
		UserService: userSvc,
	})
	authapi.Route(api, authsapi.HandlerContext{
		AuthService: authSvc,
	})

	// run proccess
	go func() {
		api.Listen(fmt.Sprintf(":%s", port))
	}()

	// listen interrupt
	quit := make(chan struct{})
	util.ListenInterrupt(quit)
	<-quit
	util.GracefullShutdown()
}
