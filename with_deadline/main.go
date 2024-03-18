package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Microsecond*10))
	defer cancel()
	compute := func() <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			// Simulate work.
			deadline, _ := ctx.Deadline()
			if deadline.Sub(time.Now().Add(time.Microsecond*50)) <= 0 {
				fmt.Println("time is not sufficient")
				return
			}
			time.Sleep(5 * time.Second)

			// Report result.
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				return
			}
		}()
		return ch
	}
	// Wait for the work to finish. If it takes too long move on.
	ch := compute()
	d, ok := <-ch
	if ok {
		fmt.Printf("work complete: %s\n", d)
	}
}
