package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		ch <- "one"
	}()

	select {
	case m := <-ch:
		fmt.Println(m)

	case <-time.After(time.Microsecond * 100000000):
		fmt.Println("timmed out")
	}
}
