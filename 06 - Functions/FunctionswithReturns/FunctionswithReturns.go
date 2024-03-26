package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func area(l int, b int) (a int) {
	a = l * b
	return
}

func rectangle(l int, b int) (a int, p int) { //Here, the function returns in the same order as mentioned after input parameters
	p = 2 * (l + b)
	a = l * b
	return
}

func addValue(x *int, y *string) {
	*x = *x + 5
	*y = *y + " World!"
	return
}

func main() {
	//Add
	sum := add(20, 30)
	fmt.Println("Sum: ", sum)

	//Area
	ar := area(5, 2)
	fmt.Println("Area: ", ar)

	//Rectangle
	rectArea, rectPeri := rectangle(10, 15)
	fmt.Println("Rectangle Area: ", rectArea)
	fmt.Println("Rectangle Perimeter: ", rectPeri)

	//AddValue
	var num = 20
	var text = "Hello"
	fmt.Println("Before: ", text, num)

	addValue(&num, &text) //If passing addresses to a function, be careful of the & and * placement
	fmt.Println("After: ", text, num)
	//Call by Reference and not Call by Value
}
