// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vértice struct {
	X, Y float64
}

func (v *Vértice) Escalar(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vértice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vértice{3, 4}
	fmt.Printf("Antes de escalar: %+v, Abs: %v\n", v, v.Abs())
	v.Escalar(5)
	fmt.Printf("Después de escalar: %+v, Abs: %v\n", v, v.Abs())
}
