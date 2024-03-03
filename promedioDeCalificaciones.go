package main

import "fmt"

func main() {
	fmt.Println(" Programa que dependiendo tu promedio te dice si aprobaste o no")
	var num1, num2, num3, promedio float32
	fmt.Println("Introduce tu primera calificaciónm")
	fmt.Scan(&num1)
	fmt.Println("Introduce  tu segunda calificación")
	fmt.Scan(&num2)
	fmt.Println("Introduce tu tercera calificación")
	fmt.Scan(&num3)
	promedio = (num1 + num2 + num3) / 3
	fmt.Println("Tu promedio es de ", promedio)
	if promedio >= 7 {
		fmt.Println(" Has aprobado el año")

	} else {
		fmt.Println("Lo siento tienes que repetir el año")
	}

}
