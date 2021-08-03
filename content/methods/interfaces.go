// +build no-build OMIT

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MiFlotante(-math.Sqrt2)
	v := Vértice{3, 4}

	a = f  // a es MiFlotante; implementa Abser
	a = &v // a es *Vértice; implements Abser

	// En la siguiente línea, v es de tipo Vértice (no *Vértice)
	// y no implementa Abser.
	a = v

	fmt.Println(a.Abs())
}

type MiFlotante float64

func (f MiFlotante) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vértice struct {
	X, Y float64
}

func (v *Vértice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
