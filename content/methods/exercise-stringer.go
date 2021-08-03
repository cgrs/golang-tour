// +build OMIT

package main

import "fmt"

type DirecciónIP [4]byte

// TODO: Añade un método "String() string" al tipo DirecciónIP.

func main() {
	servidores := map[string]DirecciónIP{
		"Interno":    {127, 0, 0, 1},
		"DNS Google": {8, 8, 8, 8},
	}
	for nombre, ip := range servidores {
		fmt.Printf("%v: %v\n", nombre, ip)
	}
}
