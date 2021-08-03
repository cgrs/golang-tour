// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Este método quiere decir que T implementa la interfaz I,
// pero no hace falta declarar la implementación explícitamente.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hola"}
	i.M()
}
