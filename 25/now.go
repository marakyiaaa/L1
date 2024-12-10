package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Привет через пару секунд")
	mySleep(2)
	fmt.Println("Привет!!")
}

func mySleep(sec int) {
	start := time.Now()
	for time.Since(start) < time.Duration(sec)*time.Second {
	}
}
