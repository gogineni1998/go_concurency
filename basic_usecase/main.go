package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, evilevilNinja := range evilNinjas {
		go attack(evilevilNinja)
	}
	time.Sleep(time.Second * 1)
}

func attack(trojon string) {
	fmt.Println("The ninjas are handeling trojen : ", trojon)
	time.Sleep(time.Second * 1)
}
