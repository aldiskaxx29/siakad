package controllers

import (
	"time"
	// "errors"
	"fmt"
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
	// "github.com/aldiskaxx29/go-fiber-crud/models/req"
	// "github.com/aldiskaxx29/go-fiber-crud/internal/appctx"
	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type JWTClaims struct{
	Email string `json:"email"`
	jwt.StandardClaims
}

type Users struct{
	Id uint `json:"id" gorm:"primaryKey"`
	// Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	// Age int `json:"age"`
}

var secretKey = []byte("secret")
var err error
func LoginControllerCheck(c *fiber.Ctx) error {
	var reqUser entity.User
	// var reqUser Users

	

	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}

	pass := reqUser.Password
	err := database.DB.Where("email = ?", reqUser.Email).First(&reqUser).Error
	if err != nil {
    if err == gorm.ErrRecordNotFound {
        // Tidak ada pengguna yang ditemukan dengan email yang diberikan
        return c.SendStatus(fiber.StatusNotFound)
    }
    // Penanganan kesalahan lainnya
    fmt.Println("Error:", err)
    return c.SendStatus(fiber.StatusInternalServerError)
	}

	storedPassword := reqUser.Password
	// Lakukan sesuatu dengan data user yang ditemukan
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(pass))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// buat token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = reqUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // contoh token selama 1 hari expired nya

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status" : true,
		"data": reqUser,
		"token" : tokenString,
	})
}	

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message":"Logout Successful",
	})
}

