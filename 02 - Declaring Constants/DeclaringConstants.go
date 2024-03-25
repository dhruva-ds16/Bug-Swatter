package main

import "fmt"

const RandVar1 float64 = 1.342352451643 //These are single constant declaration
const RandVar2 = 9999

//This is multiple constant declaration
const (
	MultiVar1 = "Multiple constant declaration 1"
	MultiVar2 = 345678
)

func main() {
	fmt.Println(RandVar1)
	fmt.Println(RandVar2)
	fmt.Println(MultiVar1)
	fmt.Println(MultiVar2)
}
