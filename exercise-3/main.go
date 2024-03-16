// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		for n := range in {
			out1 <- n * n
		}
		close(out1)
	}()

	go func() {
		defer close(out2)
		for n := range in {
			out2 <- n * n
		}
	}()
	return out1, out2
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	wg := sync.WaitGroup{}
	for _, ch := range cs {
		wg.Add(1)
		go func() {
			for val := range ch {
				out <- val
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 34, 5678)
	for val := range merge(square(in)) {
		fmt.Println(val)
	}
	// TODO: fan out square stage to run two instances.

	// TODO: fan in the results of square stages.

}
