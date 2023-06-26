package response

type Login struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
