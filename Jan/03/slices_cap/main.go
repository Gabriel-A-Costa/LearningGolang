package main

import "fmt"

func main() {
	var slice []int

	fmt.Printf("Before append: Slice: %v, Len %d, Cap %d\n", slice, len(slice), cap(slice))
	for i := 0; i <= 9; i++ {
		slice = append(slice, i)
		fmt.Printf("Apos append: %d, Slice: %v, Len %d, Cap %d\n", i, slice, len(slice), cap(slice))
	}
}
