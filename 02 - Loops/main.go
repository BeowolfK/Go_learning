package main

import "fmt"

func main() {

	// For loop is same as C syntax
    // for init; condition; post {}
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break //break a loop
    }

    for n := 0; n <= 5; n = n+1 {
        if n%2 != 0 {
            continue // go to next iteration
            // break // break the loop
        }
        fmt.Println(n)
    }

    i := 0 // For with global scope value
    for i < 3 { // same ad for ; i<3 ; {}
        fmt.Printf("i = %d\n", i)
        i++ // can also increment here
    }
    fmt.Printf("i = %d\n", i)

    // For with range
    s := [6]int{1, 2, 3, 4, 5}
    for idx, val := range s {
        fmt.Printf("idx = %d, val = %d\n", idx, val)
    }

    // For with range and _
    population := map[string]int{
		"Chine": 1.411e9,
		"Inde" : 1.372e9,
		"USA"  :  3.331e8,
	}
    for _, pop := range population {
        fmt.Println(pop)
    }

}