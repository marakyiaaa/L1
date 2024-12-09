package main

import (
	"errors"
	"fmt"
)

func deleteEl2(sl []int, num int) ([]int, error) {
	if num >= len(sl) || num < 0 {
		return nil, errors.New("неккоректный индекс")
	}

	// Копируем элементы из части слайса sl[num+1:] в часть слайса sl[num:]
	copy(sl[num:], sl[num+1:])

	sl = sl[:len(sl)-1] // Уменьшаем длину слайса на 1

	return sl, nil
}
func main() {
	sl := []int{1, 5, 6, 12, 6}

	res, err := deleteEl2(sl, -2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(res)
	}
}
