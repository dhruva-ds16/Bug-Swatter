package main

import "fmt"

func main() {
	hello("Dhruva")
	add(48, 21)
}

func hello(x string) {
	fmt.Printf("Hello %s\n", x)
}

func add(a int, b int) {
	e := a + b
	fmt.Println(e)
}
