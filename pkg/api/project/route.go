package projectapi

import (
	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data  interface{} `json:"data"`
	Meta  interface{} `json:"meta"`
	Error interface{} `json:"error"`
}

type Context struct {
	ProjectService project.Service
	Env            env.Env
}

func Route(app *fiber.App, ctx Context) {
	app.Post("/api/v1/projects", create(ctx))
}
