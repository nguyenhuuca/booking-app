package model

type Product struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Branch string  `json:"artist"`
	Price  float64 `json:"price"`
}
