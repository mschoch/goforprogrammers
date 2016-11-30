package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// START OMIT
	err := ioutil.WriteFile("file/out.txt", []byte("Goodbye"), 0777) // HL
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
