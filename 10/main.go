package main

import "fmt"

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, 11, -21.0, 32.5}
	m := make(map[int][]float64) //ключ int,значение - срез температур

	for _, val := range temperature {
		key := int(val/10.0) * 10
		m[key] = append(m[key], val)
	}
	fmt.Println(m)
}
