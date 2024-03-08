package main

import "fmt"

func main() {
	var nombre string
	fmt.Print("Introduce por favor tu nombre ")
	fmt.Scan(&nombre)
	fmt.Print("Hola mi estimado ", nombre)
}
