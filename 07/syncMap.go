package main

import (
	"fmt"
	"sync"
)

func recordSyncMap() {
	var m sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1) //запускаем счетчик
		go func(i int) {
			defer wg.Done()  //счетчик -
			m.Store(i, i+10) //записываем в map ключ-значение
		}(i)
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		value, found := m.Load(i) //читаем из map (value any, ok bool)
		if found {
			fmt.Printf("Key: %d, Value: %v\n", i, value)
		}
	}
}
func main() {
	recordSyncMap()
}
