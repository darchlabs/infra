package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/infra/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func forgotPassword(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			Email string `json:"email"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][ForgotPassword][Request] email = %v \n", payload.Email)

		token, err := ctx.AuthService.ForgotPassword(payload.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: auth.MetaToken{
				Token: token,
			},
		}

		fmt.Printf("[Auth][ForgotPassword][Response] %v \n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
