package main

import (
	"errors"
	"fmt"
)

func main() {
	var myVar string = "Hello World"
	printMe(myVar)

	numerator := 5
	denominator := 2
	result, remainder, err := intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Print(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // default value is nil

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
