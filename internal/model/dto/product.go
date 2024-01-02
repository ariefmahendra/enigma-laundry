package dto

type ProductRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price int    `json:"price"`
}
