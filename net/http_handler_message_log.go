package main

import (
	"log"
	"net/http"
)

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Before")
		h.ServeHTTP(w, r)
		log.Printf("After")
	})
}

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
	messageHandler := &MessageHandler{Message: "Hello World"}
	http.Handle("/message", logHandler(messageHandler))
	log.Fatal(http.ListenAndServe(":8083", nil))
}

// END OMIT
