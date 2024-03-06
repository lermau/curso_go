package main

import "fmt"

//Desarrollar un programa que permita ingresar un arreglo de 8 elementos, e informe:
//El valor acumulado de todos los elementos del arreglo.
//El valor acumulado de los elementos del arreglo que sean mayores a 36.
//Cantidad de valores mayores a 50.
func main() {
	var elementos [8]float64
	var acumulado float64
	var acumulado_mayores_36 float64
	var mayores_de_50 = 0
	acumulado_mayores_36 = 0
	acumulado = 0

	for i := 0; i < (len(elementos)); i++ {
		fmt.Print("Introduce tu elemento ")
		fmt.Scan(&elementos[i])
		acumulado = acumulado + elementos[i]

		if elementos[i] > 36 {
			acumulado_mayores_36 = acumulado_mayores_36 + elementos[i]
		}

		if elementos[i] > 50 {
			mayores_de_50++
		}
	}
	fmt.Println("Acumulado ", acumulado)
	fmt.Println("Acumulado mayores 36  ", acumulado_mayores_36)
	fmt.Println("Datos mayores a 50 ", mayores_de_50)

}
