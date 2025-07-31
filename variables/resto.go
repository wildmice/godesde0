package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println("El nombre es:", Nombre)
	fmt.Println("El estado es:", Estado)
	fmt.Println("El sueldo es:", Sueldo)
	fmt.Println("La fecha es:", Fecha.Format("02-01-2006 15:04:05"))
}
