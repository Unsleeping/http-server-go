package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
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

	defer conn.Close()

	buf := make([]byte, 0, 4096) 
	tmp := make([]byte, 256)     
	for {
			n, err := conn.Read(tmp)
			if err != nil {
					if err != io.EOF {
							fmt.Println("read error:", err)
					}
					break
			}
			buf = append(buf, tmp[:n]...)
    }

	fmt.Println("total size:", len(buf))
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
