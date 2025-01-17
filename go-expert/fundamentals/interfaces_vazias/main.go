package main

import "fmt"

type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World!"

	showType(x)
	showType(y)

	var minhaVar interface{} = "Alexandre Okita"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O resultado é %v e o resultado de ok é %v", res, ok)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}