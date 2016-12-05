package main

import (
	"log"
	"net"
)

// START OMIT
func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // HL
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
