package main

import (
	"encabezado"
	"fmt"
)

func main() {
	encabezado.Encabezado("Saludo personalizado")
	var nombre string
	fmt.Print("Introduce por favor tu nombre ")
	fmt.Scan(&nombre)
	fmt.Print("Hola mi estimado ", nombre)
}
