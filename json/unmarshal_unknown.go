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

	var parsedJSON map[string]interface{}
	err := json.Unmarshal(jsonBytes, &parsedJSON)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range parsedJSON {
		fmt.Printf("key: %s val: %v\n", k, v)
	}
	// END OMIT
}
