// +build OMIT

package main

import "fmt"

func separar(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(separar(17))
}
