// +build OMIT

package main

import "fmt"

func fibonacci(c, salir chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-salir:
			fmt.Println("salir")
			return
		}
	}
}

func main() {
	c := make(chan int)
	salir := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		salir <- 0
	}()
	fibonacci(c, salir)
}
