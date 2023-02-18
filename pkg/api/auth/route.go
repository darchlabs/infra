package authapi

import (
	"github.com/gofiber/fiber/v2"

	"github.com/darchlabs/infra/pkg/auth"
)

// Response ...
type Response struct {
	Data  interface{} `json:"data,omitempty"` // {}
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"` // {"error": "invalid email, name or password"}
}

type HandlerContext struct {
	AuthService auth.Service
}

func Route(app *fiber.App, ctx HandlerContext) {
	app.Post("/api/v1/auth/login", login(ctx))
	app.Post("/api/v1/auth/signup", signup(ctx))
	app.Post("/api/v1/auth/verify-email", verifyEmail(ctx))
	app.Post("/api/v1/auth/logout", logout(ctx))
	app.Post("/api/v1/auth/forgot-password", forgotPassword(ctx))
	app.Post("/api/v1/auth/recover-password", recoverPassword(ctx))
}
