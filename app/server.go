package main

import (
	"flag"
	"fmt"
	"net"
	"os"
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
		conn.Write([]byte(CreateResponse(400, nil, "")))
		return
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) < 2 {
		conn.Write([]byte(CreateResponse(400, nil, "")))
		return
	}

	method, path := parts[0], parts[1]
	headers, body := lines[1 : len(lines) - 1], lines[len(lines) - 1]

	trimmedBody := strings.TrimRight(body, "\x00")


	switch {
		case path == "/":
			conn.Write([]byte(CreateResponse(200, nil, "")))

		case strings.HasPrefix(path, "/files/"):
			status, responseHeaders, fileContent := FilesHandler(path, directory, method, trimmedBody)

			conn.Write([]byte(CreateResponse(status, responseHeaders, string(fileContent))))

		case path == "/user-agent":
			content, responseHeaders := UserAgentHandler(conn, headers)

			conn.Write([]byte(CreateResponse(200, responseHeaders, content)))

		case strings.HasPrefix(path, "/echo/"):
			content, responseHeaders := EchoHandler(path)

			conn.Write([]byte(CreateResponse(200, responseHeaders, content)))
			
		default:
			conn.Write([]byte(CreateResponse(404, nil, "")))
	}
}


