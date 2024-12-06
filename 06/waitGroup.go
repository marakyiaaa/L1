package main

import (
	"fmt"
	"sync"
	"time"
)

/*остановка с помощью канала WaitGroup*/
func stopGoroutineWaitGroup() {
	wg := sync.WaitGroup{}

	wg.Add(5) //запуск счетчика на 5
	go func() {
		for {
			fmt.Println("Working...")
			wg.Done() //счетчик -1
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait() //ждем завершения работы горутины
	fmt.Println("Stopped")
}

func main() {
	stopGoroutineWaitGroup()
}
