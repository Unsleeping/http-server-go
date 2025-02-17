package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	directory := flag.String("directory", "", "the directory to serve files from")

	flag.Parse()

	if *directory != "" {
		fmt.Printf("Serving files from directory: %s\n", *directory)
	}


	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn, *directory)
	} 


}

func handleConnection(conn net.Conn, directory string) {
	defer conn.Close()



	req := make([]byte, 1024)

	_, err := conn.Read(req)
    if err != nil {
			fmt.Println("Error reading request:", err)
			return
    }

	fmt.Printf("Received request: %s\n", string(req))

	lines := strings.Split(string(req), "\r\n")
	if len(lines) == 0 {
		conn.Write([]byte(createResponse(400, nil, "")))
		return
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) < 2 {
		conn.Write([]byte(createResponse(400, nil, "")))
		return
	}

	path := parts[1]

	fmt.Println("Path: ", path)

	headers := lines[1:]

	fmt.Println("headers: ", headers)

	switch {
		case path == "/":
			conn.Write([]byte(createResponse(200, nil, "")))

		case strings.HasPrefix(path, "/files/"):
			fileName := strings.TrimPrefix(path, "/files/")
			filePath := filepath.Join(directory, fileName)
			fileInfo, err := os.Stat(filePath)

			if err != nil {
				if os.IsNotExist(err) {
					conn.Write([]byte(createResponse(404, nil, "")))
				} else {
					fmt.Println("Error getting file info:", err)
					conn.Write([]byte(createResponse(500, nil, "")))
				}
				return
			}

			fileSize := fileInfo.Size()

			responseHeaders := map[string]string {
				"Content-Type": "application/octet-stream",
				"Content-Length": fmt.Sprintf("%d", fileSize),
			}

			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				conn.Write([]byte(createResponse(500, nil, "")))
				return
			}

			conn.Write([]byte(createResponse(200, responseHeaders, string(fileContent))))

		case path == "/user-agent":
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

			content := userAgent

			conn.Write([]byte(createResponse(200, responseHeaders, content)))



		case strings.HasPrefix(path, "/echo/"):
			content := strings.TrimPrefix(path, "/echo/")

			headers := map[string]string {
				"Content-Type": "text/plain",
				"Content-Length": fmt.Sprintf("%d", len(content)),
			}

			conn.Write([]byte(createResponse(200, headers, content)))
			
		default:
			conn.Write([]byte(createResponse(404, nil, "")))
	}
}

func createResponse(status int, headers map[string]string, body string) string {
	var statusText string
	switch status {
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
