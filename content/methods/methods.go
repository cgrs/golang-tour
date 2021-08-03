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

func main() {
	v := Vértice{3, 4}
	fmt.Println(v.Abs())
}
