package usersapi

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/infra/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func updateUser(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Printf("[Users][Update][Init]")

		// get token param
		userID := c.Params("id")

		payload := &struct {
			User *users.User `json:"user"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Users][Update][Request] payload = %v\n", payload)

		payload.User.ID = userID
		if err := ctx.UserService.Update(payload.User); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare reponse
		res := Response{
			Data: payload.User,
		}

		fmt.Printf("[Users][Update][Response] %v\n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
