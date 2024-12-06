package main

import (
	"fmt"
	"time"
)

/*остановка с помощью канала stop*/
func stopGoroutineModelCh() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop: //если сигнал остановки получен
				fmt.Println("Stopped")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
	stop <- true
	time.Sleep(100 * time.Millisecond)
}

func main() {
	stopGoroutineModelCh()
}
