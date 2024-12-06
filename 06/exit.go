package main

import (
	"fmt"
	"os"
	"time"
)

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
	stopGoroutineModeExit()
}
