package main

type MyNumber int

type Number interface {
	~int | float32
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func Imprimir[T any](a T) {
	println(a)
}

func main() {
	m := map[string]int{"okita": 1000, "jao": 2000}
	m2 := map[string]float32{"okita": 1000.20, "jao": 2000.10}
	m3 := map[string]MyNumber{"okita": 1000, "jao": 2000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
	Imprimir(1)
}
