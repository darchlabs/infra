package projectapi

import (
	"encoding/json"
	"fmt"

	"github.com/darchlabs/infra/pkg/project"
	"github.com/gofiber/fiber/v2"
)

func create(ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		// define payload struct
		payload := &struct {
			ProjectInput *project.ProjectInput `json:"project"`
		}{}

		// read and decode payload
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// parse project input to struct
		input := project.To(payload.ProjectInput)

		// debug
		fmt.Printf("[Project][Create] project=%+v \n", input)

		// create a project
		p, err := ctx.ProjectService.Create(input, ctx.Env)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Error: err.Error(),
			})
		}

		// debug
		fmt.Printf("%+v \n", p)

		// TODO(ca): maybe is neccesary to return 201 code
		return c.Status(fiber.StatusOK).JSON(Response{
			Data: input,
		})
	}
}
