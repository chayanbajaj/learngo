package main

import "log"

func main() {

	firstLine := "You are just too good to be true"

	for i, l := range firstLine {
		//this won't print the letter it will print the byte
		log.Println(i, ":", l)
	}
}
