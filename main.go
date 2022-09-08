package main

import (
	"log"

	"github.com/youruser/yourrepo/helpers"
)

const numPool = 10

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func main() {
	intChan := make(chan int)
	defer close(intChan) //defer is executed after the completion of current function

	//running a function with go routine
	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
