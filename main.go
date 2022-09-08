package main

import "log"

func main() {

	animals := []string{"cat", "dog", "fish", "horse", "buffalo"}

	// if we don't care about the index i then we can use blank identifier in place of that
	for _, animal := range animals {
		log.Println(animal)
	}
}
