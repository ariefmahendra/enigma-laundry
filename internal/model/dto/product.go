package dto

type ProductRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price int    `json:"price"`
}
