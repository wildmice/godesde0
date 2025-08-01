package main

import (
	"fmt"
	//"github.com/wildmice/godesde0/ejercicios"
	"github.com/wildmice/godesde0/ejercicios"
)

func mainx() {
	//fmt.Println("Hello, World!")
	//variables.MuestroEnteros()

	//variables.RestoVariables()

	//estado, texto := variables.ConviertoaTexto(12345)
	//fmt.Println("Estado:", estado)
	//fmt.Println("Texto:", texto)

	// Ejemplo de función que devuelve dos valores

	/*
		os := runtime.GOOS
		if os == "windows" {
			fmt.Println("Estás en Windows")
		} else if os == "linux" {
			fmt.Println("Estás en Linux")
		} else if os == "darwin" {
			fmt.Println("Estás en MacOS")
		} else {
			fmt.Println("Sistema operativo no reconocido")
		}

		switch os {
		case "windows":
			fmt.Println("Estás en Windows")
		case "linux":
			fmt.Println("Estás en Linux")
		case "darwin":
			fmt.Println("Estás en MacOS")
		default:
			fmt.Println("Sistema operativo no reconocido")
		}
	*/

	ejercicios.Devuelve2Valores("150")

	valor, texto := ejercicios.Devuelve2Valores("50")
	fmt.Println("Valor devuelto:", valor)
	fmt.Println("Texto devuelto:", texto)

}
