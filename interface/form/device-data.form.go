package form

type DeviceData struct {
	DeviceID string  `json:"device_id" form:"device_id"`
	Speed    float64 `json:"speed" form:"speed"`
	Distance float64 `json:"distance" form:"distance"`
	DateTime string  `json:"date_time" form:"date_time"`
}
