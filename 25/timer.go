package main

import (
	"fmt"
	"time"
)

/*
time.NewTimer создает новый таймер, который отправит текущее время по каналу после заданного промежутка времени

Канал C в структуре *time.Timer используется для получения сигнала о том, что таймер истек.
Этот канал позволяет вам ожидать истечения таймера и получать время истечения, если это необходимо.
Использование timer.C — это стандартный способ работы с таймерами в Go.
*/
func main() {
	fmt.Println("Привет через пару секунд")
	timer := time.NewTimer(2 * time.Second)
	<-timer.C //// Ожидаем сигнала от таймера
	fmt.Println("Привет!!")
}