package storage

import "time"

type Product struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	BranchId  int       `json:"branch_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
