package main

import (
	"errors"
	"fmt"
)

func deleteEl3(sl []int, num int) ([]int, error) {
	if num > len(sl)-1 || num < 0 {
		return nil, errors.New("неккоректный индекс")
	} else {
		sl = append(sl[:num], sl[num+1:]...)
	}
	return sl, nil
}

func main() {
	sl := []int{1, 5, 6, 12, 4, 54, 77}

	res, err := deleteEl3(sl, 5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(res)
	}
}
