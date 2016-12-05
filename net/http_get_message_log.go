package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// START OMIT
	resp, err := http.Get("http://localhost:8083/message") // HL
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", body)
	// END OMIT
}
