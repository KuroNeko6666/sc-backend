package response

type Base struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Page struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
	Page    int64       `json:"page"`
}
