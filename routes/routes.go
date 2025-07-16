package routes

import (
	"go-crud-siswa/controllers"
	"go-crud-siswa/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	siswa := api.Group("/siswa", middlewares.JWTProtected())
	siswa.Get("/", controllers.GetAllSiswa)
	siswa.Post("/", controllers.CreateSiswa)
	siswa.Put("/:id", controllers.UpdateSiswa)
	siswa.Delete("/:id", controllers.DeleteSiswa)
}
