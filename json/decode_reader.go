package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// START OMIT
	type Person struct {
		Name      string `json:"name"`
		Followers int    `json:"followers"`
		Active    bool   `json:"active"`
	}
	var person Person
	f, err := os.Open("json/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f) // HL
	err = decoder.Decode(&person) // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", person.Name)
	fmt.Printf("followers: %d\n", person.Followers)
	fmt.Printf("active: %t\n", person.Active)
	// END OMIT
}
