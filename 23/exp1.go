package main

import (
	"fmt"
	"log"
)

func deleteEl(sl []int, num int) []int {
	if num >= len(sl) || num < 0 {
		log.Fatalf("неккоректный индекас")
	}

	sl[num] = sl[len(sl)-1] //перемещаем элемент по индексу в конец слайса
	sl = sl[:len(sl)-1]     // Уменьшаем длину слайса на 1

	return sl
}
func main() {
	sl := []int{1, 5, 7, 12, 6}
	fmt.Println(deleteEl(sl, 3))
}
