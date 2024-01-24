package main

import "fmt"

func main() {
	// arrays
	intArr := [...]int32{1, 2, 3, 4} // infer length with elipsis
	fmt.Println(intArr)

	// slices -- extendable arrays
	var intSlice = []int32{5, 6, 7}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 8)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// append slice to slice
	intSlice2 := []int32{10, 11, 12}
	addSlice := []int32{14, 15, 16}
	intSlice2 = append(intSlice2, addSlice...)
	fmt.Println(intSlice2)

	// make slice with capacity (similar to new in C++)
	intSlice3 := make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Printf("Slice with length %v and capacity %v\n", len(intSlice3), cap(intSlice3))

	// maps
	var myMap map[string]uint8 = make(map[string]uint8) // empty map
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Alice": 27, "Bob": 23} // initialize map
	fmt.Println(myMap2)

	myMap2["Donna"] = 32 // add to map
	fmt.Println(myMap2)

	delete(myMap2, "Bob") // delete from map by reference
	fmt.Println(myMap2)

	var age, ok = myMap2["Cyclops"] // handle missing key from map
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// loop through maps and slices
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intSlice2 {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}
