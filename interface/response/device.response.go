package response

import (
	"time"

	"github.com/KuroNeko6666/sc-backend/interface/model"
)

type Devices struct {
	ID         string              `json:"id"`
	Model      string              `json:"model"`
	Address    model.DeviceAddress `json:"address"`
	Subcribers int64               `json:"subcribers"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

type DevicesMarket struct {
	ID         string              `json:"id"`
	Model      string              `json:"model"`
	Address    model.DeviceAddress `json:"address"`
	Subcribers int64               `json:"subcribers"`
	Subcribe   bool                `json:"subcribe"`
	Cart       bool                `json:"cart"`
	Order      bool                `json:"order"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}
