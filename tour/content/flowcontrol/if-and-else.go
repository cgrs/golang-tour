// +build OMIT

package main

import (
	"fmt"
	"math"
)

func potencia(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// la variable v no puede utilizarse aquí
	return lim
}

func main() {
	fmt.Println(
		potencia(3, 2, 10),
		potencia(3, 3, 20),
	)
}
