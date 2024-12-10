package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, canel := context.WithTimeout(context.Background(), 5*time.Second)
	defer canel()

	fmt.Println("Привет через пять секунд")
	<-ctx.Done()
	fmt.Println("Привет!!")
}
