package main

import "log"

func main() {
	//In Go we don't use arrays, we use Slice
	// var sliceName []dataType

	var mySlice []int

	//To add an element in slice we use append method
	// sliceName = append(sliceName, value)

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)
}
