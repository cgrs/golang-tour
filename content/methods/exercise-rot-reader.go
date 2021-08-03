// +build no-build OMIT

package main

import (
	"io"
	"os"
	"strings"
)

type LectorRot13 struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("¡Unf erfhrygb ry ravtzn!")
	r := LectorRot13{s}
	io.Copy(os.Stdout, &r)
}
