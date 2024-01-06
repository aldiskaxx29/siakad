package req

type Mahasiswa struct{
	Name string `json:"name" validate:"required"`
	Study string `json:"study" validate:"required"`
	Photo string `json:"photo"`
}