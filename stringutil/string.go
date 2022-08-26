package string

import "strings"

func IsHttps(url string) bool {
	return strings.Contains(url, "https://")
}
