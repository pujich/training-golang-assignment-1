package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("main process")

	fruits := []string{"apel", "banana", "ceri"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wgEnd()

	first()
	go second()
	go second()

	fmt.Println("Number : ", runtime.NumGoroutine())

	time.Sleep(time.Microsecond * 4)
	fmt.Println("main process DONE")
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Println(fruit)
	wg.Done()
}

func wgEnd() {
	var wg sync.WaitGroup
	wg.Wait()
}

func first() {
	fmt.Println("frist")
}

func second() {
	fmt.Println("second")
}
