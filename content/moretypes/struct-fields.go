// +build OMIT

package main

import "fmt"

type Punto struct {
	X int
	Y int
}

func main() {
	v := Punto{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
