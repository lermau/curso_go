package main

import "fmt"

func main() {
	var num1, num2, num3, num4 float32
	fmt.Println("Programa que suma 2 números y multiplica otros 2")
	fmt.Println("Introduce el primer número")
	fmt.Scan(&num1)
	fmt.Println("Introduce el segundo número")
	fmt.Scan(&num2)
	fmt.Println("Introduce el tercer  número")
	fmt.Scan(&num3)
	fmt.Println("Introduce el cuarto número")
	fmt.Scan(&num4)
	suma1y2 := num1 + num2
	multiplicacion3y4 := num3 * num4
	fmt.Println(" La suma de los números 1 y 2 es = ", suma1y2)
	fmt.Println(" La multiplicación de los números 3 y 4 es = ", multiplicacion3y4)

}
