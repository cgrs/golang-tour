// +build OMIT

package main

import (
	"fmt"
	"time"
)

type MiError struct {
	Cuándo time.Time
	Qué    string
}

func (e *MiError) Error() string {
	return fmt.Sprintf("en %v, %s",
		e.Cuándo, e.Qué)
}

func ejecutar() error {
	return &MiError{
		time.Now(),
		"no funcionó",
	}
}

func main() {
	if err := ejecutar(); err != nil {
		fmt.Println(err)
	}
}
