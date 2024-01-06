package controllers

import (
	"log"
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
	"github.com/aldiskaxx29/go-fiber-crud/models/req"
	"github.com/aldiskaxx29/go-fiber-crud/internal/appctx"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	// "net/http"
)

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil{
		log.Println(err)
	}

	return c.JSON(appctx.Response{
		Message: appctx.ResponseNameSuccess,
		Data: users,
	})
}

func UserControllerGetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var users []entity.User

	if err := database.DB.First(&users, id); err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status" : false,
			"message": "User Not Found",
		})
	} 

	return c.JSON(fiber.Map{
		"status" : true,
		"data" : users,
	})

}

func UserControllerSave(c *fiber.Ctx) error {
	user := new(req.UserReq)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation:= validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Failed input new user",
			"error": err.Error(),
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := entity.User{
		Name: user.Name,
		Password: string(hashPassword),
		Email: user.Email,
		Age: user.Age,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed create new user",
		})
	}

	return c.JSON(fiber.Map{
		"status":true,
		"data": newUser,
	})

}

func UserControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var users []entity.User

	err := database.DB.First(&users, id)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"User Not Found",
		})
	}
	
	database.DB.Delete(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"message": "User Success Deleted",
	})

}

func UserControllerUdate(c *fiber.Map) error {
	userReq := new(req.UserReq)
	if err := c.BodyParser(&userReq); err := nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"error",
		})
	}

	var user entity.User

	id := c.Params("id")

	if err := database.DB.First(&user, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"user not found",
		})
	}

	validation := validation.New()
	if err := validation.Struct(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error",
			"error": err.Error(),
		})
	}

	newUser = entity.User{
		Name: entity.Name,
		Email: entity.Email,
		Password: entity.Password,
		Age: entity.Age
	}

	if err := database.Create(&newUser)


}

// func handleUploadFile(c *fiber.Ctx) error {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message":"Error upload file",
// 		})
// 	}

// 	files := form.File["images"]
// 	name := form.Value["name"][0]
// 	password: form.Value["password"][0]
// 	email: form.Value["email"][0]
// 	age: form.Value["age"][0]
// }

// type File struct {
// 	ID       uint `gorm:"primary_key"`
// 	Filename string
// 	Username string
// }

// Handler untuk mengunggah file
// func UploadFile(c *fiber.Ctx) error {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 			return c.Status(http.StatusInternalServerError).SendString("Error uploading file")
// 	}

// 	files := form.File["files"]
// 	username := form.Value["username"][0] // Ambil nama pengguna dari form

// 	for _, file := range files {
// 			// Simpan file ke direktori
// 			err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
// 			if err != nil {
// 					return c.Status(http.StatusInternalServerError).SendString("Error saving file")
// 			}

// 			// Simpan informasi file (nama file, username) ke database
// 			err = SaveFileInfoToDatabase(file.Filename, username)
// 			if err != nil {
// 					return c.Status(http.StatusInternalServerError).SendString("Error saving file info to database")
// 			}
// 	}

// 	return c.SendString("Files uploaded successfully")
// }


// Fungsi untuk menyimpan informasi file ke dalam database
// func SaveFileInfoToDatabase(filename string, username string) error {
// 	db, err := gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/dbname"), &gorm.Config{})
// 	if err != nil {
// 			return err
// 	}
// 	defer db.Close()

// 	// Membuat entitas baru
// 	file := File{
// 			Filename: filename,
// 			Username: username,
// 	}

// 	// Menyimpan informasi file ke dalam tabel di database
// 	if err := db.Create(&file).Error; err != nil {
// 			return err
// 	}

// 	return nil
// }
