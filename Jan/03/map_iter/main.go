package main

import "fmt"

func main() {
	mapa := map[string]int{"Alice": 25, "Bob": 30}
	for key, value := range mapa {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
