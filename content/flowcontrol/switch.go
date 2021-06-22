// +build OMIT

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go se ejecuta en ")
	switch so := runtime.GOOS; so {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", so)
	}
}
