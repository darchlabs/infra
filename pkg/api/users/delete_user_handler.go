package usersapi

import (
	"fmt"

	"github.com/darchlabs/infra/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func deleteUser(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Printf("[Users][Delete][Init]")

		// get token param
		userID := c.Params("id")

		if err := ctx.UserService.Delete(&users.User{ID: userID}); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		fmt.Printf("[Users][Delete][Response]")
		return c.Status(fiber.StatusOK).JSON(Response{})
	}
}
