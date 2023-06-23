package model

import (
	"time"
)

type Device struct {
	ID        string        `json:"id" gorm:"primaryKey"`
	Model     string        `json:"model"`
	Address   DeviceAddress `json:"address" gorm:"foreignKey:DeviceID"`
	Users     []User        `json:"users" gorm:"many2many:user_devices"`
	Data      []DeviceData  `json:"data" gorm:"foreignKey:DeviceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
