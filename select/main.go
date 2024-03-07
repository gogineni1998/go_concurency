package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case a := <-ch1:
			fmt.Println(a)

		case b := <-ch2:
			fmt.Println(b)
		}
	}
}
