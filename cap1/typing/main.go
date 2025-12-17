package main

import "fmt"

// int é o tipo subjacente de hotdog
type hotdog int

var x hotdog = 10

func main() {
	y := 42
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v", y, y)

	// Não posso usar y = x diretamente porque são tipos diferentes
	// apesar de ambos serem inteiros e int um tipo subjacente e hotdog
	//y = x

	// para fazer isso preciso converter uma delas para ficar do mesmo tipo
	y = int(x)
}
