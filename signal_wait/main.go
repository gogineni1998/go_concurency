package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mx := sync.Mutex{}
	c := sync.NewCond(&mx)
	wg := sync.WaitGroup{}

	m := make(map[int]string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		_, ok := m[0]
		if !ok {
			c.Wait()
			fmt.Println(0)
		}
		fmt.Println(m[0])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		_, ok := m[1]
		if !ok {
			c.Wait()
		}
		fmt.Println(m[1])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		m[0] = "dheeraj"
		m[1] = "satya"
		c.L.Unlock()
		c.Signal()
		fmt.Println(2)
		time.Sleep(time.Second * 5)
		c.Signal()
	}()
	wg.Wait()
}
