// +build no-build OMIT

package main

import "golang.org/x/tour/reader"

type MiLector struct{}

// TODO: Añade el método Read([]byte) (int, error) al tipo MiLector.

func main() {
	reader.Validate(MiLector{})
}
