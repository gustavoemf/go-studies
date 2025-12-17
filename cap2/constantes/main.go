package main

// Podem ser declaradas com uma declaração de tipagem enquanto variáveis não
const (
	oi = "buenas tardes" // Definido seu tipo apenas quando for utilizado
	x  = 10
)

var tchau = "bye bye" // Definido seu tipo apenas no momento em que o programa é executado
var y = 10
var z float64
var s string

func main() {

	z = x // Consigo atribuir isso pela flexibilidade de não atribuir um tipo à constante
	//z = y // Não consigo atruir y ao z pois ele tem tipos estáticamente diferentes

	// Go não faz conversão implícita entre tipos incompatíveis
	// Constantes não tipadas podem se adaptar somente entre tipos compatíveis.
	// Numérico → numérico
	// (int, float64, int64, etc)
	// Inteiros → inteiros maiores
	// Inteiros → floats
	// s = x
}
