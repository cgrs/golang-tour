// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter es concurrentemente seguro de utilizar.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc incrementa el contador para la clave dada.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Bloquear para que sólo pueda acceder una gorrutina a la vez al mapa c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value devuelve el valor actual del contador para una clave dada.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Bloquear para que sólo pueda acceder una gorrutina a la vez al mapa c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("una clave")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("una clave"))
}
