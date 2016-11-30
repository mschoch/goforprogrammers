package main

import (
	"log"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Create("file/out.txt") // HL
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	_, err = f.Write([]byte("Goodbye")) // HL
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
