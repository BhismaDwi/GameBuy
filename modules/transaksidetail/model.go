package transaksidetail

import (
	"GameBuy/modules/game"
	"time"
)

type TransaksiDetail struct {
	ID          int       `json:"id"`
	TransaksiID int       `json:"transaksi_id"`
	GameID      int       `json:"game_id"`
	Game        game.Game `json:"game"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}
