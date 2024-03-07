package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool, 2)

	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- true
			fmt.Println("sending into channel")
		}
	}()

	for value := range ch {
		fmt.Println("receiving from channel", value)
	}
}
