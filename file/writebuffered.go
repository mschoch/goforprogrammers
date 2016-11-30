package main

import (
	"bufio"
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
	bw := bufio.NewWriter(f) // HL
	_, err = bw.Write([]byte("Goodbye"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bw.Write([]byte(", Dear Brother"))
	if err != nil {
		log.Fatal(err)
	}
	bw.Flush() // HL
	// END OMIT
}
