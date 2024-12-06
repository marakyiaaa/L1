package main

import (
	"fmt"
	"sync"
)

func SquaresNumThree(arr []int) {
	ch := make(chan int) //// Небуферизованный канал
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i * i
		}(val)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Print(val, " ")
	}
}

func main() {
	SquaresNumThree([]int{2, 4, 6, 8, 10})
}
