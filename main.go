package main

import (
	"fmt"

	"github.com/wildimice/godesde0/variables"
)

func main() {
	//fmt.Println("Hello, World!")
	//variables.MuestroEnteros()

	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(12345)
	fmt.Println("Estado:", estado)
	fmt.Println("Texto:", texto)
}
