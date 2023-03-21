package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	hello() // normal function call
	go hello() // go keyword to start a goroutine i.e. a lightweight thread
	time.Sleep(100 * time.Millisecond)
	// If no sleep, the main function will exit before the goroutine is executed

	// Another methods is with sync group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		hello()
		wg.Done()
	}()
	wg.Wait()

}

func hello() {
	fmt.Println("Hello World!")
}