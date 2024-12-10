package main

import (
	"errors"
	"fmt"
)

type point struct {
	x, y int
}

func newPoint(x int, y int) *point {
	return &point{x: x, y: y}
}

func distance(x, y int) (int, error) {
	if x < y {
		return 0, errors.New("х не может меньше y")
	}

	res := newPoint(x, y)
	return res.x - res.y, nil
}

func main() {
	x, y := 2345, 432
	res, err := distance(x, y)
	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		fmt.Println(res)
	}
}
