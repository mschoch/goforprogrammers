package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Open("file/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 20)
	br := bufio.NewReader(f) // HL
	_, err = br.Read(buf)    // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
	// END OMIT
}
