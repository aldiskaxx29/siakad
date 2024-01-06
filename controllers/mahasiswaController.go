package controllers

import (
	"fmt"
	"log"
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
	"github.com/aldiskaxx29/go-fiber-crud/models/req"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MahasiswaControllerGetAll(c *fiber.Ctx) error {
	var mahasiswa []entity.Mahasiswa
	err := database.DB.Find(&mahasiswa).Error

	if err != nil{
		log.Println(err)
	}

	return c.JSON(fiber.Map{
		"status" : true,
		"data" : mahasiswa,
	})
}

func MahasiswaControllerGetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	var mahasiswa []entity.Mahasiswa

	if err := database.DB.First(&mahasiswa, id); err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Mahasiswa Not Founc",
		})
	}

	return c.JSON(fiber.Map{
		"status":true,
		"data":mahasiswa,
	})
}


func MahasiswaControllerSave(c *fiber.Ctx) error {
	mahasiswa := new(req.Mahasiswa)

	if err := c.BodyParser(mahasiswa); err != nil {
		return err
	}

	validation := validator.New()

	if err := validation.Struct(mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Failed input new mahasiswa",
			"error": err.Error(),
		})
	}

	var filenames string

	filename := c.Locals("filename")

	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "Upload file photo is required",
			"error": err.Error(),
		})
	} else{
		filenames = fmt.Sprintf("%v", filename)
	}


	newMahasiswa := entity.Mahasiswa{
		Name : mahasiswa.Name,
		Study : mahasiswa.Study,
		Photo : filenames,
	}

	if err := database.DB.Create(&newMahasiswa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed create new mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data": newMahasiswa,
	})
}

func MahasiswaControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var mahasiswa []entity.Mahasiswa

	err := database.DB.First(&mahasiswa, id)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Mahasiswa Not Found",
		})
	}

	database.DB.Delete(&mahasiswa)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"message":"Mahasiswa Success Deleted",
	})
}