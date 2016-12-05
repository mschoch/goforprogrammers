package main

import (
	"encoding/json"
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
	person := Person{
		Name:      "Srdjan",
		Followers: 139,
		Active:    true,
	}
	jsonBytes, err := json.Marshal(person) // HL
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("json/output.json", jsonBytes, 0777) // HL
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
