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

func cuadradoDeUnNumero() {
	var numero float64
	fmt.Print("Introduce el número a elevar al cuadrado --> ")
	fmt.Scan(&numero)
	fmt.Println("Tu numero elevado al cuadrado es ", numero*numero, "\n")
}

func productoDe2Numeros() {
	var num1, num2 int
	fmt.Print("Introduce tu primer número ")
	fmt.Scan(&num1)
	fmt.Print("Introduce tu primer número ")
	fmt.Scan(&num2)
	fmt.Print("El producto de tus números es ", num1*num2, "\n")

}

func main() {
	clearConsole()
	cuadradoDeUnNumero()
	productoDe2Numeros()

}
