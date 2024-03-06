package main

import "fmt"

func main() {
	var aprobados int
	var calificacion float64
	i := 0
	aprobados = 0
	for i < 10 {
		fmt.Println(" Introduce la calificación del alumno ", i+1)
		fmt.Scan(&calificacion)
		if calificacion > 6 {
			aprobados++
		}
		i++
	}
	fmt.Println("El número de alumnos aprobados es ", aprobados)
	fmt.Println("El número de alumnos reprobados es ", 10-aprobados)

}
