package ejercicios

import ( 
	"fmt" 
	"strconv"
)


func Devuelve2Valores(string texto) (int, string) {
	// Devuelve dos valores, un entero y un string

	auxtexto := strconv.Atoi(texto)

	if auxtexto > 100 {
		texto := "El número es mayor que 100"
	} else {
		texto := "El número es menor a 100")
	}

	return 123, texto
}
