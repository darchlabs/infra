package usersapi

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/darchlabs/infra/pkg/users"
)

func createUser(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Printf("[Users][Create][Init]")

		payload := &struct {
			User *users.User `json:"user"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Users][Create][Request] payload = %v\n", payload)

		if err := ctx.UserService.Create(payload.User); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: payload.User,
		}

		fmt.Printf("[Users][Create][Response] %v\n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
