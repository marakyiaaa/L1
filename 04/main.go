package main

import (
	"context"
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

/*Способ завершения программы: context, котрый передает сигнал о завершении программы*/
func workerTwo(workerId int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Woker %d: %d\n", workerId, data)
	}
}

func modeTwo() {
	var a, workersCount int
	fmt.Scan(&workersCount)

	ch := make(chan int)
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	/*создаем контекст, ожидаюзий сигнал SIGINT и функцию stop,завершающую конткст при данном условии*/
	defer stop()

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go workerTwo(i, ch, &wg)
	}

	go func() {
		for {
			a += 1
			select {
			case ch <- a:
				time.Sleep(100 * time.Millisecond)
			case <-ctx.Done(): //если подается сигнал в поток => подан Ctrl+C и контекст завершился
				close(ch)
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("\nEND")

}

func main() {
	modeOne()
	//modeTwo()
}
