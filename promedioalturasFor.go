package main

import "fmt"

func main() {
	var numero_personas, i int
	var estaturas, suma_estaturas, promedio_estaturas float64
	fmt.Println("Programa que calcula el promedio de N personas")
	fmt.Printf("Introduce el número de personas a calcular su promedio de estaturas ")
	fmt.Scan(&numero_personas)
	i = 0
	estaturas = 0
	suma_estaturas = 0
	for i < numero_personas {
		fmt.Println("Introduce la estatura ", i+1)
		fmt.Scan(&estaturas)
		suma_estaturas = suma_estaturas + estaturas
		i++

	}

	promedio_estaturas = suma_estaturas / float64(numero_personas)
	fmt.Println("El promedio de estaturas es ", promedio_estaturas)

}
