package dto

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Branch string  `json:"branch"`
	Price  float64 `json:"price"`
}
