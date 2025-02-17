package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"strings"
)

func EchoHandler(path string, headers []string) ([]byte, map[string]string) {
	content := strings.TrimPrefix(path, "/echo/")

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}

	acceptEncoding := GetEncodingString(headers)

	if acceptEncoding != "" {
		responseHeaders["Content-Encoding"] = acceptEncoding
	}

	var responseBody []byte
	if acceptEncoding == "gzip" {
		var buf bytes.Buffer
		gzipWriter := gzip.NewWriter(&buf)

		_, err := gzipWriter.Write([]byte(content))
		if err != nil {
			fmt.Println("Error gzipping content:", err)
			return []byte(content), responseHeaders
		}
		err = gzipWriter.Close()
		if err != nil {
			fmt.Println("Error closing gzip writer:", err)
			return []byte(content), responseHeaders
		}
		responseBody = buf.Bytes()
		responseHeaders["Content-Length"] = fmt.Sprintf("%d", len(responseBody))
	} else {
		responseBody = []byte(content)
	}

	return responseBody, responseHeaders
}