package usersapi

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func listUsers(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println("[Users][List][Init]")

		listedUsers, err := ctx.UserService.List()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: listedUsers,
		}

		fmt.Printf("[Users][List][Response] %v\n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
