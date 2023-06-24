package response

import (
	"time"

	"github.com/KuroNeko6666/sc-backend/interface/model"
)

type Order struct {
	ID        string         `json:"id"`
	Items     []model.Device `json:"items"`
	Status    string         `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
