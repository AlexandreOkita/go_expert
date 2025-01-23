package main

import "fmt"

type ID int

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 100)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
