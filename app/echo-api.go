package main

import (
	"fmt"
	"strings"
)

func EchoHandler(path string, headers []string) (string, map[string]string) {
	content := strings.TrimPrefix(path, "/echo/")

	acceptEncoding := GetEncodingString(headers)


	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}

	if acceptEncoding != "" {
		responseHeaders["Content-Encoding"] = acceptEncoding
	}

	return content, responseHeaders
}