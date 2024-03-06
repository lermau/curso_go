package main

import "fmt"

func main() {
	var arreglo1 [5]float64
	var arreglo2 [5]float64
	var suma_arreglos [5]float64
	fmt.Println("Programa para sumar 2 arreglos de tamaño 5")

	for i := 0; i < (len(arreglo1)); i++ {
		fmt.Print("Introduce tú valor ", i, " del arreglo 1 ")
		fmt.Scan(&arreglo1[i])
	}

	for i := 0; i < (len(arreglo2)); i++ {
		fmt.Print("Introduce tú valor ", i, " del arreglo 2 ")
		fmt.Scan(&arreglo2[i])
	}

	for i := 0; i < (len(suma_arreglos)); i++ {
		suma_arreglos[i] = arreglo1[i] + arreglo2[i]

	}

	fmt.Print(" El arreglo suma es = ", suma_arreglos)
}
