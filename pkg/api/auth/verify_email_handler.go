package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func verifyEmail(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			Token string `json:"token"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][VerifyEmail][Request] token = %v \n", payload.Token)

		if err := ctx.AuthService.VerifyEmail(payload.Token); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Println("[Auth][VerifyEmail][Response]")
		return c.Status(fiber.StatusOK).JSON(Response{})
	}
}
