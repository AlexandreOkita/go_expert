package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 10 {
		println(10)
		i++
	}

	for {
		println("Hello, World!")
	}
}
