package common

import (
	"os"
	"time"
)

type DefaultFieldTable struct {
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy string    `json:"modified_by" db:"modified_by"`
}

func (d *DefaultFieldTable) SetDefaultField() {
	hostname, _ := os.Hostname()

	d.CreatedAt = time.Now()
	d.CreatedBy = hostname
	d.ModifiedBy = hostname
	d.ModifiedAt = time.Now()
}
