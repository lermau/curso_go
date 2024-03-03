package main

import "fmt"

func main() {
	fmt.Println(" Programa mayor de 3 números")
	var num1, num2, num3 float32
	fmt.Println("Introduce el primer número")
	fmt.Scan(&num1)
	fmt.Println("Introduce tu segundo número")
	fmt.Scan(&num2)
	fmt.Println("Introduce tu tercer número")
	fmt.Scan(&num3)
	var numeroMayor float32
	if num1 > num2 && num1 > num3 {
		numeroMayor = num1
	} else if num2 > num3 {
		numeroMayor = num2
	} else {
		numeroMayor = num3
	}
	fmt.Println("El número mayor es ", numeroMayor)

}
