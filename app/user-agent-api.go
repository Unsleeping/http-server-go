package main

import (
	"fmt"
	"net"
	"strings"
)

func UserAgentHandler(conn net.Conn, headers []string)(string, map[string]string) {
	userAgent := ""
	userAgentPrefix := "User-Agent:"
	for _, header := range headers {
		if strings.HasPrefix(header, userAgentPrefix) {
			userAgent = strings.TrimSpace(strings.TrimPrefix(header, userAgentPrefix))
			break
		}
	}

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(userAgent)),
	}


	return userAgent, responseHeaders
}