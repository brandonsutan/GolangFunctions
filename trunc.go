package main

import (
	"fmt"
)

func main() {
	var floatNumber float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&floatNumber)
	truncatedNumber := int(floatNumber)
	fmt.Println("Truncated number is", truncatedNumber)
}
