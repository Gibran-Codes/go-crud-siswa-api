package main

import (
	"go-crud-siswa/config"
	"go-crud-siswa/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal load .env")
	}
	config.ConnectDatabase()
	app := fiber.New()
	routes.Setup(app)
	log.Fatal(app.Listen(":3000"))
}