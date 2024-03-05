package main

import "fmt"

func main() {
	fmt.Println("Programa para determinar el cuadrante dada una coordenada")
	var x, y int
	fmt.Printf("Introduce la coordenada x ")
	fmt.Scan(&x)
	fmt.Printf("Introduce la coordenada y ")
	fmt.Scan(&y)
	if x > 0 && y > 0 {
		fmt.Println("Tu coordenada pertenece al cuadrante 1")
	} else if x < 0 && y > 0 {
		fmt.Println("Tu coordenada pertenece al cuadrante 2")
	} else if x < 0 && y < 0 {
		fmt.Println("Tu coordenada pertenece al cuadrante 3")
	} else {
		fmt.Println("Tu coordenada pertenece al cuadrante 4")
	}
}
