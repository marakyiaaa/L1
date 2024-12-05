package main

import (
	"fmt"
	"sync"
)

func PowNumOne() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int) //небуферизированный канал

	// Использование индекса -> приводит к непредсказуемому поведению т.к. значения из внешнего контекста
	for i := range arr {
		go func() {
			ch <- arr[i] * arr[i]
		}()
		fmt.Print(<-ch, " ")
	}
}

func SquaresNumTwo() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr)) // Буферизованный канал => гарантирует, что все горутины смогут отправить свои данные в канал без ожидания

	// Использование значения
	for _, val := range arr {
		go func(i int) {
			ch <- i * i
		}(val)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(<-ch, " ")
	}
	close(ch)
}

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
	PowNumOne()
	fmt.Print("\n")
	SquaresNumTwo()
	fmt.Print("\n")
	SquaresNumThree([]int{2, 4, 6, 8, 10})
	fmt.Print("\n")
}
