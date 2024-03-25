package main

import (
	"fmt" //This package implements real-time reflection (This is the ability of a program to inpect its values and variables at runtime and find their type). This allows the program to manipulate objects with arbitrary types
	"reflect"
	"strconv" //Implements conversions to and from string representation of basic datatypes
)

func main() {
	//Converts a string s to a float value
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Println(f)

	//Converts a string str to boolean
	str := "True"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)

	//Converts a float flo to a string value
	var flo float64 = 3.1415926535
	var strflo string = strconv.FormatFloat(flo, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(strflo))

	//Converts a float value f32 to integer
	var f32 float32 = 3.1415926535
	fmt.Println(reflect.TypeOf(f32))
	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
}
