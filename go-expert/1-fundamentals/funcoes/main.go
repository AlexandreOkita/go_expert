package main

import (
	"errors"
	"fmt"
)

func main() {
	total := func() int {
		return sum(1,2,3,4) * 2
	}

	resultado, err := div(10, 2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resultado da div: %f\n", resultado)
	fmt.Printf("Resultado da soma: %d\n", total())
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("Divisão por 0")
	}
	return a / b, nil
}
