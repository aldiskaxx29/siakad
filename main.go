package main

import (
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/routes"
	"github.com/aldiskaxx29/go-fiber-crud/database/migration"
	"github.com/gofiber/fiber/v2"
)



func main(){
	database.ConnectDB()
	migration.RunMigrate()
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"sttaus" : true,
	// 	})
	// })

	routes.RouteApp(app)

	app.Listen(":8000")
}
