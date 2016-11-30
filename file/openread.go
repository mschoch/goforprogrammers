package main

import (
	"log"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Open("file/file.txt") // HL
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 20)
	_, err = f.Read(buf) // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
	// END OMIT
}
