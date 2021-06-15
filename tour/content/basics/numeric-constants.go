// +build OMIT

package main

import "fmt"

const (
	// Crea un número enorme al desplazar un bit 100 posiciones.
	// En otras palabras, el número en binario correspondiente a un 1 seguido de 100 ceros.
	Enorme = 1 << 100
	// Desplazándolo a la derecha de nuevo 99 posiciones, nos quedaría 1<<1, o sea, 2.
	Pequeño = Enorme >> 99
)

func necesitaInt(x int) int { return x*10 + 1 }
func necesitaFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(necesitaInt(Pequeño))
	fmt.Println(necesitaFloat(Pequeño))
	fmt.Println(necesitaFloat(Enorme))
}
