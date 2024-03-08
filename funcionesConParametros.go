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

func numeroLimite(num int) {
	for i := 0; i <= num; i++ {
		fmt.Print(" ", i)
	}

}

func main() {
	clearConsole()
	var numero int
	fmt.Println("Introduce tu número límite ")
	fmt.Scan(&numero)
	numeroLimite(numero)
	fmt.Println("\n")

}
