package controllers

import (
	"go-crud-siswa/config"
	"go-crud-siswa/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllSiswa(c *fiber.Ctx) error {
	var siswa []models.Siswa
	config.DB.Find(&siswa)
	return c.JSON(siswa)
}

func CreateSiswa(c *fiber.Ctx) error {
	var siswa models.Siswa
	if err := c.BodyParser(&siswa); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}
	config.DB.Create(&siswa)
	return c.JSON(siswa)
}

func UpdateSiswa(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    var siswa models.Siswa
    if err := config.DB.First(&siswa, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
    }

    var input models.Siswa
    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
    }

    // Update hanya field yang boleh diubah
    siswa.Nama = input.Nama
    siswa.Kelas = input.Kelas
    siswa.Umur = input.Umur

    config.DB.Save(&siswa)
    return c.JSON(siswa)
}


func DeleteSiswa(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var siswa models.Siswa
	if err := config.DB.First(&siswa, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
	}
	config.DB.Delete(&siswa)
	return c.JSON(fiber.Map{"message": "Data siswa dihapus"})
}