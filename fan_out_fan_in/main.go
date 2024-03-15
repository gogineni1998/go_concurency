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
		flag := 0
		for n := range in {
			if flag == 0 {
				out1 <- n * n
				flag = 1
			} else {
				out2 <- n * n
				flag = 0
			}
		}
		close(out1)
		close(out2)
	}()
	return out1, out2
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	wg := sync.WaitGroup{}
	for _, c := range cs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := <-c
			out <- n * n
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3)
	for val := range merge(square(in)) {
		fmt.Println(val)
	}
	// TODO: fan out square stage to run two instances.

	// TODO: fan in the results of square stages.

}
