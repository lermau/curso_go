package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Encabezado(cadena string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("-" + repeatChar(len(cadena)+2, "-") + "-")
	fmt.Println(cadena)
	fmt.Println("-" + repeatChar(len(cadena)+2, "-") + "-")
}

func repeatChar(count int, char string) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
