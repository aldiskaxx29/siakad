package migration

import(
	"fmt"
	"github.com/aldiskaxx29/go-fiber-crud/database"
	"github.com/aldiskaxx29/go-fiber-crud/models/entity"
)

func RunMigrate(){
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Mahasiswa{})
	if err != nil{
		panic(err)
	}

	fmt.Println("Success To Migrate")

}