// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // apunta hacia i
	fmt.Println(*p) // lee i a través del puntero
	*p = 21         // cambia el valor de i a través del puntero
	fmt.Println(i)  // ve el nuevo valor de i

	p = &j         // apunta a j
	*p = *p / 37   // divide j a través del puntero
	fmt.Println(j) // ve el nuevo valor de j
}
