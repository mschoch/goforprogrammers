package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// START 1OMIT
	f, err := os.Open("proverbs.txt") // HL
	if err != nil {
		log.Fatal(err)
	}

	proverbs, err := loadProverbs(f) // HL
	if err != nil {
		log.Fatal(err)
	}
	// END 1OMIT

	proverbsHandler := NewProverbHandler(proverbs)
	http.Handle("/proverb", proverbsHandler) // HL
	log.Fatal(http.ListenAndServe(":8888", nil))
	// END 2OMIT
}
