// +build OMIT

package main

import "fmt"

type Vértice struct {
	X, Y float64
}

func (v *Vértice) Escalar(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func FuncEscalar(v *Vértice, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vértice{3, 4}
	v.Escalar(2)
	FuncEscalar(&v, 10)

	p := &Vértice{4, 3}
	p.Escalar(3)
	FuncEscalar(p, 8)

	fmt.Println(v, p)
}
