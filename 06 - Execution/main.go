package main

import (
	"fmt"
	"log"
)

func recovering(){
	if err := recover(); err != nil {
		log.Println(err)
	}
	// recover return nil if there is no panic
	// stops the panicking sequence and returns the value passed to the call of panic
}

func panicking() {
	defer recovering()
	fmt.Println("panicking")
	panic("error in panicking function")
	fmt.Println("end") // not executed
}

func main() {
	fmt.Println("start")
	s := "middle"
	defer fmt.Println(s) //  executed at the end of the function before return
	// defer is executed in LIFO order by loading function in a stack
	s = "end"
	fmt.Println("end")

	// defer always executed even if there is a panic
	
	panicking()
	// Always executed
	panic("throwing an error")
	fmt.Println("end") // not executed
}