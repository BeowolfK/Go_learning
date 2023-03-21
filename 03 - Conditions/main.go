package main

import "fmt"

// No ternary in GO

func main() {
    // if <declaration>; condition
    if num := 10; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    if true && true {}
    if true || false {}
    if !false {}

    // switch
    i := 10
    switch{
    case i <= 10:
        fmt.Println("10 or less")
        fallthrough // continue to next case 
    case i <= 20:
        fmt.Println("20 or less")
    default:
        fmt.Println("other")
    }

    s := "a"
    switch s {
    case "a", "b", "c":
        fmt.Println("a, b or c")
    case "d", "e", "f":
        fmt.Println("d, e or f")
    default:
        fmt.Println("other")
    }


}