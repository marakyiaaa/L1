package main

import (
	"context"
	"fmt"
	"time"
)

/*остановка с помощью канала Context*/
func stopGoroutineContext() {
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done(): //если сигнал остановки получен
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second) //замедляем главную горутину для того,чтобы дргуя успела сделать несколько итераций
	stop()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Stopped")
}
func main() {
	stopGoroutineContext()
}
