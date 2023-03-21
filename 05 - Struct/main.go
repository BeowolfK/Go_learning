package main

import "fmt"

	
type person struct {
    name string
    age  int
}

type student struct {
	person
	grade int
}

func main() {
	var p person = person{name: "John", age: 42}
	p2 := person{name: "Doe", age: 42}
	p3 := person{name: "Jane"}
	p3.age = 21

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)

	doctor := struct{name string}{name: "Dr. Brown"}
	doctor2 := &doctor
	doctor2.name = "Dr. Strange"
	fmt.Println(doctor)
	fmt.Println(*doctor2)

	s := student{} // student is a child class of person
	s.name = "John"
	s.age = 42
	s.grade = 18
	fmt.Println(s)
}