package main

import (
	"log"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Open("file/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // HL
	buf := make([]byte, 20)
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
	// END OMIT
}
