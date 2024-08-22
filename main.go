package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var person Person
	err = json.Unmarshal(file, &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", person.Name, person.Age, person.City)

}
