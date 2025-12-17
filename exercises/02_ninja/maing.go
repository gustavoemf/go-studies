package main

import "fmt"

var x int
var y string
var z bool

func main() {
	//	Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
	//
	// Identificador "x" deverá ter tipo int
	// Identificador "y" deverá ter tipo string
	// Identificador "z" deverá ter tipo bool
	// Na função main:
	// Demonstre os valores de cada identificador
	// O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

	fmt.Printf("Valor zero de inteiros: %v\n", x)
	fmt.Printf("Valor zero de strings: %v\n", y)
	fmt.Printf("Valor zero de bool: %v", z)
}
