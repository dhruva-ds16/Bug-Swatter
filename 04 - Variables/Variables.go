package main

import "fmt"

func main() {
	//Declaring variables
	var i int
	var f32 float32

	//Intializing Variables
	i = 69
	f32 = 3.1415936535

	fmt.Println(i, f32)

	//Declare and Initialize in one line
	var str string = "One Liner"
	fmt.Println(str)

	//Short Variable Declaration
	j := 96
	fmt.Println(j)

	//Declaring Multiple Variables
	firstName, lastName := "Tom", "Jenkins"
	fmt.Println(firstName, lastName)

	//Variable Declaration Block
	var (
		name = "Dhruva"
		age  = 24
	)

	fmt.Println(name)
	fmt.Println(age)

}
