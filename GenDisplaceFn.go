package main

import (
	"fmt"
)

// GenDisplaceFn generates a displacement function
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return so + vo*t + 0.5*a*t*t
	}
}

func main() {
	var a, vo, so, t float64

	// Read acceleration, initial velocity, and initial displacement from the user
	fmt.Println("Enter acceleration (a):")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity (vo):")
	fmt.Scan(&vo)
	fmt.Println("Enter initial displacement (so):")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	// Read time from the user and calculate displacement
	fmt.Println("Enter time (t):")
	fmt.Scan(&t)
	fmt.Println("The displacement is:", fn(t))

	// You can run more tests by calling fn with different time values
}
