package main

import "fmt"

var m = map[string]string{
	"c": "Cyan",
	"m": "Magenta",
	"y": "Yellow",
	"k": "Black",
}

func main() {
	//Declaring an empty map
	var shoppingList = map[string]int{}
	fmt.Println(shoppingList)

	//Initializing a map
	var people = map[string]int{"Elon": 10, "Jeff": 15}
	fmt.Println(people)

	//Map declaration using the make function
	var peopleList = make(map[string]int)
	peopleList["Elon"] = 10
	peopleList["Jeff"] = 15
	fmt.Println(peopleList)

	//Accessing items
	fmt.Println(m["c"])

	//Adding items
	m["b"] = "black"
	fmt.Println(m)

	//Updating items
	m["y"] = "Lemon Yellow"
	fmt.Println(m)

	//Deleting items
	delete(m, "b")
	fmt.Println(m)

	//Iterating over a map
	for k, v := range m {
		fmt.Printf("Key: %s value %s\n", k, v)
	}

	//Test if an item exists
	c, ok := m["y"]
	fmt.Println("\nc: ", c)
	fmt.Println("ok: ", ok)
}
