package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	type Person struct {
		Name      string `json:"name"`
		Followers int    `json:"followers"`
		Active    bool   `json:"active"`
	}
	// START OMIT
	person := Person{
		Name:      "Srdjan",
		Followers: 139,
		Active:    true,
	}
	f, err := os.Create("json/output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	encoder := json.NewEncoder(f)
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
