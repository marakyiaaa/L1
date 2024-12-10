package main

import (
	"fmt"
	"time"
)

/*
Функция time.After возвращает канал, который будет закрыт через заданное количество времени
*/
func main() {
	fmt.Println("Привет через пару секунд")
	<-time.After(2 * time.Second)
	fmt.Println("Привет!!")
}
