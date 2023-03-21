package main

import (
    "fmt"
	"math"
)

type Forme interface {
	// An interface is a set of method signatures
    Air() float64
    Perimetre() float64
}

type Rectangle struct {
    largeur  float64
    longueur float64
}

// In order to implement an interface in Go, we just need to implement all the methods of the interface

func (r Rectangle) Air() float64 {
	return r.largeur * r.longueur
}


func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur * r.longueur)
}

type Cercle struct {
	rayon float64
}

func (c Cercle) Air() float64 {
	return math.Pi * c.rayon * c.rayon
}
func (c Cercle) Perimetre() float64 {
	return 2 * math.Pi * c.rayon
}

func main() {
    var f Forme = Rectangle{5.0, 4.0} 
	// Affectation of structure Rectangle to Forme interface
    r := Rectangle{5.0, 4.0}
	// Implicit conversion of structure Rectangle to Forme interface
	// Go compiler do all the work
    fmt.Printf("Type de f : %T\n", f)
    fmt.Printf("Valeur de f : %v\n", f)
    fmt.Println("Air du rectangle r :", f.Air())
    fmt.Println("f == r ? ", f == r)

	var c Forme = Cercle{5.0}
	fmt.Printf("Type de f : %T\n", c)
    fmt.Printf("Valeur de f : %v\n", c)
    fmt.Println("Air du rectangle r :", c.Air())

}