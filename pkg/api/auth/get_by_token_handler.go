package authapi

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getByToken(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// get token param
		token := c.Params("token")

		fmt.Printf("[Auth][GetByToken][Request] token = %v \n", token)

		// get auth by token
		auth, err := ctx.AuthService.GetByToken(token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: auth,
		}

		fmt.Printf("[Auth][GetByToken][Response] %v \n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
