package main

import (
	"log"
)

func main() {
	//We can declare and initialise in a slice as
	// varName := []dataType{values}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	log.Println(numbers)

	//We can also print a range of slice

	log.Println(numbers[0:2])
	log.Println(numbers[5:8])
}
