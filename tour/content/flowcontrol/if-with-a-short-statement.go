// +build OMIT

package main

import (
	"fmt"
	"math"
)

func potencia(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		potencia(3, 2, 10),
		potencia(3, 3, 20),
	)
}
