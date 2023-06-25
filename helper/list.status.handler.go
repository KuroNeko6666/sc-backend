package helper

import "strings"

func IsStatusOrder(status string) bool {
	status = strings.ToLower(status)
	list := []string{"accept", "failure"}
	for _, v := range list {
		if status == v {
			return true
		}
	}
	return false
}
