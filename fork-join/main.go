package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		count++
	}()

	wg.Wait()
	fmt.Println(count)
}
