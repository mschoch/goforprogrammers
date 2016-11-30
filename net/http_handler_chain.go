package main

import (
	"log"
	"net/http"
)

// START OMIT
func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Before")
		h.ServeHTTP(w, r)
		log.Printf("After")
	})
}

// END OMIT
