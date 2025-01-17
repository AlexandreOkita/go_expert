package main

import "fmt"

type ID int

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Printf("O tipo de E é %T", f)

	for i, v := range meuArray {
		println(i, v)
	}

}
