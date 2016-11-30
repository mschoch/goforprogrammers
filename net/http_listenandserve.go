package main

import (
	"log"
	"net/http"
)

// START OMIT
func main() {
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// END OMIT
