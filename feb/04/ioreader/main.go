package main

import (
	"fmt"
	"io"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func LerReader(reader io.Reader) (int, []byte, error) {
	buffer := make([]byte, 1024)

	n, err := reader.Read(buffer)
	if err != nil {
		return 0, nil, err
	}

	return n, buffer, nil
}

func LerTxt(txt string) {
	file, err := os.Open(txt)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:")
		return
	}
	defer file.Close()

	n, buffer, err := LerReader(file)
	if err != nil {
		fmt.Println("Erro ao Ler do arquivo:", err)
	}

	fmt.Printf("Lidos %d bytes: %s\n", n, buffer[:n])
}

func main() {
	LerTxt("example.txt")
}
