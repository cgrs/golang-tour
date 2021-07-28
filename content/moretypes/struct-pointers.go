// +build OMIT

package main

import "fmt"

type Punto struct {
	X int
	Y int
}

func main() {
	v := Punto{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
