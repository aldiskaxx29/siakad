package req

type ProdukReq struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Description string `json:"description"`
}