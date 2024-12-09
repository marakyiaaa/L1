package main

import "fmt"

// адаптер - целевлй интерфейс
type Adapter interface {
	lala()
}

// адаптруемый
type Adaptee struct {
}

type ConcAdaptee struct {
	*Adaptee
}

// Метод адаптируемого класса
func (a *Adaptee) AdaptedLala() {
	fmt.Println("func AdaptedLala")
}

func (ca *ConcAdaptee) lala() {
	ca.AdaptedLala()
}

func NewAdptaer(adaptee *Adaptee) Adapter {
	return &ConcAdaptee{adaptee}
}

func main() {
	fmt.Println("\nAdapter demo:")
	adapter := NewAdptaer(&Adaptee{})
	adapter.lala()
}
