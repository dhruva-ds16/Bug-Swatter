//Variadic functions are functions that take a variable number of arguments

package main

import (
	"fmt"
	"reflect"
)

func main() {
	printMultiStrings("Hello", "World", "!")
	printMultiVars(1, "Green", false, 1.2123, []string{"foo", "bar", "baz"})
}

func printMultiStrings(s ...string) { //Input parameters get added as elements in an array
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func printMultiVars(i ...interface{}) { //interface{} is used to store values of unknown type
	for _, v := range i { //_ is an anonymous placeholder where the program safely discards the return value. this is a way in Go where you dont use the variable that is declared
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
