package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := generated(ctx)

	for val := range ch {
		if val == 250 {
			cancel()
		}
		fmt.Println(val)
	}
}

func generated(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				{
					close(ch)
					return
				}
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
