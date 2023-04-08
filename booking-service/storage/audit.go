package storage

import "time"

type Audit struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Identifier string    `json:"identifier"`
	Action     string    `json:"action"`
	Data       string    `json:"data"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
