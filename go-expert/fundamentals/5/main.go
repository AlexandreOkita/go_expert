package main

import "fmt"

func main() {
	salarios := map[string]int{"Okita": 100, "Jao": 160}
	fmt.Println(salarios["Okita"])
	delete(salarios, "Okita")
	salarios["okita"] = 200
	fmt.Println(salarios["Okita"])

	sal := make(map[string]int)
	sal1 := map[string]int{}

	sal["a"] = 1
	sal1["a"] = 5

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
