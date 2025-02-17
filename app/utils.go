package main

import (
	"fmt"
	"strings"
)

func CreateResponse(status int, headers map[string]string, body string) string {
	var statusText string
	switch status {
	case 201:
		statusText = "Created"
	case 200:
		statusText = "OK"
	case 400:
		statusText = "Bad Request"
	case 404:
		statusText = "Not Found"
	default:
		statusText = "Unknown"
	}

	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", status, statusText)

	for key, value := range headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	response += "\r\n"

	if body != "" {
		response += body
	}

	return response
}


func GetHeaderValue(headers []string, key string) string {
	value := ""
	for _, header := range headers {
		if strings.HasPrefix(header, key) {
			value = strings.TrimSpace(strings.TrimPrefix(header, key))
			break
		}
	}

	return value
}
