package transaksi

import (
	"GameBuy/modules/transaksidetail"
	"GameBuy/modules/users"
	"time"
)

type Transaksi struct {
	ID            int                               `son:"id"`
	TglTransakksi time.Time                         `json:"tgl_transaksi"`
	UserID        int                               `json:"user_id"`
	User          users.User                        `json:"user"`
	TotalHarga    int                               `json:"total_harga"`
	Details       []transaksidetail.TransaksiDetail `json:"details"`
	CreatedAt     time.Time                         `json:"created_at"`
	CreatedBy     string                            `json:"created_by"`
	ModifiedAt    time.Time                         `json:"modified_at"`
	ModifiedBy    string                            `json:"modified_by"`
}
