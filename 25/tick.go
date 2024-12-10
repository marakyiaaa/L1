package main

import (
	"fmt"
	"time"
)

/*
Функция time.Tick возвращает канал, который будет отправлять сигналы через заданные интервалы времени
*/
func main() {
	ticker := time.Tick(1 * time.Second)
	for i := 0; i < 6; i++ {
		<-ticker
		fmt.Println("Тик")
	}
	fmt.Println("Привет !")
}
