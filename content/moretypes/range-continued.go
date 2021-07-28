// +build OMIT

package main

import "fmt"

func main() {
	potencias := make([]int, 10)
	for i := range potencias {
		potencias[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range potencias {
		fmt.Printf("%d\n", value)
	}
}
