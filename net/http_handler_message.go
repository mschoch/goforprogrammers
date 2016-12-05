package main

import (
	"log"
	"net/http"
)

// START OMIT
type MessageHandler struct {
	Message string
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(m.Message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/message", &MessageHandler{Message: "Hello World"}) // HL
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// END OMIT
