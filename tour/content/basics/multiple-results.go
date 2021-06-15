// +build OMIT

package main

import "fmt"

func intercambiar(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := intercambiar("hola", "mundo")
	fmt.Println(a, b)
}
