package routes

import (
	"github.com/aldiskaxx29/go-fiber-crud/controllers"
	"github.com/aldiskaxx29/go-fiber-crud/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/aldiskaxx29/go-fiber-crud/config"
)


func RouteApp(c *fiber.App){
	c.Static("/public", config.ProjectRootPath+"/public/images")
	c.Static("/public", config.ProjectRootPath+"/public/produk")

	c.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":"Hello Word",
		})
	})
	
	c.Post("/api/login", controllers.LoginControllerCheck)
	c.Delete("/api/logout", controllers.Logout)

	c.Post("/api/user/save", middleware.AuthMiddleware, controllers.UserControllerSave)
	c.Get("/api/user/getAll", middleware.AuthMiddleware ,controllers.UserControllerShow)
	c.Get("/api/user/getOne/:id", middleware.AuthMiddleware, controllers.UserControllerGetOne)
	c.Delete("/api/user/delete/:id", middleware.AuthMiddleware, controllers.UserControllerDelete)

	c.Get("/api/mahasiswa/getAll", controllers.MahasiswaControllerGetAll)
	c.Get("/api/mahasiswa/getOne:id", controllers.MahasiswaControllerGetOne)
	c.Post("/api/mahasiswa/save", middleware.HandleSingleFile, controllers.MahasiswaControllerSave)
	c.Delete("/api/mahasiswa/delete/:id", controllers.MahasiswaControllerDelete)

	
}