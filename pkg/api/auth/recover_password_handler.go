package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func recoverPassword(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			NewPassword string `json:"new_password"`
			Token       string `json:"token"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][RecoverPassword][Request] newPassword = %v token = %v \n", payload.NewPassword, payload.Token)

		if err := ctx.AuthService.RecoverPassword(payload.NewPassword, payload.Token); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare reponse
		fmt.Println("[Auth][RecoverPassword][Response]")
		return c.Status(fiber.StatusOK).JSON(Response{})
	}
}
