package main

import "fmt"

//Definir un slice vacío. En una estructura repetitiva ingresar enteros
//y añadirlos al slice. Cuando se ingrese el número -1
//finalizar la carga y mostrar cuantos números ingresados son mayor al promedio de los valores ingresados.

func main() {

	fmt.Println(" Programa de slices")
	var mi_slice []float64
	var numero, suma float64
	contador := 0

	for {
		fmt.Print(" Intruduce un número o ingresa -1 para terminar la carga ")
		fmt.Scan(&numero)
		if numero == -1 {
			break
		} else {
			mi_slice = append(mi_slice, float64(numero))
			contador++
		}

	}
	fmt.Println("Los números ingresados son", mi_slice)

	//Calculo del promedio
	for i := 0; i < len(mi_slice); i++ {
		suma = suma + mi_slice[i]
	}
	promedio := suma / float64(contador)
	fmt.Println("El promedio de tu slice es ", promedio)

	//Mayores que el promedio
	numero_de_mayores := 0
	var mas_grandes []float64
	for i := 0; i < len(mi_slice); i++ {
		if mi_slice[i] > promedio {
			mas_grandes = append(mas_grandes, mi_slice[i])
			numero_de_mayores++
		}
	}

	fmt.Println("En tu arreglo existen ", numero_de_mayores, "números más grande que el promedio del mismo")
	fmt.Println(mas_grandes)
	fmt.Println("")

}
