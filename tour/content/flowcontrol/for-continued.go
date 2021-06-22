// +build OMIT

package main

import "fmt"

func main() {
	suma := 1
	for ;suma < 1000; {
		suma += suma
	}
	fmt.Println(suma)
}
