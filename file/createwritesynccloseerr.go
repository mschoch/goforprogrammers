package main

import (
	"log"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Create("file/out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	_, err = f.Write([]byte("Goodbye"))
	if err != nil {
		log.Fatal(err)
	}
	err = f.Sync() // HL
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(", Dear Brother"))
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
