package main

import (
	"fmt"
	"strings"
)

func EchoHandler(path string, headers []string) (string, map[string]string) {
	content := strings.TrimPrefix(path, "/echo/")

	acceptEncoding := GetHeaderValue(headers, "Accept-Encoding:")

	fmt.Println("acceptEncoding: ", acceptEncoding)

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}

	if strings.Contains(acceptEncoding, "gzip") {
		responseHeaders["Content-Encoding"] = acceptEncoding
	}

	return content, responseHeaders
}