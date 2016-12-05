package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// START OMIT
	jsonBytes := []byte(`
    {
      "name": "marty",
      "followers": 387,
      "active": true
    }
  `)
	type Person struct {
		Name      string `json:"name"`
		Followers int    `json:"followers"`
		Active    bool   `json:"active"`
	}
	var person Person                         // HL
	err := json.Unmarshal(jsonBytes, &person) // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", person.Name)
	fmt.Printf("followers: %d\n", person.Followers)
	fmt.Printf("active: %t\n", person.Active)
	// END OMIT
}
