package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/infra/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func signup(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			User *users.User `json:"user"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][Signup][Request] user = %v \n", payload)

		// signup created user
		resSignup, err := ctx.AuthService.Signup(payload.User)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: resSignup.Data,
			Meta: resSignup.Meta,
		}

		fmt.Printf("[Auth][Signup][Response] %v \n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
