package response

import "github.com/KuroNeko6666/sc-backend/interface/model"

type Login struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
