// +build OMIT

package main

import "fmt"

const Pi = 3.14

func main() {
	const Mundo = "🌍"
	fmt.Println("Hola", Mundo)
	fmt.Println("Feliz", "Día", Pi)

	const Verdad = true
	fmt.Println("¿Go mola?", Verdad)
}
