package main

import "fmt"

func main() {
	hello("World")
	hello2("Hello", "World")
	name := "Jane"
	name_change(&name) // pass the address of the variable will change the value in global scope
	fmt.Println(name)
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0))

	r := rectangle{width: 10, height: 5}
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())
	rp := &r 
	// Go automatically convert values and pointers for method calls
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perimeter())

	// Anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	// Anonymous function with parameters
	// Anonymous function have all parameters of the current scope 
	func(msg string) {
		fmt.Println(msg)
	}("Anonymous function with parameters")
}

func hello(msg string) {
	fmt.Printf("Hello %s!\n", msg)
}
func hello2(msg, name string) { // Define all parameters of the same type
	fmt.Printf("%s %s!\n", msg, name)
}

// Functions with pointer
func name_change(name *string) {
	*name = "John" // name is changed in global scope
}

// Functions with multiple parameters
func sum(value ...int) int{ // return type is defined after the parameters
	result := 0
	for _, v := range value {
		result += v
	}
	return result
}

// Function with static return 
func sum2(value ...int) (result int) { // return type is defined after the parameters
	// no need to initialize result
	for _, v := range value {
		result += v
	}
	return // return result
}

// Return error value
func divide(a, b int) (int, error) {
	// Don't use panic to handle error, break the flow of the program
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

// Methods : POO
type rectangle struct {
	width, height int
}

func (r rectangle) area() int { // r is a copy of the object
	return r.width * r.height
}

func (r *rectangle) perimeter() int { // r is a pointer to the object
	return 2 * (r.width + r.height)
}