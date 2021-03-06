// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vértice struct {
	X, Y float64
}

func (v Vértice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vértice) Escalar(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vértice{3, 4}
	v.Escalar(10)
	fmt.Println(v.Abs())
}
