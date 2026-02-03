package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	subslice := slice[1:4]
	fmt.Println("Slice", slice)
	fmt.Println("Sublice", subslice)

	subslice[0] = 99

	fmt.Println("Apos modificar")
	fmt.Println("Slice", slice)
	fmt.Println("Sublice", subslice)
}
