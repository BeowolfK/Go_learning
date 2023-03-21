package main

import "fmt"

// Declare a global constant
const MyString = "hello" 
/*
Naming convention
PascalCase for exported constants
camelCase for internal constants
*/

func main() {
	// Declare integer variables
	var i int
	i = 42
	var j int = i * 2

	// Dynamic type of variables
	k := i * 3

	// Mathematic notation
	l := 3e5

	fmt.Println(i, j, k, l)

	// Boolean type1
	var m bool = true
	n := 1 == 1
	fmt.Println(m, n)

	// Complex notation
	var o complex64 = 2+3i
	var p complex64 = complex(5, 12)
	fmt.Println(o, real(o), imag(o))
	fmt.Println(p, real(p), imag(p))

	// String type
	s := "Hello"
	s2 := "world!"
	fmt.Println(s, s[1], string(s[1])) // index return ascii position of char
	fmt.Println(s + " " + s2)

	byte := []byte(s)
	fmt.Println(byte) //print ascii position of all char

	r := 'a' // ascii position with single quote
	var r2 rune = 'b' // same thing as above
	fmt.Println(r, r2) 

	// Constants => Replaces at compile time
	const (
		a = iota // iota represents constant increasing by 1
		b
		c
		_ // skip a value
		d
	)

	fmt.Printf("a => %v\n", a)
	fmt.Printf("b => %v\n", b)
	fmt.Printf("c => %v\n", c)
	fmt.Printf("d => %v\n", d)

}