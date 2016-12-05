package main

import (
	"fmt"
	"log"
	"net"
)

// START OMIT
func handleConnection(conn net.Conn) {
	fmt.Printf("got connection\n")
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080") // HL
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept() // HL
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

// END OMIT
