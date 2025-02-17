package main

import (
	"fmt"
	"net"
)

func UserAgentHandler(conn net.Conn, headers []string)(string, map[string]string) {
	userAgent := GetHeaderValue(headers, "User-Agent:")

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(userAgent)),
	}


	return userAgent, responseHeaders
}