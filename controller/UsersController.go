package controller

import "github.com/gofiber/fiber/v2"

type user struct {
	Name     string `query:"name"`
	Email    string `query:"email"`
	Password int    `query:"password"`
}

func UserController(c *fiber.Ctx) error {
	user := new(user)

	if err := c.QueryParser(user); err != nil {
		return err
	}
	return c.Render("index", fiber.Map{
		"baseUrl":  c.BaseURL(),
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
}
