package main

import (
	"encoding/json"
	"fmt"
)


type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{
		Name:   "Alice",
		Age:    30,
		IsAdult: true,
	}

		fmt.Println("Person Data is :", person)

	// Convert struct to JSON using json.Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// we convert it to string for better readability 
	fmt.Println("Person Data is in JSON :" , string(jsonData))


	// Convert JSON back to struct using json.Unmarshal
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return 
	}
	fmt.Println("Person Data Unmarshaling is :", personData)


	// Accessing individual fields
	data := []byte(`{"name":"Amit","age":25,"is_adult":true}`)

	var user Person
	json.Unmarshal(data, &user)

	fmt.Println(user.Name)
	fmt.Println(user.Age)

}
