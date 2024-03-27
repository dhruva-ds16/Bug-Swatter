//Go does not have a specific keyword for while but while loop functionality can be performed by the for keyword

package main

import "fmt"

func main() {
	//Basic while
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//Do-while loop
	num := 0
	for {
		fmt.Println(num)
		if num == 10 {
			break
		}
		num++
	}
}
