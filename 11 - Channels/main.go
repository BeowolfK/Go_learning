package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int) // Create a channel for int
	wg.Add(2)
	go func() { // Receiving go routine
		i := <-ch // Receive value from channel
		close(ch)
		fmt.Println(i)
		wg.Done()
	}()
	go func() { // Sending go routine
		i := 42
		ch <- i // Send value to channel
		i = 27
		wg.Done()
	}()
	wg.Wait()
}
