package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
	"github.com/aldiskaxx29/go-fiber-crud/database"
)

func ProdukControllerGetAll(c *fiber.Ctx) error {
	var produk []entity.Produk

	err := database.DB.Find(&produk).Error
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data": produk,
	});
}


func ProdukControllerGetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	var produk []entity.Produk

	err := database.DB.First(&produk, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": false,
			"message": "Produk Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data": produk,
	})
}
