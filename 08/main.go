package main

import (
	"fmt"
	"log"
)

/*
& – побитовое И (AND)

x & 0 = 0
x & 1 = x

| – побитовое ИЛИ (OR)
0 | 0 = 0
1 | 0 = 1
0 | 1 = 1
1 | 1 = 1

x | 0 = X
X | 1 = 1

xxxxxxxxx
|
000010000
=
xxxx1xxxx

000000100

^ – побитовое исключающее ИЛИ (XOR)
&^ – побитовое И-НЕТ (AND NOT)
*/

func setBit(num int64, i int, bit int) int64 {
	if bit != 0 && bit != 1 {
		panic("must be 0 or 1")
	}
	if bit == 0 {
		return num &^ (1 << i)
	} else {
		return num | (1 << i)
	}
}

func main() {
	var n int64
	var i, bit int

	log.Println("Введите число:")
	fmt.Scan(&n)

	log.Println("Введите индекс бита, который хотите изменить:")
	fmt.Scan(&i)

	log.Println("Введите значение бита на которое хотите изменить (0 или 1):")
	fmt.Scan(&bit)

	fmt.Printf("%064b\n => %d\n", n, n)
	m := setBit(n, i, bit)
	fmt.Printf("%064b\n => %d\n", m, m)
}
