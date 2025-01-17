package main

import (
	"fmt" 
	"github.com/AlexandreOkita/go-expert-packages/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", soma)
}