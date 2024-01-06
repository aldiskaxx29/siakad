package middleware

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	file, errFile := c.FormFile("photo")
	if errFile != nil {
		log.Println("Error ", errFile)
		return errFile // tambahkan return error jika terjadi kesalahan saat memproses file
	}

	if file == nil || file.Size <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File upload is empty or null.",
		})
	}

	var filename string // Menggunakan string biasa, bukan pointer

	if file != nil {
		filename = file.Filename // Assign nilai nama file ke variabel filename

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
		if errSaveFile != nil {
			log.Println("Fail to store file")
			return errSaveFile // tambahkan return error jika gagal menyimpan file
		}
	} else {
		log.Println("Nothing file to upload")
		return fmt.Errorf("No file uploaded") // tambahkan return error jika tidak ada file yang diunggah
	}

	c.Locals("filename", filename) // Menggunakan nilai filename yang sudah didapat

	return c.Next()
}