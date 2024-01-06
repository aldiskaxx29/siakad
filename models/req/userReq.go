package req

type UserReq struct{
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required"`
	Age int `json:"age" validate:"required"`
}
