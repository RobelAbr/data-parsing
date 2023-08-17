package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	// JSON-Daten
	jsonData := `{"name": "David", "age": 19, "country": "Germany"}`

	// Parsen der JSON-Daten in das Go-Struct
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Fehler beim Parsen der JSON-Daten:", err)
		return
	}

	// Ausgabe der geparsten Daten
	fmt.Println("Name:", person.Name)
	fmt.Println("Alter:", person.Age)
	fmt.Println("Land:", person.Country)
}
