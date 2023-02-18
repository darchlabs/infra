package authapi

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func login(ctx HandlerContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// define payload struct
		payload := &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		// read body and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		fmt.Printf("[Auth][Login][Request] email = %v password = %v \n", payload.Email, payload.Password)

		// login with email and password on users-api
		resLogin, err := ctx.AuthService.Login(payload.Email, payload.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(Response{
				Error: err.Error(),
			})
		}

		// prepare response
		res := Response{
			Data: resLogin.Data,
			Meta: resLogin.Meta,
		}

		fmt.Printf("[Auth][Login][Response] %v \n", res)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
