package main

import "fmt"

func main() {
	var i int = 80
	var j int = 60
	var k int = 40

	//Arithmetic
	fmt.Printf("i + j = %d\n", i+j) //Println does not have the %d directive
	fmt.Printf("i - j = %d\n", i-j)
	fmt.Printf("i * j = %d\n", i*j)
	fmt.Printf("i / j = %d\n", i/j)
	fmt.Printf("i mod j = %d\n", i%j)

	//Comparision
	fmt.Println(i == j)
	fmt.Println(i != j)
	fmt.Println(i > j)
	fmt.Println(i >= j)
	fmt.Println(i < k)
	fmt.Println(i <= k)

	//Logical Operators
	fmt.Println(i > j && k > j)
	fmt.Println(i < j || k < i)
	fmt.Println(!(i == k && i > k))

	//Assignment Operators
	var a, b = 20, 13

	a += b
	fmt.Printf("+= %d\n", a)

	a = 20
	a -= b
	fmt.Printf("-= %d\n", a)

	a = 20
	a *= b
	fmt.Printf("*= %d\n", a)

	a = 20
	a /= b
	fmt.Printf("/= %d\n", a)

	a = 20
	a %= b
	fmt.Printf("(mod)= %d\n", a)

}
