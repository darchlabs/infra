package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func verifyToken(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			Token string `json:"token"`
			Kind  string `json:"kind"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][VerifyToken][Request] token = %v kind = %v \n", payload.Token, payload.Kind)

		t, err := ctx.AuthService.VerifyToken(payload.Token, payload.Kind)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: t,
		}

		fmt.Printf("[Auth][VerifyToken][Response] %v \n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
