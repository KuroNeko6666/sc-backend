package form

type Device struct {
	ID        string `json:"id" form:"id"`
	Model     string `json:"model" form:"model"`
	Address   string `json:"address" form:"address"`
	City      string `json:"city" form:"city"`
	Province  string `json:"province" form:"province"`
	Longitude string `json:"longitude" form:"longitude"`
	Latitude  string `json:"latitude" form:"latitude"`
}

type UpdateDevice struct {
	Model     string `json:"model" form:"model"`
	Address   string `json:"address" form:"address"`
	City      string `json:"city" form:"city"`
	Province  string `json:"province" form:"province"`
	Longitude string `json:"longitude" form:"longitude"`
	Latitude  string `json:"latitude" form:"latitude"`
}

type UserDevice struct {
	DeviceID string `json:"device_id" form:"device_id"`
	UserID   string `json:"user_id" form:"user_id"`
}
