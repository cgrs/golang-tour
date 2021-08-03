// +build OMIT

package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) String() string {
	return fmt.Sprintf("%v (%v años)", p.Nombre, p.Edad)
}

func main() {
	a := Persona{"Arthur Dent", 42}
	z := Persona{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
