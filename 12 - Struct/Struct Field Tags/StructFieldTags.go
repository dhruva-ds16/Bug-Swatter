package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name   string `json:"name"`
	Weight string `json:"weight"`
}

func main() {
	json_string := `
	{
		"name": "Rocky",
		"weight": "45"
	}`

	rocky := new(Dog)
	json.Unmarshal([]byte(json_string), rocky) //Unmarshalling is decoding a JSON string to a Go Data structure. takes a []byte slice as input which contains JSON data and a pointer to an interface{} as arguments
	fmt.Println(rocky)

	spot := new(Dog)
	spot.Name = "Spot"
	spot.Weight = "20"
	jsonStr, _ := json.Marshal(spot) //Marshalling in Go, encodes a data structure into a JSON string. This takes an interface type as an argument and returns a []byte slice and an error
	fmt.Printf("%s\n", jsonStr)
}
