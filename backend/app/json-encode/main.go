package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

func main() {
	names := []string{"brian", "habib", "malik"}
	ages := []int{25, 25, 24}
	hobbies := []string{"hiking", "touring", "traveling"}

	people := []Person{}

	// generate json
	for i := range names {
		person := Person{
			Name:  names[i],
			Age:   ages[i],
			Hobby: hobbies[i],
		}
		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		fmt.Println("Error marshal:", err)
		return
	}

	fmt.Println(string(jsonData))
}
