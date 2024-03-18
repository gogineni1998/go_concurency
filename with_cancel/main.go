package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	ctx, cancel := context.WithCancel(context.Background())
	generator := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := range 500 {
				select {
				case ch <- i:
				case <-ctx.Done():
					{
						return
					}
				}
			}
		}()
		return ch
	}

	func() {
		for val := range generator() {
			if val == 5 {
				cancel()
			}
			fmt.Println(val)
		}
	}()

	// Create a context that is cancellable.

}
