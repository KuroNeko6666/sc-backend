package response

import (
	"time"

	"github.com/KuroNeko6666/sc-backend/interface/model"
)

type Cart struct {
	ID        string         `json:"id"`
	Items     []model.Device `json:"items"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
