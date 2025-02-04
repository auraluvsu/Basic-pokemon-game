package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HairColour string `json:"hair_colour"`
	HasDog     bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_colour": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_colour": "brown",
		"has_dog": false
	}
]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println(err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Barry"
	m1.LastName = "Allen"
	m1.HairColour = "Brown"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Wally"
	m2.LastName = "West"
	m2.HairColour = "Red"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling:", err)
	}

	fmt.Println(string(newJson))
}
