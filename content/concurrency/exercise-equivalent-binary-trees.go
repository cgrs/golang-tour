// +build no-build OMIT

package main

import "golang.org/x/tour/tree"

// Recorrer recorre el árbol t enviando todos los valores
// desde el árbol al canal ch.
func Recorrer(t *tree.Tree, ch chan int)

// Igual determina si los árboles
// t1 y t2 contienen los mismos valores.
func Igual(t1, t2 *tree.Tree) bool

func main() {
}
