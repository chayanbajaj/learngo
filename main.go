package main

import "log"

func main() {

	animals := []string{"cat", "dog", "fish", "horse", "buffalo"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}
