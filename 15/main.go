package main

import (
	"fmt"
	"log"
)

/*
var justString string // глобальная переменная => проблемам с управлением состоянием и тестирования

func someFunc() {
	v := createHugeString(1 << 10) //выделать слишком большой размер
	justString = v[:100] // нет гарантии, что в строке точно есть 100 элементов => паника,если их меньше
}
func main() {
	someFunc()
}
*/

func createHugeString(len int) string {
	var str string
	for i := 0; i < len; i++ {
		str += "s"
	}
	return str
}

func someFunc() string {
	v := createHugeString(1 << 10)

	if len(v) < 100 {
		log.Fatal("Строка слишком короткая")
	}
	return v[:100]
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
