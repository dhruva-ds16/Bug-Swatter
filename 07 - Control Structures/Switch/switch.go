package main

import "fmt"

func main() {
	option := 1
	switch option {
	case 1:
		fmt.Println("Option1")
	case 2:
		fmt.Println("Option2")
	default:
		fmt.Println("Other")
	}

	//Switch with multiple cases
	num := 7
	switch num {
	case 1, 2, 3, 4, 5:
		fmt.Println("Early")
	case 6, 7, 8, 9:
		fmt.Println("On-time")
	default:
		fmt.Println("Late")
	}

	//Switch statement with fallthrough (forces execution of the next statement)

	day := 89
	switch day {
	case 1:
		fmt.Println("Do 1")
		fallthrough
	case 2:
		fmt.Println("Do 2")
		fallthrough
	case 3:
		fmt.Println("Do 3")
		fallthrough
	case 4:
		fmt.Println("Do 4")
		fallthrough
	case 5:
		fmt.Println("Do 5")
		fallthrough
	case 6:
		fmt.Println("Do 6")
	default:
		fmt.Println("Do Default")
	}

	//Switch statements with conditionals
	switch {
	case day < 3:
		fmt.Println("First Half")
	case day == 3:
		fmt.Println("Middle")
	case day > 3 && day < 7:
		fmt.Println("Second half")
	default:
		fmt.Println("Not in week")
	}
}
