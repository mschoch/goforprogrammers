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
	defer func() { // HL
		err = f.Close() // HL
		if err != nil { // HL
			log.Fatal(err) // HL
		} // HL
	}() // HL
	buf := make([]byte, 20)
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
	// END OMIT
}
