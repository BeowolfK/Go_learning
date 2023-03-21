package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // Goroutines is not executed by order => No synchronization 
var counter int = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go hello()
		go increment()
	}
	wg.Wait()
	
}


func hello() {
	fmt.Println("Hello ", counter, "!")
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}