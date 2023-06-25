package response

type Chart struct {
	Labels []string `json:"labels"`
	Data   []int64  `json:"data"`
}
