package main

import "fmt"

func main() {

	fmt.Println("Programa que calcula la suma de los últimos 5 números agregados de 10")
	var suma, numero float64
	var i, num int
	suma = 0
	//num = 1

	for i = 0; i < 10; i++ {
		fmt.Print(" Introduce tu número ", i+1)
		fmt.Scan(&numero)
		if i > 4 {
			suma = suma + numero
		}
		num++
	}
	fmt.Println("La suma de tus últimos 5 números es ", suma)
}
