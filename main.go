package main

import "log"

func main() {
	// Common Syntax to make map
	// varName := make(map[keyDataType]dataType)

	myMap := make(map[string]string)
	myMap["dog"] = "MoolchandJi"
	myMap["cat"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	//We can also override using key

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])
}
