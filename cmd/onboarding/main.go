package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kelseyhightower/envconfig"

	"github.com/darchlabs/infra/internal/env"
	projectapi "github.com/darchlabs/infra/pkg/api/project"
	projectservice "github.com/darchlabs/infra/pkg/service/project"
	"github.com/darchlabs/infra/pkg/util"
)

func main() {
	// load env values
	var env env.Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// initialize fiber
	api := fiber.New()
	api.Use(cors.New())
	api.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// initialize project service
	svc := projectservice.NewService()

	// initialize project router
	projectapi.Route(api, projectapi.Context{
		ProjectService: svc,
		Env:            env,
	})

	// run proccess
	go func() {
		api.Listen(fmt.Sprintf(":%s", env.Port))
	}()

	// listen interrupt
	quit := make(chan struct{})
	util.ListenInterrupt(quit)
	<-quit
	util.GracefullShutdown()
}
