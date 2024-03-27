package main

import (
	"errors"
	"fmt"
	"math"
)

type error interface {
	Error() string
}

// Return error using multiple return statements
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: Not possible to calculate the square root of negative numbers")
	} else {
		z := 1.0
		min_delta := 0.00000000001
		delta := 1.0
		i := 0
		for ; math.Abs(delta) > min_delta; i++ {
			delta = (z*z - x) / (2 * z)
			z = z - delta
		}

		fmt.Println("Iterations: ", i)
		return z, nil
	}
}

func main() {
	//creating an error
	err := errors.New("First Error")
	fmt.Println("Error: ", err)

	//Get error as a return value
	f, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
