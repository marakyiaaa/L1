package main

import (
	"fmt"
)

func PowNumOne() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for i := range arr {
		go func() {
			ch <- arr[i] * arr[i]
		}()
		fmt.Print(<-ch, " ")
	}
}

func main() {
	PowNumOne()
}
