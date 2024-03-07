package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	genMsg := func() {
		ch1 <- 1
	}

	relayMsg := func() {
		a := <-ch1
		ch2 <- a
	}
	go genMsg()
	go relayMsg()
	fmt.Println(<-ch2)
}
