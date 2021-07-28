// +build OMIT

package main

import "fmt"

type Punto struct {
	X int
	Y int
}

func main() {
	fmt.Println(Punto{1, 2})
}
