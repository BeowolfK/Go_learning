package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int = 0
var m sync.RWMutex // RWMutex lock Read and Write and synchronize goroutine execution

func main() {
	runtime.GOMAXPROCS(100) // Run with 100 threads
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // Lock Read of value
		go hello()
		m.Lock() // Lock Write of value
		go increment()
	}
	wg.Wait()
	
}


func hello() {
	fmt.Println("Hello ", counter, "!")
	m.RUnlock() // Unlock Read
	wg.Done()
}

func increment() {
	counter++
	m.Unlock() // Unlock Write
	wg.Done()
}