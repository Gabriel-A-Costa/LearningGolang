package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)
	fmt.Printf("Before append: Slice: %v, Len %d, Cap %d\n", slice, len(slice), cap(slice))

	for i := 1; i <= 8; i++ {
		slice = append(slice, i)
		fmt.Printf("Apos append: %d, Slice: %v, Len %d, Cap %d\n", i, slice, len(slice), cap(slice))
	}

	slice = append(slice, 6)
	fmt.Printf("Apos for: Slice: %v, Len %d, Cap %d\n", slice, len(slice), cap(slice))

}
