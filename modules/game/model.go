package game

import "time"

type Game struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Harga      int       `json:"harga"`
	CategoryId int       `json:"category_id"`
	PlatformId int       `json:"platform_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
