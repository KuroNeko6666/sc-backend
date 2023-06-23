package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	UserID    string   `json:"user_id" gorm:"size:191"`
	Items     []Device `json:"items" gorm:"foreignKey:DeviceID"`
	Status    string   `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
