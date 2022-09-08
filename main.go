package main

import "log"

func main() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "cow":
		log.Println("cat is set to cow")

	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is set to something else")
	}
}
