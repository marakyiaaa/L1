package main

import (
	"fmt"
	"sync"
)

func SumNum(arr []int) {
	ch1 := make(chan int)
	sum := 0
	wg1 := sync.WaitGroup{}

	for _, val := range arr {
		wg1.Add(1) //увеличиваем счетчик на 1
		go func(i int) {
			defer wg1.Done() //счетчик -1
			ch1 <- i * i
		}(val) // возвращаем значения квадратов
	}

	go func() {
		wg1.Wait() //ожидаем счетчик == 0 => goroutines закончили работу
		close(ch1) //закрываем канал
	}()

	for val := range ch1 { //проходимся по значениям в канале
		sum += val
	}
	fmt.Println(sum)
}

func main() {
	SumNum([]int{2, 4, 6, 8, 10})
}
