// +build OMIT

package main

import "fmt"

func main() {
	suma := 0
	for i := 0; i < 10; i++ {
		suma += i
	}
	fmt.Println(suma)
}
