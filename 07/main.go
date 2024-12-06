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
	mutex := sync.Mutex{}
	m := make(map[string]int)

	//запись
	mutex.Lock()
	m["key"] = 123
	mutex.Unlock()

	//чтение
	mutex.Lock()
	value := m["key"]
	mutex.Unlock()

	fmt.Println(m)
	fmt.Println(value)
}

func recordSyncMap() {

}
func main() {
	recordMutex()

}
