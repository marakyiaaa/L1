package main

import "fmt"

func main() {
	a, b := 10, 52
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
