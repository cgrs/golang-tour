// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vértice struct {
	X, Y float64
}

func Abs(v Vértice) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Escalar(v *Vértice, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vértice{3, 4}
	Escalar(&v, 10)
	fmt.Println(Abs(v))
}
