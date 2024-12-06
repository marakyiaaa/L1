package main

import (
	"fmt"
	"sync"
)

/*map не является потокобезопасной
При конкурентной записи (без синхронизации) в map
может произойти гонка данных
(когда несколько горутин одновременно пытаются читать и записывать значения в map),
что приведет к непредсказуемым результатам, ошибкам,потере данных,панике.

Для решения этой проблемы можно использовать sync.Mutex
или структуру данных sync.Map*/

func recordMutex() {
	m := make(map[string]int)
	mutex := sync.Mutex{}  //для защиты записи в map
	wg := sync.WaitGroup{} //для ожидания завершения всех горутин

	for i := 0; i < 5; i++ {
		wg.Add(1) //запускаем счетчик
		go func(i int) {
			defer wg.Done() //счетчик -
			mutex.Lock()    // Блокируем мьютекс перед записью в map
			m[fmt.Sprintf("key %d", i)] = i
			mutex.Unlock() // Разблокируем мьютекс после записи
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}

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
	//recordMutex()
	recordSyncMap()
}
