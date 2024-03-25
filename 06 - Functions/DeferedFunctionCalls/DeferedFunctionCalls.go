//Defer is a special statement that schedules functions to be executed after the function completes

package main

import "fmt"

func first() {
	fmt.Println("First Function")
}

func second() {
	fmt.Println("Second Function")
}

func main() {
	defer second()
	first()
}

//Here, the second function executes after the main function is executed
