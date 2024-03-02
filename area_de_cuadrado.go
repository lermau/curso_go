package main

import "fmt"

func main() {
	var lado_de_cuadrado, area_del_cuadrado float32
	fmt.Println("CALCULO DEL AREA DE UN CUADRADO")
	fmt.Print("Introduce el valo del lado de tu cuadrado ")
	fmt.Scan(&lado_de_cuadrado)
	area_del_cuadrado = lado_de_cuadrado * lado_de_cuadrado
	fmt.Println("El area de tu cuadrado es ", area_del_cuadrado)

}
