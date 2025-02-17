package main

import (
	"fmt"
	"strings"
)

func EchoHandler(path string) (string, map[string]string) {
	content := strings.TrimPrefix(path, "/echo/")

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}

	return content, responseHeaders
}