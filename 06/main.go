package main

import (
	"context"
	"fmt"
	"os"
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

/*Panict*/
func stopGoroutineWaitPanic() {
	defer func() { //recover млжет использоваться только с defer
		if r := recover(); r != nil { //перехвата паники
			fmt.Println("Stopped")
		}
	}()
	go func() {
		for {
			fmt.Println("Working...")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	panic("Stopped")
}

/*os.Exit*/
func stopGoroutineModeExit() {
	go func() {
		fmt.Println("Working..")
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
	fmt.Println("Stopped")
	os.Exit(0)

}

func main() {
	stopGoroutineWaitGroup()
	//stopGoroutineContext()
	//stopGoroutineModelCh()
	//stopGoroutineWaitPanic()
	//stopGoroutineModeExit()
}
