package main

import "fmt"

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

func main() {
	SquaresNumTwo()
}
