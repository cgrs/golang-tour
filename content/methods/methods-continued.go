// +build OMIT

package main

import (
	"fmt"
	"math"
)

type MiFlotante float64

func (f MiFlotante) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MiFlotante(-math.Sqrt2)
	fmt.Println(f.Abs())
}
