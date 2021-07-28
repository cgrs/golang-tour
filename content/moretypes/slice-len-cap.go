// +build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	pintaSlice(s)

	// "Corta" el slice para darle una longitud cero.
	s = s[:0]
	pintaSlice(s)

	// Extiende su longitud.
	s = s[:4]
	pintaSlice(s)

	// Elimina los dos primeros valores.
	s = s[2:]
	pintaSlice(s)
}

func pintaSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
