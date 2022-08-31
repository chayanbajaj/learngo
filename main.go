package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Chayan",
		LastName:  "Bajaj",
	}

	myMap["myData"] = me

	log.Println(myMap["myData"])
	log.Println(myMap["myData"].FirstName)
	log.Println(myMap["myData"].LastName)
}
