package main

import (
	"fmt"

	"github.com/wildmice/godesde0/ejercicios"
)

func main() {

	valor, texto := ejercicios.Devuelve2Valores("50")
	fmt.Println("Valor devuelto:", valor)
	fmt.Println("Texto devuelto:", texto)

}
