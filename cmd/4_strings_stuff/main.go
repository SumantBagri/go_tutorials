package main

import (
	"fmt"
	"strings"
)

func main() {
	var nonAsciiString = "résumé"
	fmt.Println(nonAsciiString)

	// strings with non-ascii characters
	fmt.Println("\nSTRING ITERATION")
	for i, v := range nonAsciiString {
		fmt.Printf("%v, %v\n", i, v)
	}

	// strings casted as runes
	fmt.Println("\nRUNE-CAST ITERATION")
	var myRuneCastedFromString = []rune(nonAsciiString)
	for i, v := range myRuneCastedFromString {
		fmt.Printf("%v, %v\n", i, v)
	}

	// string concatenation
	var strSlice = []string{"h", "e", "l", "l", "o"}

	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] // inefficient way - string recreation every call
	}
	fmt.Printf("\n%v", catStr)

	var stringBuilder strings.Builder
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}
	var catStrEfficient = stringBuilder.String()
	fmt.Printf("\n%v", catStrEfficient)
}
