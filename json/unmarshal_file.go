package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// START OMIT
	type Person struct {
		Name      string `json:"name"`
		Followers int    `json:"followers"`
		Active    bool   `json:"active"`
	}
	var person Person
	data, err := ioutil.ReadFile("json/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", person.Name)
	fmt.Printf("followers: %d\n", person.Followers)
	fmt.Printf("active: %t\n", person.Active)
	// END OMIT
}
