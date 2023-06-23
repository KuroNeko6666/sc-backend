package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceAddress struct {
	AddressID string    `json:"-" gorm:"primaryKey"`
	DeviceID  string    `json:"-" gorm:"size:191"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *DeviceAddress) BeforeCreate(tx *gorm.DB) (err error) {
	u.AddressID = uuid.NewString()
	return
}
