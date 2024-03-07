package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inc := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Println(i)
		}()
	}
	inc(&wg)
	wg.Wait()
}
