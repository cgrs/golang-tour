// +build OMIT

package main

import "fmt"

func suma(s []int, c chan int) {
	suma := 0
	for _, v := range s {
		suma += v
	}
	c <- suma // envía suma a c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go suma(s[:len(s)/2], c)
	go suma(s[len(s)/2:], c)
	x, y := <-c, <-c // recibe desde c

	fmt.Println(x, y, x+y)
}
