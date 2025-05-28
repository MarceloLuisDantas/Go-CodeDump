package main

import (
	"arquivos/files"
	"fmt"
)

func main() {
	file := files.Load("./arquivos/leitura.txt")
	for i, v := range file.Linhas {
		fmt.Printf("Linha: %d, %s\n", i, v)
	}
}
