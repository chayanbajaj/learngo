package main

import "log"

// Receiver Syntax
// func (receiver) function_name() return_type {}
func (m *myStruct) printStructFirstName() string {
	return m.FirstName
}

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Daenerys"

	myVar2 := myStruct{
		FirstName: "Khal",
	}

	log.Println("myVar is", myVar.FirstName)
	log.Println("myVar2 is", myVar2.FirstName)

	//We can also use the concept of receivers in functions for struct

	log.Println("myVar is set to", myVar.printStructFirstName())
	log.Println("myVar2 is set to", myVar2.printStructFirstName())
}
