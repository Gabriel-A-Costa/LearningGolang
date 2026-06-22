package main

import (
	"fmt"
	"time"
)

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Goroutine:", i)
// 	}
// }

// func main() {
// 	go printNumbers() // Inicia la gorutina para imprimir números

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Main:", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

func printNumber(c chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}

	close(c)
}

func main() {
	c := make(chan int)

	go printNumber(c)

	for number := range c {
		fmt.Println(number)
	}
}
