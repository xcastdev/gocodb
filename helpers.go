package gocodb

import (
	"fmt"
)

func formatQueryParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	query := "?"
	for key, value := range params {
		query += fmt.Sprintf("%s=%s&", key, value)
	}

	// Remove the trailing '&'
	return query[:len(query)-1]
}
