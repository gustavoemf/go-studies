package main

import "fmt"

const (
	naotipada     = 11 // tipo definido em runtime
	tipada    int = 10 // tipo definido em compilation time
)

func main() {
	// Crie constantes tipadas e não-tipadas.
	// Demonstre seus valores.

	fmt.Print(naotipada, tipada)
}
