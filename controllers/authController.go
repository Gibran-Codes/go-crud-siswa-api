package controllers

import (
	"go-crud-siswa/config"
	"go-crud-siswa/models"
	"go-crud-siswa/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal hash password"})
	}
	input.Password = string(hashedPassword)

	result := config.DB.Create(&input)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Email sudah digunakan"})
	}
	return c.JSON(fiber.Map{"message": "Berhasil register"})
}

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	var user models.User
	config.DB.First(&user, "email = ?", input.Email)
	if user.ID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "Email tidak ditemukan"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
	}
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal buat token"})
	}
	return c.JSON(fiber.Map{"token": token})
}
