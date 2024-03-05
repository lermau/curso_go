package main

import "fmt"

func main() {

	fmt.Println("PROGRAMA PARA EL CALCULO DE SUELDO")
	var anios_antigüedad int
	var sueldo_actual float64
	var sueldo_total float64
	fmt.Printf("Por favor ingresa tus años de antigüedad ")
	fmt.Scan(&anios_antigüedad)
	fmt.Printf("Por favor ingresa tu sueldo actual ")
	fmt.Scan(&sueldo_actual)
	if sueldo_actual >= 500 {
		fmt.Println("Tu sueldo es ", sueldo_actual)
	} else if sueldo_actual < 500 && anios_antigüedad > 10 {
		sueldo_total = sueldo_actual * 1.20
		fmt.Println("SUeldo total actualizado es ", sueldo_total)
	} else {
		sueldo_total = sueldo_actual * 1.05
		fmt.Println(" Tu sueldo actualizado es de ", sueldo_total)
	}
}
