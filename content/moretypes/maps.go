// +build OMIT

package main

import "fmt"

type Vértice struct {
	Lat, Long float64
}

var m map[string]Vértice

func main() {
	m = make(map[string]Vértice)
	m["Bell Labs"] = Vértice{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
