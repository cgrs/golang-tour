// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	imprimeSlice(s)

	// append funciona en slices nulos.
	s = append(s, 0)
	imprimeSlice(s)

	// El slice crece según lo necesite.
	s = append(s, 1)
	imprimeSlice(s)

	// Se puede añadir más de un elemento a la vez.
	s = append(s, 2, 3, 4)
	imprimeSlice(s)
}

func imprimeSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
