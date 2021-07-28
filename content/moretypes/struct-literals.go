// +build OMIT

package main

import "fmt"

type Punto struct {
	X, Y int
}

var (
	v1 = Punto{1, 2}  // es de tipo Punto
	v2 = Punto{X: 1}  // Y:0 es implícito
	v3 = Punto{}      // X:0 y Y:0
	p  = &Punto{1, 2} // es de tipo *Punto
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
