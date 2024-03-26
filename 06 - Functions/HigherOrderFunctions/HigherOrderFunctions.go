//This is where a function returns another function

package main

import "fmt"

func multi(x, y int) int {
	return x * y
}

func partialMulti(x int) func(int) int {
	return func(y int) int {
		return multi(x, y)
	}
}

func main() {
	multiple := partialMulti(10)
	fmt.Println(multiple(20))
}
