// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	type ProverbRequest struct {
		Offset int `json:"offset"`
	}

	type ProverbResponse struct {
		Proverb string `json:"proverb"`
	}

	// START OMIT
	preq := ProverbRequest{
		Offset: 0,
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(&preq)

	resp, err := http.Post("http://localhost:8888/proverb", "application/json", &buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var pres ProverbResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&pres)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", pres.Proverb)
	// END OMIT
}
