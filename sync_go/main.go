package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	balanceCheck()
	balanceCheckAtomic()
}

func balanceCheck() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	balance := 0

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mx.Lock()
			balance = balance + 1
			mx.Unlock()
		}()
		go func() {
			defer wg.Done()
			mx.Lock()
			balance = balance - 1
			mx.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}

func balanceCheckAtomic() {
	wg := sync.WaitGroup{}

	var balance int64

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&balance, 1)
		}()
		go func() {
			defer wg.Done()
			atomic.AddInt64(&balance, -1)
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}
