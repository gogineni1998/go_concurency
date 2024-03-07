package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	i := 1
	for i = 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//time.Sleep(time.Second * 3)
			fmt.Println(i)
		}(i)
	}
	fmt.Println(i)
	wg.Wait()
}
