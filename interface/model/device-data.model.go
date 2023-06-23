package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceData struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	DeviceID  string  `json:"device_id" gorm:"size:191"`
	Speed     float64 `json:"speed"`
	Distance  float64 `json:"distance"`
	DateTime  string  `json:"date_time"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *DeviceData) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
