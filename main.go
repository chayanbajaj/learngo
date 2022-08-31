package main

import "log"

func main() {
	// Common Syntax to make map
	// varName := make(map[keyDataType]dataType)

	myMap := make(map[string]int)
	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"])
	log.Println(myMap["Second"])
}
