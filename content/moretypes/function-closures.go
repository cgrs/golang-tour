// +build OMIT

package main

import "fmt"

func sumador() func(int) int {
	suma := 0
	return func(x int) int {
		suma += x
		return suma
	}
}

func main() {
	pos, neg := sumador(), sumador()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
