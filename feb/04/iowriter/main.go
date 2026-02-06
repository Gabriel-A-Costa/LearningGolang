package main

import (
	"fmt"
	"io"
	"os"
)

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

func EscrevaHelloWorld(writer io.Writer) (int, error) {
	data := []byte("Hello World!")

	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return 0, err
	}

	return n, nil
}

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	n, err := EscrevaHelloWorld(file)
	if err != nil {
		return
	}
	fmt.Printf("Gravado %d bytes no arquivo.\n", n)

}
