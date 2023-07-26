package router

import (
	"github.com/HiteshKumarMeghwar/L-M-S/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, userController *controller.UserController) *fiber.App {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	router.Post("/create_user", userController.CreateUser)
	router.Get("/user/:id", userController.GetUserById)

	return router
}
