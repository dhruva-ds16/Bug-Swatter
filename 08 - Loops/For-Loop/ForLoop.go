package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//For loop with break
	j := 0
	for {
		fmt.Println("Hello World!")
		if j == 10 {
			break
		}
		j++
	}

	//Ranges
	strings := []string{"Hello", "World", "!"}

	//Get both index and value while looping through
	for k, val := range strings {
		fmt.Printf("%d: %s\n", k, val)
	}

	//Only get the index
	for k := range strings {
		fmt.Println(k)
	}

	//Only get the value
	for _, val := range strings {
		fmt.Println(val)
	}
}
