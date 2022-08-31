package main

import "log"

func main() {
	//In Go we don't use arrays, we use Slice
	// var sliceName []dataType

	var mySlice []string

	//To add an element in slice we use append method
	// sliceName = append(sliceName, value)

	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)
}
