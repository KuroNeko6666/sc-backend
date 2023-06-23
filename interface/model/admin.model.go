package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
