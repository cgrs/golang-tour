// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("¿Hoy es sábado?")
	hoy := time.Now().Weekday()
	switch time.Saturday {
	case hoy + 0:
		fmt.Println("Sí.")
	case hoy + 1:
		fmt.Println("No, mañana.")
	case hoy + 2:
		fmt.Println("En dos días.")
	default:
		fmt.Println("Todavía queda mucho.")
	}
}
