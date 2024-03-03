package main

import "fmt"

func main() {
	var precio float32
	var cantidad int
	fmt.Println(" Programa para calcular cuenta total")
	fmt.Println(" Introduce el precio de tu artículo ")
	fmt.Scan(&precio)
	fmt.Println("Introduce la cantidad de articulos ")
	fmt.Scan(&cantidad)
	fmt.Println(" Tu cuenta total es de ", precio*float32(cantidad))

}
