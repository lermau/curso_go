package main

import "fmt"

func main() {
	var mi_matriz [2][5]int

	for col := 0; col < 5; col++ {
		for fila := 0; fila < 2; fila++ {
			fmt.Print("Introduce tu valor ", col, fila, "--> ")
			fmt.Scan(&mi_matriz[fila][col])
		}

	}

	for fila := 0; fila < 2; fila++ {
		for col := 0; col < 5; col++ {
			fmt.Print(mi_matriz[fila][col], " ")
		}
		fmt.Println("")
	}

	fmt.Println(mi_matriz)
}
