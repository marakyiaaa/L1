package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(50), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(52), nil)
	fmt.Println(new(big.Int).Add(a, b)) //+
	fmt.Println(new(big.Int).Sub(a, b)) //-
	fmt.Println(new(big.Int).Mul(a, b)) //*
	fmt.Println(new(big.Int).Div(a, b)) // /
}
