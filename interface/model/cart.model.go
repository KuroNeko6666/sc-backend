package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	UserID    string   `json:"user_id" gorm:"size:191"`
	Items     []Device `json:"items" gorm:"many2many:cart_devices"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
