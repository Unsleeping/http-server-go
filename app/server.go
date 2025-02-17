package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

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
		go HandleConnection(conn, *directory)
	} 


}



