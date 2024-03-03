package main

import "fmt"

func main() {
	var num1, num2 float32
	fmt.Println(" Programa que dependiendo que número es mayor, ejectúa operaciones")
	fmt.Println("Introduce un número")
	fmt.Scan(&num1)
	fmt.Println("Introduce otro número")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println(" La suma de tus números es ", num1+num2)
		fmt.Println(" La diferencia de tus números es ", num1-num2)
	} else {
		fmt.Println(" La multiplicación  de tus números es ", num1*num2)
		fmt.Println(" La división de tus números es ", num1/num2)
	}
}
