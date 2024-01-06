package controllers

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/models/req"
	"github.com/go-playground/validator/v10"
)

func CategoryGetAllController(c *fiber.Ctx) error {
	var category []entity.Category

	err := database.DB.Find(&category).Error
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status" : true,
		"data": category,
	})
}

func CategoryGetOneController(c *fiber.Map) error {
	id := c.Params("id")

	var category = entity.Category
	err := database.DB.Find(&category, id).Error
	if err != nil {
		return c.Status(StatusInternalServerError).JSON(fiber.Map{
			"message":"Eroro category not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":true,
		"data": category,
	})
}

func CategorySaveController(c *fiber.Ctx) error {
	category := new(req.CategoryReq)

	if err := c.BodyParser(&category); err != nil {
		return c.Status(StatusInternalServerError).JSON(fiber.Map{
			"message": "error save data",
		})
	}

	valdiation := validation.New()
	if err:= validation.Struct(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Failed input category",
			"error": err.Error(),
		})
	}

	newCategiry := entity.Category{
		Name : category.Name,
		Description : category.Description,
		Image: category.Image
	}

	if err := database.DB.Create(&newCategiry); err != nil {
		return c.Sttaus(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"error",
		})
	}

	c.JSON(fiber.Map{
		"status": true,
	})

}

func CategoryUpdateController(c *fiber.Ctx) error {
	categoryReq := c.BodyParser(req.CategoryReq)
	id := c.Params("id")

	var category entity.Category

	if err := database.DB.First(&category, id); err != nil {
		return c.JSON(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error data ot found",
		})
	}

	newCategory = {
		Name : categoryReq.Name,
		Description : categoryReq.Description,
		Image : categoryReq.Image,
	}

	err := database.DB.Create(&newCategory).Error()
	if err != nil {
		return c.JSON(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"data not found",
		})
	}
	
}

func CategoryDeleteController(c *fiber.Ctx) error {
	id := c.Params("id")

	var category []entity.Category

	if err:= database.DB.First(&category, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "Error",
		})
	}

	database.DB.Delete(&category)

	return c.JSON(fiber.Map{
		"status": true,
		"message" : "Data berhasil di hapus",
	})
}