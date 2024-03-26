package main

import "fmt"

func main() {
	x := 69
	y := 420

	//If-else
	if x > 70 {
		fmt.Println("More than NICE")
	} else {
		fmt.Println("Maybe NICE")
	}

	//If-else-If
	if y < 420 {
		fmt.Println("NOT NICE")
	} else if y > 420 {
		fmt.Println("NOT NICE")
	} else {
		fmt.Println("NICE")
	}
}
