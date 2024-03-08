package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func esVocal(letra string) {
	if len(letra) > 1 {
		fmt.Println("Por favor introduce solo una caracter ")
		esVocal(solicitaLetra())
	} else if letra == "a" || letra == "e" || letra == "i" || letra == "o" || letra == "u" ||
		letra == "A" || letra == "E" || letra == "I" || letra == "O" || letra == "U" {
		fmt.Println("Tu caracter es una vocal")
	} else {
		fmt.Println("Tu caracter no es una vocal")
	}
}

func solicitaLetra() string {
	var letra string
	fmt.Print("Introduce un caracter ")
	fmt.Scan(&letra)
	return letra
}

func main() {
	clearConsole()
	esVocal(solicitaLetra())
}
