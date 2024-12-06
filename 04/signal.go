package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Способ завершения работы: с помощью перехвата сигнала Ctrl+C - SIGINT*/
func workerOne(workerId int, ch <-chan int, wg *sync.WaitGroup) { //только канал для чтения
	defer wg.Done() //счетчик -1
	for data := range ch {
		fmt.Printf("Worker %d: %d\n", workerId, data)
	}
}

func modeOne() {
	var a, workersCount int
	fmt.Scan(&workersCount)

	ch := make(chan int)
	sigChan := make(chan os.Signal, 1)     //буферизированный канал для передачи сигналов
	signal.Notify(sigChan, syscall.SIGINT) // signal interrupt (прерывание)
	wg := sync.WaitGroup{}

	//запускаем пул воркеров
	for i := 0; i <= workersCount; i++ {
		wg.Add(1) //счетчик +1
		go workerOne(i, ch, &wg)
	}

	go func() {
		for {
			a += 1
			select {
			case ch <- a: //если происходит запись в канал
				time.Sleep(100 * time.Millisecond) //замедление для визуального понимания
			case <-sigChan: //если пойман сигнал Ctrl+C
				close(ch)
				return
			}
		}
	}()
	wg.Wait() //ждем завершения всех воркеров
	fmt.Println("\nEND")
}

func main() {
	modeOne()
}
