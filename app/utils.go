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
	var value string

	for _, header := range headers {
		if strings.HasPrefix(header, key) {
			value = strings.TrimSpace(strings.TrimPrefix(header, key))
			break
		}
	}

	return value
}


func GetEncodingString(headers []string) string {
	dirtyAcceptEncodingString := GetHeaderValue(headers, "Accept-Encoding:")

	var acceptEncoding string

	if strings.Contains(dirtyAcceptEncodingString, ",") {
		parts := strings.Split(dirtyAcceptEncodingString, ",")
		
		var validParts []string

		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "gzip" {
        validParts = append(validParts, part)
      }
		}

		return strings.Join(validParts, ",")
	}

	if strings.Contains(dirtyAcceptEncodingString, "gzip") {
		acceptEncoding = dirtyAcceptEncodingString
	}

	return acceptEncoding
}