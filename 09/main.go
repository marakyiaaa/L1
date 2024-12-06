package main

import "fmt"

func workerRecordArray(array []int, ch1 chan<- int) {
	for _, data := range array {
		ch1 <- data //записываем данные из массивы в канал
	}
	close(ch1) //закрываем канал
}

func workerOutput(ch1 <-chan int, ch2 chan int) {
	for data := range ch1 {
		ch2 <- data * 2 //записываем данные из канала в канал и * 2
	}
	close(ch2) //закрываем канал
}

func conveyor() {
	array := []int{6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go workerRecordArray(array, ch1)
	go workerOutput(ch1, ch2)
	for value := range ch2 {
		fmt.Println(value) //вывод stdout
	}
}

func main() {
	conveyor()
}
