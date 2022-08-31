package main

import "fmt"

func main() {
	//var_name := value to declare and initialise at the same time
	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
}

//func function_name() return_type

func saySomething() string {
	return "something"
}
