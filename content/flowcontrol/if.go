// +build OMIT

package main

import (
	"fmt"
	"math"
)

func raíz(x float64) string {
	if x < 0 {
		return raíz(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(raíz(2), raíz(-4))
}
