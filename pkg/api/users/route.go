package usersapi

import (
	"github.com/gofiber/fiber/v2"

	"github.com/darchlabs/infra/pkg/users"
)

// Response ...
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type HandlerContext struct {
	UserService users.Service
}

func Route(app *fiber.App, ctx HandlerContext) {
	app.Get("/api/v1/users", listUsers(ctx))
	app.Post("/api/v1/users", createUser(ctx))
	app.Post("/api/v1/users/:id", getUser(ctx))
	app.Post("/api/v1/users/:id", updateUser(ctx))
	app.Post("/api/v1/users/:id", deleteUser(ctx))
}
