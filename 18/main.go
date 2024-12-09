package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value++
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ { //запускаем 10 горутин
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
			fmt.Println(c.Value)
		}()
	}

	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Итоговое значение счетчика:", c.Value)
}
