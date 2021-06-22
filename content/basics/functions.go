// +build OMIT

package main

import "fmt"

func suma(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(suma(42, 13))
}
