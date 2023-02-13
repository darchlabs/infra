package usersapi

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getUser(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println("[Users][Get][Init]")

		// get token param
		userID := c.Params("id")

		user, err := ctx.UserService.GetByID(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: user,
		}

		fmt.Printf("[Users][Get][Response] %v\n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
