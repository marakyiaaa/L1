package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	a := 0
	var n time.Duration
	fmt.Scan(&n)
	n *= time.Second
	endCh := make(chan struct{})
	end := time.Now().Add(n)

	go func() {
		defer close(ch)
		for {
			a += 1
			ch <- a
		}
	}()

	go func() {
		for {
			if time.Now().Compare(end) > 0 {
				endCh <- struct{}{}
				return
			}
		}
	}()

	for {
		select {
		case <-ch:
		case <-endCh:
			fmt.Println("END")
			return
		}
	}
}
