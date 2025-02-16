package main

import (
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")


	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	n, err := conn.Read([]byte(make([]byte, 1024)))
	if err != nil {
    fmt.Println("Error reading from connection: ", err.Error())
    os.Exit(1)
  }

	fmt.Printf("Received %d\n", n)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
