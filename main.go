package main

import "fmt"

func main() {
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

//it is possible to return more than one thing from a function
// func function_name() (return_type_1, return_type_2, ...)

func saySomething() (string, string) {
	return "something", "else"
}
