package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	a := 42
	b := a
	fmt.Printf("a = %d, b = %d\n", a, b)
	a = 27
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Pointer
	// A pointer is a variable that stores the memory address of another variable
	// &a returns the memory address of a
	// *b returns the value stored at the memory address b
	var a2 int = 42
	var b2 *int = &a2
	fmt.Printf("a2 = %d, b2 = %d\n", a2, *b2)
	a2 = 27
	fmt.Printf("a2 = %d, b2 = %d\n", a2, *b2)

	// Pointer with struct
	var ms *myStruct
	fmt.Println(ms) // nil, struct is not initialized
	ms = new(myStruct)
	(*ms).foo = 42 //* operator as a lower priority than . operator
	fmt.Println(ms.foo) //syntax also possible with Go

	// Pointer with slice and map
	// Slice and map are reference types
	// They are pointer to the first value of array
	s := []int{1, 2, 3}
	s2 := s
	fmt.Println(s, s2)
	s[0] = 42
	fmt.Println(s, s2) // s2 is also modified



}