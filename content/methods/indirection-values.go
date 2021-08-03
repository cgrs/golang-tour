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

func FuncAbs(v Vértice) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vértice{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(FuncAbs(v))

	p := &Vértice{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(FuncAbs(*p))
}
