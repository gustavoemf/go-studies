package main

import "fmt"

// package/global level scope
// var w = 100

func main() {

	x := 42
	y := "string"
	z := true

	// Isso daria erro, pois eu estaria redeclarando a variável z
	// z := false

	fmt.Printf("x = %v, type: %T\n", x, x)
	fmt.Printf("y = %v, type: %T\n", y, y)
	fmt.Printf("z = %v, type: %T", z, z)
}
