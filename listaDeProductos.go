package main

import "fmt"

func main() {
	fmt.Println("Programa con 2 listas")
	var productos [5]string
	var precios [5]float64

	for i := 0; i < 5; i++ {
		fmt.Print("Introduce el producto -----> ", i+1, " ")
		fmt.Scan(&productos[i])
		fmt.Print("Introduce el precio del producto ---> ", i+1, " ")
		fmt.Scan(&precios[i])
	}

	var mayores int
	mayores = 0
	for i := 1; i < 5; i++ {
		if precios[0] < precios[i] {
			mayores++
		}

	}

	fmt.Println(" Existen ", mayores, " productos mas caros que el producto ", productos[0])

}
