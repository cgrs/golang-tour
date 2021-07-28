// +build OMIT

package main

import "fmt"

func main() {
	a := make([]int, 5)
	pintaSlice("a", a)

	b := make([]int, 0, 5)
	pintaSlice("b", b)

	c := b[:2]
	pintaSlice("c", c)

	d := c[2:5]
	pintaSlice("d", d)
}

func pintaSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
