package main

import "fmt"

func main() {
	fmt.Println(" Programa que te dice si un númeo es de una o 2 cifras")
	var num1 int
	fmt.Println("Introduce tu numero")
	fmt.Scan(&num1)

	if num1 > 9 {
		fmt.Println("Tú número es de 2 cifras")

	} else {
		fmt.Println("Tú número es de una cifra")
	}

}
