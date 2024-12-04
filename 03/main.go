package main

import (
	"fmt"
	"sync"
)

func SumNum(arr []int) {
	ch1 := make(chan int)
	sum := 0
	var wg1 sync.WaitGroup

	for _, val := range arr {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			ch1 <- i * i
		}(val)
	}

	go func() {
		wg1.Wait()
		close(ch1)
	}()

	for val := range ch1 {
		sum += val
	}
	fmt.Println(sum)
}

func main() {
	SumNum([]int{2, 4, 6, 8, 10})
}
