// Go provides a feature known as anonymous functions where a function does not require a name. This can be useful when a function is declared inline

//Anonymous function is also known as a function literal

package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	area := area(10, 10)
	fmt.Println(area)
}
