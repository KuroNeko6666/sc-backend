package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name"`
	Username  string   `json:"username" gorm:"unique"`
	Email     string   `json:"email" gorm:"unique"`
	Password  string   `json:"-"`
	Status    bool     `json:"status"`
	Devices   []Device `json:"devices" gorm:"many2many:user_devices"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
