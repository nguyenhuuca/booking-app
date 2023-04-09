package dto

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Branch string  `json:"branch"`
	Price  float64 `json:"price"`
}

type AuditDto struct {
	Identifier string     `json:"identifier"`
	Action     string     `json:"action"`
	Data       ProductDto `json:"data"`
}
