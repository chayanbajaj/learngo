package main

import "log"

func main() {

	cat := "cat"

	if cat == "cat" {
		log.Println("cat is", cat)
	} else {
		log.Println("cat is not a", cat)
	}
}
