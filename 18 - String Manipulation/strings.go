package main

import (
	"fmt"
	"strings"
)

func main() {
	//Creating a string
	str := "Some String"
	sentence := " I am a long sentence made up of many different words"
	numbers := []string{"one", "two", "three", "four", "five"}
	path := "/home/dhruv/project"

	//Convert string to lowercase
	lower := strings.ToLower(str)
	fmt.Println(lower)

	//Convert string to uppercase
	upper := strings.ToUpper(str)
	fmt.Println(upper)

	//Check if string contains another string
	if strings.Contains(str, "some") {
		fmt.Println("Yes, 'some' exists in the string")
	}

	//Get/print character range from string
	fmt.Println("Cars 3 -10: " + str[3:10])
	fmt.Println("First 5: " + str[:5])

	//Split a string
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	//Split a string on whitespaces using fields
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

	//Join an array of string
	joinStr := strings.Join(numbers, ",")
	fmt.Println(joinStr)

	//Replace (Takes a number of how many replacements it should do -1 = all of them)
	newStr := strings.Replace(str, "Some", "The", -1)
	fmt.Println(newStr)

	//Has Prefix
	isAbsolute := strings.HasPrefix(path, "/")
	fmt.Println(isAbsolute)

	//Has Suffix
	hasTrailingSlash := strings.HasSuffix(path, "/")
	fmt.Println(hasTrailingSlash)

	//Count characters in string
	count := strings.Count(str, "s")
	fmt.Println(count)

	//Determine string length
	length := len(str)
	fmt.Println(length)
}
