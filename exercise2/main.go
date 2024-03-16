package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()

	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	for val := range square(generator(1, 2, 3, 4, 5, 6)) {
		fmt.Println(val)
	}

}
