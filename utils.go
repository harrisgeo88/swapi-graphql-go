package swapiGraphQLGo

import (
	"fmt"
	"strings"
)

// SplitLinkID doc
func SplitLinkID(url string, separator string) string {
	id := strings.TrimRight(strings.Split(url, separator)[1], "/")

	return id
}

// ConstructURL docs
func ConstructURL(path string, id string) string {
	baseURL := "https://swapi.co/api"
	format := "format=json"

	result := fmt.Sprintf("%s/%s/%s?%s", baseURL, path, id, format)
	return result
}
