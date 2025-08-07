package ejercicios

import 	"strconv"

func Devuelve2Valores(texto string) (int, string) {
	// Devuelve dos valores, un entero y un string

	auxtexto, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Error: no se pudo convertir el texto a número" + err.Error()
	}

	if auxtexto > 100 {
		texto := "El número es mayor que 100"
	} else {
		texto := "El número es menor a 100")
	}

	return auxtexto, texto
}
