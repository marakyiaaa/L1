package main

import (
	"fmt"
	"time"
)

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

func main() {
	stopGoroutineWaitPanic()
}
