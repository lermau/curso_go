package main

import "fmt"

func main() {
	fmt.Println("PROGRAMA PARA CALCULAR EL PROMEDIO DE 4 NÚMEROS")
	var num1, num2, num3, num4 float32
	fmt.Println("Introduce el primer número")
	fmt.Scan(&num1)
	fmt.Println("Introduce el segundo número")
	fmt.Scan(&num2)
	fmt.Println("Introduce el tercer  número")
	fmt.Scan(&num3)
	fmt.Println("Introduce el cuarto número")
	fmt.Scan(&num4)
	promedio := (num1 + num2 + num3 + num4) / 4
	fmt.Println("El promedio de tus números es ", promedio)
}
