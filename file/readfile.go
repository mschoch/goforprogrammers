package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// START OMIT
	data, err := ioutil.ReadFile("file/file.txt") // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
	// END OMIT
}
